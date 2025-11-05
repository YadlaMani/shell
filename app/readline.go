package main

import (
	"fmt"
	"os"

	"github.com/YadlaMani/shell/app/utils"
)

func readLine() (string, bool) {
	var line []rune
	historyIdx := -1
	buf := make([]byte, 4)
	cursor := 0
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			continue
		}
		if n == 1 {
			switch buf[0] {
			case 3: // Ctrl-C
				fmt.Println("^C")
				return "", true
			case 4: //Ctrl-D(EOF)
				if len(line) == 0 {
					fmt.Println()
					return "", true
				}
			case 13, 10: //Enter
				fmt.Print("\r\n")
				return string(line), false
			case 127, 8: //Backspace/Delete
				if cursor > 0 {
					line = append(line[:cursor-1], line[cursor:]...)
					cursor--
					fmt.Print("\b")
					rest := string(line[cursor:])
					fmt.Print(rest + " ")
					for range len(rest) + 1 {
						fmt.Print("\b")
					}
				}
			default:
				if buf[0] >= 32 && buf[0] < 127 {
					ch := rune(buf[0])
					line = append(line[:cursor], append([]rune{ch}, line[cursor:]...)...)
					fmt.Print(string(line[cursor:]))
					cursor++
					for range len(line) - cursor {
						fmt.Print("\b")
					}
				}
			}

		} else if n == 3 && buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case 65: //UpArrow
				if len(utils.History) > 0 {
					if historyIdx == -1 {
						historyIdx = len(utils.History) - 1
					} else if historyIdx > 0 {
						historyIdx--
					} else {
						continue
					}
					clearLine(len(line))
					line = []rune(utils.History[historyIdx])
					cursor = len(line)
					fmt.Print(string(line))
				}
			case 66: //Down Arrow
				if historyIdx == -1 {
					continue
				}
				if historyIdx < len(utils.History)-1 {
					historyIdx++
					clearLine(len(line))
					line = []rune(utils.History[historyIdx])
					cursor = len(line)
					fmt.Print(string(line))
				} else {
					historyIdx = -1
					clearLine(len(line))
					line = []rune{}
					cursor = 0
				}
			case 67: //Right Arrow
				if cursor < len(line) {
					fmt.Print(string(line[cursor]))
					cursor++
				}
			case 68: //Left Arrow
				if cursor > 0 {
					fmt.Print("\b")
					cursor--
				}
			}
		}
	}
}

func clearLine(length int) {
	fmt.Print("\r")
	for i := 0; i < length+2; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r")
	fmt.Print("$ ")
}
