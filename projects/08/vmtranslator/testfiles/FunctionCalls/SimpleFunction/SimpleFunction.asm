
// function SimpleFunction.test 2
(SimpleFunction.test)
//push constant 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
//push local 0
@LCL
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
//push local 1
@LCL
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
//add
@SP
AM=M-1
D=M
@SP
AM=M-1
D=D+M

@SP
A=M
M=D
@SP
M=M+1
//not
@SP
AM=M-1
M=!M
@SP
M=M+1
//push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
//add
@SP
AM=M-1
D=M
@SP
AM=M-1
D=D+M

@SP
A=M
M=D
@SP
M=M+1
//push argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
//sub
@SP
AM=M-1
D=M
@SP
AM=M-1
D=M-D
@SP
A=M
M=D
@SP
M=M+1
//------- return start -------
// endframe = LCL
@LCL
D=M
@R13
M=D

// retAddr = *(endframe-5)
@5
A=D-A
D=M
@R14
M=D

//*ARG = pop()
@SP
AM=M-1
D=M
@ARG
A=M
M=D

//SP = ARG+1
@ARG
D=M
@SP
M=D+1

// THAT = *(endframe-1)
@R13
AM=M-1 // this avoids having to subtract endframe-n where n>1
D=M
@THAT
M=D

// THIS = *(endframe-2)
@R13
AM=M-1 // 
D=M
@THIS
M=D

// ARG = *(endframe-3)
@R13
AM=M-1 // 
D=M
@ARG
M=D

// LCL = *(endframe-4)
@R13
AM=M-1 // 
D=M
@LCL
M=D

// goto retAddr
@R14
A=M
0;JMP
//------- return end -------