package main

import "log"

// 0x00 NOP
func NOP_0x00(CPU *CPU, MMU *MMU) {
	log.Println("0x00 NOP ")
	CPU._r.M = 1
	CPU._r.T = 4
}

// 0x01 LD BC_n16
func LD_0x01_BC_n16(CPU *CPU, MMU *MMU) {
	log.Println("0x01 LD BC_n16")
}

// 0x02 LD BC_A
func LD_0x02_BC_A(CPU *CPU, MMU *MMU) {
	log.Println("0x02 LD BC_A")
}

// 0x03 INC BC
func INC_0x03_BC(CPU *CPU, MMU *MMU) {
	log.Println("0x03 INC BC")
}

// 0x04 INC B
func INC_0x04_B(CPU *CPU, MMU *MMU) {
	log.Println("0x04 INC B")
}

// 0x05 DEC B
func DEC_0x05_B(CPU *CPU, MMU *MMU) {
	log.Println("0x05 DEC B")
}

// 0x06 LD B_n8
func LD_0x06_B_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x06 LD B_n8")
}

// 0x07 RLCA
func RLCA_0x07(CPU *CPU, MMU *MMU) {
	log.Println("0x07 RLCA ")
}

// 0x08 LD a16_SP
func LD_0x08_a16_SP(CPU *CPU, MMU *MMU) {
	log.Println("0x08 LD a16_SP")
}

// 0x09 ADD HL_BC
func ADD_0x09_HL_BC(CPU *CPU, MMU *MMU) {
	log.Println("0x09 ADD HL_BC")
}

// 0x0A LD A_BC
func LD_0x0A_A_BC(CPU *CPU, MMU *MMU) {
	log.Println("0x0A LD A_BC")
}

// 0x0B DEC BC
func DEC_0x0B_BC(CPU *CPU, MMU *MMU) {
	log.Println("0x0B DEC BC")
}

// 0x0C INC C
func INC_0x0C_C(CPU *CPU, MMU *MMU) {
	log.Println("0x0C INC C")
}

// 0x0D DEC C
func DEC_0x0D_C(CPU *CPU, MMU *MMU) {
	log.Println("0x0D DEC C")
}

// 0x0E LD C_n8
func LD_0x0E_C_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x0E LD C_n8")
}

// 0x0F RRCA
func RRCA_0x0F(CPU *CPU, MMU *MMU) {
	log.Println("0x0F RRCA ")
}

// 0x10 STOP n8
func STOP_0x10_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x10 STOP n8")
}

// 0x11 LD DE_n16
func LD_0x11_DE_n16(CPU *CPU, MMU *MMU) {
	log.Println("0x11 LD DE_n16")
}

// 0x12 LD DE_A
func LD_0x12_DE_A(CPU *CPU, MMU *MMU) {
	log.Println("0x12 LD DE_A")
}

// 0x13 INC DE
func INC_0x13_DE(CPU *CPU, MMU *MMU) {
	log.Println("0x13 INC DE")
}

// 0x14 INC D
func INC_0x14_D(CPU *CPU, MMU *MMU) {
	log.Println("0x14 INC D")
}

// 0x15 DEC D
func DEC_0x15_D(CPU *CPU, MMU *MMU) {
	log.Println("0x15 DEC D")
}

// 0x16 LD D_n8
func LD_0x16_D_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x16 LD D_n8")
}

// 0x17 RLA
func RLA_0x17(CPU *CPU, MMU *MMU) {
	log.Println("0x17 RLA ")
}

// 0x18 JR e8
func JR_0x18_e8(CPU *CPU, MMU *MMU) {
	log.Println("0x18 JR e8")
}

// 0x19 ADD HL_DE
func ADD_0x19_HL_DE(CPU *CPU, MMU *MMU) {
	log.Println("0x19 ADD HL_DE")
}

// 0x1A LD A_DE
func LD_0x1A_A_DE(CPU *CPU, MMU *MMU) {
	log.Println("0x1A LD A_DE")
}

// 0x1B DEC DE
func DEC_0x1B_DE(CPU *CPU, MMU *MMU) {
	log.Println("0x1B DEC DE")
}

// 0x1C INC E
func INC_0x1C_E(CPU *CPU, MMU *MMU) {
	log.Println("0x1C INC E")
}

