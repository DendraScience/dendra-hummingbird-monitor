package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

var config Config

var defaultConfig = Config{
	// default to location for k3s kubeconfig
	KubeConfig: "/etc/rancher/k3s/k3s.yaml",
}

func init() {
	var err error
	config, err = loadConfig("/etc/dendra/hummingbird.toml")
	if err != nil {
		log.Error(err)
	}
}

func loadConfig(configFile string) (Config, error) {
	config := defaultConfig
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return defaultConfig, err
	}
	return config, nil
}
func WAN() string {
	return config.WAN
}
func LAN() string {
	return config.LAN
}

func SleepLoopTime() int {
	return config.SleepLoopTime
}
func Endpoint() string {
	return config.Endpoint
}
func AuthKey() string {
	return config.AuthKey
}

func KubeConfig() string {
	return config.KubeConfig
}
