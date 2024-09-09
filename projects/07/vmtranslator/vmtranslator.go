package main

import (
	"os"
	"vmtranslator/codeWriter"
	"vmtranslator/parser"
)

func main() {

	filename := "StackTest"
	filepath := "./testfiles/" + filename + ".vm"

	// parse the file
	parsedFile, cmdType := parser.Parser(filepath)
	outf, _ := os.Create("./testfiles/" + filename + ".asm")
	defer outf.Close()

	for el := range len(parsedFile) {
		arg1, arg2, err := parser.Args(parsedFile[el])
		if err != nil {
			panic(err)
		}
		if cmdType[el] == "C_ARITHMETIC" {
			codeWriter.WriteArithmetic(outf, arg1)
		} else {
			codeWriter.WritePushPop(outf, cmdType[el], arg1, arg2)
		}

		// terminate the programme with infinite loop
		if el == len(parsedFile)-1 {
			outf.WriteString("\n0;JEQ")
		}
	}

}
