package utils

import (
	"os"

	"strings"
)

var BuiltinCommands = map[string]string{
	"exit":    "a shell builtin",
	"echo":    "a shell builtin",
	"type":    "a shell builtin",
	"pwd":     "a shell builtin",
	"cd":      "a shell builtin",
	"history": "a shell builtin",
}
var PATH_COMMANDS = map[string]string{}

func IsBuiltin(command string) bool {
	_, ok := BuiltinCommands[command]
	return ok
}

var History []string

// ParseArguments splits a shell line into tokens. When includeWhitespace is true
func ParseArguments(input string, includeWhiteSpace bool) []string {
	res := []string{}

	isSingleQuoted := false
	isDoubleQuoted := false
	curr := ""
	index := 0
	for index < len(input) {
		r := input[index]
		if r == '\'' && !isDoubleQuoted {
			if isSingleQuoted {
				res = append(res, curr)
				curr = ""
				oldIndex := index
				for index+1 < len(input) && input[index+1] == ' ' {
					index += 1
				}
				if oldIndex < index && includeWhiteSpace {
					res = append(res, " ")
				}
			}
			isSingleQuoted = !isSingleQuoted
		} else if r == '"' && !isSingleQuoted {
			if isDoubleQuoted {
				curr = strings.ReplaceAll(curr, `\\`, `\`)
				curr = strings.ReplaceAll(curr, `\$`, `$`)
				curr = strings.ReplaceAll(curr, `\"`, `"`)
				curr = strings.ReplaceAll(curr, `\\n`, `\n`)
				res = append(res, curr)
				curr = ""
				oldIndex := index
				for index+1 < len(input) && input[index+1] == ' ' {
					index += 1
				}
				if oldIndex < index && includeWhiteSpace {
					res = append(res, " ")
				}
			}
			isDoubleQuoted = !isDoubleQuoted
		} else if r == ' ' && !isSingleQuoted && !isDoubleQuoted {
			curr = strings.Join(strings.Fields(strings.TrimSpace(curr)), " ")
			curr = strings.ReplaceAll(curr, `\`, "")
			res = append(res, curr)
			for index+1 < len(input) && input[index+1] == ' ' {
				index += 1
			}
			if includeWhiteSpace {
				res = append(res, " ")
			}
			curr = ""
		} else if r == '\\' {
			curr += string(input[index : index+2])
			index += 1
		} else {
			curr += string(r)
		}
		index += 1
	}

	if len(curr) != 0 {
		curr = strings.Join(strings.Fields(strings.TrimSpace(curr)), " ")
		curr = strings.ReplaceAll(curr, `\`, "")
		res = append(res, curr)
	}

	return res
}

// ScanPathCommands scans the PATH and populates the PATH_COMMANDS map
func ScanPathCommands() {
	PATH_COMMANDS = make(map[string]string)
	pathString := os.Getenv("PATH")
	paths := strings.Split(pathString, ":")
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			continue
		}
		commands, err := file.Readdirnames(0)
		file.Close()
		if err != nil {
			continue
		}
		for _, command := range commands {
			//Skip builtins and first occurances
			if ok := IsBuiltin(command); ok {
				continue
			}
			if _, ok := PATH_COMMANDS[command]; ok {
				continue
			}
			fullPath := path + "/" + command
			info, err := os.Stat(fullPath)
			if err != nil {
				continue
			}
			if info.Mode().IsRegular() && (info.Mode().Perm()&0111 != 0) {
				PATH_COMMANDS[command] = fullPath
			}
		}
	}
}
