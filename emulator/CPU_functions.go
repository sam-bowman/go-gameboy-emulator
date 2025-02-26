package main

//---------- Common ----------

//Set Clock Timings
func REG_CLOCK_TIMINGS(GB *GAMEBOY, PC uint16, M uint8) {
	GB.CPU._r.PC += PC
	GB.CPU._r.M = M
	GB.CPU._r.T = M * 4
}

func COMBINE_8BITS_TO_16BITS(upper uint8, lower uint8) uint16 {
	combined := uint16(upper)<<8 | uint16(lower)
	return combined
}

func SPLIT_16BITS_TO_8BITS(combined uint16) (uint8, uint8) {
	splitUpper := uint8((combined & 0xFF00) >> 8)
	splitLower := uint8(combined & 0x00FF)
	return splitUpper, splitLower
}

func SET_Z_FLAG(GB *GAMEBOY, on bool) {
	if on {
		GB.CPU._r.F.Z = 1
	} else {
		GB.CPU._r.F.Z = 0
	}
}
func GET_Z_FLAG(GB *GAMEBOY) uint8 {
	return GB.CPU._r.F.Z
}

func SET_N_FLAG(GB *GAMEBOY, on bool) {
	if on {
		GB.CPU._r.F.N = 1
	} else {
		GB.CPU._r.F.N = 0
	}
}
func GET_N_FLAG(GB *GAMEBOY) uint8 {
	return GB.CPU._r.F.N
}

func SET_H_FLAG(GB *GAMEBOY, on bool) {
	if on {
		GB.CPU._r.F.H = 1
	} else {
		GB.CPU._r.F.H = 0
	}
}
func GET_H_FLAG(GB *GAMEBOY) uint8 {
	return GB.CPU._r.F.H
}

func SET_C_FLAG(GB *GAMEBOY, on bool) {
	if on {
		GB.CPU._r.F.C = 1
	} else {
		GB.CPU._r.F.C = 0
	}
}
func GET_C_FLAG(GB *GAMEBOY) uint8 {
	return GB.CPU._r.F.C
}

func SET_Z_FLAG_result8(GB *GAMEBOY, result uint8) {
	if result == 0 {
		SET_Z_FLAG(GB, true)
	} else {
		SET_Z_FLAG(GB, false)
	}
}

func SET_Z_FLAG_result16(GB *GAMEBOY, result uint16) {
	if result == 0 {
		SET_Z_FLAG(GB, true)
	} else {
		SET_Z_FLAG(GB, false)
	}
}

func SET_H_FLAG_8(GB *GAMEBOY, left uint8, right uint8, positive bool) {
	leftLower := left & 0x0F
	rightLower := right & 0x0F
	if positive {
		if leftLower+rightLower >= 0x0F {
			SET_H_FLAG(GB, true)
		} else {
			SET_H_FLAG(GB, false)
		}
	} else {
		if rightLower > leftLower {
			SET_H_FLAG(GB, true)
		} else {
			SET_H_FLAG(GB, false)
		}
	}
}

func SET_H_FLAG_16(GB *GAMEBOY, left uint16, right uint16, positive bool) {
	leftLower := left & 0x0FFF
	rightLower := right & 0x0FFF
	if positive {
		if leftLower+rightLower >= 0x0FFF {
			SET_H_FLAG(GB, true)
		} else {
			SET_H_FLAG(GB, false)
		}
	} else {
		if rightLower > leftLower {
			SET_H_FLAG(GB, true)
		} else {
			SET_H_FLAG(GB, false)
		}
	}
}

func SET_C_FLAG_8(GB *GAMEBOY, left uint8, right uint8, positive bool) {
	leftUpper := (left & 0xF0) >> 4
	rightUpper := (right & 0xF0) >> 4
	if positive {
		if leftUpper+rightUpper >= 0x0F {
			SET_C_FLAG(GB, true)
		} else {
			SET_C_FLAG(GB, false)
		}
	} else {
		if right > left {
			SET_C_FLAG(GB, true)
		} else {
			SET_C_FLAG(GB, false)
		}
	}
}

