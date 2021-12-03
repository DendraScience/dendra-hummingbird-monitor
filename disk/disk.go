package disk

import (
	"log"

	"github.com/DendraScience/dendra_hummingbird_monitor/types"
	"github.com/jaypipes/ghw"
	"golang.org/x/sys/unix"
)

func getFreeBytes(path string) int64 {
	var stat unix.Statfs_t
	unix.Statfs(path, &stat)

	// Available blocks * size per block = available space in bytes
	return int64(stat.Bavail * uint64(stat.Bsize))
}

func GetPartitions() ([]string, int64) {
	var partitionList []string
	var totalSize int64
	block, err := ghw.Block()
	if err != nil {
		log.Printf("Error analyzing block devices: %s\n", err.Error())
		return partitionList, totalSize
	}
	for _, disk := range block.Disks {
		for _, part := range disk.Partitions {
			if part.MountPoint != "" {
				totalSize += int64(part.SizeBytes)
				partitionList = append(partitionList, part.MountPoint)
			}
		}
	}
	return partitionList, totalSize
}

func GetMountInfo() []types.MountInfo {
	mountInfo := []types.MountInfo{}
	block, err := ghw.Block()
	if err != nil {
		log.Printf("Error analyzing block devices: %s\n", err.Error())
		return mountInfo
	}
	for _, disk := range block.Disks {
		for _, part := range disk.Partitions {
			if part.MountPoint != "" {
				curMount := types.MountInfo{}
				curMount.MountPoint = part.MountPoint
				curMount.DiskAvail = float64(part.SizeBytes)
				curMount.DiskFree = float64(getFreeBytes(part.MountPoint))
				curMount.DiskUsage = curMount.DiskAvail - float64(curMount.DiskFree)
				curMount.DiskPercent = float64(curMount.DiskUsage / curMount.DiskAvail)
				curMount.DiskName = part.Name
				curMount.PartitionUUID = part.UUID
				mountInfo = append(mountInfo, curMount)
			}
		}
	}
	return mountInfo
}

func GetDiskFreeAndUsagePercentage() (int64, float64) {
	var freeSpace int64
	partitions, size := GetPartitions()
	for _, part := range partitions {
		freeSpace += getFreeBytes(part)
	}

	return freeSpace, float64(size-freeSpace) / float64(size)
}
