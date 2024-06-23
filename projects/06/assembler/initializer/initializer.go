package initializer

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initSymbolTable() (symbols map[string]int) {
	symbols = make(map[string]int)

	// pre-defined list of labels mapped to memory location
	symbols["R0"] = 0
	symbols["R1"] = 1
	symbols["R2"] = 2
	symbols["R3"] = 3
	symbols["R4"] = 4
	symbols["R5"] = 5
	symbols["R6"] = 6
	symbols["R7"] = 7
	symbols["R8"] = 8
	symbols["R9"] = 9
	symbols["R10"] = 10
	symbols["R11"] = 11
	symbols["R12"] = 12
	symbols["R13"] = 13
	symbols["R14"] = 14
	symbols["R15"] = 15
	symbols["SP"] = 0
	symbols["LCL"] = 1
	symbols["ARG"] = 2
	symbols["THIS"] = 3
	symbols["THAT"] = 4
	symbols["SCREEN"] = 16384
	symbols["KBD"] = 24576

	return symbols

}

func FirstPass(filepath string) (firstPSymbols map[string]int) {

	// open file and scan
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()

	var instructionType []string
	var sym string
	var counter int // tracks the ROM instruction number
	var lineNum int // tracks the line number, including labels

	// initialize a symboltable
	symbolTable := initSymbolTable()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// iterator used to index arrays independent of counter used for ROM
		line := scanner.Text()
		line = strings.TrimSpace((line))

		// remove comments and whitespace so symbol identifiers can classify line
		commentIdx := strings.Index(line, "//")
		if commentIdx != -1 {
			line = line[:commentIdx]
		}
		// advance to next line if current line is blank
		if len(line) == 0 {
			continue
		}

		// instruction types A and L will be  formatted @abc or (ABC)
		symbolIdx := strings.Index(line, "@")
		bracOpnIdx := strings.Index(line, "(")
		bracClsIdx := strings.Index(line, ")")

		// if line is @something, has to be A, otherwise C
		if symbolIdx == 0 {
			instructionType = append(instructionType, "A")
			sym = line[symbolIdx:]
		} else if bracOpnIdx == 0 {
			instructionType = append(instructionType, "L")
			sym = line[bracOpnIdx+1 : bracClsIdx]
		} else {
			instructionType = append(instructionType, "C")
			sym = ""
		}

		// increase counter if A or C
		//otherwise check if (LAB) is in symbolTable
		_, isFound := symbolTable[sym]
		if instructionType[lineNum] != "L" {
			counter += 1
		} else if !isFound {
			symbolTable[sym] = counter // add LABEL : nextInstructionNum
		}

		// line number increases regardless since used to track instruciton type
		lineNum += 1

	}
	return symbolTable
}
