package main

import "math/bits"

//n8  = immediate 8-bit data
//n16 = immediate little-endian 16-bit data
//a8  = 8-bit unsigned data, which is added to $FF00 in certain instructions to create a 16-bit address in HRAM (High RAM)
//a16 = little-endian 16-bit address
//e8  = means 8-bit signed data

//opcode source: https://gbdev.io/gb-opcodes/optables/
//interactive reference table: https://meganesu.github.io/generate-gb-opcodes/
//reference: https://rgbds.gbdev.io/docs/v0.6.1/gbz80.7

//0b10000000 = Z : Zero
//0b01000000 = N : Subtraction
//0b00100000 = H : Half Carry
//0b00010000 = C : Carry
//Letter means set as per the operation dictates, 1 means set, 0 means unset, - means unchanged.

// 0x00 NOP
func NOP_0x00(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x00 NOP")
	REG_CLOCK_TIMINGS(GB, 1, 1)
}

// 0x01 LD BC_n16
func LD_0x01_BC_n16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x01 LD BC_n16")
	GB.CPU._r.B, GB.CPU._r.C = LD_r8r8_n16(GB)
}

// 0x02 LD addrBC_A
func LD_0x02_addrBC_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x02 LD addrBC_A")
	LD_addrr8r8_r8(GB, GB.CPU._r.B, GB.CPU._r.C, GB.CPU._r.A)
}

// 0x03 INC BC
func INC_0x03_BC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x03 INC BC")
	INC_r8r8(GB, GB.CPU._r.B, GB.CPU._r.C)
}

// 0x04 INC B
func INC_0x04_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x04 INC B")
	GB.CPU._r.B = INC_r8(GB, GB.CPU._r.B)
}

// 0x05 DEC B
func DEC_0x05_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x05 DEC B")
	DEC_r8(GB, GB.CPU._r.B)
}

// 0x06 LD B_n8
func LD_0x06_B_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x06 LD B_n8")
	GB.CPU._r.B = LD_r8_n8(GB)
}

// 0x07 RLCA
func RLCA_0x07(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x07 RLCA")
	bits.RotateLeft8(GB.CPU._r.A, 1)
	SET_Z_FLAG(GB, false)
	SET_N_FLAG(GB, false)
	SET_H_FLAG(GB, false)
	if GB.CPU._r.A&0x01 == 0x01 {
		SET_C_FLAG(GB, true)
	} else {
		SET_C_FLAG(GB, false)
	}
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '0', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x08 LD addra16_SP
func LD_0x08_addra16_SP(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x08 LD addra16_SP")
	LD_addra16_r16(GB, GB.CPU._r.SP)
}

// 0x09 ADD HL_BC
func ADD_0x09_HL_BC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x09 ADD HL_BC")
	GB.CPU._r.H, GB.CPU._r.L = ADD_r8r8_r8r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.B, GB.CPU._r.C)
}

// 0x0A LD A_addrBC
func LD_0x0A_A_addrBC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x0A LD A_addrBC")
	GB.CPU._r.A = LD_r8_addrr8r8(GB, GB.CPU._r.B, GB.CPU._r.C)
}

// 0x0B DEC BC
func DEC_0x0B_BC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x0B DEC BC")
	GB.CPU._r.B, GB.CPU._r.C = DEC_r8r8(GB, GB.CPU._r.B, GB.CPU._r.C)
}

// 0x0C INC C
func INC_0x0C_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x0C INC C")
	GB.CPU._r.C = INC_r8(GB, GB.CPU._r.C)
}

// 0x0D DEC C
func DEC_0x0D_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x0D DEC C")
	DEC_r8(GB, GB.CPU._r.C)
}

// 0x0E LD C_n8
func LD_0x0E_C_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x0E LD C_n8")
	GB.CPU._r.C = LD_r8_n8(GB)
}

