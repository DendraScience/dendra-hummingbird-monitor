// +build linux

package wg

import (
	"os"
)

const (
	wg = "/usr/bin/wg"
)

func WGInstalled() bool {
	if _, err := os.Stat(wg); err != nil {
		return false
	}
	return true

}

//TODO
func GetIP() string {
	return ""
}

//TODO
func GetPubKey() string {
	return ""
}
