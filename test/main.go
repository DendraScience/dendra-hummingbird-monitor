package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"log"

	. "github.com/DendraScience/dendra_hummingbird_monitor/k8s"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	k8sClient *kubernetes.Clientset
)

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/ssmith/.kube/config")
	if err != nil {
		panic(err)
	}
	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}

func main() {
	//	x, _ := getContainers(context.TODO())
	//	for _, c := range x {
	//		fmt.Printf("%s\n", c.String())
	//	}
	GetClusterContainers2()
	//	fmt.Printf("%v\n", x)
	//	fmt.Printf("Number of containers: %d\n", len(x))
}

func getContainers(ctx context.Context) (containers []Container, err error) {
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
			c.Name = container.Name
			c.Image = container.Image
			c.Created = running.StartedAt.Time
			c.Uptime = int64(time.Now().Sub(c.Created).Seconds())
			x, _ := pod.Spec.Containers[i].Resources.Limits.Memory().AsInt64()
			c.MemAllowed = int(x)
			x, _ = pod.Spec.Containers[i].Resources.Limits.Cpu().AsInt64()
			c.CPU = float64(x)
			cmap[c.ID] = c
		}
	}
	for _, v := range cmap {
		containers = append(containers, v)
	}
	return
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
func GetClusterContainers2() []Container {
	var containers []Container
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	nodes, err := GetNodeIDs(ctx)
	metricLines := []string{}
	if err != nil {
		log.Printf("Error fetching node IDs: %v\n", err)
	}
	for _, nodeID := range nodes {
		d, err := k8sClient.RESTClient().
			Get().AbsPath(fmt.Sprintf("/api/v1/nodes/%s/proxy/metrics/cadvisor",
			nodeID)).DoRaw(ctx)
		if err != nil {
			fmt.Printf("Error sending rest: %v\n", err)
			continue
		}
		metricLines = append(metricLines, FilterCAdvisor(d)...)
	}
	for _, m := range metricLines {
		fmt.Println(m)
	}
	//fmt.Printf("Data received: %v\n", metricLines)
	return containers
	//	podMetrics, err := DmetricsClient.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	//	if err != nil {
	//		fmt.Println("Error collecting metrics:", err)
	//		return containers
	//	}
	//	for _, podMetric := range podMetrics.Items {
	//
	//		//var ok bool
	//		//	if c, ok = cmap[containerName]; !ok {
	//		//			continue
	//		//		}
	//		podContainers := podMetric.Containers
	//		for _, container := range podContainers {
	//			c := Container{}
	//			i, _ := container.Usage.Cpu().AsInt64()
	//			c.CPU = float64(i) / c.CPU
	//			i, _ = container.Usage.Memory().AsInt64()
	//			c.MemUsage = int(i)
	//			c.Name = container.Name
	//			c.MemPercent = float64(c.MemUsage) / float64(c.MemAllowed)
	//			containers = append(containers, c)
	//		}
	//	}
	//
	//	return containers
}

type NodeID string

// sum (rate (container_cpu_usage_seconds_total{id="/"}[1m])) / sum (machine_cpu_cores) * 100
func calculateCPUUsage() {

}

// Pull out all variables that we need and drop the rest
func FilterCAdvisor(in []byte) []string {
	input := string(in)
	split := strings.Split(input, "\n")
	variables := []string{"container_memory_usage_bytes",
		"container_spec_memory_limit_bytes",
		"machine_cpu_cores",
		"container_cpu_usage_seconds_total"}
	toKeep := []string{}
	for _, s := range split {
		for _, variable := range variables {
			if strings.HasPrefix(s, variable) {
				// ignore metrics relating to drivers, kubelet, etc. since they
				// don't relate to a pod or container
				if strings.Contains(s, "name=\"\"") || strings.Contains(s, "pod=\"\"") {
					break
				}
				toKeep = append(toKeep, s)
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
