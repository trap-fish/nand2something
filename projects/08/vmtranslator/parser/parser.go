package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// used in getCommandType to map the current line to a type
// defined as global variable to avoid creation of map for each line
var commandOps = map[string]string{
	"push":     "C_PUSH",
	"pop":      "C_POP",
	"add":      "C_ARITHMETIC",
	"sub":      "C_ARITHMETIC",
	"neg":      "C_ARITHMETIC",
	"eq":       "C_ARITHMETIC",
	"gt":       "C_ARITHMETIC",
	"lt":       "C_ARITHMETIC",
	"and":      "C_ARITHMETIC",
	"or":       "C_ARITHMETIC",
	"not":      "C_ARITHMETIC",
	"if-goto":  "C_IF",
	"goto":     "C_GOTO",
	"label":    "C_LABEL",
	"function": "C_FUNCTION",
	"return":   "C_RETURN",
	"call":     "C_CALL",
}

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
// if command is in the defined list it is mapped to the type, otherwise invalid
func GetCommandType(line string) (command string) {
	firstWord := strings.Split(line, " ")[0]
	if command, exists := commandOps[firstWord]; exists {
		return command
	}

	fmt.Printf("Unable to identify command type for: %s\n", firstWord)
	return "INVALID_COMMAND"
}

func Args(line string) (argument1 string, argument2 string, err error) {
	// TODO: Clean this up, don't need 3 if-else statements
	args := strings.Split(line, " ")
	if len(args) == 1 {
		argument1 = args[0]
		argument2 = ""
		err = nil
	} else if len(args) == 2 {
		// for example C_LABEL needs the labelname (args[1])
		argument1 = args[1]
		err = nil
	} else if len(args) == 3 {
		argument1 = args[1]
		argument2 = args[2]
		err = nil
	} else {
		argument1 = ""
		argument2 = ""
		err = fmt.Errorf("total arguments passed can only be 1, 2 or 3, got: %d\n"+
			"from - %s", len(args), line)

	}

	return argument1, argument2, err
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
			line = strings.TrimSpace((line)) // TODO: just move the trim to after commend stripping, but need to test
		}

		// removes empty lines
		if len(line) == 0 {
			continue
		}

		// define instruction type before any processing
		currentInstruction := GetCommandType(line)
		instructionType = append(instructionType, currentInstruction)

		//arg1, arg2 := args(line)
		cleanedLines = append(cleanedLines, line)

	}
	return cleanedLines, instructionType
}
