//push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 17
@17
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

@EQ1_TRUE
D;JEQ
// if condition evaluates to false
@SP
A=M
M=0
@EQ1_END
0;JMP
(EQ1_TRUE)
@SP
A=M
M=-1
(EQ1_END)
@SP
M=M+1
//push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 16
@16
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

@EQ2_TRUE
D;JEQ
// if condition evaluates to false
@SP
A=M
M=0
@EQ2_END
0;JMP
(EQ2_TRUE)
@SP
A=M
M=-1
(EQ2_END)
@SP
M=M+1
//push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 17
@17
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

@EQ3_TRUE
D;JEQ
// if condition evaluates to false
@SP
A=M
M=0
@EQ3_END
0;JMP
(EQ3_TRUE)
@SP
A=M
M=-1
(EQ3_END)
@SP
M=M+1
//push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 891
@891
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

@LT4_TRUE
D;JLT
// if condition evaluates to false
@SP
A=M
M=0
@LT4_END
0;JMP
(LT4_TRUE)
@SP
A=M
M=-1
(LT4_END)
@SP
M=M+1
//push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 892
@892
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

@LT5_TRUE
D;JLT
// if condition evaluates to false
@SP
A=M
M=0
@LT5_END
0;JMP
(LT5_TRUE)
@SP
A=M
M=-1
(LT5_END)
@SP
M=M+1
//push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 891
@891
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

@LT6_TRUE
D;JLT
// if condition evaluates to false
@SP
A=M
M=0
@LT6_END
0;JMP
(LT6_TRUE)
@SP
A=M
M=-1
(LT6_END)
@SP
M=M+1
//push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 32766
@32766
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

@GT7_TRUE
D;JGT
// if condition evaluates to false
@SP
A=M
M=0
@GT7_END
0;JMP
(GT7_TRUE)
@SP
A=M
M=-1
(GT7_END)
@SP
M=M+1
//push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 32767
@32767
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

@GT8_TRUE
D;JGT
// if condition evaluates to false
@SP
A=M
M=0
@GT8_END
0;JMP
(GT8_TRUE)
@SP
A=M
M=-1
(GT8_END)
@SP
M=M+1
//push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 32766
@32766
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

@GT9_TRUE
D;JGT
// if condition evaluates to false
@SP
A=M
M=0
@GT9_END
0;JMP
(GT9_TRUE)
@SP
A=M
M=-1
(GT9_END)
@SP
M=M+1
//push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 53
@53
D=A
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
//push constant 112
@112
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
// neg
@SP
AM=M-1
D=-M
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
D=D&M

M=D
@SP
M=M+1
//push constant 82
@82
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
D=D|M

M=D
@SP
M=M+1
//not
@SP
AM=M-1
M=!M
@SP
M=M+1

0;JEQ