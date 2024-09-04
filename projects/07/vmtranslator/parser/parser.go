package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CleanLine(line string) (cleaned string) {
	cleaned = strings.TrimSpace(line)
	return cleaned
}

// assumes this function is exclusively called line is cleaned and comment removed
// similar to parsing A instruction but included seperated so
func GetCommandType(line string) (command string) {

	arithmetOps := map[string]bool{
		"add": true,
		"sub": true,
		"neg": true,
		"eq":  true,
		"gt":  true,
		"lt":  true,
		"and": true,
		"or":  true,
		"not": true,
	}

	pushIdx := strings.Index(line, "push")
	popIdx := strings.Index(line, "pop")

	// if contains push/pop
	if pushIdx == 0 {
		command = "C_PUSH"
	} else if popIdx == 0 {
		command = "C_POP"
	} else if arithmetOps[line] {
		command = "C_ARITHMETIC"
	} else {
		fmt.Printf("Unable to identify command type")
	}
	return command
}

func args(line string) (argument1 string, argument2 string) {
	args := strings.Split(line, " ")
	if len(args) == 1 {
		argument1 = args[0]
		argument2 = ""
	} else {
		argument1 = args[1]
		argument2 = args[2]
	}

	return argument1, argument2
}

func arg2(line string) (argument string) {
	args := strings.Split(line, " ")
	if len(args) == 1 {

	} else {
		argument = args[2]
	}

	return argument
}

func Parser(filepath string) (parsed []string, instType []string) {

	// open file and scan
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()

	var cleanedLines []string
	var instructionType []string

	scanner := bufio.NewScanner(f)
	// check each line for comments, clear whitespace
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace((line))

		commentIdx := strings.Index(line, "//")
		if commentIdx != -1 {
			line = line[:commentIdx]
		}

		// removes empty lines
		if len(line) == 0 {
			continue
		}

		// define instruction type before any processing
		currentInstruction := GetCommandType(line)
		instructionType = append(instructionType, currentInstruction)

		arg1, arg2 := args(line)
		cleanedLines = append(cleanedLines, arg1)

		// arg2 will be "" here, but don't want the extra space
		if currentInstruction != "C_ARITHMETIC" {
			cleanedLines = append(cleanedLines, arg2)
		}

	}
	return cleanedLines, instructionType
}
