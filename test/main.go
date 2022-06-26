package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"log"

	. "github.com/DendraScience/dendra_hummingbird_monitor/k8s"
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
	x := GetClusterContainers2()
	fmt.Printf("%v\n", x)
	fmt.Printf("Number of containers: %d\n", len(x))
}

func GetClusterContainers2() []Container {
	var containers []Container
	cmap := make(map[string]Container)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	pods, err := k8sClient.CoreV1().Pods(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error collecting pods: %v\n", err)
		return containers
	}
	for _, pod := range pods.Items {
		for i, container := range pod.Status.ContainerStatuses {
			containerName := pod.Name + "::" + container.Name
			var c Container
			running := container.State.Running
			if running == nil {
				continue
			}
			c.Image = container.Image
			c.Created = running.StartedAt.Time
			c.Uptime = int64(time.Now().Sub(c.Created).Seconds())
			x, _ := pod.Spec.Containers[i].Resources.Limits.Memory().AsInt64()
			c.MemAllowed = int(x)
			x, _ = pod.Spec.Containers[i].Resources.Limits.Cpu().AsInt64()
			c.CPU = float64(x)
			cmap[containerName] = c
		}
	}
	fmt.Printf("Cmap: %v\n", cmap)
	fmt.Printf("Number of containers: %d\n", len(cmap))

	d, err := k8sClient.RESTClient().Get().AbsPath("/api/v1/nodes/den-shasta-k8s-cp-01/proxy/metrics/cadvisor").DoRaw(ctx)

	if err != nil {
		fmt.Printf("Error sending rest: %v\n", err)
	}
	fmt.Printf("Data received: %s\n", string(d))
	return containers
	podMetrics, err := DmetricsClient.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error collecting metrics:", err)
		return containers
	}
	for _, podMetric := range podMetrics.Items {

		//var ok bool
		//	if c, ok = cmap[containerName]; !ok {
		//			continue
		//		}
		podContainers := podMetric.Containers
		for _, container := range podContainers {
			c := Container{}
			i, _ := container.Usage.Cpu().AsInt64()
			c.CPU = float64(i) / c.CPU
			i, _ = container.Usage.Memory().AsInt64()
			c.MemUsage = int(i)
			c.Name = container.Name
			c.MemPercent = float64(c.MemUsage) / float64(c.MemAllowed)
			containers = append(containers, c)
		}
	}

	return containers
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
