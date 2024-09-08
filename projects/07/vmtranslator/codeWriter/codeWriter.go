package codeWrite

import (
	"fmt"
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

func generatePushCode(vmSegment string, vmIndex int) (assemblyCode string) {
	symbol := mapSegmentSymbol(vmSegment)
	pushCode :=
		"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n"
	if vmSegment == "constant" {
		assemblyCode += fmt.Sprintf("//push constant %d\n", vmIndex)
		assemblyCode += fmt.Sprintf("@%d\nD=A\n", vmIndex)
		assemblyCode += fmt.Sprintf(pushCode)
	} else {
		assemblyCode += fmt.Sprintf("//push %s %d\n", vmSegment, vmIndex)
		assemblyCode += fmt.Sprintf("@%s\nA=M+%d\nD=M\n", symbol, vmIndex)
		assemblyCode += fmt.Sprintf(pushCode)
	}

	return assemblyCode

}

func generatePopCode(vmSegment string, vmIndex int) (assemblyCode string) {
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
	assemblyCode += fmt.Sprintf("//pop %s %d\n", vmSegment, vmIndex)
	assemblyCode += fmt.Sprintf("@%s\nD=M\n@%d\n", symbol, vmIndex)
	assemblyCode += fmt.Sprintf(popCode)

	return assemblyCode

}

func writePushPop(cmdType string, segment string, index int) {
	//var asmLines string
	// constant can only be push command
	if cmdType == "C_PUSH" {
		pushCode := generatePushCode(segment, index)
		fmt.Printf(pushCode)
	} else if cmdType == "C_POP" {
		popCode := generatePopCode(segment, index)
		fmt.Printf(popCode)
	} else {
		fmt.Printf("do nothing %s\n", cmdType)
	}
}
