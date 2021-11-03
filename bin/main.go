package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DendraScience/dendra_hummingbird_monitor/config"
	"github.com/DendraScience/dendra_hummingbird_monitor/disk"
	"github.com/DendraScience/dendra_hummingbird_monitor/docker"
	"github.com/DendraScience/dendra_hummingbird_monitor/pkg"
	"github.com/DendraScience/dendra_hummingbird_monitor/proc"
	"github.com/DendraScience/dendra_hummingbird_monitor/publish"

	. "github.com/DendraScience/dendra_hummingbird_monitor/types"

	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
)

const (
	SemVer  string = "1.0.0"
	Package string = "dendra_hummingbird_monitor"
)

var (
	Authors   string
	BuildNo   string
	BuildTime string
	GitCommit string
	Tag       string

	hostname   string
	kubeConfig *string
	version    = flag.Bool("version", false, "Get detailed version string")
)

func init() {
	flag.Parse()

	Authors = strings.ReplaceAll(Authors, "SpAcE", " ")
	Tag = strings.ReplaceAll(Tag, ";", "; ")

	if GitCommit == "" || BuildTime == "" {
		log.Fatalf("Binary built improperly. Version variables not set!")
	}
	fmt.Printf("%s Version information:\n|| Authors: %s\n|| Commit: %s\n|| Tag: %s\n|| Build No: %s\n|| Build Date: %s\n", Package, Authors, GitCommit, Tag, BuildNo, BuildTime)

	if *version {
		os.Exit(0)
	} else {
		fmt.Printf("Initialization success...\n")
	}
	hostname = proc.GetHostname()

}
func main() {
	for {
		var stats QuarterHourly
		var err error
		startTime := time.Now()
		stats.Hostname = hostname

		stats.DiskFree = int64(disk.GetDiskUsagePercentage())
		stats.NumPackages = pkg.GetInstalledPackageCount()
		stats.Timestamp = startTime
		stats.UpdatesAvail = pkg.GetNumAvailUpdates()

		stats.WANBytesUp, stats.WANBytesDown, err = proc.GetNetworkUpDown(config.WAN())
		if err != nil {
			log.Error(err)
		}
		stats.LANBytesUp, stats.LANBytesDown, err = proc.GetNetworkUpDown(config.LAN())
		if err != nil {
			log.Error(err)
		}
		stats.MemFree, err = proc.GetFreeMemory()
		if err != nil {
			log.Error(err)
		}

		stats.Containers = docker.GetContainers()
		stats.Uptime = proc.GetUptime()
		finishTime := time.Now()
		diff := finishTime.Sub(startTime)
		stats.CollectionTime = int64(diff / time.Millisecond)

		ystats, _ := yaml.Marshal(&stats)
		fmt.Println(string(ystats))
		jstats, _ := json.Marshal(&stats)
		go publish.Post(string(jstats), config.Endpoint(), config.AuthKey(), hostname)
		time.Sleep(time.Duration(config.SleepLoopTime()) * time.Minute)
	}
}