func SET_C_FLAG_16(GB *GAMEBOY, left uint16, right uint16, positive bool) {
	leftUpper := (left & 0xF000) >> 12
	rightUpper := (right & 0xF000) >> 12
	if positive {
		if leftUpper+rightUpper >= 0x0F {
			SET_C_FLAG(GB, true)
		} else {
			SET_C_FLAG(GB, false)
		}
	} else {
		if right > left {
			SET_C_FLAG(GB, true)
		} else {
			SET_C_FLAG(GB, false)
		}
	}
}

//---------- Arithmetic ----------

//----- Increments ------

//INC 8-bit register
func INC_r8(GB *GAMEBOY, left uint8) uint8 {
	GB.InfoLogger.Println("INC_r8")

	//Perform Operation
	result := left + 1

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG_8(GB, left, 1, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': '-'}
	return result
}

//INC 2x 8-bit register (as 16 bit)
func INC_r8r8(GB *GAMEBOY, upper uint8, lower uint8) (uint8, uint8) {
	GB.InfoLogger.Println("INC_r8r8")

	//Perform Operation
	combined := COMBINE_8BITS_TO_16BITS(upper, lower)
	combined += 1

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return SPLIT_16BITS_TO_8BITS(combined)
}

//INC value ad address of 2x 8-bit register (as 16 bit)
func INC_addrr8r8(GB *GAMEBOY, upper uint8, lower uint8) {
	GB.InfoLogger.Println("INC_addrr8r8")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(upper, lower)
	addr += 1
	result := GB.MMU.readByte(GB, addr)
	updatedResult := result + 1
	GB.MMU.writeByte(GB, addr, updatedResult)

	//Set Flags
	SET_Z_FLAG_result8(GB, updatedResult)

	SET_N_FLAG(GB, false)

	SET_H_FLAG_8(GB, updatedResult, 1, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 3)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': '-'}
}

//INC 16 bit register
func INC_r16(GB *GAMEBOY, left uint16) uint16 {
	GB.InfoLogger.Println("INC_r16")

	//Perform Operation
	result := left + 1

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//----- Decrements ------

//DEC 8-bit register
func DEC_r8(GB *GAMEBOY, left uint8) uint8 {
	GB.InfoLogger.Println("DEC_r8")

	//Perform Operation
	result := left - 1

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, 1, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': '-'}
	return result
}

//DEC 2x 8-bit register (as 16 bit)
func DEC_r8r8(GB *GAMEBOY, upper uint8, lower uint8) (uint8, uint8) {
	GB.InfoLogger.Println("DEC_r8r8")

	//Perform Operation
	combined := COMBINE_8BITS_TO_16BITS(upper, lower)
	combined -= 1

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return SPLIT_16BITS_TO_8BITS(combined)
}

//DEC 2x 8-bit register (as 16 bit)
func DEC_addrr8r8(GB *GAMEBOY, upper uint8, lower uint8) {
	GB.InfoLogger.Println("DEC_addrr8r8")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(upper, lower)
	addr += 1
	result := GB.MMU.readByte(GB, addr)
	updatedResult := result + 1
	GB.MMU.writeByte(GB, addr, updatedResult)

	//Set Flags
	SET_Z_FLAG_result8(GB, updatedResult)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, result, 1, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 3)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': '-'}
}

//DEC 16 bit register
func DEC_r16(GB *GAMEBOY, left uint16) uint16 {
	GB.InfoLogger.Println("DEC_r16")

	//Perform Operation
	result := left - 1

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//----- Additions ------

//ADD 8-bit register to 8-bit register
func ADD_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("ADD_r8_r8")

	//Perform Operation
	result := left + right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG_8(GB, left, right, true)

	SET_C_FLAG_8(GB, left, right, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': 'C'}
	return result
}

//ADD following 8-bit value to 8-bit register
func ADD_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("ADD_r8_n8")

	//Perform Operation
	result := left + right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG_8(GB, left, right, true)

	SET_C_FLAG_8(GB, left, right, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': 'C'}
	return result
}

//ADD 2 8-bit registers (as 16-bit) to 2 8-bit registers (as 16-bit)
func ADD_r8r8_r8r8(GB *GAMEBOY, leftUpper uint8, leftLower uint8, rightUpper uint8, rightLower uint8) (uint8, uint8) {
	GB.InfoLogger.Println("ADD_r8r8_r8r8")

	//Perform Operation
	left := COMBINE_8BITS_TO_16BITS(leftUpper, leftLower)
	right := COMBINE_8BITS_TO_16BITS(rightUpper, rightLower)
	result := left + right

	//Set Flags
	SET_N_FLAG(GB, false)

	SET_H_FLAG_16(GB, left, right, true)

	SET_C_FLAG_16(GB, left, right, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '0', 'H': 'H', 'C': 'C'}
	return SPLIT_16BITS_TO_8BITS(result)
}

//ADD 16-bit register to 2 8-bit registers (as 16-bit)
func ADD_r8r8_r16(GB *GAMEBOY, leftUpper uint8, leftLower uint8, right uint16) (uint8, uint8) {
	GB.InfoLogger.Println("ADD_r8r8_r16")

	//Perform Operation
	left := COMBINE_8BITS_TO_16BITS(leftUpper, leftLower)
	result := left + right

	//Set Flags
	SET_N_FLAG(GB, false)

	SET_H_FLAG_16(GB, left, right, true)

	SET_C_FLAG_16(GB, left, right, true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '0', 'H': 'H', 'C': 'C'}
	return SPLIT_16BITS_TO_8BITS(result)
}

//----- Additions Carry -----

//
func ADC_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("ADC_r8_r8")

	//Perform Operation
	result := left + (right + GET_C_FLAG(GB))

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG_8(GB, left, right+GET_C_FLAG(GB), true)

	SET_C_FLAG_8(GB, left, right+GET_C_FLAG(GB), true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': 'C'}
	return result
}

//
func ADC_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("ADC_r8_n8")

	//Perform Operation
	result := left + (right + GET_C_FLAG(GB))

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG_8(GB, left, right+GET_C_FLAG(GB), true)

	SET_C_FLAG_8(GB, left, right+GET_C_FLAG(GB), true)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': 'H', 'C': 'C'}
	return result
}

//----- Subtractions ------

//
func SUB_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("SUB_r8_r8")

	//Perform Operation
	result := left - right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, right, false)

	SET_C_FLAG_8(GB, left, right, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
	return result
}