// 0x0F RRCA
func RRCA_0x0F(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x0F RRCA")
	bits.RotateLeft8(GB.CPU._r.A, -1)
	SET_Z_FLAG(GB, false)
	SET_N_FLAG(GB, false)
	SET_H_FLAG(GB, false)
	if GB.CPU._r.A&0x80 == 0x80 {
		SET_C_FLAG(GB, true)
	} else {
		SET_C_FLAG(GB, false)
	}
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '0', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x10 STOP n8
func STOP_0x10_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x10 STOP n8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x11 LD DE_n16
func LD_0x11_DE_n16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x11 LD DE_n16")
	GB.CPU._r.D, GB.CPU._r.E = LD_r8r8_n16(GB)
}

// 0x12 LD addrDE_A
func LD_0x12_addrDE_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x12 LD addrDE_A")
	LD_addrr8r8_r8(GB, GB.CPU._r.D, GB.CPU._r.E, GB.CPU._r.A)
}

// 0x13 INC DE
func INC_0x13_DE(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x13 INC DE")
	INC_r8r8(GB, GB.CPU._r.D, GB.CPU._r.E)
}

// 0x14 INC D
func INC_0x14_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x14 INC D")
	GB.CPU._r.D = INC_r8(GB, GB.CPU._r.D)
}

// 0x15 DEC D
func DEC_0x15_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x15 DEC D")
	DEC_r8(GB, GB.CPU._r.D)
}

// 0x16 LD D_n8
func LD_0x16_D_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x16 LD D_n8")
	GB.CPU._r.D = LD_r8_n8(GB)
}

// 0x17 RLA
func RLA_0x17(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x17 RLA")
	bits.RotateLeft8(GB.CPU._r.A, 1)
	SET_Z_FLAG(GB, false)
	SET_N_FLAG(GB, false)
	SET_H_FLAG(GB, false)
	cFlag := GB.CPU._r.F.C
	if GB.CPU._r.A&0x01 == 0x01 {
		SET_C_FLAG(GB, true)
	} else {
		SET_C_FLAG(GB, false)
	}
	GB.CPU._r.A = (GB.CPU._r.A & 0xFE) | (cFlag & 0x01)
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '0', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x18 JR e8
func JR_0x18_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x18 JR e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x19 ADD HL_DE
func ADD_0x19_HL_DE(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x19 ADD HL_DE")
	GB.CPU._r.H, GB.CPU._r.L = ADD_r8r8_r8r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.D, GB.CPU._r.E)
}

// 0x1A LD A_addrDE
func LD_0x1A_A_addrDE(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x1A LD A_addrDE")
	GB.CPU._r.A = LD_r8_addrr8r8(GB, GB.CPU._r.D, GB.CPU._r.E)
}

// 0x1B DEC DE
func DEC_0x1B_DE(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x1B DEC DE")
	GB.CPU._r.D, GB.CPU._r.E = DEC_r8r8(GB, GB.CPU._r.D, GB.CPU._r.E)
}

// 0x1C INC E
func INC_0x1C_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x1C INC E")
	GB.CPU._r.E = INC_r8(GB, GB.CPU._r.E)
}

// 0x1D DEC E
func DEC_0x1D_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x1D DEC E")
	DEC_r8(GB, GB.CPU._r.E)
}

// 0x1E LD E_n8
func LD_0x1E_E_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x1E LD E_n8")
	GB.CPU._r.E = LD_r8_n8(GB)
}

// 0x1F RRA
func RRA_0x1F(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x1F RRA ")
	bits.RotateLeft8(GB.CPU._r.A, -1)
	SET_Z_FLAG(GB, false)
	SET_N_FLAG(GB, false)
	SET_H_FLAG(GB, false)
	cFlag := GB.CPU._r.F.C
	if GB.CPU._r.A&0x80 == 0x80 {
		SET_C_FLAG(GB, true)
	} else {
		SET_C_FLAG(GB, false)
	}
	GB.CPU._r.A = (GB.CPU._r.A & 0x7F) | ((cFlag * 0x01) << 7)
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '0', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x20 JR NZ_e8
func JR_0x20_NZ_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x20 JR NZ_e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3 //[3, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x21 LD HL_n16
func LD_0x21_HL_n16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x21 LD HL_n16")
	GB.CPU._r.H, GB.CPU._r.L = LD_r8r8_n16(GB)
}

