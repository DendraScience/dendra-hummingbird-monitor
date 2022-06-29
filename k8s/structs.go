package k8s

import (
	"fmt"
	"time"
)

type Metric struct {
	ID        string
	Node      string
	Value     int64
	TimeStamp time.Time
}

type Container struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Node            string    `json:"node"`
	CPUUsage        float64   `json:"cpu"`
	CPUAllowed      float64   `json:"-"`
	MemUsage        int       `json:"mem_usage"`
	MemAllowed      int       `json:"mem_allowed"`
	MemPercent      float64   `json:"mem_percent"`
	Created         time.Time `json:"time_created"`
	Image           string    `json:"image"`
	Uptime          int64     `json:"uptime"`
	PrevCPUMS       int64     `json:"-"`
	PrevReadingTime time.Time `json:"-"`
	NewCPUMS        int64     `json:"-"`
	NewReadingTime  time.Time `json:"-"`
}

func (c Container) String() string {
	return fmt.Sprintf("\nName: %s\tImage: %s\nID: %s\nCreated: %v\nMemPercent: %.2f\tMemAllowed: %d\tMemUsage: %d\nCPUUsage: %.2f\tCPUAllowed: %.2f\tCPUMS: %d", c.Name, c.Image, c.ID, c.Created, c.MemPercent, c.MemAllowed, c.MemUsage, c.CPUUsage, c.CPUAllowed, c.NewCPUMS)
}
