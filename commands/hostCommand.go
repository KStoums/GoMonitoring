package commands

import (
	"GoMonitoring/messages"
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"log"
)

func HostCommand() {
	osInfo, err := host.Info()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(messages.HeaderHostMessage)
	fmt.Println(fmt.Sprintf("OS: %s", osInfo.Platform))
	fmt.Println(fmt.Sprintf("Hostname: %s", osInfo.Hostname))
	fmt.Println(fmt.Sprintf("KernalARCH: %s", osInfo.KernelArch))
	fmt.Println(fmt.Sprintf("KernelVersion: %s", osInfo.KernelVersion))
	fmt.Println(messages.BottomHostMessage)
}