// 0x22 LD addrHL_A
func LD_0x22_addrHLi_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x22 LD addrHLi_A")
	GB.CPU._r.H, GB.CPU._r.L = LD_addrr8r8i_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.A)
}

// 0x23 INC HL
func INC_0x23_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x23 INC HL")
	INC_r8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x24 INC H
func INC_0x24_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x24 INC H")
	GB.CPU._r.H = INC_r8(GB, GB.CPU._r.H)
}

// 0x25 DEC H
func DEC_0x25_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x25 DEC H")
	DEC_r8(GB, GB.CPU._r.H)
}

// 0x26 LD H_n8
func LD_0x26_H_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x26 LD H_n8")
	GB.CPU._r.H = LD_r8_n8(GB)
}

// 0x27 DAA
func DAA_0x27(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x27 DAA ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '-', 'H': '0', 'C': 'C'}
}

// 0x28 JR Z_e8
func JR_0x28_Z_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x28 JR Z_e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3 //[3, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x29 ADD HL_HL
func ADD_0x29_HL_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x29 ADD HL_HL")
	GB.CPU._r.H, GB.CPU._r.L = ADD_r8r8_r8r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x2A LD A_addrHL
func LD_0x2A_A_addrHLi(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x2A LD A_addrHL")
	GB.CPU._r.A, GB.CPU._r.H, GB.CPU._r.L = LD_r8_addrr8r8i(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x2B DEC HL
func DEC_0x2B_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x2B DEC HL")
	GB.CPU._r.H, GB.CPU._r.L = DEC_r8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x2C INC L
func INC_0x2C_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x2C INC L")
	GB.CPU._r.L = INC_r8(GB, GB.CPU._r.L)
}

// 0x2D DEC L
func DEC_0x2D_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x2D DEC L")
	DEC_r8(GB, GB.CPU._r.L)
}

// 0x2E LD L_n8
func LD_0x2E_L_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x2E LD L_n8")
	GB.CPU._r.L = LD_r8_n8(GB)
}

// 0x2F CPL
func CPL_0x2F(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x2F CPL ")
	GB.CPU._r.A = ^GB.CPU._r.A
	SET_N_FLAG(GB, true)
	SET_H_FLAG(GB, true)
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '-', 'N': '1', 'H': '1', 'C': '-'}
}

// 0x30 JR NC_e8
func JR_0x30_NC_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x30 JR NC_e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3 //[3, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x31 LD SP_n16
func LD_0x31_SP_n16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x31 LD SP_n16")
	GB.CPU._r.SP = LD_r16_n16(GB)
}

// 0x32 LD addrHLd_A
func LD_0x32_addrHLd_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x32 LD addrHLd_A")
	GB.CPU._r.H, GB.CPU._r.L = LD_addrr8r8d_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.A)
}

// 0x33 INC SP
func INC_0x33_SP(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x33 INC SP")
	GB.CPU._r.SP = INC_r16(GB, GB.CPU._r.SP)
}

// 0x34 INC addrHL
func INC_0x34_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x34 INC addrHL")
	INC_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x35 DEC addrHL
func DEC_0x35_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x35 DEC addrHL")
	DEC_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x36 LD addrHL_n8
func LD_0x36_addrHL_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x36 LD addrHL_n8")
	LD_addrr8r8_r8(GB, GB.CPU._r.B, GB.CPU._r.C, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0x37 SCF
func SCF_0x37(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x37 SCF ")
	SET_N_FLAG(GB, false)
	SET_H_FLAG(GB, false)
	SET_C_FLAG(GB, true)
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '-', 'N': '0', 'H': '0', 'C': '1'}
}

// 0x38 JR C_e8
func JR_0x38_C_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x38 JR C_e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3 //[3, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x39 ADD HL_SP
func ADD_0x39_HL_SP(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x39 ADD HL_SP")
	GB.CPU._r.H, GB.CPU._r.L = ADD_r8r8_r16(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.SP)
}