// 0x1D DEC E
func DEC_0x1D_E(CPU *CPU, MMU *MMU) {
	log.Println("0x1D DEC E")
}

// 0x1E LD E_n8
func LD_0x1E_E_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x1E LD E_n8")
}

// 0x1F RRA
func RRA_0x1F(CPU *CPU, MMU *MMU) {
	log.Println("0x1F RRA ")
}

// 0x20 JR NZ_e8
func JR_0x20_NZ_e8(CPU *CPU, MMU *MMU) {
	log.Println("0x20 JR NZ_e8")
}

// 0x21 LD HL_n16
func LD_0x21_HL_n16(CPU *CPU, MMU *MMU) {
	log.Println("0x21 LD HL_n16")
}

// 0x22 LD HL_A
func LD_0x22_HL_A(CPU *CPU, MMU *MMU) {
	log.Println("0x22 LD HL_A")
}

// 0x23 INC HL
func INC_0x23_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x23 INC HL")
}

// 0x24 INC H
func INC_0x24_H(CPU *CPU, MMU *MMU) {
	log.Println("0x24 INC H")
}

// 0x25 DEC H
func DEC_0x25_H(CPU *CPU, MMU *MMU) {
	log.Println("0x25 DEC H")
}

// 0x26 LD H_n8
func LD_0x26_H_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x26 LD H_n8")
}

// 0x27 DAA
func DAA_0x27(CPU *CPU, MMU *MMU) {
	log.Println("0x27 DAA ")
}

// 0x28 JR Z_e8
func JR_0x28_Z_e8(CPU *CPU, MMU *MMU) {
	log.Println("0x28 JR Z_e8")
}

// 0x29 ADD HL_HL
func ADD_0x29_HL_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x29 ADD HL_HL")
}

// 0x2A LD A_HL
func LD_0x2A_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x2A LD A_HL")
}

// 0x2B DEC HL
func DEC_0x2B_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x2B DEC HL")
}

// 0x2C INC L
func INC_0x2C_L(CPU *CPU, MMU *MMU) {
	log.Println("0x2C INC L")
}

// 0x2D DEC L
func DEC_0x2D_L(CPU *CPU, MMU *MMU) {
	log.Println("0x2D DEC L")
}

// 0x2E LD L_n8
func LD_0x2E_L_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x2E LD L_n8")
}

// 0x2F CPL
func CPL_0x2F(CPU *CPU, MMU *MMU) {
	log.Println("0x2F CPL ")
}

// 0x30 JR NC_e8
func JR_0x30_NC_e8(CPU *CPU, MMU *MMU) {
	log.Println("0x30 JR NC_e8")
}

// 0x31 LD SP_n16
func LD_0x31_SP_n16(CPU *CPU, MMU *MMU) {
	log.Println("0x31 LD SP_n16")
}

// 0x32 LD HL_A
func LD_0x32_HL_A(CPU *CPU, MMU *MMU) {
	log.Println("0x32 LD HL_A")
}

// 0x33 INC SP
func INC_0x33_SP(CPU *CPU, MMU *MMU) {
	log.Println("0x33 INC SP")
}

// 0x34 INC HL
func INC_0x34_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x34 INC HL")
}

// 0x35 DEC HL
func DEC_0x35_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x35 DEC HL")
}

// 0x36 LD HL_n8
func LD_0x36_HL_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x36 LD HL_n8")
}

// 0x37 SCF
func SCF_0x37(CPU *CPU, MMU *MMU) {
	log.Println("0x37 SCF ")
}

// 0x38 JR C_e8
func JR_0x38_C_e8(CPU *CPU, MMU *MMU) {
	log.Println("0x38 JR C_e8")
}

// 0x39 ADD HL_SP
func ADD_0x39_HL_SP(CPU *CPU, MMU *MMU) {
	log.Println("0x39 ADD HL_SP")
}

// 0x3A LD A_HL
func LD_0x3A_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x3A LD A_HL")
}

// 0x3B DEC SP
func DEC_0x3B_SP(CPU *CPU, MMU *MMU) {
	log.Println("0x3B DEC SP")
}

// 0x3C INC A
func INC_0x3C_A(CPU *CPU, MMU *MMU) {
	log.Println("0x3C INC A")
}

