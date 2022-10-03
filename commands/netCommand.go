package commands

import (
	"GoMonitoring/messages"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
)

func NetCommand() {
	network := net.ConnectionStat{}

	//Debug Incoming...
	fmt.Println(messages.HeaderNetMessage)
	fmt.Println(fmt.Sprintf("IP: %v", network))
	fmt.Println(messages.BottomNetMessage)
}
