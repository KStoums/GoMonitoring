package commands

import (
	"GoMonitoring/messages"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
)

func NetCommand() {
	network, _ := net.Connections("all")

	fmt.Println(messages.HeaderNetMessage)
	for _, value := range network {
		if value.Raddr.Port != 0 {
			fmt.Println(fmt.Sprintf("IP: %s:%d   |   Status: %s", value.Raddr.IP, value.Raddr.Port, value.Status))
		}
	}
	fmt.Println(messages.BottomNetMessage)
}
