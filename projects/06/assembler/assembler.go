package main

import (
	"os"
	"strconv"

	"assembler/code"
	"assembler/initializer"
	"assembler/parser"
)

func main() {

	filepath := "./test_files/PongL.asm"
	// first pass produces symbol table with labels included
	fpSymTab := initializer.FirstPass(filepath)

	// parse the file
	parsedFile, instType := parser.Parser(filepath)
	outf, _ := os.Create("./test_files/output_files/PongL.hack")
	defer outf.Close()

	var coded string // stores the converted instruction
	memAddress := 16 // first available memory address, increments each use

	for i := range parsedFile {
		currentLine, currentInst := parsedFile[i], instType[i]

		// check if the symbol is in the table for A instructions
		if currentInst == "A" {
			_, isFound := fpSymTab[currentLine]
			_, err := strconv.Atoi(currentLine)
			// not in table and can't convert, must be new variable name
			if !isFound && err != nil {
				fpSymTab[currentLine] = memAddress
				currentLine = strconv.Itoa(fpSymTab[currentLine]) // parser needs to take in string
				memAddress += 1
			} else if isFound {
				//if found map and convert to str
				currentLine = strconv.Itoa(fpSymTab[currentLine])
			}

		}

		if currentInst == "C" {
			destination := parser.Dest(currentLine)
			compute := parser.Comp(currentLine)
			jump := parser.Jump(currentLine)
			coded = code.GetBinary(compute, destination, jump)
		} else if currentInst == "A" {

			coded = parser.ToBinary(currentLine)
		} else {
			//fmt.Printf("ROM Instruction: %d : \t %s \t - Type: %s \t encoding: %s\n", i, parsedFile[i], instType[i], coded)
			continue
		}

		outf.WriteString(coded + "\n")
		//fmt.Printf("ROM Instruction: %d : \t %s \t - Type: %s \t encoding: %s\n", i, parsedFile[i], instType[i], coded)
	}
	outf.Sync()
	//fmt.Printf("DM Out: %s + %s + %s\n", code.GetBinary(parser.Comp("D+M"), parser.Jump("D+M"), parser.Dest("D+M")))
}
