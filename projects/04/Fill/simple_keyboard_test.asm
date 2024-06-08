
// listen for keyboard input
// if key pressed, goto ACTION
(LOOP)
    @KBD
    D=M
    
    //Assign value in RAM[5] to Keyboard char value
    @R5
    M=D
    
    // goto ACTION if KDB > 0
    @ACTION
    D;JGT

    // if KBD = 0 return to LOOP
    @LOOP
    D;JEQ

// return to LOOP if key released, loop over ACTION if depressed
(ACTION)
    @KBD
    D=M

    // goto ACTION if KBD > 0
    @ACTION
    D;JGT

    // if KBD = 0 return to LOOP
    @LOOP
    D;JEQ