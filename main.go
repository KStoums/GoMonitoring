package main

import (
	"GoMonitoring/messages"
	"fmt"
	cpu2 "github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"log"
	"math"
	"os"
	"strconv"
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

		sliceFreeMemory := strconv.Itoa(int(memory.Free))
		sliceTotalMemory := strconv.Itoa(int(memory.Total))

		fmt.Println(messages.HeaderMemoryMessage)
		fmt.Println(fmt.Sprintf("Used: %v%%", memory.UsedPercent))
		fmt.Println(fmt.Sprintf("Free: %v%% (Fix soon)", sliceFreeMemory[:2]))
		fmt.Println(fmt.Sprintf("Total: %vGo (Fix soon) (Take the number and do -1)", sliceTotalMemory[:2]))
		fmt.Println(messages.BottomMemoryMessage)
		return
	}

	if args[1] == "cpu" && len(args) <= 2 {
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
			fmt.Println(fmt.Sprintf("Frequency: %vMhz", v.Mhz))
			fmt.Println(messages.BottomCPUMessage)
		}
		return
	}

	if args[1] == "host" && len(args) <= 2 {
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

	if args[1] == "disk" && len(args) <= 2 {
		fmt.Print(messages.SelectHardDrive)
		var selectedHardDrive string
		fmt.Scanln(&selectedHardDrive)

		diskInfo, err := disk.Usage(selectedHardDrive)
		if err != nil {
			log.Fatal(err)
		}
		sliceFreeSpace := strconv.Itoa(int(diskInfo.Free))
		sliceTotalSpace := strconv.Itoa(int(diskInfo.Total))

		fmt.Println(messages.HeaderDiskMessage)
		fmt.Println(fmt.Sprintf("Used: %v%%", math.Round(diskInfo.UsedPercent)))
		fmt.Println(fmt.Sprintf("Free: %v%%", sliceFreeSpace[:2]))

		if diskInfo.Total > 1000000000000 {
			fmt.Println(fmt.Sprintf("Total: %v%Gb", sliceTotalSpace[:4]))
			fmt.Println(messages.BottomDiskMessage)
			return
		}

		fmt.Println(fmt.Sprintf("Total: %vGb", sliceTotalSpace[:3]))
		fmt.Println(messages.BottomDiskMessage)
		return
	}

	fmt.Println("Error: Type 'help' for view all commands.")
}
