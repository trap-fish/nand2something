@256
D=A
@SP
M=D
// call Sys.init 0
//create function return address and push address to stack
@Sys.init$ret0
D=A
@SP
A=M
M=D
@SP
M=M+1
// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
//set LCL to SP, reposition ARG, then go to function
@SP
D=M
@LCL
M=D
@5
D=D-A
@0
D=D-A
@ARG
M=D
@Sys.init
0;JMP
(Sys.init$ret0)
// function Main.fibonacci 0
(Main.fibonacci)
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
//push constant 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
// get two values from stack to compare
@SP
AM=M-1
D=M
@SP
AM=M-1
D=M-D

@LT1_TRUE
D;JLT
// if condition evaluates to false
@SP
A=M
M=0
@LT1_END
0;JMP
(LT1_TRUE)
@SP
A=M
M=-1
(LT1_END)
@SP
M=M+1
@SP
AM=M-1
D=M
@Main.fibonacciN_LT_2
D;JLT
@Main.fibonacciN_GE_2
0;JMP
// label for Main.fibonacciN_LT_2 loop
(Main.fibonacciN_LT_2)
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
//------- return end -------// label for Main.fibonacciN_GE_2 loop
(Main.fibonacciN_GE_2)
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
//push constant 2
@2
D=A
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
// call Main.fibonacci 1
//create function return address and push address to stack
@Main.fibonacci$ret1
D=A
@SP
A=M
M=D
@SP
M=M+1
// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
//set LCL to SP, reposition ARG, then go to function
@SP
D=M
@LCL
M=D
@5
D=D-A
@1
D=D-A
@ARG
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci$ret1)
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
//push constant 1
@1
D=A
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
// call Main.fibonacci 1
//create function return address and push address to stack
@Main.fibonacci$ret2
D=A
@SP
A=M
M=D
@SP
M=M+1
// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
//set LCL to SP, reposition ARG, then go to function
@SP
D=M
@LCL
M=D
@5
D=D-A
@1
D=D-A
@ARG
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci$ret2)
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
//------- return end -------// function Sys.init 0
(Sys.init)
//push constant 4
@4
D=A
@SP
A=M
M=D
@SP
M=M+1
// call Main.fibonacci 1
//create function return address and push address to stack
@Main.fibonacci$ret3
D=A
@SP
A=M
M=D
@SP
M=M+1
// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
//set LCL to SP, reposition ARG, then go to function
@SP
D=M
@LCL
M=D
@5
D=D-A
@1
D=D-A
@ARG
M=D
@Main.fibonacci
0;JMP
(Main.fibonacci$ret3)
// label for Sys.initEND loop
(Sys.initEND)
@Sys.initEND
0;JMP
