package k8s

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	k8sClient    *kubernetes.Clientset
	containerMap map[string]Container
)

func Init() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	containerMap = make(map[string]Container)
}

func getContainerMap(ctx context.Context) (containers map[string]Container, err error) {
	cmap := make(map[string]Container)
	var pods *v1.PodList
	pods, err = k8sClient.CoreV1().Pods(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error collecting pods: %v\n", err)
		return
	}
	for _, pod := range pods.Items {
		for i, container := range pod.Status.ContainerStatuses {
			var c Container
			running := container.State.Running
			// we only want to consider running containers
			if running == nil {
				continue
			}
			c.ID = strings.ReplaceAll(container.ContainerID, "containerd://", "")
			c.Name = pod.Name + "_" + container.Name + "_" + pod.Namespace
			c.Image = container.Image
			c.Created = running.StartedAt.Time
			c.Uptime = int64(time.Now().Sub(c.Created).Seconds())
			x, _ := pod.Spec.Containers[i].Resources.Limits.Memory().AsInt64()
			c.MemAllowed = int(x)
			x, _ = pod.Spec.Containers[i].Resources.Limits.Cpu().AsInt64()
			c.CPUAllowed = float64(x)
			cmap[c.ID] = c
		}
	}
	return cmap, nil
}

type NodeMetric struct {
	Metric string
	Node   string
}

// Get a listing of all the nodes
// Then get a list of all the pods
// extract each container from each pod
// fill in all definitions for the containers
// for each node, grab the RAW data using a rest client
// filter all data out to the required parameters for node
// append all strings into a single slice
// parse all filtered strings into individual types
// create maps for filtered strings and match against containers using IDs or names
func GetClusterContainers() []Container {
	var containers []Container
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	nodes, err := GetNodeIDs(ctx)
	metricLines := []NodeMetric{}
	if err != nil {
		log.Printf("Error fetching node IDs: %v\n", err)
		return containers
	}
	// for each node, hit the RAW api and collect the metrics we need
	for _, nodeID := range nodes {
		d, err := k8sClient.RESTClient().
			Get().AbsPath(fmt.Sprintf("/api/v1/nodes/%s/proxy/metrics/cadvisor",
			nodeID)).DoRaw(ctx)
		if err != nil {
			fmt.Printf("Error sending rest: %v\n", err)
			continue
		}
		// filter the results and add them to our slice of metrics
		metricLines = append(metricLines, FilterCAdvisor(d, string(nodeID))...)
	}
	memUsage, memSpec, cpuTime := categorizeMetrics(metricLines)
	cmap, err := getContainerMap(ctx)
	if err != nil {
		log.Printf("Error retrieving containers: %v\n", err)
		return containers
	}
	// pull the prev variables into our new map
	for k, v := range containerMap {
		if c, ok := cmap[k]; ok {
			c.PrevCPUMS = v.PrevCPUMS
			c.PrevReadingTime = v.PrevReadingTime
			cmap[k] = c
		}
	}
	// now range over the new map and pull in metrics
	for k, v := range cmap {
		if val, ok := memUsage[k]; ok {
			v.MemUsage = int(val.Value)
			v.Node = val.Node
		}
		if val, ok := memSpec[k]; ok {
			v.MemAllowed = int(val.Value)
			v.Node = val.Node
		}
		if val, ok := cpuTime[k]; ok {
			v.NewCPUMS = val.Value
			v.NewReadingTime = val.TimeStamp
			v.Node = val.Node
		}
		v.MemPercent = float64(v.MemUsage) / float64(v.MemAllowed)
		if v.MemAllowed == 0 {
			v.MemPercent = 0
		}
		v.CPUUsage = calculateCPUUsage(v.PrevReadingTime,
			v.NewReadingTime,
			v.PrevCPUMS,
			v.NewCPUMS,
			v.CPUAllowed)
		diff := v.NewCPUMS - v.PrevCPUMS
		v.PrevCPUMS = v.NewCPUMS
		v.NewCPUMS = diff
		v.PrevReadingTime = v.NewReadingTime
		cmap[k] = v
		containers = append(containers, v)
	}
	containerMap = cmap
	return containers
}

type NodeID string

// sum (rate (container_cpu_usage_seconds_total{id="/"}[1m])) / sum (machine_cpu_cores) * 100
// https://stackoverflow.com/questions/40327062/how-to-calculate-containers-cpu-usage-in-kubernetes-with-prometheus-as-monitori#40391872
func calculateCPUUsage(prevTime, curTime time.Time, prevSeconds, curSeconds int64, allowed float64) float64 {
	// don't do CPU calculations if things haven't changed or this is our first read
	if curTime.Equal(prevTime) || prevTime.Equal(time.Time{}) {
		return 0
	}
	diff := curSeconds - prevSeconds
	duration := curTime.Sub(prevTime)

	return float64(diff) / float64(duration.Milliseconds())
}