// 0x3D DEC A
func DEC_0x3D_A(CPU *CPU, MMU *MMU) {
	log.Println("0x3D DEC A")
}

// 0x3E LD A_n8
func LD_0x3E_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0x3E LD A_n8")
}

// 0x3F CCF
func CCF_0x3F(CPU *CPU, MMU *MMU) {
	log.Println("0x3F CCF ")
}

// 0x40 LD B_B
func LD_0x40_B_B(CPU *CPU, MMU *MMU) {
	log.Println("0x40 LD B_B")
}

// 0x41 LD B_C
func LD_0x41_B_C(CPU *CPU, MMU *MMU) {
	log.Println("0x41 LD B_C")
}

// 0x42 LD B_D
func LD_0x42_B_D(CPU *CPU, MMU *MMU) {
	log.Println("0x42 LD B_D")
}

// 0x43 LD B_E
func LD_0x43_B_E(CPU *CPU, MMU *MMU) {
	log.Println("0x43 LD B_E")
}

// 0x44 LD B_H
func LD_0x44_B_H(CPU *CPU, MMU *MMU) {
	log.Println("0x44 LD B_H")
}

// 0x45 LD B_L
func LD_0x45_B_L(CPU *CPU, MMU *MMU) {
	log.Println("0x45 LD B_L")
}

// 0x46 LD B_HL
func LD_0x46_B_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x46 LD B_HL")
}

// 0x47 LD B_A
func LD_0x47_B_A(CPU *CPU, MMU *MMU) {
	log.Println("0x47 LD B_A")
}

// 0x48 LD C_B
func LD_0x48_C_B(CPU *CPU, MMU *MMU) {
	log.Println("0x48 LD C_B")
}

// 0x49 LD C_C
func LD_0x49_C_C(CPU *CPU, MMU *MMU) {
	log.Println("0x49 LD C_C")
}

// 0x4A LD C_D
func LD_0x4A_C_D(CPU *CPU, MMU *MMU) {
	log.Println("0x4A LD C_D")
}

// 0x4B LD C_E
func LD_0x4B_C_E(CPU *CPU, MMU *MMU) {
	log.Println("0x4B LD C_E")
}

// 0x4C LD C_H
func LD_0x4C_C_H(CPU *CPU, MMU *MMU) {
	log.Println("0x4C LD C_H")
}

// 0x4D LD C_L
func LD_0x4D_C_L(CPU *CPU, MMU *MMU) {
	log.Println("0x4D LD C_L")
}

// 0x4E LD C_HL
func LD_0x4E_C_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x4E LD C_HL")
}

// 0x4F LD C_A
func LD_0x4F_C_A(CPU *CPU, MMU *MMU) {
	log.Println("0x4F LD C_A")
}

// 0x50 LD D_B
func LD_0x50_D_B(CPU *CPU, MMU *MMU) {
	log.Println("0x50 LD D_B")
}

// 0x51 LD D_C
func LD_0x51_D_C(CPU *CPU, MMU *MMU) {
	log.Println("0x51 LD D_C")
}

// 0x52 LD D_D
func LD_0x52_D_D(CPU *CPU, MMU *MMU) {
	log.Println("0x52 LD D_D")
}

// 0x53 LD D_E
func LD_0x53_D_E(CPU *CPU, MMU *MMU) {
	log.Println("0x53 LD D_E")
}

// 0x54 LD D_H
func LD_0x54_D_H(CPU *CPU, MMU *MMU) {
	log.Println("0x54 LD D_H")
}

// 0x55 LD D_L
func LD_0x55_D_L(CPU *CPU, MMU *MMU) {
	log.Println("0x55 LD D_L")
}

// 0x56 LD D_HL
func LD_0x56_D_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x56 LD D_HL")
}

// 0x57 LD D_A
func LD_0x57_D_A(CPU *CPU, MMU *MMU) {
	log.Println("0x57 LD D_A")
}

// 0x58 LD E_B
func LD_0x58_E_B(CPU *CPU, MMU *MMU) {
	log.Println("0x58 LD E_B")
}

// 0x59 LD E_C
func LD_0x59_E_C(CPU *CPU, MMU *MMU) {
	log.Println("0x59 LD E_C")
}

