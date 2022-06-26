package k8s

import (
	"context"
	"fmt"
	"time"

	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

var (
	metricsClient *metrics.Clientset
	k8sClient     *kubernetes.Clientset
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
	metricsClient, err = metrics.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}

func GetClusterContainers() []Container {
	var containers []Container
	cmap := make(map[string]Container)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error collecting metrics:", err)
		//	return containers
	}
	for _, podMetric := range podMetrics.Items {
		podContainers := podMetric.Containers
		for _, container := range podContainers {
			c := Container{}
			i, _ := container.Usage.Cpu().AsInt64()
			c.CPUUsage = float64(i)
			i, _ = container.Usage.Memory().AsInt64()
			c.MemUsage = int(i)
			c.Name = container.Name
			cmap[container.Name] = c
		}
	}
	fmt.Printf("Cmap: %v\n", cmap)
	pods, err := k8sClient.CoreV1().Pods(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error collecting pods: %v\n", err)
		return containers
	}
	for _, pod := range pods.Items {
		for i, container := range pod.Status.ContainerStatuses {
			containerName := container.Name
			var ok bool
			var c Container
			if c, ok = cmap[containerName]; !ok {
				continue
			}
			running := container.State.Running
			if running == nil {
				continue
			}
			c.Image = container.Image
			c.Created = running.StartedAt.Time
			c.Uptime = int64(time.Now().Sub(c.Created).Seconds())
			x, _ := pod.Spec.Containers[i].Resources.Limits.Memory().AsInt64()
			c.MemAllowed = int(x)
			c.MemPercent = float64(c.MemUsage) / float64(c.MemAllowed)
			x, _ = pod.Spec.Containers[i].Resources.Limits.Cpu().AsInt64()
			c.CPUUsage = c.CPUUsage / float64(x)
			containers = append(containers, c)
		}
	}
	return containers
}
