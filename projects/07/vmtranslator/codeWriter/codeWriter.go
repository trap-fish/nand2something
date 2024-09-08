package codeWriter

import (
	"fmt"
	"os"
)

// takes the command's segment and maps to the assembly symbol
func mapSegmentSymbol(segment string) (segSymbol string) {

	segmentSymbols := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		"temp":     "TEMP",
	}

	segSymbol = segmentSymbols[segment]

	return segSymbol
}

func generatePushCode(vmSegment string, vmIndex string) (assemblyCode string) {
	symbol := mapSegmentSymbol(vmSegment)
	pushCode :=
		"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n"
	// constant can only be push command
	if vmSegment == "constant" {
		assemblyCode += fmt.Sprintf("//push constant %s\n", vmIndex)
		assemblyCode += fmt.Sprintf("@%s\nD=A\n", vmIndex)
		assemblyCode += fmt.Sprintf(pushCode)
	} else {
		assemblyCode += fmt.Sprintf("//push %s %s\n", vmSegment, vmIndex)
		assemblyCode += fmt.Sprintf("@%s\nA=M+%s\nD=M\n", symbol, vmIndex)
		assemblyCode += fmt.Sprintf(pushCode)
	}

	return assemblyCode

}

func generatePopCode(vmSegment string, vmIndex string) (assemblyCode string) {
	symbol := mapSegmentSymbol(vmSegment)
	popCode :=
		"D=D+A\n" +
			"@R13\n" +
			"M=D\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@R13\n" +
			"A=M\n" +
			"M=D\n"
	assemblyCode += fmt.Sprintf("//pop %s %s\n", vmSegment, vmIndex)
	assemblyCode += fmt.Sprintf("@%s\nD=M\n@%s\n", symbol, vmIndex)
	assemblyCode += fmt.Sprintf(popCode)

	return assemblyCode

}

func WritePushPop(file *os.File, cmdType string, segment string, index string) error {
	var assemblyCode string

	if cmdType == "C_PUSH" {
		assemblyCode = generatePushCode(segment, index)
	} else if cmdType == "C_POP" {
		assemblyCode = generatePopCode(segment, index)
	} else {
		fmt.Errorf("failed to write to file: Type:%s - %s %s", cmdType, segment, index)
	}

	_, err := file.WriteString(assemblyCode)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil

}
