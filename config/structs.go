package config

type Config struct {
	AuthKey       string `toml:"authkey"`
	Endpoint      string `toml:"endpoint"`
	KubeConfig    string `toml:"kubeconfig"`
	LAN           string `toml:"lan"`
	SleepLoopTime int    `toml:"sleeplooptime"`
	WAN           string `toml:"wan"`
}