// 0x3A LD A_addrHL
func LD_0x3A_A_addrHLd(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x3A LD A_addrHLd")
	GB.CPU._r.A, GB.CPU._r.H, GB.CPU._r.L = LD_r8_addrr8r8d(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x3B DEC SP
func DEC_0x3B_SP(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x3B DEC SP")
	GB.CPU._r.SP = DEC_r16(GB, GB.CPU._r.SP)
}

// 0x3C INC A
func INC_0x3C_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x3C INC A")
	GB.CPU._r.A = INC_r8(GB, GB.CPU._r.A)
}

// 0x3D DEC A
func DEC_0x3D_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x3D DEC A")
	DEC_r8(GB, GB.CPU._r.A)
}

// 0x3E LD A_n8
func LD_0x3E_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x3E LD A_n8")
	GB.CPU._r.A = LD_r8_n8(GB)
}

// 0x3F CCF
func CCF_0x3F(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x3F CCF ")
	SET_N_FLAG(GB, false)
	SET_H_FLAG(GB, false)
	GB.CPU._r.F.C = ^GB.CPU._r.F.C
	REG_CLOCK_TIMINGS(GB, 1, 1)
	//FLAGS AFFECTED : {'Z': '-', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x40 LD B_B
func LD_0x40_B_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x40 LD B_B")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x41 LD B_C
func LD_0x41_B_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x41 LD B_C")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x42 LD B_D
func LD_0x42_B_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x42 LD B_D")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x43 LD B_E
func LD_0x43_B_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x43 LD B_E")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x44 LD B_H
func LD_0x44_B_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x44 LD B_H")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x45 LD B_L
func LD_0x45_B_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x45 LD B_L")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x46 LD B_addrHL
func LD_0x46_B_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x46 LD B_addrHL")
	GB.CPU._r.B = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x47 LD B_A
func LD_0x47_B_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x47 LD B_A")
	GB.CPU._r.B = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x48 LD C_B
func LD_0x48_C_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x48 LD C_B")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x49 LD C_C
func LD_0x49_C_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x49 LD C_C")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x4A LD C_D
func LD_0x4A_C_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x4A LD C_D")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x4B LD C_E
func LD_0x4B_C_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x4B LD C_E")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x4C LD C_H
func LD_0x4C_C_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x4C LD C_H")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x4D LD C_L
func LD_0x4D_C_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x4D LD C_L")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x4E LD C_addrHL
func LD_0x4E_C_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x4E LD C_addrHL")
	GB.CPU._r.C = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x4F LD C_A
func LD_0x4F_C_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x4F LD C_A")
	GB.CPU._r.C = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x50 LD D_B
func LD_0x50_D_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x50 LD D_B")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x51 LD D_C
func LD_0x51_D_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x51 LD D_C")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x52 LD D_D
func LD_0x52_D_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x52 LD D_D")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x53 LD D_E
func LD_0x53_D_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x53 LD D_E")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x54 LD D_H
func LD_0x54_D_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x54 LD D_H")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x55 LD D_L
func LD_0x55_D_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x55 LD D_L")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x56 LD D_addrHL
func LD_0x56_D_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x56 LD D_addrHL")
	GB.CPU._r.D = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x57 LD D_A
func LD_0x57_D_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x57 LD D_A")
	GB.CPU._r.D = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x58 LD E_B
func LD_0x58_E_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x58 LD E_B")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x59 LD E_C
func LD_0x59_E_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x59 LD E_C")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x5A LD E_D
func LD_0x5A_E_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x5A LD E_D")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x5B LD E_E
func LD_0x5B_E_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x5B LD E_E")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x5C LD E_H
func LD_0x5C_E_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x5C LD E_H")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x5D LD E_L
func LD_0x5D_E_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x5D LD E_L")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x5E LD E_addrHL
func LD_0x5E_E_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x5E LD E_addrHL")
	GB.CPU._r.E = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x5F LD E_A
func LD_0x5F_E_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x5F LD E_A")
	GB.CPU._r.E = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x60 LD H_B
func LD_0x60_H_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x60 LD H_B")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x61 LD H_C
func LD_0x61_H_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x61 LD H_C")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x62 LD H_D
func LD_0x62_H_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x62 LD H_D")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x63 LD H_E
func LD_0x63_H_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x63 LD H_E")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x64 LD H_H
func LD_0x64_H_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x64 LD H_H")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x65 LD H_L
func LD_0x65_H_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x65 LD H_L")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x66 LD H_addrHL
func LD_0x66_H_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x66 LD H_addrHL")
	GB.CPU._r.H = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x67 LD H_A
func LD_0x67_H_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x67 LD H_A")
	GB.CPU._r.H = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x68 LD L_B
func LD_0x68_L_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x68 LD L_B")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x69 LD L_C
func LD_0x69_L_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x69 LD L_C")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x6A LD L_D
func LD_0x6A_L_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x6A LD L_D")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x6B LD L_E
func LD_0x6B_L_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x6B LD L_E")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x6C LD L_H
func LD_0x6C_L_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x6C LD L_H")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x6D LD L_L
func LD_0x6D_L_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x6D LD L_L")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x6E LD L_addrHL
func LD_0x6E_L_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x6E LD L_addrHL")
	GB.CPU._r.L = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x6F LD L_A
func LD_0x6F_L_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x6F LD L_A")
	GB.CPU._r.L = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x70 LD addrHL_B
func LD_0x70_addrHL_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x70 LD addrHL_B")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.B)
}

