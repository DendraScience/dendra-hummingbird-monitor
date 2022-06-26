package k8s

import (
	"fmt"
	"time"
)

type Container struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
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
	return fmt.Sprintf("Name: %s\nImage: %s\nID: %s\n", c.Name, c.Image, c.ID)
}
