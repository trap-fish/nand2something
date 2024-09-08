package main

import (
	"os"
	"vmtranslator/codeWriter"
	"vmtranslator/parser"
)

func main() {

	filepath := "../BasicTest.vm"

	// parse the file
	parsedFile, cmdType := parser.Parser(filepath)
	outf, _ := os.Create("./Basic.asm")
	defer outf.Close()

	for el := range len(parsedFile) {
		arg1, arg2, err := parser.Args(parsedFile[el])
		if err != nil {
			panic(err)
		}
		codeWriter.WritePushPop(outf, cmdType[el], arg1, arg2)
	}

}
