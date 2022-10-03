package main

import (
	"GoMonitoring/messages"
	"fmt"
	cpu2 "github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"log"
	"os"
)

func main() {
	args := os.Args

	if args[1] == "help" && len(args) <= 2 {
		fmt.Println(messages.HeaderHelpMessage)
		fmt.Println(messages.CommandOne)
		fmt.Println(messages.CommandTwo)
		fmt.Println(messages.CommandThree)
		fmt.Println(messages.CommandFour)
		fmt.Println(messages.CommandFive)
		fmt.Println(messages.BottomHelpMessage)
		return
	}

	if args[1] == "memory" && len(args) <= 2 {
		memory, err := mem.VirtualMemory()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(messages.HeaderMemoryMessage)
		fmt.Println(fmt.Sprintf("Used: %v%s", memory.UsedPercent, "%"))
		fmt.Println(fmt.Sprintf("Free: %v%s", memory.Free, "%"))
		fmt.Println(fmt.Sprintf("Total: %v%s", memory.Total, "%"))
		fmt.Println(messages.BottomMemoryMessage)
		return
	}

	if args[1] == "cpu" && len(args) <= 2 {
		cpu, err := cpu2.Info()
		if err != nil {
			log.Fatal(err)
			return
		}

		for i, v := range cpu {
			fmt.Println(messages.HeaderCPUMessage)
			fmt.Println("CPU:", i)
			fmt.Println("Model:", v.ModelName)
			fmt.Println("Cores:", v.Cores)
			fmt.Println("CacheSize:", v.CacheSize)
			fmt.Println("Mhz:", v.Mhz)
			fmt.Println(messages.BottomCPUMessage)
		}
		return
	}

	if args[1] == "host" && len(args) <= 2 {
		userOs, err := host.Info()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(messages.HeaderHostMessage)
		fmt.Println(fmt.Sprintf("OS: %s", userOs.Platform))
		fmt.Println(fmt.Sprintf("Hostname: %s", userOs.Hostname))
		fmt.Println(fmt.Sprintf("KernalARCH: %s", userOs.KernelArch))
		fmt.Println(fmt.Sprintf("KernelVersion: %s", userOs.KernelVersion))
		fmt.Println(messages.BottomHostMessage)
		return
	}

	if args[1] == "net" && len(args) <= 2 {
		network := net.ConnectionStat{}

		//Debug Incoming...
		fmt.Println(messages.HeaderNetMessage)
		fmt.Println(fmt.Sprintf("IP: %v", network))
		fmt.Println(messages.BottomNetMessage)
		return
	}

	fmt.Println("Error: Type 'help' for view all commands.")
}
