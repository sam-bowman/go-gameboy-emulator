package main

import "log"

//CUSTOM FUNC - INC_r8
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

//CUSTOM FUNC - DEC_r8
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

//CUSTOM FUNC - ADD_r8_r8
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

//CUSTOM FUNC - ADD_r8_n8
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

//CUSTOM FUNC - ADD_r8_n8
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
