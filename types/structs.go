package types

import (
	"encoding/json"
	"time"

	"github.com/DendraScience/dendra_hummingbird_monitor/docker"
	log "github.com/sirupsen/logrus"
)

type MountInfo struct {
	DiskAvail     float64 `json:"disk_avail"`    // Disk usage
	DiskFree      float64 `json:"disk_free"`     // Disk free
	DiskName      string  `json:"disk_name"`     // name like '/dev/sda1'
	DiskUsage     float64 `json:"disk_usage"`    // Disk usage
	DiskPercent   float64 `json:"usage_percent"` // Disk usage
	MountPoint    string  `json:"mount_point"`   // folder the disk is mounted to
	PartitionUUID string  `json:"partition_uuid"`
}
type QuarterHourly struct {
	CPU_Percent    float64            `json:"cpu_percent"`       // CPU Percentage usage
	CollectionTime int64              `json:"collection_time"`   // Time it takes for metrics collection
	Containers     []docker.Container `json:"containers"`        // Containers and statuses
	DiskFree       int64              `json:"disk_free"`         // Disk free
	DiskUsage      float64            `json:"disk_usage"`        // Disk usage
	MountInfo      []MountInfo        `json:"mounts"`            // mounted disks information
	Hostname       string             `json:"hostname"`          // Hostname
	LANBytesDown   int64              `json:"lan_bytes_down"`    // Lan Network bytes down
	LANBytesUp     int64              `json:"lan_bytes_up"`      // Lan Network bytes up
	LANIP          string             `json:"lan_ip"`            // IP of lan interface
	LoadAverage    float64            `json:"load_average"`      // Load Average
	MemAvail       int64              `json:"memory_avail"`      // Memory avail
	MemBuffered    int64              `json:"memory_buffered"`   // Memory buffered
	MemFree        int64              `json:"memory_free"`       // Memory free
	MemPercent     float64            `json:"memory_percent"`    // Memory percentage
	MemTotal       int64              `json:"memory_total"`      // Memory total
	NumPackages    int                `json:"num_packages"`      // Number of installed packages
	ProcessorCount int                `json:"processor_count"`   // Thread count
	Timestamp      time.Time          `json:"timestamp"`         // Timestamp
	UpdatesAvail   int                `json:"updates_available"` // Package updates avail count
	Uptime         int64              `json:"uptime"`            // Uptime
	Version        string             `json:"version"`           // Hummingbird version
	WANBytesDown   int64              `json:"wan_bytes_down"`    // Wan Network bytes down
	WANBytesUp     int64              `json:"wan_bytes_up"`      // Wan Network bytes up
	WANIP          string             `json:"wan_ip"`            // IP of wan interface
}

func (h QuarterHourly) String() string {
	str, err := json.MarshalIndent(h, "", "\t")
	if err != nil {
		log.Error(err)
	}
	return string(str)
}
