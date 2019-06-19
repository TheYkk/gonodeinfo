package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {

	/*t, err := getLastSystemBootTime()
	if err != nil {
		panic(err)
	}*/
	//uptime := time.Now().Unix() - t.Unix()
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

	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60
	fmt.Printf("Uptime:     %d days, %d hours, %d minutes \n", days, hours, minutes)
}

func getLastBootTime() string {
	out, err := exec.Command("who", "-b").Output()
	if err != nil {
		panic(err)
	}
	t := strings.TrimSpace(string(out))
	t = strings.TrimPrefix(t, "system boot")
	t = strings.TrimSpace(t)
	return t
}

func getTimezone() string {
	out, err := exec.Command("date", "+%Z").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func getLastSystemBootTime() (time.Time, error) {
	return time.Parse(`2006-01-02 15:04MST`, getLastBootTime()+getTimezone())
}
