package commands

import (
	"fmt"
	"os"
)

// CdCommand changes the current working directory to the target directory
func CdCommand(target string) {
	var dir string
	var err error
	if target == "~" {
		dir, err = os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
		}
	} else {
		dir = target
	}
	_, err = os.Stat(dir)
	if err != nil {
		fmt.Println("cd: no such file or directory:", target)
		return
	}
	err = os.Chdir(dir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
	}

}
