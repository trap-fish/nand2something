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
// function Class1.set 0
(Class1.set)
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
//pop static 0
@StaticsTest.0
D=M
@0
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
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
//pop static 1
@StaticsTest.1
D=M
@1
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
//push constant 0
@0
D=A
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
//------- return end -------// function Class1.get 0
(Class1.get)
//push static 0
@StaticsTest.0
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
//push static 1
@StaticsTest.1
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
//------- return end -------// function Class2.set 0
(Class2.set)
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
//pop static 0
@StaticsTest.0
D=M
@0
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
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
//pop static 1
@StaticsTest.1
D=M
@1
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
//push constant 0
@0
D=A
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
//------- return end -------// function Class2.get 0
(Class2.get)
//push static 0
@StaticsTest.0
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
//push static 1
@StaticsTest.1
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
//------- return end -------// function Sys.init 0
(Sys.init)
//push constant 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 8
@8
D=A
@SP
A=M
M=D
@SP
M=M+1
// call Class1.set 2
//create function return address and push address to stack
@Class1.set$ret1
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
@2
D=D-A
@ARG
M=D
@Class1.set
0;JMP
(Class1.set$ret1)
//pop temp 0
@SP
AM=M-1
D=M
@R5
M=D
//push constant 23
@23
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 15
@15
D=A
@SP
A=M
M=D
@SP
M=M+1
// call Class2.set 2
//create function return address and push address to stack
@Class2.set$ret2
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
@2
D=D-A
@ARG
M=D
@Class2.set
0;JMP
(Class2.set$ret2)
//pop temp 0
@SP
AM=M-1
D=M
@R5
M=D
// call Class1.get 0
//create function return address and push address to stack
@Class1.get$ret3
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
@Class1.get
0;JMP
(Class1.get$ret3)
// call Class2.get 0
//create function return address and push address to stack
@Class2.get$ret4
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
@Class2.get
0;JMP
(Class2.get$ret4)
// label for Sys.initEND loop
(Sys.initEND)
@Sys.initEND
0;JMP
