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
