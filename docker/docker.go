package docker

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
	mobyClient, _ = client.NewClientWithOpts()
}

func GetContainers() []Container {
	var containers []Container
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	containerSet, err := mobyClient.ContainerList(ctx, types.ContainerListOptions{All: false})
	if err != nil {
		log.Printf("Error fetching the container list: %s", err.Error())
		return containers
	}
	for _, c := range containerSet {
		var container Container
		var v *types.StatsJSON
		var mem, memLimit float64

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
		container.Created = time.Unix(c.Created*1000, 0)
		container.CPU = calculateCPUPercentUnix(previousCPU, previousSystem, v)
		container.MemUsage = calculateMemUsageUnixNoCache(v.MemoryStats)
		container.MemAllowed = float64(v.MemoryStats.Limit)
		container.MemPercent = calculateMemPercentUnixNoCache(memLimit, mem)
		// netRx, netTx := calculateNetwork(v.Networks)
		container.Uptime = time.Since(time.UnixMilli(c.Created * 1000)).Milliseconds()
		containers = append(containers, container)
	}

	return containers
}