// 0x5A LD E_D
func LD_0x5A_E_D(CPU *CPU, MMU *MMU) {
	log.Println("0x5A LD E_D")
}

// 0x5B LD E_E
func LD_0x5B_E_E(CPU *CPU, MMU *MMU) {
	log.Println("0x5B LD E_E")
}

// 0x5C LD E_H
func LD_0x5C_E_H(CPU *CPU, MMU *MMU) {
	log.Println("0x5C LD E_H")
}

// 0x5D LD E_L
func LD_0x5D_E_L(CPU *CPU, MMU *MMU) {
	log.Println("0x5D LD E_L")
}

// 0x5E LD E_HL
func LD_0x5E_E_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x5E LD E_HL")
}

// 0x5F LD E_A
func LD_0x5F_E_A(CPU *CPU, MMU *MMU) {
	log.Println("0x5F LD E_A")
}

// 0x60 LD H_B
func LD_0x60_H_B(CPU *CPU, MMU *MMU) {
	log.Println("0x60 LD H_B")
}

// 0x61 LD H_C
func LD_0x61_H_C(CPU *CPU, MMU *MMU) {
	log.Println("0x61 LD H_C")
}

// 0x62 LD H_D
func LD_0x62_H_D(CPU *CPU, MMU *MMU) {
	log.Println("0x62 LD H_D")
}

// 0x63 LD H_E
func LD_0x63_H_E(CPU *CPU, MMU *MMU) {
	log.Println("0x63 LD H_E")
}

// 0x64 LD H_H
func LD_0x64_H_H(CPU *CPU, MMU *MMU) {
	log.Println("0x64 LD H_H")
}

// 0x65 LD H_L
func LD_0x65_H_L(CPU *CPU, MMU *MMU) {
	log.Println("0x65 LD H_L")
}

// 0x66 LD H_HL
func LD_0x66_H_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x66 LD H_HL")
}

// 0x67 LD H_A
func LD_0x67_H_A(CPU *CPU, MMU *MMU) {
	log.Println("0x67 LD H_A")
}

// 0x68 LD L_B
func LD_0x68_L_B(CPU *CPU, MMU *MMU) {
	log.Println("0x68 LD L_B")
}

// 0x69 LD L_C
func LD_0x69_L_C(CPU *CPU, MMU *MMU) {
	log.Println("0x69 LD L_C")
}

// 0x6A LD L_D
func LD_0x6A_L_D(CPU *CPU, MMU *MMU) {
	log.Println("0x6A LD L_D")
}

// 0x6B LD L_E
func LD_0x6B_L_E(CPU *CPU, MMU *MMU) {
	log.Println("0x6B LD L_E")
}

// 0x6C LD L_H
func LD_0x6C_L_H(CPU *CPU, MMU *MMU) {
	log.Println("0x6C LD L_H")
}

// 0x6D LD L_L
func LD_0x6D_L_L(CPU *CPU, MMU *MMU) {
	log.Println("0x6D LD L_L")
}

// 0x6E LD L_HL
func LD_0x6E_L_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x6E LD L_HL")
}

// 0x6F LD L_A
func LD_0x6F_L_A(CPU *CPU, MMU *MMU) {
	log.Println("0x6F LD L_A")
}

// 0x70 LD HL_B
func LD_0x70_HL_B(CPU *CPU, MMU *MMU) {
	log.Println("0x70 LD HL_B")
}

// 0x71 LD HL_C
func LD_0x71_HL_C(CPU *CPU, MMU *MMU) {
	log.Println("0x71 LD HL_C")
}

// 0x72 LD HL_D
func LD_0x72_HL_D(CPU *CPU, MMU *MMU) {
	log.Println("0x72 LD HL_D")
}

// 0x73 LD HL_E
func LD_0x73_HL_E(CPU *CPU, MMU *MMU) {
	log.Println("0x73 LD HL_E")
}

// 0x74 LD HL_H
func LD_0x74_HL_H(CPU *CPU, MMU *MMU) {
	log.Println("0x74 LD HL_H")
}

// 0x75 LD HL_L
func LD_0x75_HL_L(CPU *CPU, MMU *MMU) {
	log.Println("0x75 LD HL_L")
}