// 0x71 LD addrHL_C
func LD_0x71_addrHL_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x71 LD addrHL_C")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.C)
}

// 0x72 LD addrHL_D
func LD_0x72_addrHL_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x72 LD addrHL_D")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.D)
}

// 0x73 LD addrHL_E
func LD_0x73_addrHL_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x73 LD addrHL_E")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.E)
}

// 0x74 LD addrHL_H
func LD_0x74_addrHL_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x74 LD addrHL_H")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.H)
}

// 0x75 LD addrHL_L
func LD_0x75_addrHL_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x75 LD addrHL_L")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.L)
}

// 0x76 HALT
func HALT_0x76(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x76 HALT ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x77 LD addrHL_A
func LD_0x77_addrHL_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x77 LD addrHL_A")
	LD_addrr8r8_r8(GB, GB.CPU._r.H, GB.CPU._r.L, GB.CPU._r.A)
}

// 0x78 LD A_B
func LD_0x78_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x78 LD A_B")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.B)
}

// 0x79 LD A_C
func LD_0x79_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x79 LD A_C")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.C)
}

// 0x7A LD A_D
func LD_0x7A_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x7A LD A_D")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.D)
}

// 0x7B LD A_E
func LD_0x7B_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x7B LD A_E")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.E)
}

// 0x7C LD A_H
func LD_0x7C_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x7C LD A_H")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.H)
}

// 0x7D LD A_L
func LD_0x7D_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x7D LD A_L")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.L)
}

// 0x7E LD A_addrHL
func LD_0x7E_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x7E LD A_addrHL")
	GB.CPU._r.A = LD_r8_addrr8r8(GB, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x7F LD A_A
func LD_0x7F_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x7F LD A_A")
	GB.CPU._r.A = LD_r8_r8(GB, GB.CPU._r.A)
}

// 0x80 ADD A_B
func ADD_0x80_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x80 ADD A_B")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0x81 ADD A_C
func ADD_0x81_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x81 ADD A_C")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0x82 ADD A_D
func ADD_0x82_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x82 ADD A_D")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0x83 ADD A_E
func ADD_0x83_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x83 ADD A_E")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0x84 ADD A_H
func ADD_0x84_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x84 ADD A_H")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0x85 ADD A_L
func ADD_0x85_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x85 ADD A_L")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0x86 ADD A_addrHL
func ADD_0x86_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x86 ADD A_addrHL")
	GB.CPU._r.A = ADD_r8_addrr8r8(GB, GB.CPU._r.A, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x87 ADD A_A
func ADD_0x87_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x87 ADD A_A")
	GB.CPU._r.A = ADD_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0x88 ADC A_B
func ADC_0x88_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x88 ADC A_B")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0x89 ADC A_C
func ADC_0x89_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x89 ADC A_C")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0x8A ADC A_D
func ADC_0x8A_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x8A ADC A_D")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0x8B ADC A_E
func ADC_0x8B_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x8B ADC A_E")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0x8C ADC A_H
func ADC_0x8C_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x8C ADC A_H")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0x8D ADC A_L
func ADC_0x8D_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x8D ADC A_L")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0x8E ADC A_addrHL
func ADC_0x8E_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x8E ADC A_addrHL")
	GB.CPU._r.A = ADC_r8_addrr8r8(GB, GB.CPU._r.A, GB.CPU._r.H, GB.CPU._r.L)
}

