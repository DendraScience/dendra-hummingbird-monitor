package docker

import "time"

type Container struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CPU        float64   `json:"cpu"`
	MemUsage   int       `json:"mem_usage"`
	MemAllowed int       `json:"mem_allowed"`
	MemPercent float64   `json:"mem_percent"`
	Created    time.Time `json:"time_created"`
	Image      string    `json:"image"`
	Uptime     int64     `json:"uptime"`
}
