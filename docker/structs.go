package docker

type Container struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CPU        float64 `json:"cpu"`
	MemUsage   float64 `json:"mem_usage"`
	MemAllowed float64 `json:"mem_allowed"`
	MemPercent float64 `json:"mem_percent"`
	Created    int64   `json:"time_created"`
	Image      string  `json:"image"`
	Tag        string  `json:"tag"`
	Uptime     int64   `json:"uptime"`
}
