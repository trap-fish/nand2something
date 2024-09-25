package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"vmtranslator/codeWriter"
	"vmtranslator/parser"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isFile(filename string) bool {
	return strings.HasSuffix(filename, ".vm")
}

func getFileList(dirName string) []string {

	var fileList []string
	c, err := os.ReadDir(dirName)
	check(err)

	for _, entry := range c {
		if isFile(entry.Name()) {
			fullPath := filepath.Join(dirName, entry.Name())
			fileList = append(fileList, fullPath)
		}
	}
	return fileList
}

func main() {
	arg := os.Args[1]
	filename := arg
	var outf *os.File

	fileList := []string{filename}

	if !isFile(arg) {
		fileList = getFileList(arg)
		outname := filepath.Base(arg)
		outf, _ = os.Create(filename + outname + ".asm")
		defer outf.Close()
	} else {
		outf, _ = os.Create(strings.TrimSuffix(filename, ".vm") + ".asm")
		defer outf.Close()
	}

	for _, filepath := range fileList {

		// parse the file
		parsedFile, cmdType := parser.Parser(filepath)

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

}
