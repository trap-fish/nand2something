package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// assumes this function is exclusively called after comments are ignored
// similar to parsing A instruction but included seperated so
func GetInstructionType(line string) (instruction string) {
	line = strings.TrimSpace(line)
	symbolIdx := strings.Index(line, "@")
	bracketIdx := strings.Index(line, "(")

	// if line is @something, has to be A, otherwise C
	if symbolIdx == 0 {
		instruction = "A"
	} else if bracketIdx == 0 {
		instruction = "L"
	} else {
		instruction = "C"
	}
	return instruction
}

func Symbol(line string) (address string) {
	line = strings.TrimSpace(line)
	symbolIdx := strings.Index(line, "@")
	bracketIdx := strings.Index(line, "(")
	// if line is @something, remove the @, otherwise ignore
	if symbolIdx == 0 {
		line = line[1:]
	} else if bracketIdx == 0 {
		line = line[1 : len(line)-1]
	} else {
		line = ""
	}
	line = strings.TrimSpace(line)
	return line

}

func Dest(line string) (destination string) {
	eqIdx := strings.Index(line, "=")
	// if line has x=y, extract x
	if eqIdx != -1 {
		line = line[:eqIdx]
	} else {
		line = ""
	}
	line = strings.TrimSpace(line)
	return line
}

func Comp(line string) (compute string) {

	semiIdx := strings.Index(line, ";")
	eqIdx := strings.Index(line, "=")
	// if line has x;y, extract x
	if semiIdx != -1 {
		line = line[eqIdx+1 : semiIdx]
		// if format x=y, extract y
	} else if eqIdx > 0 && semiIdx == -1 {
		line = line[eqIdx+1:]
	} else {
		line = ""
		//TODO: shouldn't have any D+M style instruction
		//since this is wasted compute not stored anywhere
		//need an error thrown if that is encountered this code will not handle it
	}
	line = strings.TrimSpace(line)
	return line
}

func Jump(line string) (compute string) {
	semiIdx := strings.Index(line, ";")
	// if line has x=y;z, extract
	if semiIdx != -1 {
		line = line[semiIdx+1:]
	} else {
		line = ""
	}
	line = strings.TrimSpace(line)
	return line
}

func CleanLine(line string) (cleaned string) {
	cleaned = strings.TrimSpace(line)
	return cleaned
}

// base2Conv converts a base 10 number to a binary string, capped at 17-bits
// could easily use fmt.Sprintf(%b, int), but wanted to implement manually
// even though this way uses fmt.Sprint to concat the string
func ToBinary(base10 string) (binary string) {
	// parser outputs strings, need to convert
	base10int, err := strconv.Atoi(base10)
	if err != nil {
		fmt.Printf("Error is with: %s\n\n", base10)
		panic(err)
	}
	// dealing with 16-bit computer, we don't care if values above are wrong
	// TODO: add error handler
	if base10int > 32767 {
		return "10000000000000000"
	}
	binary = ""
	for base10int > 0 && base10int < 32768 {
		remainder := base10int % 2
		base10int /= 2
		binary = fmt.Sprint(remainder, binary)
	}

	// needs to be 16-bit format
	bitCount := len(binary)
	if bitCount < 16 {
		pad := 16 - bitCount
		binary = strings.Repeat("0", pad) + binary
	}

	return binary
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
		currentInstruction := GetInstructionType(line)
		instructionType = append(instructionType, currentInstruction)

		if currentInstruction != "C" {
			aInst := Symbol(line)
			cleanedLines = append(cleanedLines, aInst)
		} else {

			cInst := line //destination + compute + jump
			cleanedLines = append(cleanedLines, cInst)
		}

	}
	return cleanedLines, instructionType
}
