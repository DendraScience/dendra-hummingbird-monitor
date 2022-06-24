package k8s

import (
	"context"
	"encoding/json"
	"time"

	"log"

	"github.com/docker/docker/api/types"
	"github.com/moby/moby/client"
)

var (
	configSet  bool = false
	mobyClient *client.Client
)

func init() {
	// 	version string
	//	80     // custom http headers configured by users.
	//	81     customHTTPHeaders map[string]string
	//	82     // manualOverride is set to true when the version was set by users.
	//	83     manualOverride bool
	mobyClient, _ = client.NewClientWithOpts(client.WithVersion("1.40"))
}

func GetContainers() []Container {
	var containers []Container
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

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
