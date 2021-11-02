package cpu

import (
	"log"

	"github.com/jaypipes/ghw"
)

func GetCPUThreads() int {
	cpu, err := ghw.CPU()
	if err != nil {
		log.Printf("Error analyzing CPU: %s\n", err.Error())
		return 0
	}
	return int(cpu.TotalThreads)
}
