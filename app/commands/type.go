package commands

import (
	"fmt"

	"github.com/YadlaMani/shell/app/utils"
)

// TypeCommand identifies whether a command is built-in or an external command
func TypeCommand(cmd string) {
	if description, ok := utils.BuiltinCommands[cmd]; ok {
		fmt.Printf("%s is %s\n", cmd, description)
		return
	}
	if path, ok := utils.PATH_COMMANDS[cmd]; ok {
		fmt.Printf("%s is %s\n", cmd, path)
		return
	}
	fmt.Printf("%s: not found\n", cmd)

}
