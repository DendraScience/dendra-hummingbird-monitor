//go:build linux
// +build linux

package pkg

import (
	"context"
	"fmt"
	"math/rand/v2"
	"syscall"
	"time"

	"github.com/gogrlx/snack"
	"github.com/gogrlx/snack/detect"
	log "github.com/sirupsen/logrus"
)

var (
	availUpdates      int
	manuallyInstalled int
	pkgManager        snack.Manager
)

func init() {
	availUpdates = -1
	manuallyInstalled = -1

	mgr, err := detect.Default()
	if err != nil {
		log.Warnf("No supported package manager detected: %v", err)
		return
	}
	pkgManager = mgr
	log.Infof("Detected package manager: %s", mgr.Name())
	go refreshPackages()
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
		return fmt.Sprintf("%s:::%s:::%s", int8ToStr(uname.Sysname[:]),
			int8ToStr(uname.Release[:]),
			int8ToStr(uname.Version[:]))
	}
	return ""
}

func GetInstalledPackageCount() int {
	if pkgManager == nil {
		return -1
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	packages, err := pkgManager.List(ctx)
	if err != nil {
		log.Error(err)
		return -1
	}
	return len(packages)
}

func GetNumAvailUpdates() int {
	return availUpdates
}

func refreshPackages() {
	if pkgManager == nil {
		return
	}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

		// Check if the package manager supports listing upgrades
		if vq, ok := pkgManager.(snack.VersionQuerier); ok {
			upgradable, err := vq.ListUpgrades(ctx)
			if err != nil {
				log.Warnf("Failed to check upgradable packages: %v", err)
			} else {
				availUpdates = len(upgradable)
			}
		}
		cancel()

		// Sleep for 2-4 days to be nice to package mirrors.
		// Machines will naturally drift over time.
		n := rand.IntN(2)
		time.Sleep(time.Duration(n+2) * time.Hour * 24)

		// Update the package lists
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Minute)
		err := pkgManager.Update(ctx)
		if err != nil {
			log.Warnf("Failed to update package lists: %v", err)
		}
		cancel()
	}
}
