package main

import "log"

//n8  = immediate 8-bit data
//n16 = immediate little-endian 16-bit data
//a8  = 8-bit unsigned data, which is added to $FF00 in certain instructions to create a 16-bit address in HRAM (High RAM)
//a16 = little-endian 16-bit address
//e8  = means 8-bit signed data

//opcode source: https://gbdev.io/gb-opcodes/optables/
//interactive reference table: https://meganesu.github.io/generate-gb-opcodes/

//0b10000000 = Z : Zero
//0b01000000 = N : Subtraction
//0b00100000 = H : Half Carry
//0b00010000 = C : Carry
//Letter means set as per the operation dictates, 1 means set, 0 means unset, - means unchanged.

// 0x00 RLC B
func CB_RLC_0x00_B(GB *GAMEBOY) {
	log.Println("0x00 RLC B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x01 RLC C
func CB_RLC_0x01_C(GB *GAMEBOY) {
	log.Println("0x01 RLC C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x02 RLC D
func CB_RLC_0x02_D(GB *GAMEBOY) {
	log.Println("0x02 RLC D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x03 RLC E
func CB_RLC_0x03_E(GB *GAMEBOY) {
	log.Println("0x03 RLC E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x04 RLC H
func CB_RLC_0x04_H(GB *GAMEBOY) {
	log.Println("0x04 RLC H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x05 RLC L
func CB_RLC_0x05_L(GB *GAMEBOY) {
	log.Println("0x05 RLC L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x06 RLC addrHL
func CB_RLC_0x06_addrHL(GB *GAMEBOY) {
	log.Println("0x06 RLC addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x07 RLC A
func CB_RLC_0x07_A(GB *GAMEBOY) {
	log.Println("0x07 RLC A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x08 RRC B
func CB_RRC_0x08_B(GB *GAMEBOY) {
	log.Println("0x08 RRC B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x09 RRC C
func CB_RRC_0x09_C(GB *GAMEBOY) {
	log.Println("0x09 RRC C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x0A RRC D
func CB_RRC_0x0A_D(GB *GAMEBOY) {
	log.Println("0x0A RRC D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x0B RRC E
func CB_RRC_0x0B_E(GB *GAMEBOY) {
	log.Println("0x0B RRC E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x0C RRC H
func CB_RRC_0x0C_H(GB *GAMEBOY) {
	log.Println("0x0C RRC H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x0D RRC L
func CB_RRC_0x0D_L(GB *GAMEBOY) {
	log.Println("0x0D RRC L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x0E RRC addrHL
func CB_RRC_0x0E_addrHL(GB *GAMEBOY) {
	log.Println("0x0E RRC addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x0F RRC A
func CB_RRC_0x0F_A(GB *GAMEBOY) {
	log.Println("0x0F RRC A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x10 RL B
func CB_RL_0x10_B(GB *GAMEBOY) {
	log.Println("0x10 RL B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x11 RL C
func CB_RL_0x11_C(GB *GAMEBOY) {
	log.Println("0x11 RL C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x12 RL D
func CB_RL_0x12_D(GB *GAMEBOY) {
	log.Println("0x12 RL D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x13 RL E
func CB_RL_0x13_E(GB *GAMEBOY) {
	log.Println("0x13 RL E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x14 RL H
func CB_RL_0x14_H(GB *GAMEBOY) {
	log.Println("0x14 RL H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x15 RL L
func CB_RL_0x15_L(GB *GAMEBOY) {
	log.Println("0x15 RL L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x16 RL addrHL
func CB_RL_0x16_addrHL(GB *GAMEBOY) {
	log.Println("0x16 RL addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x17 RL A
func CB_RL_0x17_A(GB *GAMEBOY) {
	log.Println("0x17 RL A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x18 RR B
func CB_RR_0x18_B(GB *GAMEBOY) {
	log.Println("0x18 RR B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x19 RR C
func CB_RR_0x19_C(GB *GAMEBOY) {
	log.Println("0x19 RR C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x1A RR D
func CB_RR_0x1A_D(GB *GAMEBOY) {
	log.Println("0x1A RR D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x1B RR E
func CB_RR_0x1B_E(GB *GAMEBOY) {
	log.Println("0x1B RR E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x1C RR H
func CB_RR_0x1C_H(GB *GAMEBOY) {
	log.Println("0x1C RR H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x1D RR L
func CB_RR_0x1D_L(GB *GAMEBOY) {
	log.Println("0x1D RR L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x1E RR addrHL
func CB_RR_0x1E_addrHL(GB *GAMEBOY) {
	log.Println("0x1E RR addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x1F RR A
func CB_RR_0x1F_A(GB *GAMEBOY) {
	log.Println("0x1F RR A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x20 SLA B
func CB_SLA_0x20_B(GB *GAMEBOY) {
	log.Println("0x20 SLA B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x21 SLA C
func CB_SLA_0x21_C(GB *GAMEBOY) {
	log.Println("0x21 SLA C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x22 SLA D
func CB_SLA_0x22_D(GB *GAMEBOY) {
	log.Println("0x22 SLA D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x23 SLA E
func CB_SLA_0x23_E(GB *GAMEBOY) {
	log.Println("0x23 SLA E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x24 SLA H
func CB_SLA_0x24_H(GB *GAMEBOY) {
	log.Println("0x24 SLA H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x25 SLA L
func CB_SLA_0x25_L(GB *GAMEBOY) {
	log.Println("0x25 SLA L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x26 SLA addrHL
func CB_SLA_0x26_addrHL(GB *GAMEBOY) {
	log.Println("0x26 SLA addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x27 SLA A
func CB_SLA_0x27_A(GB *GAMEBOY) {
	log.Println("0x27 SLA A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x28 SRA B
func CB_SRA_0x28_B(GB *GAMEBOY) {
	log.Println("0x28 SRA B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x29 SRA C
func CB_SRA_0x29_C(GB *GAMEBOY) {
	log.Println("0x29 SRA C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x2A SRA D
func CB_SRA_0x2A_D(GB *GAMEBOY) {
	log.Println("0x2A SRA D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x2B SRA E
func CB_SRA_0x2B_E(GB *GAMEBOY) {
	log.Println("0x2B SRA E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x2C SRA H
func CB_SRA_0x2C_H(GB *GAMEBOY) {
	log.Println("0x2C SRA H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x2D SRA L
func CB_SRA_0x2D_L(GB *GAMEBOY) {
	log.Println("0x2D SRA L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x2E SRA addrHL
func CB_SRA_0x2E_addrHL(GB *GAMEBOY) {
	log.Println("0x2E SRA addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x2F SRA A
func CB_SRA_0x2F_A(GB *GAMEBOY) {
	log.Println("0x2F SRA A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x30 SWAP B
func CB_SWAP_0x30_B(GB *GAMEBOY) {
	log.Println("0x30 SWAP B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x31 SWAP C
func CB_SWAP_0x31_C(GB *GAMEBOY) {
	log.Println("0x31 SWAP C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x32 SWAP D
func CB_SWAP_0x32_D(GB *GAMEBOY) {
	log.Println("0x32 SWAP D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x33 SWAP E
func CB_SWAP_0x33_E(GB *GAMEBOY) {
	log.Println("0x33 SWAP E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x34 SWAP H
func CB_SWAP_0x34_H(GB *GAMEBOY) {
	log.Println("0x34 SWAP H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x35 SWAP L
func CB_SWAP_0x35_L(GB *GAMEBOY) {
	log.Println("0x35 SWAP L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x36 SWAP addrHL
func CB_SWAP_0x36_addrHL(GB *GAMEBOY) {
	log.Println("0x36 SWAP addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x37 SWAP A
func CB_SWAP_0x37_A(GB *GAMEBOY) {
	log.Println("0x37 SWAP A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': '0'}
}

// 0x38 SRL B
func CB_SRL_0x38_B(GB *GAMEBOY) {
	log.Println("0x38 SRL B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x39 SRL C
func CB_SRL_0x39_C(GB *GAMEBOY) {
	log.Println("0x39 SRL C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x3A SRL D
func CB_SRL_0x3A_D(GB *GAMEBOY) {
	log.Println("0x3A SRL D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x3B SRL E
func CB_SRL_0x3B_E(GB *GAMEBOY) {
	log.Println("0x3B SRL E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x3C SRL H
func CB_SRL_0x3C_H(GB *GAMEBOY) {
	log.Println("0x3C SRL H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x3D SRL L
func CB_SRL_0x3D_L(GB *GAMEBOY) {
	log.Println("0x3D SRL L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x3E SRL addrHL
func CB_SRL_0x3E_addrHL(GB *GAMEBOY) {
	log.Println("0x3E SRL addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x3F SRL A
func CB_SRL_0x3F_A(GB *GAMEBOY) {
	log.Println("0x3F SRL A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '0', 'C': 'C'}
}

// 0x40 BIT 0_B
func CB_BIT_0x40_0_B(GB *GAMEBOY) {
	log.Println("0x40 BIT 0_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x41 BIT 0_C
func CB_BIT_0x41_0_C(GB *GAMEBOY) {
	log.Println("0x41 BIT 0_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x42 BIT 0_D
func CB_BIT_0x42_0_D(GB *GAMEBOY) {
	log.Println("0x42 BIT 0_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x43 BIT 0_E
func CB_BIT_0x43_0_E(GB *GAMEBOY) {
	log.Println("0x43 BIT 0_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x44 BIT 0_H
func CB_BIT_0x44_0_H(GB *GAMEBOY) {
	log.Println("0x44 BIT 0_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x45 BIT 0_L
func CB_BIT_0x45_0_L(GB *GAMEBOY) {
	log.Println("0x45 BIT 0_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x46 BIT 0_addrHL
func CB_BIT_0x46_0_addrHL(GB *GAMEBOY) {
	log.Println("0x46 BIT 0_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x47 BIT 0_A
func CB_BIT_0x47_0_A(GB *GAMEBOY) {
	log.Println("0x47 BIT 0_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x48 BIT 1_B
func CB_BIT_0x48_1_B(GB *GAMEBOY) {
	log.Println("0x48 BIT 1_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x49 BIT 1_C
func CB_BIT_0x49_1_C(GB *GAMEBOY) {
	log.Println("0x49 BIT 1_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x4A BIT 1_D
func CB_BIT_0x4A_1_D(GB *GAMEBOY) {
	log.Println("0x4A BIT 1_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x4B BIT 1_E
func CB_BIT_0x4B_1_E(GB *GAMEBOY) {
	log.Println("0x4B BIT 1_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x4C BIT 1_H
func CB_BIT_0x4C_1_H(GB *GAMEBOY) {
	log.Println("0x4C BIT 1_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x4D BIT 1_L
func CB_BIT_0x4D_1_L(GB *GAMEBOY) {
	log.Println("0x4D BIT 1_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x4E BIT 1_addrHL
func CB_BIT_0x4E_1_addrHL(GB *GAMEBOY) {
	log.Println("0x4E BIT 1_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x4F BIT 1_A
func CB_BIT_0x4F_1_A(GB *GAMEBOY) {
	log.Println("0x4F BIT 1_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x50 BIT 2_B
func CB_BIT_0x50_2_B(GB *GAMEBOY) {
	log.Println("0x50 BIT 2_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x51 BIT 2_C
func CB_BIT_0x51_2_C(GB *GAMEBOY) {
	log.Println("0x51 BIT 2_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x52 BIT 2_D
func CB_BIT_0x52_2_D(GB *GAMEBOY) {
	log.Println("0x52 BIT 2_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x53 BIT 2_E
func CB_BIT_0x53_2_E(GB *GAMEBOY) {
	log.Println("0x53 BIT 2_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x54 BIT 2_H
func CB_BIT_0x54_2_H(GB *GAMEBOY) {
	log.Println("0x54 BIT 2_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x55 BIT 2_L
func CB_BIT_0x55_2_L(GB *GAMEBOY) {
	log.Println("0x55 BIT 2_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x56 BIT 2_addrHL
func CB_BIT_0x56_2_addrHL(GB *GAMEBOY) {
	log.Println("0x56 BIT 2_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x57 BIT 2_A
func CB_BIT_0x57_2_A(GB *GAMEBOY) {
	log.Println("0x57 BIT 2_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x58 BIT 3_B
func CB_BIT_0x58_3_B(GB *GAMEBOY) {
	log.Println("0x58 BIT 3_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x59 BIT 3_C
func CB_BIT_0x59_3_C(GB *GAMEBOY) {
	log.Println("0x59 BIT 3_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x5A BIT 3_D
func CB_BIT_0x5A_3_D(GB *GAMEBOY) {
	log.Println("0x5A BIT 3_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x5B BIT 3_E
func CB_BIT_0x5B_3_E(GB *GAMEBOY) {
	log.Println("0x5B BIT 3_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x5C BIT 3_H
func CB_BIT_0x5C_3_H(GB *GAMEBOY) {
	log.Println("0x5C BIT 3_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x5D BIT 3_L
func CB_BIT_0x5D_3_L(GB *GAMEBOY) {
	log.Println("0x5D BIT 3_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x5E BIT 3_addrHL
func CB_BIT_0x5E_3_addrHL(GB *GAMEBOY) {
	log.Println("0x5E BIT 3_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x5F BIT 3_A
func CB_BIT_0x5F_3_A(GB *GAMEBOY) {
	log.Println("0x5F BIT 3_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x60 BIT 4_B
func CB_BIT_0x60_4_B(GB *GAMEBOY) {
	log.Println("0x60 BIT 4_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x61 BIT 4_C
func CB_BIT_0x61_4_C(GB *GAMEBOY) {
	log.Println("0x61 BIT 4_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x62 BIT 4_D
func CB_BIT_0x62_4_D(GB *GAMEBOY) {
	log.Println("0x62 BIT 4_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x63 BIT 4_E
func CB_BIT_0x63_4_E(GB *GAMEBOY) {
	log.Println("0x63 BIT 4_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x64 BIT 4_H
func CB_BIT_0x64_4_H(GB *GAMEBOY) {
	log.Println("0x64 BIT 4_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x65 BIT 4_L
func CB_BIT_0x65_4_L(GB *GAMEBOY) {
	log.Println("0x65 BIT 4_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x66 BIT 4_addrHL
func CB_BIT_0x66_4_addrHL(GB *GAMEBOY) {
	log.Println("0x66 BIT 4_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x67 BIT 4_A
func CB_BIT_0x67_4_A(GB *GAMEBOY) {
	log.Println("0x67 BIT 4_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x68 BIT 5_B
func CB_BIT_0x68_5_B(GB *GAMEBOY) {
	log.Println("0x68 BIT 5_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x69 BIT 5_C
func CB_BIT_0x69_5_C(GB *GAMEBOY) {
	log.Println("0x69 BIT 5_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x6A BIT 5_D
func CB_BIT_0x6A_5_D(GB *GAMEBOY) {
	log.Println("0x6A BIT 5_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x6B BIT 5_E
func CB_BIT_0x6B_5_E(GB *GAMEBOY) {
	log.Println("0x6B BIT 5_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x6C BIT 5_H
func CB_BIT_0x6C_5_H(GB *GAMEBOY) {
	log.Println("0x6C BIT 5_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x6D BIT 5_L
func CB_BIT_0x6D_5_L(GB *GAMEBOY) {
	log.Println("0x6D BIT 5_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x6E BIT 5_addrHL
func CB_BIT_0x6E_5_addrHL(GB *GAMEBOY) {
	log.Println("0x6E BIT 5_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x6F BIT 5_A
func CB_BIT_0x6F_5_A(GB *GAMEBOY) {
	log.Println("0x6F BIT 5_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x70 BIT 6_B
func CB_BIT_0x70_6_B(GB *GAMEBOY) {
	log.Println("0x70 BIT 6_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x71 BIT 6_C
func CB_BIT_0x71_6_C(GB *GAMEBOY) {
	log.Println("0x71 BIT 6_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x72 BIT 6_D
func CB_BIT_0x72_6_D(GB *GAMEBOY) {
	log.Println("0x72 BIT 6_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x73 BIT 6_E
func CB_BIT_0x73_6_E(GB *GAMEBOY) {
	log.Println("0x73 BIT 6_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x74 BIT 6_H
func CB_BIT_0x74_6_H(GB *GAMEBOY) {
	log.Println("0x74 BIT 6_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x75 BIT 6_L
func CB_BIT_0x75_6_L(GB *GAMEBOY) {
	log.Println("0x75 BIT 6_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x76 BIT 6_addrHL
func CB_BIT_0x76_6_addrHL(GB *GAMEBOY) {
	log.Println("0x76 BIT 6_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x77 BIT 6_A
func CB_BIT_0x77_6_A(GB *GAMEBOY) {
	log.Println("0x77 BIT 6_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x78 BIT 7_B
func CB_BIT_0x78_7_B(GB *GAMEBOY) {
	log.Println("0x78 BIT 7_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x79 BIT 7_C
func CB_BIT_0x79_7_C(GB *GAMEBOY) {
	log.Println("0x79 BIT 7_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x7A BIT 7_D
func CB_BIT_0x7A_7_D(GB *GAMEBOY) {
	log.Println("0x7A BIT 7_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x7B BIT 7_E
func CB_BIT_0x7B_7_E(GB *GAMEBOY) {
	log.Println("0x7B BIT 7_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x7C BIT 7_H
func CB_BIT_0x7C_7_H(GB *GAMEBOY) {
	log.Println("0x7C BIT 7_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x7D BIT 7_L
func CB_BIT_0x7D_7_L(GB *GAMEBOY) {
	log.Println("0x7D BIT 7_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x7E BIT 7_addrHL
func CB_BIT_0x7E_7_addrHL(GB *GAMEBOY) {
	log.Println("0x7E BIT 7_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 3
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x7F BIT 7_A
func CB_BIT_0x7F_7_A(GB *GAMEBOY) {
	log.Println("0x7F BIT 7_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': 'Z', 'N': '0', 'H': '1', 'C': '-'}
}

// 0x80 RES 0_B
func CB_RES_0x80_0_B(GB *GAMEBOY) {
	log.Println("0x80 RES 0_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x81 RES 0_C
func CB_RES_0x81_0_C(GB *GAMEBOY) {
	log.Println("0x81 RES 0_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x82 RES 0_D
func CB_RES_0x82_0_D(GB *GAMEBOY) {
	log.Println("0x82 RES 0_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x83 RES 0_E
func CB_RES_0x83_0_E(GB *GAMEBOY) {
	log.Println("0x83 RES 0_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x84 RES 0_H
func CB_RES_0x84_0_H(GB *GAMEBOY) {
	log.Println("0x84 RES 0_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x85 RES 0_L
func CB_RES_0x85_0_L(GB *GAMEBOY) {
	log.Println("0x85 RES 0_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x86 RES 0_addrHL
func CB_RES_0x86_0_addrHL(GB *GAMEBOY) {
	log.Println("0x86 RES 0_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x87 RES 0_A
func CB_RES_0x87_0_A(GB *GAMEBOY) {
	log.Println("0x87 RES 0_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x88 RES 1_B
func CB_RES_0x88_1_B(GB *GAMEBOY) {
	log.Println("0x88 RES 1_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x89 RES 1_C
func CB_RES_0x89_1_C(GB *GAMEBOY) {
	log.Println("0x89 RES 1_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x8A RES 1_D
func CB_RES_0x8A_1_D(GB *GAMEBOY) {
	log.Println("0x8A RES 1_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x8B RES 1_E
func CB_RES_0x8B_1_E(GB *GAMEBOY) {
	log.Println("0x8B RES 1_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x8C RES 1_H
func CB_RES_0x8C_1_H(GB *GAMEBOY) {
	log.Println("0x8C RES 1_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x8D RES 1_L
func CB_RES_0x8D_1_L(GB *GAMEBOY) {
	log.Println("0x8D RES 1_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x8E RES 1_addrHL
func CB_RES_0x8E_1_addrHL(GB *GAMEBOY) {
	log.Println("0x8E RES 1_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x8F RES 1_A
func CB_RES_0x8F_1_A(GB *GAMEBOY) {
	log.Println("0x8F RES 1_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x90 RES 2_B
func CB_RES_0x90_2_B(GB *GAMEBOY) {
	log.Println("0x90 RES 2_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x91 RES 2_C
func CB_RES_0x91_2_C(GB *GAMEBOY) {
	log.Println("0x91 RES 2_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x92 RES 2_D
func CB_RES_0x92_2_D(GB *GAMEBOY) {
	log.Println("0x92 RES 2_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x93 RES 2_E
func CB_RES_0x93_2_E(GB *GAMEBOY) {
	log.Println("0x93 RES 2_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x94 RES 2_H
func CB_RES_0x94_2_H(GB *GAMEBOY) {
	log.Println("0x94 RES 2_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x95 RES 2_L
func CB_RES_0x95_2_L(GB *GAMEBOY) {
	log.Println("0x95 RES 2_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x96 RES 2_addrHL
func CB_RES_0x96_2_addrHL(GB *GAMEBOY) {
	log.Println("0x96 RES 2_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x97 RES 2_A
func CB_RES_0x97_2_A(GB *GAMEBOY) {
	log.Println("0x97 RES 2_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x98 RES 3_B
func CB_RES_0x98_3_B(GB *GAMEBOY) {
	log.Println("0x98 RES 3_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x99 RES 3_C
func CB_RES_0x99_3_C(GB *GAMEBOY) {
	log.Println("0x99 RES 3_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x9A RES 3_D
func CB_RES_0x9A_3_D(GB *GAMEBOY) {
	log.Println("0x9A RES 3_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x9B RES 3_E
func CB_RES_0x9B_3_E(GB *GAMEBOY) {
	log.Println("0x9B RES 3_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x9C RES 3_H
func CB_RES_0x9C_3_H(GB *GAMEBOY) {
	log.Println("0x9C RES 3_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x9D RES 3_L
func CB_RES_0x9D_3_L(GB *GAMEBOY) {
	log.Println("0x9D RES 3_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x9E RES 3_addrHL
func CB_RES_0x9E_3_addrHL(GB *GAMEBOY) {
	log.Println("0x9E RES 3_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0x9F RES 3_A
func CB_RES_0x9F_3_A(GB *GAMEBOY) {
	log.Println("0x9F RES 3_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA0 RES 4_B
func CB_RES_0xA0_4_B(GB *GAMEBOY) {
	log.Println("0xA0 RES 4_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA1 RES 4_C
func CB_RES_0xA1_4_C(GB *GAMEBOY) {
	log.Println("0xA1 RES 4_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA2 RES 4_D
func CB_RES_0xA2_4_D(GB *GAMEBOY) {
	log.Println("0xA2 RES 4_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA3 RES 4_E
func CB_RES_0xA3_4_E(GB *GAMEBOY) {
	log.Println("0xA3 RES 4_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA4 RES 4_H
func CB_RES_0xA4_4_H(GB *GAMEBOY) {
	log.Println("0xA4 RES 4_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA5 RES 4_L
func CB_RES_0xA5_4_L(GB *GAMEBOY) {
	log.Println("0xA5 RES 4_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA6 RES 4_addrHL
func CB_RES_0xA6_4_addrHL(GB *GAMEBOY) {
	log.Println("0xA6 RES 4_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA7 RES 4_A
func CB_RES_0xA7_4_A(GB *GAMEBOY) {
	log.Println("0xA7 RES 4_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA8 RES 5_B
func CB_RES_0xA8_5_B(GB *GAMEBOY) {
	log.Println("0xA8 RES 5_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xA9 RES 5_C
func CB_RES_0xA9_5_C(GB *GAMEBOY) {
	log.Println("0xA9 RES 5_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xAA RES 5_D
func CB_RES_0xAA_5_D(GB *GAMEBOY) {
	log.Println("0xAA RES 5_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xAB RES 5_E
func CB_RES_0xAB_5_E(GB *GAMEBOY) {
	log.Println("0xAB RES 5_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xAC RES 5_H
func CB_RES_0xAC_5_H(GB *GAMEBOY) {
	log.Println("0xAC RES 5_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xAD RES 5_L
func CB_RES_0xAD_5_L(GB *GAMEBOY) {
	log.Println("0xAD RES 5_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xAE RES 5_addrHL
func CB_RES_0xAE_5_addrHL(GB *GAMEBOY) {
	log.Println("0xAE RES 5_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xAF RES 5_A
func CB_RES_0xAF_5_A(GB *GAMEBOY) {
	log.Println("0xAF RES 5_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB0 RES 6_B
func CB_RES_0xB0_6_B(GB *GAMEBOY) {
	log.Println("0xB0 RES 6_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB1 RES 6_C
func CB_RES_0xB1_6_C(GB *GAMEBOY) {
	log.Println("0xB1 RES 6_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB2 RES 6_D
func CB_RES_0xB2_6_D(GB *GAMEBOY) {
	log.Println("0xB2 RES 6_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB3 RES 6_E
func CB_RES_0xB3_6_E(GB *GAMEBOY) {
	log.Println("0xB3 RES 6_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB4 RES 6_H
func CB_RES_0xB4_6_H(GB *GAMEBOY) {
	log.Println("0xB4 RES 6_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB5 RES 6_L
func CB_RES_0xB5_6_L(GB *GAMEBOY) {
	log.Println("0xB5 RES 6_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB6 RES 6_addrHL
func CB_RES_0xB6_6_addrHL(GB *GAMEBOY) {
	log.Println("0xB6 RES 6_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB7 RES 6_A
func CB_RES_0xB7_6_A(GB *GAMEBOY) {
	log.Println("0xB7 RES 6_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB8 RES 7_B
func CB_RES_0xB8_7_B(GB *GAMEBOY) {
	log.Println("0xB8 RES 7_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xB9 RES 7_C
func CB_RES_0xB9_7_C(GB *GAMEBOY) {
	log.Println("0xB9 RES 7_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xBA RES 7_D
func CB_RES_0xBA_7_D(GB *GAMEBOY) {
	log.Println("0xBA RES 7_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xBB RES 7_E
func CB_RES_0xBB_7_E(GB *GAMEBOY) {
	log.Println("0xBB RES 7_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xBC RES 7_H
func CB_RES_0xBC_7_H(GB *GAMEBOY) {
	log.Println("0xBC RES 7_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xBD RES 7_L
func CB_RES_0xBD_7_L(GB *GAMEBOY) {
	log.Println("0xBD RES 7_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xBE RES 7_addrHL
func CB_RES_0xBE_7_addrHL(GB *GAMEBOY) {
	log.Println("0xBE RES 7_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xBF RES 7_A
func CB_RES_0xBF_7_A(GB *GAMEBOY) {
	log.Println("0xBF RES 7_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC0 SET 0_B
func CB_SET_0xC0_0_B(GB *GAMEBOY) {
	log.Println("0xC0 SET 0_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC1 SET 0_C
func CB_SET_0xC1_0_C(GB *GAMEBOY) {
	log.Println("0xC1 SET 0_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC2 SET 0_D
func CB_SET_0xC2_0_D(GB *GAMEBOY) {
	log.Println("0xC2 SET 0_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC3 SET 0_E
func CB_SET_0xC3_0_E(GB *GAMEBOY) {
	log.Println("0xC3 SET 0_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC4 SET 0_H
func CB_SET_0xC4_0_H(GB *GAMEBOY) {
	log.Println("0xC4 SET 0_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC5 SET 0_L
func CB_SET_0xC5_0_L(GB *GAMEBOY) {
	log.Println("0xC5 SET 0_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC6 SET 0_addrHL
func CB_SET_0xC6_0_addrHL(GB *GAMEBOY) {
	log.Println("0xC6 SET 0_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC7 SET 0_A
func CB_SET_0xC7_0_A(GB *GAMEBOY) {
	log.Println("0xC7 SET 0_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC8 SET 1_B
func CB_SET_0xC8_1_B(GB *GAMEBOY) {
	log.Println("0xC8 SET 1_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xC9 SET 1_C
func CB_SET_0xC9_1_C(GB *GAMEBOY) {
	log.Println("0xC9 SET 1_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCA SET 1_D
func CB_SET_0xCA_1_D(GB *GAMEBOY) {
	log.Println("0xCA SET 1_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCB SET 1_E
func CB_SET_0xCB_1_E(GB *GAMEBOY) {
	log.Println("0xCB SET 1_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCC SET 1_H
func CB_SET_0xCC_1_H(GB *GAMEBOY) {
	log.Println("0xCC SET 1_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCD SET 1_L
func CB_SET_0xCD_1_L(GB *GAMEBOY) {
	log.Println("0xCD SET 1_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCE SET 1_addrHL
func CB_SET_0xCE_1_addrHL(GB *GAMEBOY) {
	log.Println("0xCE SET 1_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xCF SET 1_A
func CB_SET_0xCF_1_A(GB *GAMEBOY) {
	log.Println("0xCF SET 1_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD0 SET 2_B
func CB_SET_0xD0_2_B(GB *GAMEBOY) {
	log.Println("0xD0 SET 2_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD1 SET 2_C
func CB_SET_0xD1_2_C(GB *GAMEBOY) {
	log.Println("0xD1 SET 2_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD2 SET 2_D
func CB_SET_0xD2_2_D(GB *GAMEBOY) {
	log.Println("0xD2 SET 2_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD3 SET 2_E
func CB_SET_0xD3_2_E(GB *GAMEBOY) {
	log.Println("0xD3 SET 2_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD4 SET 2_H
func CB_SET_0xD4_2_H(GB *GAMEBOY) {
	log.Println("0xD4 SET 2_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD5 SET 2_L
func CB_SET_0xD5_2_L(GB *GAMEBOY) {
	log.Println("0xD5 SET 2_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD6 SET 2_addrHL
func CB_SET_0xD6_2_addrHL(GB *GAMEBOY) {
	log.Println("0xD6 SET 2_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD7 SET 2_A
func CB_SET_0xD7_2_A(GB *GAMEBOY) {
	log.Println("0xD7 SET 2_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD8 SET 3_B
func CB_SET_0xD8_3_B(GB *GAMEBOY) {
	log.Println("0xD8 SET 3_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xD9 SET 3_C
func CB_SET_0xD9_3_C(GB *GAMEBOY) {
	log.Println("0xD9 SET 3_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDA SET 3_D
func CB_SET_0xDA_3_D(GB *GAMEBOY) {
	log.Println("0xDA SET 3_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDB SET 3_E
func CB_SET_0xDB_3_E(GB *GAMEBOY) {
	log.Println("0xDB SET 3_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDC SET 3_H
func CB_SET_0xDC_3_H(GB *GAMEBOY) {
	log.Println("0xDC SET 3_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDD SET 3_L
func CB_SET_0xDD_3_L(GB *GAMEBOY) {
	log.Println("0xDD SET 3_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDE SET 3_addrHL
func CB_SET_0xDE_3_addrHL(GB *GAMEBOY) {
	log.Println("0xDE SET 3_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xDF SET 3_A
func CB_SET_0xDF_3_A(GB *GAMEBOY) {
	log.Println("0xDF SET 3_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE0 SET 4_B
func CB_SET_0xE0_4_B(GB *GAMEBOY) {
	log.Println("0xE0 SET 4_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE1 SET 4_C
func CB_SET_0xE1_4_C(GB *GAMEBOY) {
	log.Println("0xE1 SET 4_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE2 SET 4_D
func CB_SET_0xE2_4_D(GB *GAMEBOY) {
	log.Println("0xE2 SET 4_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE3 SET 4_E
func CB_SET_0xE3_4_E(GB *GAMEBOY) {
	log.Println("0xE3 SET 4_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE4 SET 4_H
func CB_SET_0xE4_4_H(GB *GAMEBOY) {
	log.Println("0xE4 SET 4_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE5 SET 4_L
func CB_SET_0xE5_4_L(GB *GAMEBOY) {
	log.Println("0xE5 SET 4_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE6 SET 4_addrHL
func CB_SET_0xE6_4_addrHL(GB *GAMEBOY) {
	log.Println("0xE6 SET 4_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE7 SET 4_A
func CB_SET_0xE7_4_A(GB *GAMEBOY) {
	log.Println("0xE7 SET 4_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE8 SET 5_B
func CB_SET_0xE8_5_B(GB *GAMEBOY) {
	log.Println("0xE8 SET 5_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xE9 SET 5_C
func CB_SET_0xE9_5_C(GB *GAMEBOY) {
	log.Println("0xE9 SET 5_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xEA SET 5_D
func CB_SET_0xEA_5_D(GB *GAMEBOY) {
	log.Println("0xEA SET 5_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xEB SET 5_E
func CB_SET_0xEB_5_E(GB *GAMEBOY) {
	log.Println("0xEB SET 5_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xEC SET 5_H
func CB_SET_0xEC_5_H(GB *GAMEBOY) {
	log.Println("0xEC SET 5_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xED SET 5_L
func CB_SET_0xED_5_L(GB *GAMEBOY) {
	log.Println("0xED SET 5_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xEE SET 5_addrHL
func CB_SET_0xEE_5_addrHL(GB *GAMEBOY) {
	log.Println("0xEE SET 5_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xEF SET 5_A
func CB_SET_0xEF_5_A(GB *GAMEBOY) {
	log.Println("0xEF SET 5_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF0 SET 6_B
func CB_SET_0xF0_6_B(GB *GAMEBOY) {
	log.Println("0xF0 SET 6_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF1 SET 6_C
func CB_SET_0xF1_6_C(GB *GAMEBOY) {
	log.Println("0xF1 SET 6_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF2 SET 6_D
func CB_SET_0xF2_6_D(GB *GAMEBOY) {
	log.Println("0xF2 SET 6_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF3 SET 6_E
func CB_SET_0xF3_6_E(GB *GAMEBOY) {
	log.Println("0xF3 SET 6_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF4 SET 6_H
func CB_SET_0xF4_6_H(GB *GAMEBOY) {
	log.Println("0xF4 SET 6_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF5 SET 6_L
func CB_SET_0xF5_6_L(GB *GAMEBOY) {
	log.Println("0xF5 SET 6_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF6 SET 6_addrHL
func CB_SET_0xF6_6_addrHL(GB *GAMEBOY) {
	log.Println("0xF6 SET 6_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF7 SET 6_A
func CB_SET_0xF7_6_A(GB *GAMEBOY) {
	log.Println("0xF7 SET 6_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF8 SET 7_B
func CB_SET_0xF8_7_B(GB *GAMEBOY) {
	log.Println("0xF8 SET 7_B")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xF9 SET 7_C
func CB_SET_0xF9_7_C(GB *GAMEBOY) {
	log.Println("0xF9 SET 7_C")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFA SET 7_D
func CB_SET_0xFA_7_D(GB *GAMEBOY) {
	log.Println("0xFA SET 7_D")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFB SET 7_E
func CB_SET_0xFB_7_E(GB *GAMEBOY) {
	log.Println("0xFB SET 7_E")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFC SET 7_H
func CB_SET_0xFC_7_H(GB *GAMEBOY) {
	log.Println("0xFC SET 7_H")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFD SET 7_L
func CB_SET_0xFD_7_L(GB *GAMEBOY) {
	log.Println("0xFD SET 7_L")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFE SET 7_addrHL
func CB_SET_0xFE_7_addrHL(GB *GAMEBOY) {
	log.Println("0xFE SET 7_addrHL")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 4
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}

// 0xFF SET 7_A
func CB_SET_0xFF_7_A(GB *GAMEBOY) {
	log.Println("0xFF SET 7_A")
	//NEEDS CODE
	GB.CPU._r.PC += 2
	GB.CPU._r.M = 2
	GB.CPU._r.T = GB.CPU._r.M * 4
	//FLAGS AFFECTED : {'Z': '-', 'N': '-', 'H': '-', 'C': '-'}
}