// 0x76 HALT
func HALT_0x76(CPU *CPU, MMU *MMU) {
	log.Println("0x76 HALT ")
}

// 0x77 LD HL_A
func LD_0x77_HL_A(CPU *CPU, MMU *MMU) {
	log.Println("0x77 LD HL_A")
}

// 0x78 LD A_B
func LD_0x78_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0x78 LD A_B")
}

// 0x79 LD A_C
func LD_0x79_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0x79 LD A_C")
}

// 0x7A LD A_D
func LD_0x7A_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0x7A LD A_D")
}

// 0x7B LD A_E
func LD_0x7B_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0x7B LD A_E")
}

// 0x7C LD A_H
func LD_0x7C_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0x7C LD A_H")
}

// 0x7D LD A_L
func LD_0x7D_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0x7D LD A_L")
}

// 0x7E LD A_HL
func LD_0x7E_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x7E LD A_HL")
}

// 0x7F LD A_A
func LD_0x7F_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0x7F LD A_A")
}

// 0x80 ADD A_B
func ADD_0x80_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0x80 ADD A_B")
}

// 0x81 ADD A_C
func ADD_0x81_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0x81 ADD A_C")
}

// 0x82 ADD A_D
func ADD_0x82_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0x82 ADD A_D")
}

// 0x83 ADD A_E
func ADD_0x83_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0x83 ADD A_E")
}

// 0x84 ADD A_H
func ADD_0x84_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0x84 ADD A_H")
}

// 0x85 ADD A_L
func ADD_0x85_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0x85 ADD A_L")
}

// 0x86 ADD A_HL
func ADD_0x86_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x86 ADD A_HL")
}

// 0x87 ADD A_A
func ADD_0x87_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0x87 ADD A_A")
}

// 0x88 ADC A_B
func ADC_0x88_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0x88 ADC A_B")
}

// 0x89 ADC A_C
func ADC_0x89_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0x89 ADC A_C")
}

// 0x8A ADC A_D
func ADC_0x8A_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0x8A ADC A_D")
}

// 0x8B ADC A_E
func ADC_0x8B_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0x8B ADC A_E")
}

// 0x8C ADC A_H
func ADC_0x8C_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0x8C ADC A_H")
}

// 0x8D ADC A_L
func ADC_0x8D_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0x8D ADC A_L")
}

// 0x8E ADC A_HL
func ADC_0x8E_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x8E ADC A_HL")
}

// 0x8F ADC A_A
func ADC_0x8F_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0x8F ADC A_A")
}

// 0x90 SUB A_B
func SUB_0x90_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0x90 SUB A_B")
}

// 0x91 SUB A_C
func SUB_0x91_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0x91 SUB A_C")
}

// 0x92 SUB A_D
func SUB_0x92_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0x92 SUB A_D")
}

// 0x93 SUB A_E
func SUB_0x93_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0x93 SUB A_E")
}

// 0x94 SUB A_H
func SUB_0x94_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0x94 SUB A_H")
}

// 0x95 SUB A_L
func SUB_0x95_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0x95 SUB A_L")
}

// 0x96 SUB A_HL
func SUB_0x96_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x96 SUB A_HL")
}

// 0x97 SUB A_A
func SUB_0x97_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0x97 SUB A_A")
}

// 0x98 SBC A_B
func SBC_0x98_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0x98 SBC A_B")
}

// 0x99 SBC A_C
func SBC_0x99_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0x99 SBC A_C")
}

// 0x9A SBC A_D
func SBC_0x9A_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0x9A SBC A_D")
}

// 0x9B SBC A_E
func SBC_0x9B_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0x9B SBC A_E")
}

// 0x9C SBC A_H
func SBC_0x9C_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0x9C SBC A_H")
}

// 0x9D SBC A_L
func SBC_0x9D_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0x9D SBC A_L")
}

// 0x9E SBC A_HL
func SBC_0x9E_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0x9E SBC A_HL")
}

// 0x9F SBC A_A
func SBC_0x9F_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0x9F SBC A_A")
}

// 0xA0 AND A_B
func AND_0xA0_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0xA0 AND A_B")
}

// 0xA1 AND A_C
func AND_0xA1_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0xA1 AND A_C")
}

