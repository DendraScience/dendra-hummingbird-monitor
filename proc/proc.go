package proc

import (
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/prometheus/procfs"
	log "github.com/sirupsen/logrus"
)

var (
	networkFS       procfs.FS
	networkFSOnce   sync.Once
	networkLastUp   map[string](int64)
	networkLastDown map[string](int64)
	bootTime        uint64
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return hostname
}
func init() {
	networkFSOnce.Do(func() {
		var err error
		networkFS, err = procfs.NewDefaultFS()
		if err != nil {
			log.Fatal("Cannot create DefaultFS" + err.Error())
		}
	})
	var err error
	bootStat, err := procfs.NewStat()
	if err != nil {
		log.Fatal("Cannot create DefaultFS" + err.Error())
	}
	bootTime = bootStat.BootTime
	networkLastDown = make(map[string]int64)
	networkLastUp = make(map[string]int64)
}
func GetNetworkUpDown(interface_name string) (newUpBytes int64, newDownBytes int64, err error) {
	netDev, err := networkFS.NetDev()
	if err != nil {
		return -1, -1, err
	}
	if interface_data, ok := netDev[interface_name]; ok {
		if networkLastDown[interface_name] == 0 || networkLastDown[interface_name] > int64(interface_data.RxBytes) {
			networkLastDown[interface_name] = int64(interface_data.RxBytes)
		} else {
			newDownBytes = int64(interface_data.RxBytes) - networkLastDown[interface_name]
			networkLastDown[interface_name] = int64(interface_data.RxBytes)
		}
		if networkLastUp[interface_name] == 0 || networkLastUp[interface_name] > int64(interface_data.TxBytes) {
			networkLastUp[interface_name] = int64(interface_data.TxBytes)
		} else {
			newUpBytes = int64(interface_data.TxBytes) - networkLastUp[interface_name]
			networkLastUp[interface_name] = int64(interface_data.TxBytes)
		}
		return newUpBytes, newDownBytes, err
	} else {
		err = errors.New(fmt.Sprintf("Network interface %s not found!", interface_name))
		return -1, -1, err
	}
}
func GetTotalMemory() (int64, error) {
	avail, err := networkFS.Meminfo()
	if err != nil {
		return -1, err
	}
	return int64(*avail.MemTotal), err
}
func GetFreeMemory() (int64, error) {
	avail, err := networkFS.Meminfo()
	if err != nil {
		return -1, err
	}
	return int64(*avail.MemAvailable), err
}

func GetUptime() int64 {
	return (time.Now().UnixNano() / int64(time.Second)) - int64(bootTime)
}

func GetIPForInterface(interface_name string) (string, error) {
	iface, err := net.InterfaceByName(interface_name)
	if err != nil {
		return "", err
	}
	addrs, err := iface.Addrs()
	if len(addrs) < 1 {
		return "", err
	}
	return addrs[0].String(), err
}

func GetLoad() (float64, error) {
	load, err := networkFS.LoadAvg()
	if err != nil {
		return -1, err
	}
	return load.Load15, err

}

func GetMacForInterface(interface_name string) (string, error) {
	iface, err := net.InterfaceByName(interface_name)
	if err != nil {
		return "", err
	}
	return iface.HardwareAddr.String(), err
}
