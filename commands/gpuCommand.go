package commands

import (
	"GoMonitoring/messages"
	"fmt"
)

func GPUCommand() {
	fmt.Println(messages.HeaderGPUMessage)
	fmt.Println("» Feature in coming.")
	fmt.Println(messages.BottomGPUMessage)
}
