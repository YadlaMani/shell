package commands

import (
	"fmt"

	"github.com/YadlaMani/shell/app/utils"
)

func HistoryCommand(arg string) {
	var limit int

	if arg == "" {
		limit = len(utils.History)
	} else {
		_, err := fmt.Sscanf(arg, "%d", &limit)
		if err != nil || limit < 1 {
			fmt.Println("history: invalid number:", arg)
			return
		}
		if limit > len(utils.History) {
			limit = len(utils.History)
		}
	}
	start := len(utils.History) - limit

	for i := start; i < len(utils.History); i++ {
		fmt.Printf("\t%d\t%s\n", i+1, utils.History[i])
	}
}
