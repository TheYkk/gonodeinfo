package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"
)

func main() {

	hostname, _ := os.Hostname()
	si := &syscall.Sysinfo_t{}

	// XXX is a raw syscall thread safe?
	err := syscall.Sysinfo(si)
	if err != nil {
		panic("Commander, we have a problem. syscall.Sysinfo:" + err.Error())
	}

	uptime := time.Duration(si.Uptime)
	fmt.Println("Hostname:  ", hostname)
	fmt.Println("Platform:  ", runtime.GOOS)
	fmt.Println("Arch:      ", runtime.GOARCH)
	fmt.Println("CPU count: ", runtime.NumCPU())
	fmt.Println("Total Ram: ", ByteCountIEC(uint64(si.Totalram) * uint64(si.Unit)))
	

	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60
	fmt.Printf("Uptime:     %d days, %d hours, %d minutes \n", days, hours, minutes)
}

func ByteCountIEC(b uint64) string {
    const unit = 1024
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %ciB",
        float64(b)/float64(div), "KMGTPE"[exp])
}