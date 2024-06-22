package code

func dest(destination string) (destBits string) {
	// TODO: need parser to sort alphabetically
	switch destination {
	case "":
		destBits = "000"
	case "M":
		destBits = "001"
	case "D":
		destBits = "010"
	case "DM":
		destBits = "011"
	case "A":
		destBits = "100"
	case "AM":
		destBits = "101"
	case "AD":
		destBits = "110"
	case "ADM":
		destBits = "111"
	}

	return destBits
}

func jump(jumpValue string) (jumpBits string) {
	// TODO: need parser to sort alphabetically
	switch jumpValue {
	case "":
		jumpBits = "000"
	case "JGT":
		jumpBits = "001"
	case "JEQ":
		jumpBits = "010"
	case "JGE":
		jumpBits = "011"
	case "JLT":
		jumpBits = "100"
	case "JNE":
		jumpBits = "101"
	case "JLE":
		jumpBits = "110"
	case "JMP":
		jumpBits = "111"
	}

	return jumpBits
}

func comp(compute string) (computeBits string) {
	// TODO: need parser to sort alphabetically
	switch compute {
	case "0":
		computeBits = "0101010"
	case "1":
		computeBits = "0111111"
	case "-1":
		computeBits = "0111010"
	case "D":
		computeBits = "0001100"
	case "A":
		computeBits = "0110000"
	case "M":
		computeBits = "1110000"
	case "!D":
		computeBits = "0001101"
	case "!A":
		computeBits = "0110001"
	case "!M":
		computeBits = "1110001"
	case "-D":
		computeBits = "0001111"
	case "-A":
		computeBits = "0110011"
	case "-M":
		computeBits = "1110011"
	case "D+1":
		computeBits = "0011111"
	case "A+1":
		computeBits = "0110111"
	case "M+1":
		computeBits = "1110111"
	case "D-1":
		computeBits = "0001110"
	case "A-1":
		computeBits = "0110010"
	case "M-1":
		computeBits = "1110010"
	case "D+A":
		computeBits = "0000010"
	case "D+M":
		computeBits = "1000010"
	case "D-A":
		computeBits = "0010011"
	case "D-M":
		computeBits = "1010011"
	case "A-D":
		computeBits = "0000111"
	case "M-D":
		computeBits = "1000111"
	case "D&A":
		computeBits = "0000000"
	case "D&M":
		computeBits = "1000000"
	case "A|D":
		computeBits = "0010101"
	case "M|D":
		computeBits = "1010101"
	}

	return computeBits
}

func GetBinary(compute string, destination string, jumpVal string) (computeBits string) {
	computeBits = "111" + comp(compute) + dest(destination) + jump(jumpVal)
	return computeBits
}
