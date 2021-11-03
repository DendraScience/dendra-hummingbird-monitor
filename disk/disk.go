package disk

import (
	"log"

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

func GetDiskFreeAndUsagePercentage() (int64, float64) {
	var freeSpace int64
	partitions, size := GetPartitions()
	for _, part := range partitions {
		freeSpace += getFreeBytes(part)
	}

	return freeSpace, float64(freeSpace) / float64(size)
}
