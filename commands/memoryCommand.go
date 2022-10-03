package commands

import (
	"GoMonitoring/messages"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"strconv"
)

func MemoryCommand() {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
		return
	}

	sliceFreeMemory := strconv.Itoa(int(memory.Free))
	sliceTotalMemory := strconv.Itoa(int(memory.Total))

	fmt.Println(messages.HeaderMemoryMessage)
	fmt.Println(fmt.Sprintf("Used: %v%%", memory.UsedPercent))
	fmt.Println(fmt.Sprintf("Free: %v%% (Fix soon)", sliceFreeMemory[:2]))
	fmt.Println(fmt.Sprintf("Total: %vGo (Fix soon) (Take the number and do -1)", sliceTotalMemory[:2]))
	fmt.Println(messages.BottomMemoryMessage)
}
