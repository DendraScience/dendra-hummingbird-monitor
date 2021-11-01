//go:build linux
// +build linux

package pkg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	availUpdates      int
	manuallyInstalled int
)

const (
	aptGet  = "/usr/bin/apt-get"
	aptMark = "/usr/bin/apt-mark"
)

func init() {
	availUpdates = -1
	manuallyInstalled = -1
	if _, err := os.Stat(aptGet); err != nil {
		log.Error(err)
	} else {
		go refreshPackages()
	}

}

// A utility to convert the values to proper strings.
// From https://stackoverflow.com/a/53197771
func int8ToStr(arr []int8) string {
	b := make([]byte, 0, len(arr))
	for _, v := range arr {
		if v == 0x00 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}

func GetKernelVersion() string {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err == nil {
		// extract members:
		// type Utsname struct {
		//  Sysname    [65]int8
		//  Nodename   [65]int8
		//  Release    [65]int8
		//  Version    [65]int8
		//  Machine    [65]int8
		//  Domainname [65]int8
		// }

		return fmt.Sprintf("%s:::%s:::%s", int8ToStr(uname.Sysname[:]),
			int8ToStr(uname.Release[:]),
			int8ToStr(uname.Version[:]))

	}
	return ""
}

func GetInstalledPackageCount() int {
	if _, err := os.Stat(aptMark); err != nil {
		log.Error(err)
		return -1
	}

	cmd := exec.Command(aptMark, "showmanual")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error(err)
	}
	if err = cmd.Start(); err != nil {
		log.Error(err)
	}
	buf, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Error(err)
	}
	count := 0
	for _, c := range string(buf) {
		if c == '\n' {
			count++
		}
	}
	go cmd.Wait()
	return count
}

func GetNumAvailUpdates() int {
	return availUpdates
}

//TODO: later, pass this function a logger instead of assuming we want stuff written out
func refreshPackages() {
	for {
		// First, update the value of the available updates. -s simulates an upgradde without taking a lock.
		cmd := exec.Command(aptGet, "-s", "upgrade")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Error(err)
		}
		if err = cmd.Start(); err != nil {
			log.Error(err)
		}
		buf, err := ioutil.ReadAll(stdout)
		if err != nil {
			log.Error(err)
		}
		re := regexp.MustCompile(`[[:digit:]]+ upgraded`)
		// If we got a match
		if matches := re.Find(buf); matches != nil {
			re = regexp.MustCompile(`^[[:digit:]]+`)
			if packageCount := re.Find(matches); packageCount != nil {
				var err error
				availUpdates, err = strconv.Atoi(string(packageCount))
				if err != nil {
					log.Error(err)
				}
			}
		}
		cmd.Wait()
		// Then, sleep for 2-5 days to be nice to the ubuntu mirrors.
		// there's no need to add an additional random offset, since
		// the machines will naturally drift over time
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(2) // n will be between 0 and 4
		time.Sleep(time.Duration(n+2) * time.Hour * 24)
		// Now try to update the package lists
		cmd = exec.Command(aptGet, "update")
		if err = cmd.Start(); err != nil {
			log.Error(err)
		}
		// Wait for the command to finish so we get a proper number back
		// when the loop starts again
		if err = cmd.Wait(); err != nil {
			log.Error(err)
		}
	}
}