//
func SUB_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("SUB_r8_n8")

	//Perform Operation
	result := left - right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, right, false)

	SET_C_FLAG_8(GB, left, right, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
	return result
}

//----- Subtractions Carry -----

//
func SBC_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("SBC_r8_r8")

	//Perform Operation
	result := left - right - GET_C_FLAG(GB)

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, right+GET_C_FLAG(GB), false)

	SET_C_FLAG_8(GB, left, right+GET_C_FLAG(GB), false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
	return result
}

//
func SBC_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("SUB_r8_n8")

	//Perform Operation
	result := left - right - GET_C_FLAG(GB)

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, right+GET_C_FLAG(GB), false)

	SET_C_FLAG_8(GB, left, right+GET_C_FLAG(GB), false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
	return result
}

//----- Compare ------

//
func CP_r8_r8(GB *GAMEBOY, left uint8, right uint8) {
	GB.InfoLogger.Println("CP_r8_r8")

	//Perform Operation
	result := left - right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, right, false)

	SET_C_FLAG_8(GB, left, right, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
}

//
func CP_r8_n8(GB *GAMEBOY, left uint8, right uint8) {
	GB.InfoLogger.Println("CP_r8_n8")

	//Perform Operation
	result := left - right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, true)

	SET_H_FLAG_8(GB, left, right, false)

	SET_C_FLAG_8(GB, left, right, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
}

