package codeWriter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// takes the command's segment and maps to the assembly symbol
func mapSegmentSymbol(segment string, vmIndex string) (segSymbol string) {
	var tempRAM = [8]string{"R5", "R6", "R7", "R8", "R9", "R10", "R11", "R12"}
	var pointer = [2]string{"THIS", "THAT"}

	if segment == "temp" {
		index, _ := strconv.Atoi(vmIndex)
		return tempRAM[index]
	} else if segment == "pointer" {
		index, _ := strconv.Atoi(vmIndex)
		return pointer[index]
	}
	segmentSymbols := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		"eq":       "JEQ",
		"gt":       "JGT",
		"lt":       "JLT",
		"and":      "&",
		"or":       "|",
		"not":      "!",
	}

	segSymbol = segmentSymbols[segment]

	return segSymbol
}

func generatePushCode(vmSegment string, vmIndex string) (assemblyCode string) {
	symbol := mapSegmentSymbol(vmSegment, vmIndex)
	pushCode :=
		"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n"
	// constant can only be push command
	if vmSegment == "constant" {
		assemblyCode += "//push constant " + vmIndex + "\n"
		assemblyCode += "@" + vmIndex + "\nD=A\n"
		assemblyCode += pushCode
	} else if vmSegment == "temp" || vmSegment == "pointer" {
		assemblyCode += "//push " + vmSegment + " " + vmIndex + "\n"
		assemblyCode += "@" + symbol + "\nD=M\n"
		assemblyCode += pushCode
	} else {
		assemblyCode += "//push " + vmSegment + " " + vmIndex + "\n"
		assemblyCode += "@" + symbol + "\nD=M\n" + "@" + vmIndex + "\nA=D+A\nD=M\n"
		assemblyCode += pushCode
	}

	return assemblyCode

}

func generatePopCode(vmSegment string, vmIndex string) (assemblyCode string) {
	symbol := mapSegmentSymbol(vmSegment, vmIndex)
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

	tempPopCode :=
		"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@" + symbol + "\n" +
			"M=D\n"

	if vmSegment == "temp" || vmSegment == "pointer" {
		assemblyCode = "//pop " + vmSegment + " " + vmIndex + "\n"
		assemblyCode += tempPopCode
		return assemblyCode
	}
	assemblyCode = "//pop " + vmSegment + " " + vmIndex + "\n"
	assemblyCode += "@" + symbol + "\nD=M\n@" + vmIndex + "\n"
	assemblyCode += popCode

	return assemblyCode

}

func WritePushPop(file *os.File, cmdType string, segment string, index string) error {
	var assemblyCode string

	if cmdType == "C_PUSH" {
		assemblyCode = generatePushCode(segment, index)
	} else if cmdType == "C_POP" {
		assemblyCode = generatePopCode(segment, index)
	} else {
		err1 := fmt.Errorf("writePushPop recieved a non C_PUSH/C_POP command: Type:%s - %s %s", cmdType, segment, index)
		return err1
	}

	_, err2 := file.WriteString(assemblyCode)
	if err2 != nil {
		return fmt.Errorf("failed to write to file: %v", err2)
	}

	return nil

}

func getLogicalAssembly(operator string) (logAssemblyCode string) {
	opSym := mapSegmentSymbol(operator, "0") // operator does not need an index
	logAssemblyCode =
		"// get two values from stack to compare\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=D" + opSym + "M\n\n" +
			"M=D\n" +

			"@SP\n" +
			"M=M+1\n"
	return logAssemblyCode
}

// labels defined require a unique name, this is used as an index
var labelCounter int

func getConditionalAssembly(operator string) (condAssemblyCode string) {
	opSym := mapSegmentSymbol(operator, "0") // operator does not need an index
	op := strings.ToUpper(operator)
	id := strconv.Itoa((labelCounter))
	labelId := op + id

	condAssemblyCode =
		"// get two values from stack to compare\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=M-D\n\n" +

			"@" + labelId + "_TRUE\n" +
			"D;" + opSym + "\n" +
			"// if condition evaluates to false\n" +
			"@SP\n" +
			"A=M\n" +
			"M=0\n" +
			"@" + labelId + "_END\n" +
			"0;JMP\n" +

			"(" + labelId + "_TRUE)\n" +
			"@SP\n" +
			"A=M\n" +
			"M=-1\n" +

			"(" + labelId + "_END)\n" +
			"@SP\n" +
			"M=M+1\n"
	return condAssemblyCode
}

func generateArithmeticCode(operator string) (assemblyCode string) {
	var asmOperation string
	spInc := "@SP\nM=M+1\n" // repetitive code to increment stack pointer

	switch operator {
	case "add":
		asmOperation = "D=D+M\n"
	case "sub":
		asmOperation = "D=M-D"
	case "neg":
		return "// neg\n@SP\nAM=M-1\nD=-M\n@SP\nA=M\nM=D\n" + spInc
	case "eq":
		labelCounter++
		return getConditionalAssembly(operator)
	case "gt":
		labelCounter++
		return getConditionalAssembly(operator)
	case "lt":
		labelCounter++
		return getConditionalAssembly(operator)
	case "and":
		return getLogicalAssembly(operator)
	case "or":
		return getLogicalAssembly(operator)
	case "not":
		return "//not\n@SP\nAM=M-1\nM=!M\n" + spInc

	default:
		return "// function not yet defined"
	}
	assemblyCode =
		"// add\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@SP\n" +
			"AM=M-1\n" +
			asmOperation +
			"\n@SP\n" +
			"A=M\n" +
			"M=D\n" +
			spInc

	return assemblyCode

}

func WriteArithmetic(file *os.File, operator string) (err error) {
	assemblyCode := generateArithmeticCode(operator)

	_, err = file.WriteString(assemblyCode)
	if err != nil {
		return fmt.Errorf("failed writing arithmetic operation to file: %v", err)
	}

	return nil
}
