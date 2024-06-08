// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

// this tracks the pixel state from last action (black or white)
// initially 0 (white screen)
@pixelvalue
M=0

// listen for keyboard input
// if not pressed but screen is black, goto ACTION so screen goes white
(KBDLOOP)
    @KBD
    D=M
    
    // if KBD > 0 (key pressed) goto ACTION, so screen goes black
    @ACTION
    D;JGT

    //if KBD==0 but previous pixelvalue!=0 goto ACTION, so switch back to white
    @pixelvalue
    D=M
    @ACTION
    D;JLE

    // if KBD==0 continue listening for keyboard input
    @KBDLOOP
    D;JEQ

// setup the constants and variables used in SCREENLOOP
(ACTION)
    // n=8192 registers in the SCREEN memory map, need to loop over n times
    @8192
    D=A  
    @n
    M=D

    // set addr value will point @SCREEN address initially 
    @SCREEN
    D=A 
    @addr
    M=D

    // initialise iterator, i used to exit SCREENLOOP when i==n
    @i
    M=0

// update pixel target value depending on whether keyboard is pressed
(CHANGEPIXEL)
    @KBD
    D=M

    //if KBD > 0, use pixelvalue -1 (black) then goto screenloop
    @pixelvalue
    M=-1
    @SCREENLOOP
    D;JGT

    @pixelvalue
    M=0
    
(SCREENLOOP)
    // if i > n then exit loop
    @i 
    D=M
    @n 
    D=D-M 
    @KBDLOOP
    D;JEQ

    // fetch pixel value for black/white screen
    @pixelvalue
    D=M
    // turn RAM[screen+i] to pixelvalue
    @addr 
    A=M  //fetch address of the value in addr
    M=D // set RAM[addr] = pixelvalue

    @i 
    M=M+1

    // increment addr by 1, i.e. next register
    @1 
    D=A
    @addr 
    M=D+M
    
    // next iteration
    @SCREENLOOP
    0;JMP

    // goto ACTION if KBD > 0
    //@ACTION
    //D;JGT

    // if KBD = 0 return to KBDLOOP
    @KBDLOOP
    D;JEQ