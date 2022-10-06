package commands

import (
	"GoMonitoring/messages"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"log"
	"math"
	"strconv"
)

func DiskCommand() {
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
		fmt.Println(fmt.Sprintf("Total: %vGb", sliceTotalSpace[:4]))
		fmt.Println(messages.BottomDiskMessage)
		return
	}

	fmt.Println(fmt.Sprintf("Total: %vGb", sliceTotalSpace[:3]))
	fmt.Println(messages.BottomDiskMessage)
}
