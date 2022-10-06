package commands

import (
	"GoMonitoring/messages"
	"fmt"
	cpu2 "github.com/shirou/gopsutil/cpu"
	"log"
)

func CPUCommand() {
	cpuInfo, err := cpu2.Info()
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, v := range cpuInfo {
		fmt.Println(messages.HeaderCPUMessage)
		fmt.Println("CPU:", i)
		fmt.Println("Model:", v.ModelName)
		fmt.Println("Cores:", v.Cores)
		fmt.Println("CacheSize:", v.CacheSize)
		fmt.Println(fmt.Sprintf("Temperature: 0° (COMING SOON)")) //ADD TEMPERATURE
		fmt.Println(fmt.Sprintf("Frequency: %vMhz", v.Mhz))
		fmt.Println(messages.BottomCPUMessage)
		return
	}
}