// 0xA2 AND A_D
func AND_0xA2_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0xA2 AND A_D")
}

// 0xA3 AND A_E
func AND_0xA3_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0xA3 AND A_E")
}

// 0xA4 AND A_H
func AND_0xA4_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0xA4 AND A_H")
}

// 0xA5 AND A_L
func AND_0xA5_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0xA5 AND A_L")
}

// 0xA6 AND A_HL
func AND_0xA6_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xA6 AND A_HL")
}

// 0xA7 AND A_A
func AND_0xA7_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0xA7 AND A_A")
}

// 0xA8 XOR A_B
func XOR_0xA8_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0xA8 XOR A_B")
}

// 0xA9 XOR A_C
func XOR_0xA9_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0xA9 XOR A_C")
}

// 0xAA XOR A_D
func XOR_0xAA_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0xAA XOR A_D")
}

// 0xAB XOR A_E
func XOR_0xAB_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0xAB XOR A_E")
}

// 0xAC XOR A_H
func XOR_0xAC_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0xAC XOR A_H")
}

// 0xAD XOR A_L
func XOR_0xAD_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0xAD XOR A_L")
}

// 0xAE XOR A_HL
func XOR_0xAE_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xAE XOR A_HL")
}

// 0xAF XOR A_A
func XOR_0xAF_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0xAF XOR A_A")
}

// 0xB0 OR A_B
func OR_0xB0_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0xB0 OR A_B")
}

// 0xB1 OR A_C
func OR_0xB1_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0xB1 OR A_C")
}

// 0xB2 OR A_D
func OR_0xB2_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0xB2 OR A_D")
}

// 0xB3 OR A_E
func OR_0xB3_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0xB3 OR A_E")
}

// 0xB4 OR A_H
func OR_0xB4_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0xB4 OR A_H")
}

// 0xB5 OR A_L
func OR_0xB5_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0xB5 OR A_L")
}

// 0xB6 OR A_HL
func OR_0xB6_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xB6 OR A_HL")
}

// 0xB7 OR A_A
func OR_0xB7_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0xB7 OR A_A")
}

// 0xB8 CP A_B
func CP_0xB8_A_B(CPU *CPU, MMU *MMU) {
	log.Println("0xB8 CP A_B")
}

// 0xB9 CP A_C
func CP_0xB9_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0xB9 CP A_C")
}

// 0xBA CP A_D
func CP_0xBA_A_D(CPU *CPU, MMU *MMU) {
	log.Println("0xBA CP A_D")
}

// 0xBB CP A_E
func CP_0xBB_A_E(CPU *CPU, MMU *MMU) {
	log.Println("0xBB CP A_E")
}

// 0xBC CP A_H
func CP_0xBC_A_H(CPU *CPU, MMU *MMU) {
	log.Println("0xBC CP A_H")
}

// 0xBD CP A_L
func CP_0xBD_A_L(CPU *CPU, MMU *MMU) {
	log.Println("0xBD CP A_L")
}

// 0xBE CP A_HL
func CP_0xBE_A_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xBE CP A_HL")
}

// 0xBF CP A_A
func CP_0xBF_A_A(CPU *CPU, MMU *MMU) {
	log.Println("0xBF CP A_A")
}

// 0xC0 RET NZ
func RET_0xC0_NZ(CPU *CPU, MMU *MMU) {
	log.Println("0xC0 RET NZ")
}

// 0xC1 POP BC
func POP_0xC1_BC(CPU *CPU, MMU *MMU) {
	log.Println("0xC1 POP BC")
}

// 0xC2 JP NZ_a16
func JP_0xC2_NZ_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xC2 JP NZ_a16")
}

// 0xC3 JP a16
func JP_0xC3_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xC3 JP a16")
}

// 0xC4 CALL NZ_a16
func CALL_0xC4_NZ_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xC4 CALL NZ_a16")
}

// 0xC5 PUSH BC
func PUSH_0xC5_BC(CPU *CPU, MMU *MMU) {
	log.Println("0xC5 PUSH BC")
}

// 0xC6 ADD A_n8
func ADD_0xC6_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xC6 ADD A_n8")
}

// 0xC7 RST $00
func RST_0xC7_dollar00(CPU *CPU, MMU *MMU) {
	log.Println("0xC7 RST $00")
}