// 0x8F ADC A_A
func ADC_0x8F_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x8F ADC A_A")
	GB.CPU._r.A = ADC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0x90 SUB A_B
func SUB_0x90_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x90 SUB A_B")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0x91 SUB A_C
func SUB_0x91_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x91 SUB A_C")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0x92 SUB A_D
func SUB_0x92_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x92 SUB A_D")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0x93 SUB A_E
func SUB_0x93_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x93 SUB A_E")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0x94 SUB A_H
func SUB_0x94_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x94 SUB A_H")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0x95 SUB A_L
func SUB_0x95_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x95 SUB A_L")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0x96 SUB A_addrHL
func SUB_0x96_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x96 SUB A_addrHL")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
}

// 0x97 SUB A_A
func SUB_0x97_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x97 SUB A_A")
	GB.CPU._r.A = SUB_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0x98 SBC A_B
func SBC_0x98_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x98 SBC A_B")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0x99 SBC A_C
func SBC_0x99_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x99 SBC A_C")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0x9A SBC A_D
func SBC_0x9A_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x9A SBC A_D")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0x9B SBC A_E
func SBC_0x9B_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x9B SBC A_E")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0x9C SBC A_H
func SBC_0x9C_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x9C SBC A_H")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0x9D SBC A_L
func SBC_0x9D_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x9D SBC A_L")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0x9E SBC A_addrHL
func SBC_0x9E_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x9E SBC A_addrHL")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
}

// 0x9F SBC A_A
func SBC_0x9F_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0x9F SBC A_A")
	GB.CPU._r.A = SBC_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0xA0 AND A_B
func AND_0xA0_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA0 AND A_B")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0xA1 AND A_C
func AND_0xA1_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA1 AND A_C")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0xA2 AND A_D
func AND_0xA2_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA2 AND A_D")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0xA3 AND A_E
func AND_0xA3_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA3 AND A_E")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0xA4 AND A_H
func AND_0xA4_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA4 AND A_H")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0xA5 AND A_L
func AND_0xA5_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA5 AND A_L")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0xA6 AND A_addrHL
func AND_0xA6_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA6 AND A_addrHL")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '0'}
}

// 0xA7 AND A_A
func AND_0xA7_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA7 AND A_A")
	GB.CPU._r.A = AND_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0xA8 XOR A_B
func XOR_0xA8_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA8 XOR A_B")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0xA9 XOR A_C
func XOR_0xA9_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xA9 XOR A_C")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0xAA XOR A_D
func XOR_0xAA_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xAA XOR A_D")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0xAB XOR A_E
func XOR_0xAB_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xAB XOR A_E")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0xAC XOR A_H
func XOR_0xAC_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xAC XOR A_H")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0xAD XOR A_L
func XOR_0xAD_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xAD XOR A_L")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0xAE XOR A_addrHL
func XOR_0xAE_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xAE XOR A_addrHL")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0xAF XOR A_A
func XOR_0xAF_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xAF XOR A_A")
	GB.CPU._r.A = XOR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0xB0 OR A_B
func OR_0xB0_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB0 OR A_B")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0xB1 OR A_C
func OR_0xB1_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB1 OR A_C")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0xB2 OR A_D
func OR_0xB2_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB2 OR A_D")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0xB3 OR A_E
func OR_0xB3_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB3 OR A_E")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0xB4 OR A_H
func OR_0xB4_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB4 OR A_H")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0xB5 OR A_L
func OR_0xB5_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB5 OR A_L")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0xB6 OR A_addrHL
func OR_0xB6_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB6 OR A_addrHL")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0xB7 OR A_A
func OR_0xB7_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB7 OR A_A")
	GB.CPU._r.A = OR_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0xB8 CP A_B