//---------- Memory ----------

//----- Loads ------

//LOAD next 16-bits into 2 8-bit registers
func LD_r8r8_n16(GB *GAMEBOY) (uint8, uint8) {
	GB.InfoLogger.Println("LD_r8r8_n16")

	//Perform Operation
	upper := GB.MMU.readByte(GB, GB.CPU._r.PC+2)
	lower := GB.MMU.readByte(GB, GB.CPU._r.PC+1)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 3, 3)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return upper, lower
}

//LOAD next 16-bits into 16-bit register
func LD_r16_n16(GB *GAMEBOY) uint16 {
	GB.InfoLogger.Println("LD_r8r8_n16")

	//Perform Operation
	upper := GB.MMU.readByte(GB, GB.CPU._r.PC+2)
	lower := GB.MMU.readByte(GB, GB.CPU._r.PC+1)
	combined := COMBINE_8BITS_TO_16BITS(upper, lower)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 3, 3)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return combined
}

//LOAD next 8-bits into 8-bit register
func LD_r8_n8(GB *GAMEBOY) uint8 {
	GB.InfoLogger.Println("LD_r8_n8")

	//Perform Operation
	result := GB.MMU.readByte(GB, GB.CPU._r.PC+1)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//LOAD 8-bit register into 8-bit register
func LD_r8_r8(GB *GAMEBOY, right uint8) uint8 {
	GB.InfoLogger.Println("LD_r8_r8")

	//Perform Operation
	result := right

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//LOAD 8-bit register into the memory address of 2 8-bit registers (as 16-bit)
func LD_addrr8r8_r8(GB *GAMEBOY, addrLeft uint8, addrRight uint8, right uint8) {
	GB.InfoLogger.Println("LD_addrr8r8_r8")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(addrLeft, addrRight)
	GB.MMU.writeByte(GB, addr, right)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

//LOAD 8-bit register into the memory address of 2 8-bit registers (as 16-bit) and increment
func LD_addrr8r8i_r8(GB *GAMEBOY, addrLeft uint8, addrRight uint8, right uint8) (uint8, uint8) {
	GB.InfoLogger.Println("LD_addrr8r8i_r8")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(addrLeft, addrRight)
	GB.MMU.writeByte(GB, addr, right)
	addr += 1
	newAddrLeft := uint8((addr & 0xFF00) >> 8)
	newAddrRight := uint8(addr & 0x00FF)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return newAddrLeft, newAddrRight
}

//LOAD 8-bit register into the memory address of 2 8-bit registers (as 16-bit) and decrement
func LD_addrr8r8d_r8(GB *GAMEBOY, addrLeft uint8, addrRight uint8, right uint8) (uint8, uint8) {
	GB.InfoLogger.Println("LD_addrr8r8d_r8")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(addrLeft, addrRight)
	GB.MMU.writeByte(GB, addr, right)
	addr -= 1
	newAddrLeft := uint8((addr & 0xFF00) >> 8)
	newAddrRight := uint8(addr & 0x00FF)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return newAddrLeft, newAddrRight
}

//LOAD value at memory address of 2 8-bit registers (as 16-bit) into 8-bit register
func LD_r8_addrr8r8(GB *GAMEBOY, addrLeft uint8, addrRight uint8) uint8 {
	GB.InfoLogger.Println("LD_r8_addrr8r8")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(addrLeft, addrRight)
	result := GB.MMU.readByte(GB, addr)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//LOAD value at memory address of 2 8-bit registers (as 16-bit) and increment, into 8-bit register
func LD_r8_addrr8r8i(GB *GAMEBOY, addrLeft uint8, addrRight uint8) (uint8, uint8, uint8) {
	GB.InfoLogger.Println("LD_r8_addrr8r8i")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(addrLeft, addrRight)
	result := GB.MMU.readByte(GB, addr)
	addr += 1
	newAddrLeft := uint8((addr & 0xFF00) >> 8)
	newAddrRight := uint8(addr & 0x00FF)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result, newAddrLeft, newAddrRight
}

//LOAD value at memory address of 2 8-bit registers (as 16-bit)and decrement, into 8-bit register
func LD_r8_addrr8r8d(GB *GAMEBOY, addrLeft uint8, addrRight uint8) (uint8, uint8, uint8) {
	GB.InfoLogger.Println("LD_r8_addrr8r8d")

	//Perform Operation
	addr := COMBINE_8BITS_TO_16BITS(addrLeft, addrRight)
	result := GB.MMU.readByte(GB, addr)
	addr -= 1
	newAddrLeft := uint8((addr & 0xFF00) >> 8)
	newAddrRight := uint8(addr & 0x00FF)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 2)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result, newAddrLeft, newAddrRight
}

//LOAD 8-bit register into memory address of next 16-bits
func LD_addra16_r8(GB *GAMEBOY, val uint8) {
	GB.InfoLogger.Println("LD_addra16_r8")

	//Perform Operation
	addr := uint16(GB.MMU.read2Bytes(GB, GB.CPU._r.PC+1))
	GB.InfoLogger.Println(addr)
	GB.InfoLogger.Println(val)
	GB.MMU.writeByte(GB, addr, val)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 3, 5)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

//LOAD 16-bit register into memory address of next 16-bits
func LD_addra16_r16(GB *GAMEBOY, val uint16) {
	GB.InfoLogger.Println("LD_addra16_r16")

	//Perform Operation
	addr := uint16(GB.MMU.read2Bytes(GB, GB.CPU._r.PC+1))
	GB.MMU.write2Bytes(GB, addr, val)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 3, 5)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

//LOAD value at memory address of next 16-bits into 8-bit register
func LD_r8_addra16(GB *GAMEBOY) uint8 {
	GB.InfoLogger.Println("LD_r8_addra16")

	//Perform Operation
	addr := uint16(GB.MMU.read2Bytes(GB, GB.CPU._r.PC+1))
	result := GB.MMU.readByte(GB, addr)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 3, 4)

	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
	return result
}

//---------- Bitwise Operations ----------

//----- AND -----

//
func AND_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("AND_r8_r8")

	//Perform Operation
	result := left & right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG(GB, true)

	SET_C_FLAG(GB, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '0'}
	return result
}

//
func AND_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("AND_r8_n8")

	//Perform Operation
	result := left & right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG(GB, true)

	SET_C_FLAG(GB, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '0'}
	return result
}

//----- XOR -----

//
func XOR_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("XOR_r8_r8")

	//Perform Operation
	result := left ^ right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG(GB, false)

	SET_C_FLAG(GB, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
	return result
}

//
func XOR_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("XOR_r8_n8")

	//Perform Operation
	result := left ^ right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG(GB, false)

	SET_C_FLAG(GB, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
	return result
}

//----- OR -----

//
func OR_r8_r8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("OR_r8_r8")

	//Perform Operation
	result := left | right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG(GB, false)

	SET_C_FLAG(GB, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 1, 1)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
	return result
}

//
func OR_r8_n8(GB *GAMEBOY, left uint8, right uint8) uint8 {
	GB.InfoLogger.Println("OR_r8_n8")

	//Perform Operation
	result := left | right

	//Set Flags
	SET_Z_FLAG_result8(GB, result)

	SET_N_FLAG(GB, false)

	SET_H_FLAG(GB, false)

	SET_C_FLAG(GB, false)

	//Set PC & Timings
	REG_CLOCK_TIMINGS(GB, 2, 2)

	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
	return result
}
