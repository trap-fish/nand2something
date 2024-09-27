package codeWriter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// labels defined require a unique name, this is used as an index
var labelCounter int
var returnCounter int

// global variable to store function name each time a function command is parsed
var globalFilename string

// global variable to store function name each time a function command is parsed
var GlobalFuncName = ""

// global variable to avoid repetitive creation of push code
var pushCode = "@SP\n" +
	"A=M\n" +
	"M=D\n" +
	"@SP\n" +
	"M=M+1\n"

func SetFilename(filename string) {
	globalFilename = strings.TrimSuffix(filename, ".vm")

}

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

	var symbol string

	// static has no symbol, but a unique label based on filename.index
	if vmSegment == "static" {
		symbol = generateStaticLabel(vmIndex)
	} else {
		symbol = mapSegmentSymbol(vmSegment, vmIndex)
	}

	// constant can only be push command
	if vmSegment == "constant" {
		assemblyCode += "//push constant " + vmIndex + "\n"
		assemblyCode += "@" + vmIndex + "\nD=A\n"
		assemblyCode += pushCode
	} else if vmSegment == "temp" || vmSegment == "pointer" || vmSegment == "static" {
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
	var symbol string

	// static has no symbol, but a unique label based on filename.index
	if vmSegment == "static" {
		symbol = generateStaticLabel(vmIndex)
	} else {
		symbol = mapSegmentSymbol(vmSegment, vmIndex)
	}
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

	if vmSegment == "temp" || vmSegment == "pointer" || vmSegment == "static" {
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

// static variables are to have a format Foo.i where Foo is the filename and i the index
func generateStaticLabel(index string) (staticLabel string) {
	staticLabel = globalFilename + "." + index
	return staticLabel
}

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
		"//" + operator + "\n" +
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

func generateGoTo(labelName string) (assemblyCode string) {
	// unconditional jump to @argument
	assemblyCode =
		"@" + generateFuncLabelName(labelName) + "\n" +
			"0;JMP\n"
	return assemblyCode
}

func generatePushSymbol(symbol string) (assembleyCode string) {
	assembleyCode =
		"// push " + symbol + "\n" +
			"@" + symbol + "\n" +
			"D=M\n" +
			pushCode

	return assembleyCode

}

func generateCallCode(function string, args string) (assemblyCode string) {
	returnLabel := getReturnLabel(function)
	assemblyCode += "// call " + function + " " + args + "\n"
	//save the return address and current memory segments
	assemblyCode += generateReturnAddr(returnLabel)
	assemblyCode += generatePushSymbol("LCL")
	assemblyCode += generatePushSymbol("ARG")
	assemblyCode += generatePushSymbol("THIS")
	assemblyCode += generatePushSymbol("THAT")
	assemblyCode +=
		"//set LCL to SP, reposition ARG, then go to function\n" +
			"@SP\n" +
			"D=M\n" +
			"@LCL\n" +
			"M=D\n" +
			"@5\n" +
			"D=D-A\n" +
			"@" + args + "\n" +
			"D=D-A\n" +
			"@ARG\n" +
			"M=D\n" +
			"@" + function + "\n" +
			"0;JMP\n"
	assemblyCode += "(" + returnLabel + ")\n"

	return assemblyCode
}

func WriteCall(file *os.File, function string, args string) (err error) {

	assemblyCode := generateCallCode(function, args)

	_, err = file.WriteString(assemblyCode)

	if err != nil {
		return fmt.Errorf("failed writing call operation to file: %v", err)
	}

	return nil

}

func WriteFunction(file *os.File, function string, nVars string) (err error) {

	functionName := function
	localVars := generatePushCode("constant", "0")
	numLocalVars, _ := strconv.Atoi(nVars)

	assemblyCode := "// function " + function + " " + nVars + "\n" +
		"(" + functionName + ")\n" +
		strings.Repeat(localVars, numLocalVars)

	_, err = file.WriteString(assemblyCode)

	if err != nil {
		return fmt.Errorf("failed writing function operation to file: %v", err)
	}

	return nil
}

func getReturnLabel(functionName string) (returnLabel string) {
	// create a unique label for function return address
	returnId := strconv.Itoa(returnCounter)
	returnLabel = functionName + "$ret" + returnId
	returnCounter += 1

	return returnLabel
}

func generateReturnAddr(returnLabelName string) (assemblyCode string) {
	// creates label for filename.functionName return address
	// then pushes the address of this label onto the stack
	assemblyCode =
		"//create function return address and push address to stack\n" +
			//TODO - confirm that this code isn't needed!"(" + returnLabelName + ")\n" +
			"@" + returnLabelName + "\n" +
			//"0;JMP\n" +
			"D=A\n" +
			pushCode

	return assemblyCode
}

func WriteGoTo(file *os.File, argument string) (err error) {
	assemblyCode := generateGoTo(argument)

	_, err = file.WriteString(assemblyCode)

	if err != nil {
		return fmt.Errorf("failed writing goto operation to file: %v", err)
	}

	return nil

}

func WriteIf(file *os.File, label string) (err error) {
	// if the top most value on stack is true/1, jump to @label
	// Compiler will handle prior logic to ensure the value in SP-1 addr
	// was the result of conditional logic
	label = generateFuncLabelName(label)
	assemblyCode :=
		"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@" + label + "\n" +
			"D;JNE\n"

	_, err = file.WriteString(assemblyCode)

	if err != nil {
		return fmt.Errorf("failed writing if-goto operation to file: %v", err)
	}

	return nil

}

func generateFuncLabelName(label string) string {
	// a label must always be defined from inside a function
	// the function name of the function being parsed will be stored in
	// global variable GlobalFuncName
	return GlobalFuncName + label
}

func WriteLabel(file *os.File, label string) (err error) {
	label = generateFuncLabelName(label)
	assemblyCode :=
		"// label for " + label + " loop\n" +
			"(" + label + ")\n"

	_, err = file.WriteString(assemblyCode)
	if err != nil {
		return fmt.Errorf("failed writing loop operation to file: %v", err)
	}

	return nil

}

func generateReturnCode() (assembleyCode string) {
	// endframe = LCL
	assembleyCode =
		"//------- return start -------\n" +
			"// endframe = LCL\n" +
			"@LCL\n" +
			"D=M\n" +
			"@R13\n" +
			"M=D\n" +
			"\n" +
			"// retAddr = *(endframe-5)\n" +
			"@5\n" +
			"A=D-A\n" +
			"D=M\n" +
			"@R14\n" +
			"M=D\n" +
			"\n" +
			"//*ARG = pop()\n" +
			"@SP\n" +
			"AM=M-1\n" +
			"D=M\n" +
			"@ARG\n" +
			"A=M\n" +
			"M=D\n" +
			"\n" +
			"//SP = ARG+1\n" +
			"@ARG\n" +
			"D=M\n" +
			"@SP\n" +
			"M=D+1\n" +
			"\n" +
			"// THAT = *(endframe-1)\n" +
			"@R13\n" +
			"AM=M-1 // this avoids having to subtract endframe-n where n>1\n" +
			"D=M\n" +
			"@THAT\n" +
			"M=D\n" +
			"\n" +
			"// THIS = *(endframe-2)\n" +
			"@R13\n" +
			"AM=M-1 // \n" +
			"D=M\n" +
			"@THIS\n" +
			"M=D\n" +
			"\n" +
			"// ARG = *(endframe-3)\n" +
			"@R13\n" +
			"AM=M-1 // \n" +
			"D=M\n" +
			"@ARG\n" +
			"M=D\n" +
			"\n" +
			"// LCL = *(endframe-4)\n" +
			"@R13\n" +
			"AM=M-1 // \n" +
			"D=M\n" +
			"@LCL\n" +
			"M=D\n" +
			"\n" +
			"// goto retAddr\n" +
			"@R14\n" +
			"A=M\n" +
			"0;JMP\n" +
			"//------- return end -------\n" +
			"\n"

	return strings.TrimSpace(assembleyCode)
}

func WriteReturn(file *os.File) (err error) {

	assemblyCode := generateReturnCode()

	_, err = file.WriteString(assemblyCode)

	if err != nil {
		return fmt.Errorf("failed writing return operation to file: %v", err)
	}

	return nil
}

func WriteInit(file *os.File) (err error) {
	functionName := "Sys.init"
	assemblyCode :=
		"@256\n" +
			"D=A\n" +
			"@SP\n" +
			"M=D\n" +
			generateCallCode(functionName, "0")

	_, err = file.WriteString(assemblyCode)

	if err != nil {
		return fmt.Errorf("failed writing sys.init  to file: %v", err)
	}

	fmt.Print(assemblyCode)

	return nil
}
