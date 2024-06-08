// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

// setup variables for loop iterator and sum to store values
    @i
    M=0

    @sum
    M=0

// fetches value stored in R1 and adds R0 to R0, R1 times
(LOOP)
    // if i == R1 then stop
    @i
    D=M
    @R1
    D=D-M
    @STOP
    D;JEQ

    // Otherwise sum = sum + R0
    @R0
    D=M

    // add value stored in R0 to sum
    @sum
    D=D+M
    M=D

    // increment iterator by 1 and return to LOOP
    @i
    M=M+1
    @LOOP
    0;JMP

// loop have completed, store result in R2
(STOP)
    @sum
    D=M
    @R2
    M=D

// infinite loop to prevent computer executing every ROM instruction after STOP
(END)
    @END
    0;JMP
