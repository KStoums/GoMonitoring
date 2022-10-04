package commands

import (
	"GoMonitoring/messages"
	"fmt"
)

func HelpCommand() {
	fmt.Println(messages.HeaderHelpMessage)
	fmt.Println(messages.CommandOne)
	fmt.Println(messages.CommandTwo)
	fmt.Println(messages.CommandThree)
	fmt.Println(messages.CommandFour)
	fmt.Println(messages.CommandFive)
	fmt.Println(messages.CommandSix)
	fmt.Println(messages.CommandSeven)
	fmt.Println(messages.BottomHelpMessage)
}
