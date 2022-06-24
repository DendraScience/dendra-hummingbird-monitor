package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"log"

	"github.com/docker/docker/api/types"
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
	config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	k8sClient, _ = kubernetes.NewForConfig(config)
	metricsClient, _ = metrics.NewForConfig(config)
}

func GetClusterContainers() []Container {
	var containers []Container
	cmap := make(map[string]Container)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error:", err)
		return containers
	}
	for _, podMetric := range podMetrics.Items {
		podContainers := podMetric.Containers
		for _, container := range podContainers {
			c := Container{}
			i, _ := container.Usage.Cpu().AsInt64()
			c.CPU = float64(i)
			i, _ = container.Usage.Memory().AsInt64()
			c.MemUsage = int(i)
			c.Name = container.Name
			cmap[container.Name] = c
		}

	}

	containerSet, err := mobyClient.ContainerList(ctx, types.ContainerListOptions{All: false, Limit: -1})
	if err != nil {
		log.Printf("Error fetching the container list: %s", err.Error())
		return containers
	}
	for _, c := range containerSet {
		var container Container
		var v *types.StatsJSON

		stats, err := mobyClient.ContainerStatsOneShot(ctx, c.ID)
		if err != nil {
			log.Printf("Error fetching the container stats: %s", err.Error())
			continue
		}

		dec := json.NewDecoder(stats.Body)
		if err := dec.Decode(&v); err != nil {
			log.Printf("Error decoding the container stats: %s", err.Error())
		}
		previousCPU := v.PreCPUStats.CPUUsage.TotalUsage
		previousSystem := v.PreCPUStats.SystemUsage

		container.ID = c.ID
		if len(c.Names) > 0 {
			container.Name = c.Names[0]
		}
		container.Image = c.Image
		container.Created = time.Unix(c.Created, 0)
		container.CPU = calculateCPUPercentUnix(previousCPU, previousSystem, v) / 100
		container.MemUsage = int(calculateMemUsageUnixNoCache(v.MemoryStats))
		container.MemAllowed = int(v.MemoryStats.Limit)
		container.MemPercent = calculateMemPercentUnixNoCache(float64(v.MemoryStats.Limit), calculateMemUsageUnixNoCache(v.MemoryStats)) / 100
		// netRx, netTx := calculateNetwork(v.Networks)
		container.Uptime = int64(time.Since(container.Created).Seconds())
		containers = append(containers, container)
	}

	return containers
}
