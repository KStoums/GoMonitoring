package main

import (
	"GoMonitoring/commands"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if args[1] == "help" && len(args) <= 2 {
		commands.HelpCommand()
		return
	}

	if args[1] == "memory" && len(args) <= 2 {
		commands.MemoryCommand()
		return
	}

	if args[1] == "cpu" && len(args) <= 2 {
		commands.CPUCommand()
		return
	}

	if args[1] == "host" && len(args) <= 2 {
		commands.HostCommand()
		return
	}

	if args[1] == "net" && len(args) <= 2 {
		commands.NetCommand()
		return
	}

	if args[1] == "disk" && len(args) <= 2 {
		commands.DiskCommand()
		return
	}

	if args[1] == "gpu" && len(args) <= 2 {
		commands.GPUCommand()
		return
	}

	fmt.Println("Error: Type 'help' for view all commands.")
}
