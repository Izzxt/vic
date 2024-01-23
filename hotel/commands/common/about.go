package commands

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Izzxt/vic/core"
)

type AboutCommand struct{}

func (c *AboutCommand) Execute(client core.HabboClient, args []string) {
	runtime.GC()

	msgs := make([]string, 0)

	msgs = append(msgs, "<div style='font-size: 16px;'>")
	msgs = append(msgs, "<b>Hotel Statistics</b>")
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "- Online users: 0") // TODO: add online users
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "</br>")

	cpu := getCPUUsage()
	totalAlloc, freeMemory, sysMemory := getMemoryUsage()
	msgs = append(msgs, "<b>Server Statistics</b>")
	msgs = append(msgs, "</br>")
	msgs = append(msgs, fmt.Sprintf("- CPU: %d", cpu))
	msgs = append(msgs, "</br>")
	msgs = append(msgs, fmt.Sprintf("- Memory Usage: %.fMB/%.fMB", totalAlloc, freeMemory))
	msgs = append(msgs, "</br>")
	msgs = append(msgs, fmt.Sprintf("- Total Memory: %.fMB", sysMemory))
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "</br>")

	msgs = append(msgs, "<b>Hotel Information</b>")
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "- Version: 1.0.0")
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "- Release: 0.0.1")
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "- Uptime: 0:00:00") // TODO: add uptime
	msgs = append(msgs, "</br>")
	msgs = append(msgs, "</br>")

	msgs = append(msgs, "<b><i>Special thanks to Izzxt 💖</i></b>")
	msgs = append(msgs, "</div>")

	msg := strings.Join(msgs, "")

	client.SendAlert(msg)
}

func getCPUUsage() int {
	return runtime.NumCPU()
}

func getMemoryUsage() (float64, float64, float64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Get total allocated memory
	totalAlloc := float64(m.TotalAlloc) / 1024 / 1024

	// Get total memory obtained from the OS
	sysMemory := float64(m.Sys) / 1024 / 1024

	// Calculate free memory
	freeMemory := sysMemory - totalAlloc
	return totalAlloc, freeMemory, sysMemory
}