func categorizeMetrics(metrics []NodeMetric) (memUsage, memSpec, cpuTime map[string]Metric) {
	// See the following example metric:
	//
	// container_memory_usage_bytes{container="main",id="/system.slice/containerd.service/kubepods-burstable-pod811bce1a_b926_49c6_af1c_115c2a06df25.slice:cri-containerd:f81d390115234186549def0dd6aa82f15c4afa264b8394d6e7a1559bc04b753f",image="docker.io/library/influxdb:1.8.10",name="f81d390115234186549def0dd6aa82f15c4afa264b8394d6e7a1559bc04b753f",namespace="default",pod="influxdb-v1-cdfw-m1-0"} 3.269632e+08 1656209471742
	//
	// initialize maps
	memUsage = make(map[string]Metric)
	memSpec = make(map[string]Metric)
	cpuTime = make(map[string]Metric)
	// define some regular expressions here for extracting names and stripping metadata
	nameFinder := regexp.MustCompile(`name="(\w+)"`)
	valueStripper := regexp.MustCompile(`.*{.*}`)
	for _, m := range metrics {
		// first, find and extract the containerd id for the metric
		matches := nameFinder.FindStringSubmatch(m.Metric)
		if len(matches) != 2 {
			log.Printf("Error: could not find name in %s\n", m)
			continue
		}
		name := matches[1]
		if strings.HasPrefix(m.Metric, "container_spec_memory_limit_bytes") {
			var metric Metric
			metric.ID = name
			metric.Node = m.Node
			m.Metric = valueStripper.ReplaceAllLiteralString(m.Metric, "")
			m.Metric = strings.TrimSpace(m.Metric)
			flt, _, err := big.ParseFloat(m.Metric, 10, 0, big.ToNearestEven)
			if err != nil {
				log.Printf("Error parsing value for mem spec: %v :: %s\n", err, m)
				continue
			}
			i := new(big.Int)
			i, _ = flt.Int(i)
			metric.Value = i.Int64()
			memSpec[name] = metric
		} else if strings.HasPrefix(m.Metric, "container_memory_usage_bytes") {
			var metric Metric
			metric.ID = name
			metric.Node = m.Node
			m.Metric = valueStripper.ReplaceAllLiteralString(m.Metric, "")
			m.Metric = strings.TrimSpace(m.Metric)
			values := strings.Split(m.Metric, " ")
			if len(values) != 2 {
				log.Printf("Error: Expected usage string to have two values, but one was found: %s\n", m)
				continue
			}
			millis, err := strconv.Atoi(values[1])
			if err != nil {
				log.Printf("Error extracting ts: %v: %s\n", err, m)
				continue
			}
			metric.TimeStamp = time.UnixMilli(int64(millis))
			flt, _, err := big.ParseFloat(values[0], 10, 0, big.ToNearestEven)
			if err != nil {
				log.Printf("Error parsing value for mem usage: %v :: %s\n", err, m)
				continue
			}
			i := new(big.Int)
			i, _ = flt.Int(i)
			metric.Value = i.Int64()
			memUsage[name] = metric
		} else if strings.HasPrefix(m.Metric, "container_cpu_usage_seconds_total") {
			var metric Metric
			metric.ID = name
			metric.Node = m.Node
			m.Metric = valueStripper.ReplaceAllLiteralString(m.Metric, "")
			m.Metric = strings.TrimSpace(m.Metric)
			values := strings.Split(m.Metric, " ")
			if len(values) != 2 {
				log.Printf("Error: Expected usage string to have two values, but one was found: %s\n", m)
				continue
			}
			millis, err := strconv.Atoi(values[1])
			if err != nil {
				log.Printf("Error extracting ts: %v: %s\n", err, m)
				continue
			}
			metric.TimeStamp = time.UnixMilli(int64(millis))
			seconds, err := strconv.ParseFloat(values[0], 32)
			if err != nil {
				log.Printf("Error extracting seconds: %v: %s\n", err, m)
				continue
			}
			metric.Value = int64(seconds * 1000)
			cpuTime[name] = metric
		}
	}
	return
}

// Pull out all variables that we need and drop the rest
func FilterCAdvisor(in []byte, node string) []NodeMetric {
	input := string(in)
	split := strings.Split(input, "\n")
	variables := []string{
		"container_memory_usage_bytes",
		"container_spec_memory_limit_bytes",
		//		"machine_cpu_cores",
		"container_cpu_usage_seconds_total",
	}
	toKeep := []NodeMetric{}
	for _, s := range split {
		for _, variable := range variables {
			if strings.HasPrefix(s, variable) {
				// ignore metrics relating to drivers, kubelet, etc. since they
				// don't relate to a pod or container
				if strings.Contains(s, "name=\"\"") || strings.Contains(s, "pod=\"\"") {
					break
				}
				toKeep = append(toKeep, NodeMetric{Node: node, Metric: s})
				break
			}
		}
	}
	return toKeep
}

func GetNodeIDs(ctx context.Context) ([]NodeID, error) {
	var nodeList []NodeID
	nodes, err := k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nodeList, err
	}
	for _, n := range nodes.Items {
		nodeList = append(nodeList, NodeID(n.Name))
	}
	return nodeList, nil
}
