package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/YadlaMani/shell/app/commands"
	"github.com/YadlaMani/shell/app/utils"
	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	utils.ScanPathCommands()

	for {
		input, err := line.Prompt("$ ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				continue
			}
			fmt.Println("exit")
			break
		}

		if strings.TrimSpace(input) != "" {
			utils.History = append(utils.History, input)
			line.AppendHistory(input)

		}

		args := utils.ParseArguments(input, false)
		if len(args) == 0 {
			continue
		}
		command := args[0]
		utils.ScanPathCommands()

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			message := strings.Join(utils.ParseArguments(input[len(command)+1:], true), "")
			fmt.Println(message)
		case "type":
			commands.TypeCommand(args[1])
		case "pwd":
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				continue
			}
			fmt.Println(cwd)
		case "cd":
			commands.CdCommand(args[1])
		case "history":
			if len(args) > 1 {
				commands.HistoryCommand(args[1])
			} else {
				commands.HistoryCommand("")
			}
		default:
			_, ok := utils.PATH_COMMANDS[command]
			if ok {
				if len(input) > len(command) {
					input = string(input[len(command)+1:])
				} else {
					input = ""
				}
				commands.PathCommand(command, input)
			} else {
				fmt.Printf("%s: command not found\n", command)
			}

		}

	}

}