// 0xC8 RET Z
func RET_0xC8_Z(CPU *CPU, MMU *MMU) {
	log.Println("0xC8 RET Z")
}

// 0xC9 RET
func RET_0xC9(CPU *CPU, MMU *MMU) {
	log.Println("0xC9 RET ")
}

// 0xCA JP Z_a16
func JP_0xCA_Z_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xCA JP Z_a16")
}

// 0xCB PREFIX
func PREFIX_0xCB(CPU *CPU, MMU *MMU) {
	log.Println("0xCB PREFIX ")
}

// 0xCC CALL Z_a16
func CALL_0xCC_Z_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xCC CALL Z_a16")
}

// 0xCD CALL a16
func CALL_0xCD_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xCD CALL a16")
}

// 0xCE ADC A_n8
func ADC_0xCE_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xCE ADC A_n8")
}

// 0xCF RST $08
func RST_0xCF_dollar08(CPU *CPU, MMU *MMU) {
	log.Println("0xCF RST $08")
}

// 0xD0 RET NC
func RET_0xD0_NC(CPU *CPU, MMU *MMU) {
	log.Println("0xD0 RET NC")
}

// 0xD1 POP DE
func POP_0xD1_DE(CPU *CPU, MMU *MMU) {
	log.Println("0xD1 POP DE")
}

// 0xD2 JP NC_a16
func JP_0xD2_NC_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xD2 JP NC_a16")
}

// 0xD3 ILLEGAL_D3
func ILLEGAL_D3_0xD3(CPU *CPU, MMU *MMU) {
	log.Println("0xD3 ILLEGAL_D3 ")
}

// 0xD4 CALL NC_a16
func CALL_0xD4_NC_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xD4 CALL NC_a16")
}

// 0xD5 PUSH DE
func PUSH_0xD5_DE(CPU *CPU, MMU *MMU) {
	log.Println("0xD5 PUSH DE")
}

// 0xD6 SUB A_n8
func SUB_0xD6_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xD6 SUB A_n8")
}

// 0xD7 RST $10
func RST_0xD7_dollar10(CPU *CPU, MMU *MMU) {
	log.Println("0xD7 RST $10")
}

// 0xD8 RET C
func RET_0xD8_C(CPU *CPU, MMU *MMU) {
	log.Println("0xD8 RET C")
}

// 0xD9 RETI
func RETI_0xD9(CPU *CPU, MMU *MMU) {
	log.Println("0xD9 RETI ")
}

// 0xDA JP C_a16
func JP_0xDA_C_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xDA JP C_a16")
}

// 0xDB ILLEGAL_DB
func ILLEGAL_DB_0xDB(CPU *CPU, MMU *MMU) {
	log.Println("0xDB ILLEGAL_DB ")
}

// 0xDC CALL C_a16
func CALL_0xDC_C_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xDC CALL C_a16")
}

// 0xDD ILLEGAL_DD
func ILLEGAL_DD_0xDD(CPU *CPU, MMU *MMU) {
	log.Println("0xDD ILLEGAL_DD ")
}

// 0xDE SBC A_n8
func SBC_0xDE_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xDE SBC A_n8")
}

// 0xDF RST $18
func RST_0xDF_dollar18(CPU *CPU, MMU *MMU) {
	log.Println("0xDF RST $18")
}

// 0xE0 LDH a8_A
func LDH_0xE0_a8_A(CPU *CPU, MMU *MMU) {
	log.Println("0xE0 LDH a8_A")
}

// 0xE1 POP HL
func POP_0xE1_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xE1 POP HL")
}

// 0xE2 LDH C_A
func LDH_0xE2_C_A(CPU *CPU, MMU *MMU) {
	log.Println("0xE2 LDH C_A")
}

// 0xE3 ILLEGAL_E3
func ILLEGAL_E3_0xE3(CPU *CPU, MMU *MMU) {
	log.Println("0xE3 ILLEGAL_E3 ")
}

// 0xE4 ILLEGAL_E4
func ILLEGAL_E4_0xE4(CPU *CPU, MMU *MMU) {
	log.Println("0xE4 ILLEGAL_E4 ")
}

