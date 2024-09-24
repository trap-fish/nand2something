package main

import (
	"fmt"
	"os"
	"vmtranslator/codeWriter"
	"vmtranslator/parser"
)

func main() {

	filename := "Sys"
	filepath := "./testfiles/" + filename + ".vm"

	// parse the file
	parsedFile, cmdType := parser.Parser(filepath)
	outf, _ := os.Create("./testfiles/" + filename + ".asm")
	defer outf.Close()

	// write initialiser code to file
	codeWriter.WriteInit(outf)

	for el := range len(parsedFile) {
		arg1, arg2, err := parser.Args(parsedFile[el])
		if err != nil {
			panic(err)
		}
		fmt.Printf("line: %s - \t\t\t\tType: %s\n", parsedFile[el], cmdType[el])
		if cmdType[el] == "C_ARITHMETIC" {
			codeWriter.WriteArithmetic(outf, arg1)
		} else if cmdType[el] == "C_LABEL" {
			codeWriter.WriteLabel(outf, arg1)
		} else if cmdType[el] == "C_GOTO" {
			codeWriter.WriteGoTo(outf, arg1)
		} else if cmdType[el] == "C_IF" {
			codeWriter.WriteIf(outf, arg1)
		} else if cmdType[el] == "C_FUNCTION" {
			codeWriter.WriteFunction(outf, arg1, arg2)
		} else if cmdType[el] == "C_CALL" {
			codeWriter.WriteCall(outf, arg1, arg2)
		} else if cmdType[el] == "C_RETURN" {
			codeWriter.WriteReturn(outf)
		} else {
			codeWriter.WritePushPop(outf, cmdType[el], arg1, arg2)
		}

		// terminate the programme with infinite loop
		if el == len(parsedFile)-1 {
			outf.WriteString("\n(END)\n@END\n0;JEQ")
		}
	}

}
