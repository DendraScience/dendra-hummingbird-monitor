package types

import (
	"encoding/json"
	"time"

	"github.com/DendraScience/dendra_hummingbird_monitor/docker"
	log "github.com/sirupsen/logrus"
)

type QuarterHourly struct {
	Timestamp time.Time `json:"timestamp"` // - Timestamp
	Hostname  string    `json:"hostname"`  // - Hostname
	Version   Version   `json:"version"`   // - Tickbird version

	CollectionTime int64              `json:"collection_time"`   // - Time it takes for metrics collection
	Containers     []docker.Container `json:"containers"`        // - Containers and statuses
	DiskFree       int64              `json:"disk_free"`         // - Disk free
	LANBytesDown   int64              `json:"lan_bytes_down"`    // - Lan Network bytes down
	LANBytesUp     int64              `json:"lan_bytes_up"`      // - Lan Network bytes up
	MemFree        int64              `json:"memory_free"`       // - Memory free
	NumPackages    int                `json:"num_packages"`      // - Number of installed packages
	UpdatesAvail   int                `json:"updates_available"` // - Package updates avail count
	Uptime         int64              `json:"uptime"`            // - Uptime
	WANBytesDown   int64              `json:"wan_bytes_down"`    // - Wan Network bytes down
	WANBytesUp     int64              `json:"wan_bytes_up"`      // - Wan Network bytes up
}

type HNetworking struct {
	LANBytesDown int64 `json:"lan_bytes_down"` // - Lan Network bytes down
	LANBytesUp   int64 `json:"lan_bytes_up"`   // - Lan Network bytes up
	WANBytesDown int64 `json:"wan_bytes_down"` // - Wan Network bytes down
	WANBytesUp   int64 `json:"wan_bytes_up"`   // - Wan Network bytes up

}

type Version struct {
	Semantic string `json:"semantic"`
	Commit   string `json:"commit"`
}

func (h QuarterHourly) String() string {
	str, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		log.Error(err)
	}
	return string(str)
}
