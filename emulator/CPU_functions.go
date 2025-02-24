package main

import "log"

//Increments ------
func INC_r8(GB *GAMEBOY, left uint8) uint8 {
	log.Println("INC_r8")

	//Perform Operation
	original := left
	result := left + 1

	//Set Flags
	if result == 0 {
		GB.CPU._r.F.Z = 1
	} else {
		GB.CPU._r.F.Z = 0
	}

	GB.CPU._r.F.N = 0

	if original&0x0F == 0x0F && original&0xFF != 0xFF {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': '-'}
	return result
}

func INC_r8r8(GB *GAMEBOY, upper uint8, lower uint8) (uint8, uint8) {
	log.Println("INC_r8r8")

	//Perform Operation
	combined := uint16(upper)<<8 | uint16(lower)
	combined += 1

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return byte(combined >> 8), byte(combined & 0xFF)
}

func INC_r16(GB *GAMEBOY, left uint16) uint16 {
	log.Println("INC_r16")

	//Perform Operation
	result := left + 1

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//Decrements ------
func DEC_r8(GB *GAMEBOY, left uint8) uint8 {
	log.Println("DEC_r8")

	//Perform Operation
	result := left - 1

	if GB.CPU._r.B == 0 {
		GB.CPU._r.F.Z = 1
	} else {
		GB.CPU._r.F.Z = 0
	}

	GB.CPU._r.F.N = 1

	if GB.CPU._r.B&0x0F == 0x0F {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}

	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': '-'}
	return result
}

func DEC_r8r8(GB *GAMEBOY, upper uint8, lower uint8) (uint8, uint8) {
	log.Println("DEC_r8r8")

	//Perform Operation
	combined := uint16(upper)<<8 | uint16(lower)
	combined -= 1

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return byte(combined >> 8), byte(combined & 0xFF)
}

func DEC_r16(GB *GAMEBOY, left uint16) uint16 {
	log.Println("DEC_r16")

	//Perform Operation
	result := left - 1

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//Additions ------
func ADD_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	log.Println("ADD_r8_r8")

	//Perform Operation
	result := left + right
	result16 := uint16(left) + uint16(left)

	//Set Flags
	if result == 0 {
		GB.CPU._r.F.Z = 1
	} else {
		GB.CPU._r.F.Z = 0
	}

	GB.CPU._r.F.N = 0

	if (((left & 0xF) + (right & 0xF)) & 0x10) == 0x10 {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}

	if result16 >= 0x100 {
		GB.CPU._r.F.C = 1
	} else {
		GB.CPU._r.F.C = 0
	}

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': 'C'}
	return result
}

func ADD_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	log.Println("ADD_r8_n8")

	//Perform Operation
	result := left + right
	result16 := uint16(left) + uint16(left)

	//Set Flags
	if result == 0 {
		GB.CPU._r.F.Z = 1
	} else {
		GB.CPU._r.F.Z = 0
	}

	GB.CPU._r.F.N = 0

	if (((left & 0xF) + (right & 0xF)) & 0x10) == 0x10 {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}

	if result16 >= 0x100 {
		GB.CPU._r.F.C = 1
	} else {
		GB.CPU._r.F.C = 0
	}

	//Set PC & Timings
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': 'C'}
	return result
}

func ADD_r8r8_r8r8(GB *GAMEBOY, leftUpper uint8, leftLower uint8, rightUpper uint8, rightLower uint8) (uint8, uint8) {
	log.Println("ADD_r8r8_r8r8")

	//Perform Operation
	left := uint16(leftUpper)<<8 | uint16(leftLower)
	right := uint16(rightUpper)<<8 | uint16(rightLower)
	result := left + right
	result32 := uint32(left) + uint32(right)

	//Set Flags
	GB.CPU._r.F.N = 0

	if (((left & 0xFFF) + (right & 0xFFF)) & 0x1000) == 0x1000 {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}

	if result32 >= 0x00010000 {
		GB.CPU._r.F.C = 1
	} else {
		GB.CPU._r.F.C = 0
	}

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '0', 'H': 'H', 'C': 'C'}
	return uint8(result >> 8), uint8(result & 0xFF)
}

func ADD_r8r8_r16(GB *GAMEBOY, leftUpper uint8, leftLower uint8, right uint16) (uint8, uint8) {
	log.Println("ADD_r8r8_r16")

	//Perform Operation
	left := uint16(leftUpper)<<8 | uint16(leftLower)
	result := left + right
	result32 := uint32(left) + uint32(right)

	//Set Flags
	GB.CPU._r.F.N = 0

	if (((left & 0xFFF) + (right & 0xFFF)) & 0x1000) == 0x1000 {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}

	if result32 >= 0x00010000 {
		GB.CPU._r.F.C = 1
	} else {
		GB.CPU._r.F.C = 0
	}

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '0', 'H': 'H', 'C': 'C'}
	return uint8(result >> 8), uint8(result & 0xFF)
}

//Loads ------
func LD_r8r8_n16(GB *GAMEBOY) (uint8, uint8) {
	log.Println("LD_r8r8_n16")

	//Perform Operation
	upper := GB.MMU._rom[GB.CPU._r.PC+2]
	lower := GB.MMU._rom[GB.CPU._r.PC+1]

	//Set PC & Timings
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return upper, lower
}

func LD_r16_n16(GB *GAMEBOY) uint16 {
	log.Println("LD_r8r8_n16")

	//Perform Operation
	upper := GB.MMU._rom[GB.CPU._r.PC+2]
	lower := GB.MMU._rom[GB.CPU._r.PC+1]
	combined := uint16(upper)<<8 | uint16(lower)

	//Set PC & Timings
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return combined
}

func LD_r8_n8(GB *GAMEBOY) uint8 {
	log.Println("LD_r8_n8")

	//Perform Operation
	result := GB.MMU._rom[GB.CPU._r.PC+1]

	//Set PC & Timings
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

func LD_r8_r8(GB *GAMEBOY, right uint8) uint8 {
	log.Println("LD_r8_r8")

	//Perform Operation
	result := right

	//Set PC & Timings
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}