// 0xE5 PUSH HL
func PUSH_0xE5_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xE5 PUSH HL")
}

// 0xE6 AND A_n8
func AND_0xE6_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xE6 AND A_n8")
}

// 0xE7 RST $20
func RST_0xE7_dollar20(CPU *CPU, MMU *MMU) {
	log.Println("0xE7 RST $20")
}

// 0xE8 ADD SP_e8
func ADD_0xE8_SP_e8(CPU *CPU, MMU *MMU) {
	log.Println("0xE8 ADD SP_e8")
}

// 0xE9 JP HL
func JP_0xE9_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xE9 JP HL")
}

// 0xEA LD a16_A
func LD_0xEA_a16_A(CPU *CPU, MMU *MMU) {
	log.Println("0xEA LD a16_A")
}

// 0xEB ILLEGAL_EB
func ILLEGAL_EB_0xEB(CPU *CPU, MMU *MMU) {
	log.Println("0xEB ILLEGAL_EB ")
}

// 0xEC ILLEGAL_EC
func ILLEGAL_EC_0xEC(CPU *CPU, MMU *MMU) {
	log.Println("0xEC ILLEGAL_EC ")
}

// 0xED ILLEGAL_ED
func ILLEGAL_ED_0xED(CPU *CPU, MMU *MMU) {
	log.Println("0xED ILLEGAL_ED ")
}

// 0xEE XOR A_n8
func XOR_0xEE_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xEE XOR A_n8")
}

// 0xEF RST $28
func RST_0xEF_dollar28(CPU *CPU, MMU *MMU) {
	log.Println("0xEF RST $28")
}

// 0xF0 LDH A_a8
func LDH_0xF0_A_a8(CPU *CPU, MMU *MMU) {
	log.Println("0xF0 LDH A_a8")
}

// 0xF1 POP AF
func POP_0xF1_AF(CPU *CPU, MMU *MMU) {
	log.Println("0xF1 POP AF")
}

// 0xF2 LDH A_C
func LDH_0xF2_A_C(CPU *CPU, MMU *MMU) {
	log.Println("0xF2 LDH A_C")
}

// 0xF3 DI
func DI_0xF3(CPU *CPU, MMU *MMU) {
	log.Println("0xF3 DI ")
}

// 0xF4 ILLEGAL_F4
func ILLEGAL_F4_0xF4(CPU *CPU, MMU *MMU) {
	log.Println("0xF4 ILLEGAL_F4 ")
}

// 0xF5 PUSH AF
func PUSH_0xF5_AF(CPU *CPU, MMU *MMU) {
	log.Println("0xF5 PUSH AF")
}

// 0xF6 OR A_n8
func OR_0xF6_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xF6 OR A_n8")
}

// 0xF7 RST $30
func RST_0xF7_dollar30(CPU *CPU, MMU *MMU) {
	log.Println("0xF7 RST $30")
}

// 0xF8 LD HL_SP_e8
func LD_0xF8_HL_SP_e8(CPU *CPU, MMU *MMU) {
	log.Println("0xF8 LD HL_SP_e8")
}

// 0xF9 LD SP_HL
func LD_0xF9_SP_HL(CPU *CPU, MMU *MMU) {
	log.Println("0xF9 LD SP_HL")
}

// 0xFA LD A_a16
func LD_0xFA_A_a16(CPU *CPU, MMU *MMU) {
	log.Println("0xFA LD A_a16")
}

// 0xFB EI
func EI_0xFB(CPU *CPU, MMU *MMU) {
	log.Println("0xFB EI ")
}

// 0xFC ILLEGAL_FC
func ILLEGAL_FC_0xFC(CPU *CPU, MMU *MMU) {
	log.Println("0xFC ILLEGAL_FC ")
}

// 0xFD ILLEGAL_FD
func ILLEGAL_FD_0xFD(CPU *CPU, MMU *MMU) {
	log.Println("0xFD ILLEGAL_FD ")
}

// 0xFE CP A_n8
func CP_0xFE_A_n8(CPU *CPU, MMU *MMU) {
	log.Println("0xFE CP A_n8")
}

// 0xFF RST $38
func RST_0xFF_dollar38(CPU *CPU, MMU *MMU) {
	log.Println("0xFF RST $38")
}
