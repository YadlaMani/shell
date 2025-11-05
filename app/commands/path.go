package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/YadlaMani/shell/app/utils"
)

// PathCommand executes an external command found in PATH
func PathCommand(command string, input string) {
	cmd := exec.Command(command, utils.ParseArguments(input, false)...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(out))
		log.Fatalf("error running command %v", err)
	}
	if len(out) == 0 {
		fmt.Println(string(out))
	} else {
		fmt.Println(string(out[:len(out)-1]))
	}
}