func CP_0xB8_A_B(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB8 CP A_B")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.B)
}

// 0xB9 CP A_C
func CP_0xB9_A_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xB9 CP A_C")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.C)
}

// 0xBA CP A_D
func CP_0xBA_A_D(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xBA CP A_D")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.D)
}

// 0xBB CP A_E
func CP_0xBB_A_E(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xBB CP A_E")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.E)
}

// 0xBC CP A_H
func CP_0xBC_A_H(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xBC CP A_H")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.H)
}

// 0xBD CP A_L
func CP_0xBD_A_L(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xBD CP A_L")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.L)
}

// 0xBE CP A_addrHL
func CP_0xBE_A_addrHL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xBE CP A_addrHL")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '1', 'H': 'H', 'C': 'C'}
}

// 0xBF CP A_A
func CP_0xBF_A_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xBF CP A_A")
	CP_r8_r8(GB, GB.CPU._r.A, GB.CPU._r.A)
}

// 0xC0 RET NZ
func RET_0xC0_NZ(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC0 RET NZ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 5 //[5, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC1 POP BC
func POP_0xC1_BC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC1 POP BC")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC2 JP NZ_a16
func JP_0xC2_NZ_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC2 JP NZ_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 4 //[4, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC3 JP a16
func JP_0xC3_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC3 JP a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC4 CALL NZ_a16
func CALL_0xC4_NZ_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC4 CALL NZ_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 6 //[6, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC5 PUSH BC
func PUSH_0xC5_BC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC5 PUSH BC")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC6 ADD A_n8
func ADD_0xC6_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC6 ADD A_n8")
	ADD_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xC7 RST $00
func RST_0xC7_dollar00(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC7 RST $00")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC8 RET Z
func RET_0xC8_Z(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC8 RET Z")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 5 //[5, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC9 RET
func RET_0xC9(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xC9 RET ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCA JP Z_a16
func JP_0xCA_Z_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xCA JP Z_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 4 //[4, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCB PREFIX
func PREFIX_0xCB(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xCB PREFIX ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCC CALL Z_a16
func CALL_0xCC_Z_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xCC CALL Z_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 6 //[6, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCD CALL a16
func CALL_0xCD_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xCD CALL a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 6
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCE ADC A_n8
func ADC_0xCE_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xCE ADC A_n8")
	GB.CPU._r.A = ADC_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xCF RST $08
func RST_0xCF_dollar08(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xCF RST $08")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD0 RET NC
func RET_0xD0_NC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD0 RET NC")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 5 //[5, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD1 POP DE
func POP_0xD1_DE(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD1 POP DE")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD2 JP NC_a16
func JP_0xD2_NC_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD2 JP NC_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 4 //[4, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD3 ILLEGAL_D3
func ILLEGAL_D3_0xD3(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD3 ILLEGAL_D3 ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xD4 CALL NC_a16
func CALL_0xD4_NC_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD4 CALL NC_a16")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xD5 PUSH DE
func PUSH_0xD5_DE(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD5 PUSH DE")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD6 SUB A_n8
func SUB_0xD6_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD6 SUB A_n8")
	GB.CPU._r.A = SUB_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xD7 RST $10
func RST_0xD7_dollar10(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD7 RST $10")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD8 RET C
func RET_0xD8_C(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD8 RET C")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 5 //[5, 2]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD9 RETI
func RETI_0xD9(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xD9 RETI ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDA JP C_a16
func JP_0xDA_C_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xDA JP C_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 4 //[4, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDB ILLEGAL_DB
func ILLEGAL_DB_0xDB(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xDB ILLEGAL_DB ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xDC CALL C_a16
func CALL_0xDC_C_a16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xDC CALL C_a16")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 3
	GB.CPU._r.M = 6 //[6, 3]
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDD ILLEGAL_DD
func ILLEGAL_DD_0xDD(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xDD ILLEGAL_DD ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xDE SBC A_n8
func SBC_0xDE_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xDE SBC A_n8")
	GB.CPU._r.A = SBC_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xDF RST $18
func RST_0xDF_dollar18(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xDF RST $18")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE0 LDH addra8_A
func LDH_0xE0_addra8_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE0 LDH addra8_A")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE1 POP HL
func POP_0xE1_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE1 POP HL")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE2 LDH addrC_A
func LDH_0xE2_addrC_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE2 LDH addrC_A")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE3 ILLEGAL_E3
func ILLEGAL_E3_0xE3(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE3 ILLEGAL_E3 ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xE4 ILLEGAL_E4
func ILLEGAL_E4_0xE4(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE4 ILLEGAL_E4 ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xE5 PUSH HL
func PUSH_0xE5_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE5 PUSH HL")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE6 AND A_n8
func AND_0xE6_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE6 AND A_n8")
	GB.CPU._r.A = AND_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xE7 RST $20
func RST_0xE7_dollar20(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE7 RST $20")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE8 ADD SP_e8
func ADD_0xE8_SP_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE8 ADD SP_e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '0', 'N': '0', 'H': 'H', 'C': 'C'}
}

// 0xE9 JP HL
func JP_0xE9_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xE9 JP HL")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xEA LD addra16_A
func LD_0xEA_addra16_A(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xEA LD addra16_A")
	LD_addra16_r8(GB, GB.CPU._r.A)
}

// 0xEB ILLEGAL_EB
func ILLEGAL_EB_0xEB(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xEB ILLEGAL_EB ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xEC ILLEGAL_EC
func ILLEGAL_EC_0xEC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xEC ILLEGAL_EC ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xED ILLEGAL_ED
func ILLEGAL_ED_0xED(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xED ILLEGAL_ED ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xEE XOR A_n8
func XOR_0xEE_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xEE XOR A_n8")
	GB.CPU._r.A = XOR_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xEF RST $28
func RST_0xEF_dollar28(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xEF RST $28")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF0 LDH A_addra8
func LDH_0xF0_A_addra8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF0 LDH A_addra8")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF1 POP AF
func POP_0xF1_AF(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF1 POP AF")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': 'N', 'H': 'H', 'C': 'C'}
}

// 0xF2 LDH A_addrC
func LDH_0xF2_A_addrC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF2 LDH A_addrC")
	//NEEDS CODE - ADDR
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF3 DI
func DI_0xF3(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF3 DI ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF4 ILLEGAL_F4
func ILLEGAL_F4_0xF4(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF4 ILLEGAL_F4 ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xF5 PUSH AF
func PUSH_0xF5_AF(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF5 PUSH AF")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF6 OR A_n8
func OR_0xF6_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF6 OR A_n8")
	GB.CPU._r.A = OR_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xF7 RST $30
func RST_0xF7_dollar30(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF7 RST $30")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF8 LD HL_SP_e8
func LD_0xF8_HL_SP_e8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF8 LD HL_SP_e8")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '0', 'N': '0', 'H': 'H', 'C': 'C'}
}

// 0xF9 LD SP_HL
func LD_0xF9_SP_HL(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xF9 LD SP_HL")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFA LD A_addra16
func LD_0xFA_A_addra16(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xFA LD A_addra16")
	GB.CPU._r.A = LD_r8_addra16(GB)
}

// 0xFB EI
func EI_0xFB(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xFB EI ")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 1
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFC ILLEGAL_FC
func ILLEGAL_FC_0xFC(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xFC ILLEGAL_FC ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xFD ILLEGAL_FD
func ILLEGAL_FD_0xFD(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xFD ILLEGAL_FD ")
	GB.ErrorLogger.Println("HIT ILLEGAL INSTRUCTION")
}

// 0xFE CP A_n8
func CP_0xFE_A_n8(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xFE CP A_n8")
	CP_r8_n8(GB, GB.CPU._r.A, GB.MMU.readByte(GB, GB.CPU._r.PC+1))
}

// 0xFF RST $38
func RST_0xFF_dollar38(GB *GAMEBOY) {
	GB.InfoLogger.Println("0xFF RST $38")
	//NEEDS CODE
	GB.CPU._r.PC += 1
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}
