package main

import (
	"fmt"
	"os"

	"assembler/code"
	"assembler/parser"
)

func main() {
	testFile := "./MaxL.asm"
	parsedFile, instType := parser.Parser(testFile)
	outf, _ := os.Create("./Max.hack")
	defer outf.Close()

	for i := range parsedFile {
		currentLine, currentInst := parsedFile[i], instType[i]
		var coded string
		if currentInst == "C" {
			destination := parser.Dest(currentLine)
			compute := parser.Comp(currentLine)
			jump := parser.Jump(currentLine)
			coded = code.GetBinary(compute, destination, jump)
		} else if currentInst == "A" {
			coded = parser.ToBinary(currentLine)
		}

		outf.WriteString(coded + "\n")
		fmt.Printf("wrote %s bytes\n", coded)

		fmt.Printf("ROM Instruction: %d : \t %s \t - Type: %s \t encoding: %s\n", i, parsedFile[i], instType[i], coded)
	}
	outf.Sync()
	//fmt.Printf("DM Out: %s + %s + %s\n", code.GetBinary(parser.Comp("D+M"), parser.Jump("D+M"), parser.Dest("D+M")))
}
