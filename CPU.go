package main

type CLOCK struct {
	M, T uint8
}

type REGISTER struct {
	A, B, C, D, E, H, L, F uint8
	PC, SP                 uint16
	M, T                   uint8
}

type OPS struct {
	// Unprefixed Opcodes
	NOP_0x00          func(CPU *CPU, MMU *MMU)
	LD_0x01_BC_n16    func(CPU *CPU, MMU *MMU)
	LD_0x02_BC_A      func(CPU *CPU, MMU *MMU)
	INC_0x03_BC       func(CPU *CPU, MMU *MMU)
	INC_0x04_B        func(CPU *CPU, MMU *MMU)
	DEC_0x05_B        func(CPU *CPU, MMU *MMU)
	LD_0x06_B_n8      func(CPU *CPU, MMU *MMU)
	RLCA_0x07         func(CPU *CPU, MMU *MMU)
	LD_0x08_a16_SP    func(CPU *CPU, MMU *MMU)
	ADD_0x09_HL_BC    func(CPU *CPU, MMU *MMU)
	LD_0x0A_A_BC      func(CPU *CPU, MMU *MMU)
	DEC_0x0B_BC       func(CPU *CPU, MMU *MMU)
	INC_0x0C_C        func(CPU *CPU, MMU *MMU)
	DEC_0x0D_C        func(CPU *CPU, MMU *MMU)
	LD_0x0E_C_n8      func(CPU *CPU, MMU *MMU)
	RRCA_0x0F         func(CPU *CPU, MMU *MMU)
	STOP_0x10_n8      func(CPU *CPU, MMU *MMU)
	LD_0x11_DE_n16    func(CPU *CPU, MMU *MMU)
	LD_0x12_DE_A      func(CPU *CPU, MMU *MMU)
	INC_0x13_DE       func(CPU *CPU, MMU *MMU)
	INC_0x14_D        func(CPU *CPU, MMU *MMU)
	DEC_0x15_D        func(CPU *CPU, MMU *MMU)
	LD_0x16_D_n8      func(CPU *CPU, MMU *MMU)
	RLA_0x17          func(CPU *CPU, MMU *MMU)
	JR_0x18_e8        func(CPU *CPU, MMU *MMU)
	ADD_0x19_HL_DE    func(CPU *CPU, MMU *MMU)
	LD_0x1A_A_DE      func(CPU *CPU, MMU *MMU)
	DEC_0x1B_DE       func(CPU *CPU, MMU *MMU)
	INC_0x1C_E        func(CPU *CPU, MMU *MMU)
	DEC_0x1D_E        func(CPU *CPU, MMU *MMU)
	LD_0x1E_E_n8      func(CPU *CPU, MMU *MMU)
	RRA_0x1F          func(CPU *CPU, MMU *MMU)
	JR_0x20_NZ_e8     func(CPU *CPU, MMU *MMU)
	LD_0x21_HL_n16    func(CPU *CPU, MMU *MMU)
	LD_0x22_HL_A      func(CPU *CPU, MMU *MMU)
	INC_0x23_HL       func(CPU *CPU, MMU *MMU)
	INC_0x24_H        func(CPU *CPU, MMU *MMU)
	DEC_0x25_H        func(CPU *CPU, MMU *MMU)
	LD_0x26_H_n8      func(CPU *CPU, MMU *MMU)
	DAA_0x27          func(CPU *CPU, MMU *MMU)
	JR_0x28_Z_e8      func(CPU *CPU, MMU *MMU)
	ADD_0x29_HL_HL    func(CPU *CPU, MMU *MMU)
	LD_0x2A_A_HL      func(CPU *CPU, MMU *MMU)
	DEC_0x2B_HL       func(CPU *CPU, MMU *MMU)
	INC_0x2C_L        func(CPU *CPU, MMU *MMU)
	DEC_0x2D_L        func(CPU *CPU, MMU *MMU)
	LD_0x2E_L_n8      func(CPU *CPU, MMU *MMU)
	CPL_0x2F          func(CPU *CPU, MMU *MMU)
	JR_0x30_NC_e8     func(CPU *CPU, MMU *MMU)
	LD_0x31_SP_n16    func(CPU *CPU, MMU *MMU)
	LD_0x32_HL_A      func(CPU *CPU, MMU *MMU)
	INC_0x33_SP       func(CPU *CPU, MMU *MMU)
	INC_0x34_HL       func(CPU *CPU, MMU *MMU)
	DEC_0x35_HL       func(CPU *CPU, MMU *MMU)
	LD_0x36_HL_n8     func(CPU *CPU, MMU *MMU)
	SCF_0x37          func(CPU *CPU, MMU *MMU)
	JR_0x38_C_e8      func(CPU *CPU, MMU *MMU)
	ADD_0x39_HL_SP    func(CPU *CPU, MMU *MMU)
	LD_0x3A_A_HL      func(CPU *CPU, MMU *MMU)
	DEC_0x3B_SP       func(CPU *CPU, MMU *MMU)
	INC_0x3C_A        func(CPU *CPU, MMU *MMU)
	DEC_0x3D_A        func(CPU *CPU, MMU *MMU)
	LD_0x3E_A_n8      func(CPU *CPU, MMU *MMU)
	CCF_0x3F          func(CPU *CPU, MMU *MMU)
	LD_0x40_B_B       func(CPU *CPU, MMU *MMU)
	LD_0x41_B_C       func(CPU *CPU, MMU *MMU)
	LD_0x42_B_D       func(CPU *CPU, MMU *MMU)
	LD_0x43_B_E       func(CPU *CPU, MMU *MMU)
	LD_0x44_B_H       func(CPU *CPU, MMU *MMU)
	LD_0x45_B_L       func(CPU *CPU, MMU *MMU)
	LD_0x46_B_HL      func(CPU *CPU, MMU *MMU)
	LD_0x47_B_A       func(CPU *CPU, MMU *MMU)
	LD_0x48_C_B       func(CPU *CPU, MMU *MMU)
	LD_0x49_C_C       func(CPU *CPU, MMU *MMU)
	LD_0x4A_C_D       func(CPU *CPU, MMU *MMU)
	LD_0x4B_C_E       func(CPU *CPU, MMU *MMU)
	LD_0x4C_C_H       func(CPU *CPU, MMU *MMU)
	LD_0x4D_C_L       func(CPU *CPU, MMU *MMU)
	LD_0x4E_C_HL      func(CPU *CPU, MMU *MMU)
	LD_0x4F_C_A       func(CPU *CPU, MMU *MMU)
	LD_0x50_D_B       func(CPU *CPU, MMU *MMU)
	LD_0x51_D_C       func(CPU *CPU, MMU *MMU)
	LD_0x52_D_D       func(CPU *CPU, MMU *MMU)
	LD_0x53_D_E       func(CPU *CPU, MMU *MMU)
	LD_0x54_D_H       func(CPU *CPU, MMU *MMU)
	LD_0x55_D_L       func(CPU *CPU, MMU *MMU)
	LD_0x56_D_HL      func(CPU *CPU, MMU *MMU)
	LD_0x57_D_A       func(CPU *CPU, MMU *MMU)
	LD_0x58_E_B       func(CPU *CPU, MMU *MMU)
	LD_0x59_E_C       func(CPU *CPU, MMU *MMU)
	LD_0x5A_E_D       func(CPU *CPU, MMU *MMU)
	LD_0x5B_E_E       func(CPU *CPU, MMU *MMU)
	LD_0x5C_E_H       func(CPU *CPU, MMU *MMU)
	LD_0x5D_E_L       func(CPU *CPU, MMU *MMU)
	LD_0x5E_E_HL      func(CPU *CPU, MMU *MMU)
	LD_0x5F_E_A       func(CPU *CPU, MMU *MMU)
	LD_0x60_H_B       func(CPU *CPU, MMU *MMU)
	LD_0x61_H_C       func(CPU *CPU, MMU *MMU)
	LD_0x62_H_D       func(CPU *CPU, MMU *MMU)
	LD_0x63_H_E       func(CPU *CPU, MMU *MMU)
	LD_0x64_H_H       func(CPU *CPU, MMU *MMU)
	LD_0x65_H_L       func(CPU *CPU, MMU *MMU)
	LD_0x66_H_HL      func(CPU *CPU, MMU *MMU)
	LD_0x67_H_A       func(CPU *CPU, MMU *MMU)
	LD_0x68_L_B       func(CPU *CPU, MMU *MMU)
	LD_0x69_L_C       func(CPU *CPU, MMU *MMU)
	LD_0x6A_L_D       func(CPU *CPU, MMU *MMU)
	LD_0x6B_L_E       func(CPU *CPU, MMU *MMU)
	LD_0x6C_L_H       func(CPU *CPU, MMU *MMU)
	LD_0x6D_L_L       func(CPU *CPU, MMU *MMU)
	LD_0x6E_L_HL      func(CPU *CPU, MMU *MMU)
	LD_0x6F_L_A       func(CPU *CPU, MMU *MMU)
	LD_0x70_HL_B      func(CPU *CPU, MMU *MMU)
	LD_0x71_HL_C      func(CPU *CPU, MMU *MMU)
	LD_0x72_HL_D      func(CPU *CPU, MMU *MMU)
	LD_0x73_HL_E      func(CPU *CPU, MMU *MMU)
	LD_0x74_HL_H      func(CPU *CPU, MMU *MMU)
	LD_0x75_HL_L      func(CPU *CPU, MMU *MMU)
	HALT_0x76         func(CPU *CPU, MMU *MMU)
	LD_0x77_HL_A      func(CPU *CPU, MMU *MMU)
	LD_0x78_A_B       func(CPU *CPU, MMU *MMU)
	LD_0x79_A_C       func(CPU *CPU, MMU *MMU)
	LD_0x7A_A_D       func(CPU *CPU, MMU *MMU)
	LD_0x7B_A_E       func(CPU *CPU, MMU *MMU)
	LD_0x7C_A_H       func(CPU *CPU, MMU *MMU)
	LD_0x7D_A_L       func(CPU *CPU, MMU *MMU)
	LD_0x7E_A_HL      func(CPU *CPU, MMU *MMU)
	LD_0x7F_A_A       func(CPU *CPU, MMU *MMU)
	ADD_0x80_A_B      func(CPU *CPU, MMU *MMU)
	ADD_0x81_A_C      func(CPU *CPU, MMU *MMU)
	ADD_0x82_A_D      func(CPU *CPU, MMU *MMU)
	ADD_0x83_A_E      func(CPU *CPU, MMU *MMU)
	ADD_0x84_A_H      func(CPU *CPU, MMU *MMU)
	ADD_0x85_A_L      func(CPU *CPU, MMU *MMU)
	ADD_0x86_A_HL     func(CPU *CPU, MMU *MMU)
	ADD_0x87_A_A      func(CPU *CPU, MMU *MMU)
	ADC_0x88_A_B      func(CPU *CPU, MMU *MMU)
	ADC_0x89_A_C      func(CPU *CPU, MMU *MMU)
	ADC_0x8A_A_D      func(CPU *CPU, MMU *MMU)
	ADC_0x8B_A_E      func(CPU *CPU, MMU *MMU)
	ADC_0x8C_A_H      func(CPU *CPU, MMU *MMU)
	ADC_0x8D_A_L      func(CPU *CPU, MMU *MMU)
	ADC_0x8E_A_HL     func(CPU *CPU, MMU *MMU)
	ADC_0x8F_A_A      func(CPU *CPU, MMU *MMU)
	SUB_0x90_A_B      func(CPU *CPU, MMU *MMU)
	SUB_0x91_A_C      func(CPU *CPU, MMU *MMU)
	SUB_0x92_A_D      func(CPU *CPU, MMU *MMU)
	SUB_0x93_A_E      func(CPU *CPU, MMU *MMU)
	SUB_0x94_A_H      func(CPU *CPU, MMU *MMU)
	SUB_0x95_A_L      func(CPU *CPU, MMU *MMU)
	SUB_0x96_A_HL     func(CPU *CPU, MMU *MMU)
	SUB_0x97_A_A      func(CPU *CPU, MMU *MMU)
	SBC_0x98_A_B      func(CPU *CPU, MMU *MMU)
	SBC_0x99_A_C      func(CPU *CPU, MMU *MMU)
	SBC_0x9A_A_D      func(CPU *CPU, MMU *MMU)
	SBC_0x9B_A_E      func(CPU *CPU, MMU *MMU)
	SBC_0x9C_A_H      func(CPU *CPU, MMU *MMU)
	SBC_0x9D_A_L      func(CPU *CPU, MMU *MMU)
	SBC_0x9E_A_HL     func(CPU *CPU, MMU *MMU)
	SBC_0x9F_A_A      func(CPU *CPU, MMU *MMU)
	AND_0xA0_A_B      func(CPU *CPU, MMU *MMU)
	AND_0xA1_A_C      func(CPU *CPU, MMU *MMU)
	AND_0xA2_A_D      func(CPU *CPU, MMU *MMU)
	AND_0xA3_A_E      func(CPU *CPU, MMU *MMU)
	AND_0xA4_A_H      func(CPU *CPU, MMU *MMU)
	AND_0xA5_A_L      func(CPU *CPU, MMU *MMU)
	AND_0xA6_A_HL     func(CPU *CPU, MMU *MMU)
	AND_0xA7_A_A      func(CPU *CPU, MMU *MMU)
	XOR_0xA8_A_B      func(CPU *CPU, MMU *MMU)
	XOR_0xA9_A_C      func(CPU *CPU, MMU *MMU)
	XOR_0xAA_A_D      func(CPU *CPU, MMU *MMU)
	XOR_0xAB_A_E      func(CPU *CPU, MMU *MMU)
	XOR_0xAC_A_H      func(CPU *CPU, MMU *MMU)
	XOR_0xAD_A_L      func(CPU *CPU, MMU *MMU)
	XOR_0xAE_A_HL     func(CPU *CPU, MMU *MMU)
	XOR_0xAF_A_A      func(CPU *CPU, MMU *MMU)
	OR_0xB0_A_B       func(CPU *CPU, MMU *MMU)
	OR_0xB1_A_C       func(CPU *CPU, MMU *MMU)
	OR_0xB2_A_D       func(CPU *CPU, MMU *MMU)
	OR_0xB3_A_E       func(CPU *CPU, MMU *MMU)
	OR_0xB4_A_H       func(CPU *CPU, MMU *MMU)
	OR_0xB5_A_L       func(CPU *CPU, MMU *MMU)
	OR_0xB6_A_HL      func(CPU *CPU, MMU *MMU)
	OR_0xB7_A_A       func(CPU *CPU, MMU *MMU)
	CP_0xB8_A_B       func(CPU *CPU, MMU *MMU)
	CP_0xB9_A_C       func(CPU *CPU, MMU *MMU)
	CP_0xBA_A_D       func(CPU *CPU, MMU *MMU)
	CP_0xBB_A_E       func(CPU *CPU, MMU *MMU)
	CP_0xBC_A_H       func(CPU *CPU, MMU *MMU)
	CP_0xBD_A_L       func(CPU *CPU, MMU *MMU)
	CP_0xBE_A_HL      func(CPU *CPU, MMU *MMU)
	CP_0xBF_A_A       func(CPU *CPU, MMU *MMU)
	RET_0xC0_NZ       func(CPU *CPU, MMU *MMU)
	POP_0xC1_BC       func(CPU *CPU, MMU *MMU)
	JP_0xC2_NZ_a16    func(CPU *CPU, MMU *MMU)
	JP_0xC3_a16       func(CPU *CPU, MMU *MMU)
	CALL_0xC4_NZ_a16  func(CPU *CPU, MMU *MMU)
	PUSH_0xC5_BC      func(CPU *CPU, MMU *MMU)
	ADD_0xC6_A_n8     func(CPU *CPU, MMU *MMU)
	RST_0xC7_dollar00 func(CPU *CPU, MMU *MMU)
	RET_0xC8_Z        func(CPU *CPU, MMU *MMU)
	RET_0xC9          func(CPU *CPU, MMU *MMU)
	JP_0xCA_Z_a16     func(CPU *CPU, MMU *MMU)
	PREFIX_0xCB       func(CPU *CPU, MMU *MMU)
	CALL_0xCC_Z_a16   func(CPU *CPU, MMU *MMU)
	CALL_0xCD_a16     func(CPU *CPU, MMU *MMU)
	ADC_0xCE_A_n8     func(CPU *CPU, MMU *MMU)
	RST_0xCF_dollar08 func(CPU *CPU, MMU *MMU)
	RET_0xD0_NC       func(CPU *CPU, MMU *MMU)
	POP_0xD1_DE       func(CPU *CPU, MMU *MMU)
	JP_0xD2_NC_a16    func(CPU *CPU, MMU *MMU)
	ILLEGAL_D3_0xD3   func(CPU *CPU, MMU *MMU)
	CALL_0xD4_NC_a16  func(CPU *CPU, MMU *MMU)
	PUSH_0xD5_DE      func(CPU *CPU, MMU *MMU)
	SUB_0xD6_A_n8     func(CPU *CPU, MMU *MMU)
	RST_0xD7_dollar10 func(CPU *CPU, MMU *MMU)
	RET_0xD8_C        func(CPU *CPU, MMU *MMU)
	RETI_0xD9         func(CPU *CPU, MMU *MMU)
	JP_0xDA_C_a16     func(CPU *CPU, MMU *MMU)
	ILLEGAL_DB_0xDB   func(CPU *CPU, MMU *MMU)
	CALL_0xDC_C_a16   func(CPU *CPU, MMU *MMU)
	ILLEGAL_DD_0xDD   func(CPU *CPU, MMU *MMU)
	SBC_0xDE_A_n8     func(CPU *CPU, MMU *MMU)
	RST_0xDF_dollar18 func(CPU *CPU, MMU *MMU)
	LDH_0xE0_a8_A     func(CPU *CPU, MMU *MMU)
	POP_0xE1_HL       func(CPU *CPU, MMU *MMU)
	LDH_0xE2_C_A      func(CPU *CPU, MMU *MMU)
	ILLEGAL_E3_0xE3   func(CPU *CPU, MMU *MMU)
	ILLEGAL_E4_0xE4   func(CPU *CPU, MMU *MMU)
	PUSH_0xE5_HL      func(CPU *CPU, MMU *MMU)
	AND_0xE6_A_n8     func(CPU *CPU, MMU *MMU)
	RST_0xE7_dollar20 func(CPU *CPU, MMU *MMU)
	ADD_0xE8_SP_e8    func(CPU *CPU, MMU *MMU)
	JP_0xE9_HL        func(CPU *CPU, MMU *MMU)
	LD_0xEA_a16_A     func(CPU *CPU, MMU *MMU)
	ILLEGAL_EB_0xEB   func(CPU *CPU, MMU *MMU)
	ILLEGAL_EC_0xEC   func(CPU *CPU, MMU *MMU)
	ILLEGAL_ED_0xED   func(CPU *CPU, MMU *MMU)
	XOR_0xEE_A_n8     func(CPU *CPU, MMU *MMU)
	RST_0xEF_dollar28 func(CPU *CPU, MMU *MMU)
	LDH_0xF0_A_a8     func(CPU *CPU, MMU *MMU)
	POP_0xF1_AF       func(CPU *CPU, MMU *MMU)
	LDH_0xF2_A_C      func(CPU *CPU, MMU *MMU)
	DI_0xF3           func(CPU *CPU, MMU *MMU)
	ILLEGAL_F4_0xF4   func(CPU *CPU, MMU *MMU)
	PUSH_0xF5_AF      func(CPU *CPU, MMU *MMU)
	OR_0xF6_A_n8      func(CPU *CPU, MMU *MMU)
	RST_0xF7_dollar30 func(CPU *CPU, MMU *MMU)
	LD_0xF8_HL_SP_e8  func(CPU *CPU, MMU *MMU)
	LD_0xF9_SP_HL     func(CPU *CPU, MMU *MMU)
	LD_0xFA_A_a16     func(CPU *CPU, MMU *MMU)
	EI_0xFB           func(CPU *CPU, MMU *MMU)
	ILLEGAL_FC_0xFC   func(CPU *CPU, MMU *MMU)
	ILLEGAL_FD_0xFD   func(CPU *CPU, MMU *MMU)
	CP_0xFE_A_n8      func(CPU *CPU, MMU *MMU)
	RST_0xFF_dollar38 func(CPU *CPU, MMU *MMU)

	// CBPrefixed Opcodes
	CB_RLC_0x00_B    func(CPU *CPU, MMU *MMU)
	CB_RLC_0x01_C    func(CPU *CPU, MMU *MMU)
	CB_RLC_0x02_D    func(CPU *CPU, MMU *MMU)
	CB_RLC_0x03_E    func(CPU *CPU, MMU *MMU)
	CB_RLC_0x04_H    func(CPU *CPU, MMU *MMU)
	CB_RLC_0x05_L    func(CPU *CPU, MMU *MMU)
	CB_RLC_0x06_HL   func(CPU *CPU, MMU *MMU)
	CB_RLC_0x07_A    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x08_B    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x09_C    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x0A_D    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x0B_E    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x0C_H    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x0D_L    func(CPU *CPU, MMU *MMU)
	CB_RRC_0x0E_HL   func(CPU *CPU, MMU *MMU)
	CB_RRC_0x0F_A    func(CPU *CPU, MMU *MMU)
	CB_RL_0x10_B     func(CPU *CPU, MMU *MMU)
	CB_RL_0x11_C     func(CPU *CPU, MMU *MMU)
	CB_RL_0x12_D     func(CPU *CPU, MMU *MMU)
	CB_RL_0x13_E     func(CPU *CPU, MMU *MMU)
	CB_RL_0x14_H     func(CPU *CPU, MMU *MMU)
	CB_RL_0x15_L     func(CPU *CPU, MMU *MMU)
	CB_RL_0x16_HL    func(CPU *CPU, MMU *MMU)
	CB_RL_0x17_A     func(CPU *CPU, MMU *MMU)
	CB_RR_0x18_B     func(CPU *CPU, MMU *MMU)
	CB_RR_0x19_C     func(CPU *CPU, MMU *MMU)
	CB_RR_0x1A_D     func(CPU *CPU, MMU *MMU)
	CB_RR_0x1B_E     func(CPU *CPU, MMU *MMU)
	CB_RR_0x1C_H     func(CPU *CPU, MMU *MMU)
	CB_RR_0x1D_L     func(CPU *CPU, MMU *MMU)
	CB_RR_0x1E_HL    func(CPU *CPU, MMU *MMU)
	CB_RR_0x1F_A     func(CPU *CPU, MMU *MMU)
	CB_SLA_0x20_B    func(CPU *CPU, MMU *MMU)
	CB_SLA_0x21_C    func(CPU *CPU, MMU *MMU)
	CB_SLA_0x22_D    func(CPU *CPU, MMU *MMU)
	CB_SLA_0x23_E    func(CPU *CPU, MMU *MMU)
	CB_SLA_0x24_H    func(CPU *CPU, MMU *MMU)
	CB_SLA_0x25_L    func(CPU *CPU, MMU *MMU)
	CB_SLA_0x26_HL   func(CPU *CPU, MMU *MMU)
	CB_SLA_0x27_A    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x28_B    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x29_C    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x2A_D    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x2B_E    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x2C_H    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x2D_L    func(CPU *CPU, MMU *MMU)
	CB_SRA_0x2E_HL   func(CPU *CPU, MMU *MMU)
	CB_SRA_0x2F_A    func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x30_B   func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x31_C   func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x32_D   func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x33_E   func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x34_H   func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x35_L   func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x36_HL  func(CPU *CPU, MMU *MMU)
	CB_SWAP_0x37_A   func(CPU *CPU, MMU *MMU)
	CB_SRL_0x38_B    func(CPU *CPU, MMU *MMU)
	CB_SRL_0x39_C    func(CPU *CPU, MMU *MMU)
	CB_SRL_0x3A_D    func(CPU *CPU, MMU *MMU)
	CB_SRL_0x3B_E    func(CPU *CPU, MMU *MMU)
	CB_SRL_0x3C_H    func(CPU *CPU, MMU *MMU)
	CB_SRL_0x3D_L    func(CPU *CPU, MMU *MMU)
	CB_SRL_0x3E_HL   func(CPU *CPU, MMU *MMU)
	CB_SRL_0x3F_A    func(CPU *CPU, MMU *MMU)
	CB_BIT_0x40_0_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x41_0_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x42_0_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x43_0_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x44_0_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x45_0_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x46_0_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x47_0_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x48_1_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x49_1_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x4A_1_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x4B_1_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x4C_1_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x4D_1_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x4E_1_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x4F_1_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x50_2_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x51_2_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x52_2_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x53_2_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x54_2_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x55_2_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x56_2_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x57_2_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x58_3_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x59_3_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x5A_3_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x5B_3_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x5C_3_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x5D_3_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x5E_3_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x5F_3_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x60_4_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x61_4_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x62_4_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x63_4_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x64_4_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x65_4_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x66_4_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x67_4_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x68_5_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x69_5_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x6A_5_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x6B_5_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x6C_5_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x6D_5_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x6E_5_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x6F_5_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x70_6_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x71_6_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x72_6_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x73_6_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x74_6_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x75_6_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x76_6_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x77_6_A  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x78_7_B  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x79_7_C  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x7A_7_D  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x7B_7_E  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x7C_7_H  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x7D_7_L  func(CPU *CPU, MMU *MMU)
	CB_BIT_0x7E_7_HL func(CPU *CPU, MMU *MMU)
	CB_BIT_0x7F_7_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0x80_0_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0x81_0_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0x82_0_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0x83_0_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0x84_0_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0x85_0_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0x86_0_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0x87_0_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0x88_1_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0x89_1_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0x8A_1_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0x8B_1_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0x8C_1_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0x8D_1_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0x8E_1_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0x8F_1_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0x90_2_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0x91_2_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0x92_2_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0x93_2_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0x94_2_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0x95_2_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0x96_2_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0x97_2_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0x98_3_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0x99_3_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0x9A_3_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0x9B_3_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0x9C_3_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0x9D_3_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0x9E_3_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0x9F_3_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA0_4_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA1_4_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA2_4_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA3_4_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA4_4_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA5_4_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA6_4_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0xA7_4_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA8_5_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0xA9_5_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0xAA_5_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0xAB_5_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0xAC_5_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0xAD_5_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0xAE_5_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0xAF_5_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB0_6_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB1_6_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB2_6_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB3_6_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB4_6_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB5_6_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB6_6_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0xB7_6_A  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB8_7_B  func(CPU *CPU, MMU *MMU)
	CB_RES_0xB9_7_C  func(CPU *CPU, MMU *MMU)
	CB_RES_0xBA_7_D  func(CPU *CPU, MMU *MMU)
	CB_RES_0xBB_7_E  func(CPU *CPU, MMU *MMU)
	CB_RES_0xBC_7_H  func(CPU *CPU, MMU *MMU)
	CB_RES_0xBD_7_L  func(CPU *CPU, MMU *MMU)
	CB_RES_0xBE_7_HL func(CPU *CPU, MMU *MMU)
	CB_RES_0xBF_7_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC0_0_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC1_0_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC2_0_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC3_0_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC4_0_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC5_0_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC6_0_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xC7_0_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC8_1_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xC9_1_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xCA_1_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xCB_1_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xCC_1_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xCD_1_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xCE_1_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xCF_1_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD0_2_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD1_2_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD2_2_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD3_2_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD4_2_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD5_2_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD6_2_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xD7_2_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD8_3_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xD9_3_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xDA_3_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xDB_3_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xDC_3_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xDD_3_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xDE_3_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xDF_3_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE0_4_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE1_4_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE2_4_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE3_4_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE4_4_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE5_4_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE6_4_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xE7_4_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE8_5_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xE9_5_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xEA_5_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xEB_5_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xEC_5_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xED_5_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xEE_5_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xEF_5_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF0_6_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF1_6_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF2_6_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF3_6_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF4_6_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF5_6_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF6_6_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xF7_6_A  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF8_7_B  func(CPU *CPU, MMU *MMU)
	CB_SET_0xF9_7_C  func(CPU *CPU, MMU *MMU)
	CB_SET_0xFA_7_D  func(CPU *CPU, MMU *MMU)
	CB_SET_0xFB_7_E  func(CPU *CPU, MMU *MMU)
	CB_SET_0xFC_7_H  func(CPU *CPU, MMU *MMU)
	CB_SET_0xFD_7_L  func(CPU *CPU, MMU *MMU)
	CB_SET_0xFE_7_HL func(CPU *CPU, MMU *MMU)
	CB_SET_0xFF_7_A  func(CPU *CPU, MMU *MMU)
}

type CPU struct {
	_r REGISTER
	_c CLOCK

	reset func(CPU *CPU)

	_ops   OPS
	_map   [256]func(CPU *CPU, MMU *MMU)
	_cbmap [256]func(CPU *CPU, MMU *MMU)
}

func reset(CPU *CPU) {
	// Reset Registers
	CPU._r.A = 0
	CPU._r.B = 0
	CPU._r.C = 0
	CPU._r.D = 0
	CPU._r.E = 0
	CPU._r.H = 0
	CPU._r.L = 0
	CPU._r.F = 0
	CPU._r.PC = 0
	CPU._r.SP = 0
	CPU._r.M = 0
	CPU._r.T = 0

	// Reset Clock
	CPU._c.M = 0
	CPU._c.T = 0
}

func newCPU() *CPU {
	CPU := &CPU{}

	CPU._r = REGISTER{
		A:  0,
		B:  0,
		C:  0,
		D:  0,
		E:  0,
		H:  0,
		L:  0,
		F:  0,
		PC: 0,
		SP: 0,
		M:  0,
		T:  0,
	}

	CPU._c = CLOCK{
		M: 0,
		T: 0,
	}

	CPU.reset = reset

	// Assign Unprefixed functions and map them.
	CPU._ops.NOP_0x00 = NOP_0x00
	CPU._map[0x00] = CPU._ops.NOP_0x00
	CPU._ops.LD_0x01_BC_n16 = LD_0x01_BC_n16
	CPU._map[0x01] = CPU._ops.LD_0x01_BC_n16
	CPU._ops.LD_0x02_BC_A = LD_0x02_BC_A
	CPU._map[0x02] = CPU._ops.LD_0x02_BC_A
	CPU._ops.INC_0x03_BC = INC_0x03_BC
	CPU._map[0x03] = CPU._ops.INC_0x03_BC
	CPU._ops.INC_0x04_B = INC_0x04_B
	CPU._map[0x04] = CPU._ops.INC_0x04_B
	CPU._ops.DEC_0x05_B = DEC_0x05_B
	CPU._map[0x05] = CPU._ops.DEC_0x05_B
	CPU._ops.LD_0x06_B_n8 = LD_0x06_B_n8
	CPU._map[0x06] = CPU._ops.LD_0x06_B_n8
	CPU._ops.RLCA_0x07 = RLCA_0x07
	CPU._map[0x07] = CPU._ops.RLCA_0x07
	CPU._ops.LD_0x08_a16_SP = LD_0x08_a16_SP
	CPU._map[0x08] = CPU._ops.LD_0x08_a16_SP
	CPU._ops.ADD_0x09_HL_BC = ADD_0x09_HL_BC
	CPU._map[0x09] = CPU._ops.ADD_0x09_HL_BC
	CPU._ops.LD_0x0A_A_BC = LD_0x0A_A_BC
	CPU._map[0x0A] = CPU._ops.LD_0x0A_A_BC
	CPU._ops.DEC_0x0B_BC = DEC_0x0B_BC
	CPU._map[0x0B] = CPU._ops.DEC_0x0B_BC
	CPU._ops.INC_0x0C_C = INC_0x0C_C
	CPU._map[0x0C] = CPU._ops.INC_0x0C_C
	CPU._ops.DEC_0x0D_C = DEC_0x0D_C
	CPU._map[0x0D] = CPU._ops.DEC_0x0D_C
	CPU._ops.LD_0x0E_C_n8 = LD_0x0E_C_n8
	CPU._map[0x0E] = CPU._ops.LD_0x0E_C_n8
	CPU._ops.RRCA_0x0F = RRCA_0x0F
	CPU._map[0x0F] = CPU._ops.RRCA_0x0F
	CPU._ops.STOP_0x10_n8 = STOP_0x10_n8
	CPU._map[0x10] = CPU._ops.STOP_0x10_n8
	CPU._ops.LD_0x11_DE_n16 = LD_0x11_DE_n16
	CPU._map[0x11] = CPU._ops.LD_0x11_DE_n16
	CPU._ops.LD_0x12_DE_A = LD_0x12_DE_A
	CPU._map[0x12] = CPU._ops.LD_0x12_DE_A
	CPU._ops.INC_0x13_DE = INC_0x13_DE
	CPU._map[0x13] = CPU._ops.INC_0x13_DE
	CPU._ops.INC_0x14_D = INC_0x14_D
	CPU._map[0x14] = CPU._ops.INC_0x14_D
	CPU._ops.DEC_0x15_D = DEC_0x15_D
	CPU._map[0x15] = CPU._ops.DEC_0x15_D
	CPU._ops.LD_0x16_D_n8 = LD_0x16_D_n8
	CPU._map[0x16] = CPU._ops.LD_0x16_D_n8
	CPU._ops.RLA_0x17 = RLA_0x17
	CPU._map[0x17] = CPU._ops.RLA_0x17
	CPU._ops.JR_0x18_e8 = JR_0x18_e8
	CPU._map[0x18] = CPU._ops.JR_0x18_e8
	CPU._ops.ADD_0x19_HL_DE = ADD_0x19_HL_DE
	CPU._map[0x19] = CPU._ops.ADD_0x19_HL_DE
	CPU._ops.LD_0x1A_A_DE = LD_0x1A_A_DE
	CPU._map[0x1A] = CPU._ops.LD_0x1A_A_DE
	CPU._ops.DEC_0x1B_DE = DEC_0x1B_DE
	CPU._map[0x1B] = CPU._ops.DEC_0x1B_DE
	CPU._ops.INC_0x1C_E = INC_0x1C_E
	CPU._map[0x1C] = CPU._ops.INC_0x1C_E
	CPU._ops.DEC_0x1D_E = DEC_0x1D_E
	CPU._map[0x1D] = CPU._ops.DEC_0x1D_E
	CPU._ops.LD_0x1E_E_n8 = LD_0x1E_E_n8
	CPU._map[0x1E] = CPU._ops.LD_0x1E_E_n8
	CPU._ops.RRA_0x1F = RRA_0x1F
	CPU._map[0x1F] = CPU._ops.RRA_0x1F
	CPU._ops.JR_0x20_NZ_e8 = JR_0x20_NZ_e8
	CPU._map[0x20] = CPU._ops.JR_0x20_NZ_e8
	CPU._ops.LD_0x21_HL_n16 = LD_0x21_HL_n16
	CPU._map[0x21] = CPU._ops.LD_0x21_HL_n16
	CPU._ops.LD_0x22_HL_A = LD_0x22_HL_A
	CPU._map[0x22] = CPU._ops.LD_0x22_HL_A
	CPU._ops.INC_0x23_HL = INC_0x23_HL
	CPU._map[0x23] = CPU._ops.INC_0x23_HL
	CPU._ops.INC_0x24_H = INC_0x24_H
	CPU._map[0x24] = CPU._ops.INC_0x24_H
	CPU._ops.DEC_0x25_H = DEC_0x25_H
	CPU._map[0x25] = CPU._ops.DEC_0x25_H
	CPU._ops.LD_0x26_H_n8 = LD_0x26_H_n8
	CPU._map[0x26] = CPU._ops.LD_0x26_H_n8
	CPU._ops.DAA_0x27 = DAA_0x27
	CPU._map[0x27] = CPU._ops.DAA_0x27
	CPU._ops.JR_0x28_Z_e8 = JR_0x28_Z_e8
	CPU._map[0x28] = CPU._ops.JR_0x28_Z_e8
	CPU._ops.ADD_0x29_HL_HL = ADD_0x29_HL_HL
	CPU._map[0x29] = CPU._ops.ADD_0x29_HL_HL
	CPU._ops.LD_0x2A_A_HL = LD_0x2A_A_HL
	CPU._map[0x2A] = CPU._ops.LD_0x2A_A_HL
	CPU._ops.DEC_0x2B_HL = DEC_0x2B_HL
	CPU._map[0x2B] = CPU._ops.DEC_0x2B_HL
	CPU._ops.INC_0x2C_L = INC_0x2C_L
	CPU._map[0x2C] = CPU._ops.INC_0x2C_L
	CPU._ops.DEC_0x2D_L = DEC_0x2D_L
	CPU._map[0x2D] = CPU._ops.DEC_0x2D_L
	CPU._ops.LD_0x2E_L_n8 = LD_0x2E_L_n8
	CPU._map[0x2E] = CPU._ops.LD_0x2E_L_n8
	CPU._ops.CPL_0x2F = CPL_0x2F
	CPU._map[0x2F] = CPU._ops.CPL_0x2F
	CPU._ops.JR_0x30_NC_e8 = JR_0x30_NC_e8
	CPU._map[0x30] = CPU._ops.JR_0x30_NC_e8
	CPU._ops.LD_0x31_SP_n16 = LD_0x31_SP_n16
	CPU._map[0x31] = CPU._ops.LD_0x31_SP_n16
	CPU._ops.LD_0x32_HL_A = LD_0x32_HL_A
	CPU._map[0x32] = CPU._ops.LD_0x32_HL_A
	CPU._ops.INC_0x33_SP = INC_0x33_SP
	CPU._map[0x33] = CPU._ops.INC_0x33_SP
	CPU._ops.INC_0x34_HL = INC_0x34_HL
	CPU._map[0x34] = CPU._ops.INC_0x34_HL
	CPU._ops.DEC_0x35_HL = DEC_0x35_HL
	CPU._map[0x35] = CPU._ops.DEC_0x35_HL
	CPU._ops.LD_0x36_HL_n8 = LD_0x36_HL_n8
	CPU._map[0x36] = CPU._ops.LD_0x36_HL_n8
	CPU._ops.SCF_0x37 = SCF_0x37
	CPU._map[0x37] = CPU._ops.SCF_0x37
	CPU._ops.JR_0x38_C_e8 = JR_0x38_C_e8
	CPU._map[0x38] = CPU._ops.JR_0x38_C_e8
	CPU._ops.ADD_0x39_HL_SP = ADD_0x39_HL_SP
	CPU._map[0x39] = CPU._ops.ADD_0x39_HL_SP
	CPU._ops.LD_0x3A_A_HL = LD_0x3A_A_HL
	CPU._map[0x3A] = CPU._ops.LD_0x3A_A_HL
	CPU._ops.DEC_0x3B_SP = DEC_0x3B_SP
	CPU._map[0x3B] = CPU._ops.DEC_0x3B_SP
	CPU._ops.INC_0x3C_A = INC_0x3C_A
	CPU._map[0x3C] = CPU._ops.INC_0x3C_A
	CPU._ops.DEC_0x3D_A = DEC_0x3D_A
	CPU._map[0x3D] = CPU._ops.DEC_0x3D_A
	CPU._ops.LD_0x3E_A_n8 = LD_0x3E_A_n8
	CPU._map[0x3E] = CPU._ops.LD_0x3E_A_n8
	CPU._ops.CCF_0x3F = CCF_0x3F
	CPU._map[0x3F] = CPU._ops.CCF_0x3F
	CPU._ops.LD_0x40_B_B = LD_0x40_B_B
	CPU._map[0x40] = CPU._ops.LD_0x40_B_B
	CPU._ops.LD_0x41_B_C = LD_0x41_B_C
	CPU._map[0x41] = CPU._ops.LD_0x41_B_C
	CPU._ops.LD_0x42_B_D = LD_0x42_B_D
	CPU._map[0x42] = CPU._ops.LD_0x42_B_D
	CPU._ops.LD_0x43_B_E = LD_0x43_B_E
	CPU._map[0x43] = CPU._ops.LD_0x43_B_E
	CPU._ops.LD_0x44_B_H = LD_0x44_B_H
	CPU._map[0x44] = CPU._ops.LD_0x44_B_H
	CPU._ops.LD_0x45_B_L = LD_0x45_B_L
	CPU._map[0x45] = CPU._ops.LD_0x45_B_L
	CPU._ops.LD_0x46_B_HL = LD_0x46_B_HL
	CPU._map[0x46] = CPU._ops.LD_0x46_B_HL
	CPU._ops.LD_0x47_B_A = LD_0x47_B_A
	CPU._map[0x47] = CPU._ops.LD_0x47_B_A
	CPU._ops.LD_0x48_C_B = LD_0x48_C_B
	CPU._map[0x48] = CPU._ops.LD_0x48_C_B
	CPU._ops.LD_0x49_C_C = LD_0x49_C_C
	CPU._map[0x49] = CPU._ops.LD_0x49_C_C
	CPU._ops.LD_0x4A_C_D = LD_0x4A_C_D
	CPU._map[0x4A] = CPU._ops.LD_0x4A_C_D
	CPU._ops.LD_0x4B_C_E = LD_0x4B_C_E
	CPU._map[0x4B] = CPU._ops.LD_0x4B_C_E
	CPU._ops.LD_0x4C_C_H = LD_0x4C_C_H
	CPU._map[0x4C] = CPU._ops.LD_0x4C_C_H
	CPU._ops.LD_0x4D_C_L = LD_0x4D_C_L
	CPU._map[0x4D] = CPU._ops.LD_0x4D_C_L
	CPU._ops.LD_0x4E_C_HL = LD_0x4E_C_HL
	CPU._map[0x4E] = CPU._ops.LD_0x4E_C_HL
	CPU._ops.LD_0x4F_C_A = LD_0x4F_C_A
	CPU._map[0x4F] = CPU._ops.LD_0x4F_C_A
	CPU._ops.LD_0x50_D_B = LD_0x50_D_B
	CPU._map[0x50] = CPU._ops.LD_0x50_D_B
	CPU._ops.LD_0x51_D_C = LD_0x51_D_C
	CPU._map[0x51] = CPU._ops.LD_0x51_D_C
	CPU._ops.LD_0x52_D_D = LD_0x52_D_D
	CPU._map[0x52] = CPU._ops.LD_0x52_D_D
	CPU._ops.LD_0x53_D_E = LD_0x53_D_E
	CPU._map[0x53] = CPU._ops.LD_0x53_D_E
	CPU._ops.LD_0x54_D_H = LD_0x54_D_H
	CPU._map[0x54] = CPU._ops.LD_0x54_D_H
	CPU._ops.LD_0x55_D_L = LD_0x55_D_L
	CPU._map[0x55] = CPU._ops.LD_0x55_D_L
	CPU._ops.LD_0x56_D_HL = LD_0x56_D_HL
	CPU._map[0x56] = CPU._ops.LD_0x56_D_HL
	CPU._ops.LD_0x57_D_A = LD_0x57_D_A
	CPU._map[0x57] = CPU._ops.LD_0x57_D_A
	CPU._ops.LD_0x58_E_B = LD_0x58_E_B
	CPU._map[0x58] = CPU._ops.LD_0x58_E_B
	CPU._ops.LD_0x59_E_C = LD_0x59_E_C
	CPU._map[0x59] = CPU._ops.LD_0x59_E_C
	CPU._ops.LD_0x5A_E_D = LD_0x5A_E_D
	CPU._map[0x5A] = CPU._ops.LD_0x5A_E_D
	CPU._ops.LD_0x5B_E_E = LD_0x5B_E_E
	CPU._map[0x5B] = CPU._ops.LD_0x5B_E_E
	CPU._ops.LD_0x5C_E_H = LD_0x5C_E_H
	CPU._map[0x5C] = CPU._ops.LD_0x5C_E_H
	CPU._ops.LD_0x5D_E_L = LD_0x5D_E_L
	CPU._map[0x5D] = CPU._ops.LD_0x5D_E_L
	CPU._ops.LD_0x5E_E_HL = LD_0x5E_E_HL
	CPU._map[0x5E] = CPU._ops.LD_0x5E_E_HL
	CPU._ops.LD_0x5F_E_A = LD_0x5F_E_A
	CPU._map[0x5F] = CPU._ops.LD_0x5F_E_A
	CPU._ops.LD_0x60_H_B = LD_0x60_H_B
	CPU._map[0x60] = CPU._ops.LD_0x60_H_B
	CPU._ops.LD_0x61_H_C = LD_0x61_H_C
	CPU._map[0x61] = CPU._ops.LD_0x61_H_C
	CPU._ops.LD_0x62_H_D = LD_0x62_H_D
	CPU._map[0x62] = CPU._ops.LD_0x62_H_D
	CPU._ops.LD_0x63_H_E = LD_0x63_H_E
	CPU._map[0x63] = CPU._ops.LD_0x63_H_E
	CPU._ops.LD_0x64_H_H = LD_0x64_H_H
	CPU._map[0x64] = CPU._ops.LD_0x64_H_H
	CPU._ops.LD_0x65_H_L = LD_0x65_H_L
	CPU._map[0x65] = CPU._ops.LD_0x65_H_L
	CPU._ops.LD_0x66_H_HL = LD_0x66_H_HL
	CPU._map[0x66] = CPU._ops.LD_0x66_H_HL
	CPU._ops.LD_0x67_H_A = LD_0x67_H_A
	CPU._map[0x67] = CPU._ops.LD_0x67_H_A
	CPU._ops.LD_0x68_L_B = LD_0x68_L_B
	CPU._map[0x68] = CPU._ops.LD_0x68_L_B
	CPU._ops.LD_0x69_L_C = LD_0x69_L_C
	CPU._map[0x69] = CPU._ops.LD_0x69_L_C
	CPU._ops.LD_0x6A_L_D = LD_0x6A_L_D
	CPU._map[0x6A] = CPU._ops.LD_0x6A_L_D
	CPU._ops.LD_0x6B_L_E = LD_0x6B_L_E
	CPU._map[0x6B] = CPU._ops.LD_0x6B_L_E
	CPU._ops.LD_0x6C_L_H = LD_0x6C_L_H
	CPU._map[0x6C] = CPU._ops.LD_0x6C_L_H
	CPU._ops.LD_0x6D_L_L = LD_0x6D_L_L
	CPU._map[0x6D] = CPU._ops.LD_0x6D_L_L
	CPU._ops.LD_0x6E_L_HL = LD_0x6E_L_HL
	CPU._map[0x6E] = CPU._ops.LD_0x6E_L_HL
	CPU._ops.LD_0x6F_L_A = LD_0x6F_L_A
	CPU._map[0x6F] = CPU._ops.LD_0x6F_L_A
	CPU._ops.LD_0x70_HL_B = LD_0x70_HL_B
	CPU._map[0x70] = CPU._ops.LD_0x70_HL_B
	CPU._ops.LD_0x71_HL_C = LD_0x71_HL_C
	CPU._map[0x71] = CPU._ops.LD_0x71_HL_C
	CPU._ops.LD_0x72_HL_D = LD_0x72_HL_D
	CPU._map[0x72] = CPU._ops.LD_0x72_HL_D
	CPU._ops.LD_0x73_HL_E = LD_0x73_HL_E
	CPU._map[0x73] = CPU._ops.LD_0x73_HL_E
	CPU._ops.LD_0x74_HL_H = LD_0x74_HL_H
	CPU._map[0x74] = CPU._ops.LD_0x74_HL_H
	CPU._ops.LD_0x75_HL_L = LD_0x75_HL_L
	CPU._map[0x75] = CPU._ops.LD_0x75_HL_L
	CPU._ops.HALT_0x76 = HALT_0x76
	CPU._map[0x76] = CPU._ops.HALT_0x76
	CPU._ops.LD_0x77_HL_A = LD_0x77_HL_A
	CPU._map[0x77] = CPU._ops.LD_0x77_HL_A
	CPU._ops.LD_0x78_A_B = LD_0x78_A_B
	CPU._map[0x78] = CPU._ops.LD_0x78_A_B
	CPU._ops.LD_0x79_A_C = LD_0x79_A_C
	CPU._map[0x79] = CPU._ops.LD_0x79_A_C
	CPU._ops.LD_0x7A_A_D = LD_0x7A_A_D
	CPU._map[0x7A] = CPU._ops.LD_0x7A_A_D
	CPU._ops.LD_0x7B_A_E = LD_0x7B_A_E
	CPU._map[0x7B] = CPU._ops.LD_0x7B_A_E
	CPU._ops.LD_0x7C_A_H = LD_0x7C_A_H
	CPU._map[0x7C] = CPU._ops.LD_0x7C_A_H
	CPU._ops.LD_0x7D_A_L = LD_0x7D_A_L
	CPU._map[0x7D] = CPU._ops.LD_0x7D_A_L
	CPU._ops.LD_0x7E_A_HL = LD_0x7E_A_HL
	CPU._map[0x7E] = CPU._ops.LD_0x7E_A_HL
	CPU._ops.LD_0x7F_A_A = LD_0x7F_A_A
	CPU._map[0x7F] = CPU._ops.LD_0x7F_A_A
	CPU._ops.ADD_0x80_A_B = ADD_0x80_A_B
	CPU._map[0x80] = CPU._ops.ADD_0x80_A_B
	CPU._ops.ADD_0x81_A_C = ADD_0x81_A_C
	CPU._map[0x81] = CPU._ops.ADD_0x81_A_C
	CPU._ops.ADD_0x82_A_D = ADD_0x82_A_D
	CPU._map[0x82] = CPU._ops.ADD_0x82_A_D
	CPU._ops.ADD_0x83_A_E = ADD_0x83_A_E
	CPU._map[0x83] = CPU._ops.ADD_0x83_A_E
	CPU._ops.ADD_0x84_A_H = ADD_0x84_A_H
	CPU._map[0x84] = CPU._ops.ADD_0x84_A_H
	CPU._ops.ADD_0x85_A_L = ADD_0x85_A_L
	CPU._map[0x85] = CPU._ops.ADD_0x85_A_L
	CPU._ops.ADD_0x86_A_HL = ADD_0x86_A_HL
	CPU._map[0x86] = CPU._ops.ADD_0x86_A_HL
	CPU._ops.ADD_0x87_A_A = ADD_0x87_A_A
	CPU._map[0x87] = CPU._ops.ADD_0x87_A_A
	CPU._ops.ADC_0x88_A_B = ADC_0x88_A_B
	CPU._map[0x88] = CPU._ops.ADC_0x88_A_B
	CPU._ops.ADC_0x89_A_C = ADC_0x89_A_C
	CPU._map[0x89] = CPU._ops.ADC_0x89_A_C
	CPU._ops.ADC_0x8A_A_D = ADC_0x8A_A_D
	CPU._map[0x8A] = CPU._ops.ADC_0x8A_A_D
	CPU._ops.ADC_0x8B_A_E = ADC_0x8B_A_E
	CPU._map[0x8B] = CPU._ops.ADC_0x8B_A_E
	CPU._ops.ADC_0x8C_A_H = ADC_0x8C_A_H
	CPU._map[0x8C] = CPU._ops.ADC_0x8C_A_H
	CPU._ops.ADC_0x8D_A_L = ADC_0x8D_A_L
	CPU._map[0x8D] = CPU._ops.ADC_0x8D_A_L
	CPU._ops.ADC_0x8E_A_HL = ADC_0x8E_A_HL
	CPU._map[0x8E] = CPU._ops.ADC_0x8E_A_HL
	CPU._ops.ADC_0x8F_A_A = ADC_0x8F_A_A
	CPU._map[0x8F] = CPU._ops.ADC_0x8F_A_A
	CPU._ops.SUB_0x90_A_B = SUB_0x90_A_B
	CPU._map[0x90] = CPU._ops.SUB_0x90_A_B
	CPU._ops.SUB_0x91_A_C = SUB_0x91_A_C
	CPU._map[0x91] = CPU._ops.SUB_0x91_A_C
	CPU._ops.SUB_0x92_A_D = SUB_0x92_A_D
	CPU._map[0x92] = CPU._ops.SUB_0x92_A_D
	CPU._ops.SUB_0x93_A_E = SUB_0x93_A_E
	CPU._map[0x93] = CPU._ops.SUB_0x93_A_E
	CPU._ops.SUB_0x94_A_H = SUB_0x94_A_H
	CPU._map[0x94] = CPU._ops.SUB_0x94_A_H
	CPU._ops.SUB_0x95_A_L = SUB_0x95_A_L
	CPU._map[0x95] = CPU._ops.SUB_0x95_A_L
	CPU._ops.SUB_0x96_A_HL = SUB_0x96_A_HL
	CPU._map[0x96] = CPU._ops.SUB_0x96_A_HL
	CPU._ops.SUB_0x97_A_A = SUB_0x97_A_A
	CPU._map[0x97] = CPU._ops.SUB_0x97_A_A
	CPU._ops.SBC_0x98_A_B = SBC_0x98_A_B
	CPU._map[0x98] = CPU._ops.SBC_0x98_A_B
	CPU._ops.SBC_0x99_A_C = SBC_0x99_A_C
	CPU._map[0x99] = CPU._ops.SBC_0x99_A_C
	CPU._ops.SBC_0x9A_A_D = SBC_0x9A_A_D
	CPU._map[0x9A] = CPU._ops.SBC_0x9A_A_D
	CPU._ops.SBC_0x9B_A_E = SBC_0x9B_A_E
	CPU._map[0x9B] = CPU._ops.SBC_0x9B_A_E
	CPU._ops.SBC_0x9C_A_H = SBC_0x9C_A_H
	CPU._map[0x9C] = CPU._ops.SBC_0x9C_A_H
	CPU._ops.SBC_0x9D_A_L = SBC_0x9D_A_L
	CPU._map[0x9D] = CPU._ops.SBC_0x9D_A_L
	CPU._ops.SBC_0x9E_A_HL = SBC_0x9E_A_HL
	CPU._map[0x9E] = CPU._ops.SBC_0x9E_A_HL
	CPU._ops.SBC_0x9F_A_A = SBC_0x9F_A_A
	CPU._map[0x9F] = CPU._ops.SBC_0x9F_A_A
	CPU._ops.AND_0xA0_A_B = AND_0xA0_A_B
	CPU._map[0xA0] = CPU._ops.AND_0xA0_A_B
	CPU._ops.AND_0xA1_A_C = AND_0xA1_A_C
	CPU._map[0xA1] = CPU._ops.AND_0xA1_A_C
	CPU._ops.AND_0xA2_A_D = AND_0xA2_A_D
	CPU._map[0xA2] = CPU._ops.AND_0xA2_A_D
	CPU._ops.AND_0xA3_A_E = AND_0xA3_A_E
	CPU._map[0xA3] = CPU._ops.AND_0xA3_A_E
	CPU._ops.AND_0xA4_A_H = AND_0xA4_A_H
	CPU._map[0xA4] = CPU._ops.AND_0xA4_A_H
	CPU._ops.AND_0xA5_A_L = AND_0xA5_A_L
	CPU._map[0xA5] = CPU._ops.AND_0xA5_A_L
	CPU._ops.AND_0xA6_A_HL = AND_0xA6_A_HL
	CPU._map[0xA6] = CPU._ops.AND_0xA6_A_HL
	CPU._ops.AND_0xA7_A_A = AND_0xA7_A_A
	CPU._map[0xA7] = CPU._ops.AND_0xA7_A_A
	CPU._ops.XOR_0xA8_A_B = XOR_0xA8_A_B
	CPU._map[0xA8] = CPU._ops.XOR_0xA8_A_B
	CPU._ops.XOR_0xA9_A_C = XOR_0xA9_A_C
	CPU._map[0xA9] = CPU._ops.XOR_0xA9_A_C
	CPU._ops.XOR_0xAA_A_D = XOR_0xAA_A_D
	CPU._map[0xAA] = CPU._ops.XOR_0xAA_A_D
	CPU._ops.XOR_0xAB_A_E = XOR_0xAB_A_E
	CPU._map[0xAB] = CPU._ops.XOR_0xAB_A_E
	CPU._ops.XOR_0xAC_A_H = XOR_0xAC_A_H
	CPU._map[0xAC] = CPU._ops.XOR_0xAC_A_H
	CPU._ops.XOR_0xAD_A_L = XOR_0xAD_A_L
	CPU._map[0xAD] = CPU._ops.XOR_0xAD_A_L
	CPU._ops.XOR_0xAE_A_HL = XOR_0xAE_A_HL
	CPU._map[0xAE] = CPU._ops.XOR_0xAE_A_HL
	CPU._ops.XOR_0xAF_A_A = XOR_0xAF_A_A
	CPU._map[0xAF] = CPU._ops.XOR_0xAF_A_A
	CPU._ops.OR_0xB0_A_B = OR_0xB0_A_B
	CPU._map[0xB0] = CPU._ops.OR_0xB0_A_B
	CPU._ops.OR_0xB1_A_C = OR_0xB1_A_C
	CPU._map[0xB1] = CPU._ops.OR_0xB1_A_C
	CPU._ops.OR_0xB2_A_D = OR_0xB2_A_D
	CPU._map[0xB2] = CPU._ops.OR_0xB2_A_D
	CPU._ops.OR_0xB3_A_E = OR_0xB3_A_E
	CPU._map[0xB3] = CPU._ops.OR_0xB3_A_E
	CPU._ops.OR_0xB4_A_H = OR_0xB4_A_H
	CPU._map[0xB4] = CPU._ops.OR_0xB4_A_H
	CPU._ops.OR_0xB5_A_L = OR_0xB5_A_L
	CPU._map[0xB5] = CPU._ops.OR_0xB5_A_L
	CPU._ops.OR_0xB6_A_HL = OR_0xB6_A_HL
	CPU._map[0xB6] = CPU._ops.OR_0xB6_A_HL
	CPU._ops.OR_0xB7_A_A = OR_0xB7_A_A
	CPU._map[0xB7] = CPU._ops.OR_0xB7_A_A
	CPU._ops.CP_0xB8_A_B = CP_0xB8_A_B
	CPU._map[0xB8] = CPU._ops.CP_0xB8_A_B
	CPU._ops.CP_0xB9_A_C = CP_0xB9_A_C
	CPU._map[0xB9] = CPU._ops.CP_0xB9_A_C
	CPU._ops.CP_0xBA_A_D = CP_0xBA_A_D
	CPU._map[0xBA] = CPU._ops.CP_0xBA_A_D
	CPU._ops.CP_0xBB_A_E = CP_0xBB_A_E
	CPU._map[0xBB] = CPU._ops.CP_0xBB_A_E
	CPU._ops.CP_0xBC_A_H = CP_0xBC_A_H
	CPU._map[0xBC] = CPU._ops.CP_0xBC_A_H
	CPU._ops.CP_0xBD_A_L = CP_0xBD_A_L
	CPU._map[0xBD] = CPU._ops.CP_0xBD_A_L
	CPU._ops.CP_0xBE_A_HL = CP_0xBE_A_HL
	CPU._map[0xBE] = CPU._ops.CP_0xBE_A_HL
	CPU._ops.CP_0xBF_A_A = CP_0xBF_A_A
	CPU._map[0xBF] = CPU._ops.CP_0xBF_A_A
	CPU._ops.RET_0xC0_NZ = RET_0xC0_NZ
	CPU._map[0xC0] = CPU._ops.RET_0xC0_NZ
	CPU._ops.POP_0xC1_BC = POP_0xC1_BC
	CPU._map[0xC1] = CPU._ops.POP_0xC1_BC
	CPU._ops.JP_0xC2_NZ_a16 = JP_0xC2_NZ_a16
	CPU._map[0xC2] = CPU._ops.JP_0xC2_NZ_a16
	CPU._ops.JP_0xC3_a16 = JP_0xC3_a16
	CPU._map[0xC3] = CPU._ops.JP_0xC3_a16
	CPU._ops.CALL_0xC4_NZ_a16 = CALL_0xC4_NZ_a16
	CPU._map[0xC4] = CPU._ops.CALL_0xC4_NZ_a16
	CPU._ops.PUSH_0xC5_BC = PUSH_0xC5_BC
	CPU._map[0xC5] = CPU._ops.PUSH_0xC5_BC
	CPU._ops.ADD_0xC6_A_n8 = ADD_0xC6_A_n8
	CPU._map[0xC6] = CPU._ops.ADD_0xC6_A_n8
	CPU._ops.RST_0xC7_dollar00 = RST_0xC7_dollar00
	CPU._map[0xC7] = CPU._ops.RST_0xC7_dollar00
	CPU._ops.RET_0xC8_Z = RET_0xC8_Z
	CPU._map[0xC8] = CPU._ops.RET_0xC8_Z
	CPU._ops.RET_0xC9 = RET_0xC9
	CPU._map[0xC9] = CPU._ops.RET_0xC9
	CPU._ops.JP_0xCA_Z_a16 = JP_0xCA_Z_a16
	CPU._map[0xCA] = CPU._ops.JP_0xCA_Z_a16
	CPU._ops.PREFIX_0xCB = PREFIX_0xCB
	CPU._map[0xCB] = CPU._ops.PREFIX_0xCB
	CPU._ops.CALL_0xCC_Z_a16 = CALL_0xCC_Z_a16
	CPU._map[0xCC] = CPU._ops.CALL_0xCC_Z_a16
	CPU._ops.CALL_0xCD_a16 = CALL_0xCD_a16
	CPU._map[0xCD] = CPU._ops.CALL_0xCD_a16
	CPU._ops.ADC_0xCE_A_n8 = ADC_0xCE_A_n8
	CPU._map[0xCE] = CPU._ops.ADC_0xCE_A_n8
	CPU._ops.RST_0xCF_dollar08 = RST_0xCF_dollar08
	CPU._map[0xCF] = CPU._ops.RST_0xCF_dollar08
	CPU._ops.RET_0xD0_NC = RET_0xD0_NC
	CPU._map[0xD0] = CPU._ops.RET_0xD0_NC
	CPU._ops.POP_0xD1_DE = POP_0xD1_DE
	CPU._map[0xD1] = CPU._ops.POP_0xD1_DE
	CPU._ops.JP_0xD2_NC_a16 = JP_0xD2_NC_a16
	CPU._map[0xD2] = CPU._ops.JP_0xD2_NC_a16
	CPU._ops.ILLEGAL_D3_0xD3 = ILLEGAL_D3_0xD3
	CPU._map[0xD3] = CPU._ops.ILLEGAL_D3_0xD3
	CPU._ops.CALL_0xD4_NC_a16 = CALL_0xD4_NC_a16
	CPU._map[0xD4] = CPU._ops.CALL_0xD4_NC_a16
	CPU._ops.PUSH_0xD5_DE = PUSH_0xD5_DE
	CPU._map[0xD5] = CPU._ops.PUSH_0xD5_DE
	CPU._ops.SUB_0xD6_A_n8 = SUB_0xD6_A_n8
	CPU._map[0xD6] = CPU._ops.SUB_0xD6_A_n8
	CPU._ops.RST_0xD7_dollar10 = RST_0xD7_dollar10
	CPU._map[0xD7] = CPU._ops.RST_0xD7_dollar10
	CPU._ops.RET_0xD8_C = RET_0xD8_C
	CPU._map[0xD8] = CPU._ops.RET_0xD8_C
	CPU._ops.RETI_0xD9 = RETI_0xD9
	CPU._map[0xD9] = CPU._ops.RETI_0xD9
	CPU._ops.JP_0xDA_C_a16 = JP_0xDA_C_a16
	CPU._map[0xDA] = CPU._ops.JP_0xDA_C_a16
	CPU._ops.ILLEGAL_DB_0xDB = ILLEGAL_DB_0xDB
	CPU._map[0xDB] = CPU._ops.ILLEGAL_DB_0xDB
	CPU._ops.CALL_0xDC_C_a16 = CALL_0xDC_C_a16
	CPU._map[0xDC] = CPU._ops.CALL_0xDC_C_a16
	CPU._ops.ILLEGAL_DD_0xDD = ILLEGAL_DD_0xDD
	CPU._map[0xDD] = CPU._ops.ILLEGAL_DD_0xDD
	CPU._ops.SBC_0xDE_A_n8 = SBC_0xDE_A_n8
	CPU._map[0xDE] = CPU._ops.SBC_0xDE_A_n8
	CPU._ops.RST_0xDF_dollar18 = RST_0xDF_dollar18
	CPU._map[0xDF] = CPU._ops.RST_0xDF_dollar18
	CPU._ops.LDH_0xE0_a8_A = LDH_0xE0_a8_A
	CPU._map[0xE0] = CPU._ops.LDH_0xE0_a8_A
	CPU._ops.POP_0xE1_HL = POP_0xE1_HL
	CPU._map[0xE1] = CPU._ops.POP_0xE1_HL
	CPU._ops.LDH_0xE2_C_A = LDH_0xE2_C_A
	CPU._map[0xE2] = CPU._ops.LDH_0xE2_C_A
	CPU._ops.ILLEGAL_E3_0xE3 = ILLEGAL_E3_0xE3
	CPU._map[0xE3] = CPU._ops.ILLEGAL_E3_0xE3
	CPU._ops.ILLEGAL_E4_0xE4 = ILLEGAL_E4_0xE4
	CPU._map[0xE4] = CPU._ops.ILLEGAL_E4_0xE4
	CPU._ops.PUSH_0xE5_HL = PUSH_0xE5_HL
	CPU._map[0xE5] = CPU._ops.PUSH_0xE5_HL
	CPU._ops.AND_0xE6_A_n8 = AND_0xE6_A_n8
	CPU._map[0xE6] = CPU._ops.AND_0xE6_A_n8
	CPU._ops.RST_0xE7_dollar20 = RST_0xE7_dollar20
	CPU._map[0xE7] = CPU._ops.RST_0xE7_dollar20
	CPU._ops.ADD_0xE8_SP_e8 = ADD_0xE8_SP_e8
	CPU._map[0xE8] = CPU._ops.ADD_0xE8_SP_e8
	CPU._ops.JP_0xE9_HL = JP_0xE9_HL
	CPU._map[0xE9] = CPU._ops.JP_0xE9_HL
	CPU._ops.LD_0xEA_a16_A = LD_0xEA_a16_A
	CPU._map[0xEA] = CPU._ops.LD_0xEA_a16_A
	CPU._ops.ILLEGAL_EB_0xEB = ILLEGAL_EB_0xEB
	CPU._map[0xEB] = CPU._ops.ILLEGAL_EB_0xEB
	CPU._ops.ILLEGAL_EC_0xEC = ILLEGAL_EC_0xEC
	CPU._map[0xEC] = CPU._ops.ILLEGAL_EC_0xEC
	CPU._ops.ILLEGAL_ED_0xED = ILLEGAL_ED_0xED
	CPU._map[0xED] = CPU._ops.ILLEGAL_ED_0xED
	CPU._ops.XOR_0xEE_A_n8 = XOR_0xEE_A_n8
	CPU._map[0xEE] = CPU._ops.XOR_0xEE_A_n8
	CPU._ops.RST_0xEF_dollar28 = RST_0xEF_dollar28
	CPU._map[0xEF] = CPU._ops.RST_0xEF_dollar28
	CPU._ops.LDH_0xF0_A_a8 = LDH_0xF0_A_a8
	CPU._map[0xF0] = CPU._ops.LDH_0xF0_A_a8
	CPU._ops.POP_0xF1_AF = POP_0xF1_AF
	CPU._map[0xF1] = CPU._ops.POP_0xF1_AF
	CPU._ops.LDH_0xF2_A_C = LDH_0xF2_A_C
	CPU._map[0xF2] = CPU._ops.LDH_0xF2_A_C
	CPU._ops.DI_0xF3 = DI_0xF3
	CPU._map[0xF3] = CPU._ops.DI_0xF3
	CPU._ops.ILLEGAL_F4_0xF4 = ILLEGAL_F4_0xF4
	CPU._map[0xF4] = CPU._ops.ILLEGAL_F4_0xF4
	CPU._ops.PUSH_0xF5_AF = PUSH_0xF5_AF
	CPU._map[0xF5] = CPU._ops.PUSH_0xF5_AF
	CPU._ops.OR_0xF6_A_n8 = OR_0xF6_A_n8
	CPU._map[0xF6] = CPU._ops.OR_0xF6_A_n8
	CPU._ops.RST_0xF7_dollar30 = RST_0xF7_dollar30
	CPU._map[0xF7] = CPU._ops.RST_0xF7_dollar30
	CPU._ops.LD_0xF8_HL_SP_e8 = LD_0xF8_HL_SP_e8
	CPU._map[0xF8] = CPU._ops.LD_0xF8_HL_SP_e8
	CPU._ops.LD_0xF9_SP_HL = LD_0xF9_SP_HL
	CPU._map[0xF9] = CPU._ops.LD_0xF9_SP_HL
	CPU._ops.LD_0xFA_A_a16 = LD_0xFA_A_a16
	CPU._map[0xFA] = CPU._ops.LD_0xFA_A_a16
	CPU._ops.EI_0xFB = EI_0xFB
	CPU._map[0xFB] = CPU._ops.EI_0xFB
	CPU._ops.ILLEGAL_FC_0xFC = ILLEGAL_FC_0xFC
	CPU._map[0xFC] = CPU._ops.ILLEGAL_FC_0xFC
	CPU._ops.ILLEGAL_FD_0xFD = ILLEGAL_FD_0xFD
	CPU._map[0xFD] = CPU._ops.ILLEGAL_FD_0xFD
	CPU._ops.CP_0xFE_A_n8 = CP_0xFE_A_n8
	CPU._map[0xFE] = CPU._ops.CP_0xFE_A_n8
	CPU._ops.RST_0xFF_dollar38 = RST_0xFF_dollar38
	CPU._map[0xFF] = CPU._ops.RST_0xFF_dollar38

	// Assign CBPrefixed functions and map them.
	CPU._ops.CB_RLC_0x00_B = CB_RLC_0x00_B
	CPU._cbmap[0x00] = CPU._ops.CB_RLC_0x00_B
	CPU._ops.CB_RLC_0x01_C = CB_RLC_0x01_C
	CPU._cbmap[0x01] = CPU._ops.CB_RLC_0x01_C
	CPU._ops.CB_RLC_0x02_D = CB_RLC_0x02_D
	CPU._cbmap[0x02] = CPU._ops.CB_RLC_0x02_D
	CPU._ops.CB_RLC_0x03_E = CB_RLC_0x03_E
	CPU._cbmap[0x03] = CPU._ops.CB_RLC_0x03_E
	CPU._ops.CB_RLC_0x04_H = CB_RLC_0x04_H
	CPU._cbmap[0x04] = CPU._ops.CB_RLC_0x04_H
	CPU._ops.CB_RLC_0x05_L = CB_RLC_0x05_L
	CPU._cbmap[0x05] = CPU._ops.CB_RLC_0x05_L
	CPU._ops.CB_RLC_0x06_HL = CB_RLC_0x06_HL
	CPU._cbmap[0x06] = CPU._ops.CB_RLC_0x06_HL
	CPU._ops.CB_RLC_0x07_A = CB_RLC_0x07_A
	CPU._cbmap[0x07] = CPU._ops.CB_RLC_0x07_A
	CPU._ops.CB_RRC_0x08_B = CB_RRC_0x08_B
	CPU._cbmap[0x08] = CPU._ops.CB_RRC_0x08_B
	CPU._ops.CB_RRC_0x09_C = CB_RRC_0x09_C
	CPU._cbmap[0x09] = CPU._ops.CB_RRC_0x09_C
	CPU._ops.CB_RRC_0x0A_D = CB_RRC_0x0A_D
	CPU._cbmap[0x0A] = CPU._ops.CB_RRC_0x0A_D
	CPU._ops.CB_RRC_0x0B_E = CB_RRC_0x0B_E
	CPU._cbmap[0x0B] = CPU._ops.CB_RRC_0x0B_E
	CPU._ops.CB_RRC_0x0C_H = CB_RRC_0x0C_H
	CPU._cbmap[0x0C] = CPU._ops.CB_RRC_0x0C_H
	CPU._ops.CB_RRC_0x0D_L = CB_RRC_0x0D_L
	CPU._cbmap[0x0D] = CPU._ops.CB_RRC_0x0D_L
	CPU._ops.CB_RRC_0x0E_HL = CB_RRC_0x0E_HL
	CPU._cbmap[0x0E] = CPU._ops.CB_RRC_0x0E_HL
	CPU._ops.CB_RRC_0x0F_A = CB_RRC_0x0F_A
	CPU._cbmap[0x0F] = CPU._ops.CB_RRC_0x0F_A
	CPU._ops.CB_RL_0x10_B = CB_RL_0x10_B
	CPU._cbmap[0x10] = CPU._ops.CB_RL_0x10_B
	CPU._ops.CB_RL_0x11_C = CB_RL_0x11_C
	CPU._cbmap[0x11] = CPU._ops.CB_RL_0x11_C
	CPU._ops.CB_RL_0x12_D = CB_RL_0x12_D
	CPU._cbmap[0x12] = CPU._ops.CB_RL_0x12_D
	CPU._ops.CB_RL_0x13_E = CB_RL_0x13_E
	CPU._cbmap[0x13] = CPU._ops.CB_RL_0x13_E
	CPU._ops.CB_RL_0x14_H = CB_RL_0x14_H
	CPU._cbmap[0x14] = CPU._ops.CB_RL_0x14_H
	CPU._ops.CB_RL_0x15_L = CB_RL_0x15_L
	CPU._cbmap[0x15] = CPU._ops.CB_RL_0x15_L
	CPU._ops.CB_RL_0x16_HL = CB_RL_0x16_HL
	CPU._cbmap[0x16] = CPU._ops.CB_RL_0x16_HL
	CPU._ops.CB_RL_0x17_A = CB_RL_0x17_A
	CPU._cbmap[0x17] = CPU._ops.CB_RL_0x17_A
	CPU._ops.CB_RR_0x18_B = CB_RR_0x18_B
	CPU._cbmap[0x18] = CPU._ops.CB_RR_0x18_B
	CPU._ops.CB_RR_0x19_C = CB_RR_0x19_C
	CPU._cbmap[0x19] = CPU._ops.CB_RR_0x19_C
	CPU._ops.CB_RR_0x1A_D = CB_RR_0x1A_D
	CPU._cbmap[0x1A] = CPU._ops.CB_RR_0x1A_D
	CPU._ops.CB_RR_0x1B_E = CB_RR_0x1B_E
	CPU._cbmap[0x1B] = CPU._ops.CB_RR_0x1B_E
	CPU._ops.CB_RR_0x1C_H = CB_RR_0x1C_H
	CPU._cbmap[0x1C] = CPU._ops.CB_RR_0x1C_H
	CPU._ops.CB_RR_0x1D_L = CB_RR_0x1D_L
	CPU._cbmap[0x1D] = CPU._ops.CB_RR_0x1D_L
	CPU._ops.CB_RR_0x1E_HL = CB_RR_0x1E_HL
	CPU._cbmap[0x1E] = CPU._ops.CB_RR_0x1E_HL
	CPU._ops.CB_RR_0x1F_A = CB_RR_0x1F_A
	CPU._cbmap[0x1F] = CPU._ops.CB_RR_0x1F_A
	CPU._ops.CB_SLA_0x20_B = CB_SLA_0x20_B
	CPU._cbmap[0x20] = CPU._ops.CB_SLA_0x20_B
	CPU._ops.CB_SLA_0x21_C = CB_SLA_0x21_C
	CPU._cbmap[0x21] = CPU._ops.CB_SLA_0x21_C
	CPU._ops.CB_SLA_0x22_D = CB_SLA_0x22_D
	CPU._cbmap[0x22] = CPU._ops.CB_SLA_0x22_D
	CPU._ops.CB_SLA_0x23_E = CB_SLA_0x23_E
	CPU._cbmap[0x23] = CPU._ops.CB_SLA_0x23_E
	CPU._ops.CB_SLA_0x24_H = CB_SLA_0x24_H
	CPU._cbmap[0x24] = CPU._ops.CB_SLA_0x24_H
	CPU._ops.CB_SLA_0x25_L = CB_SLA_0x25_L
	CPU._cbmap[0x25] = CPU._ops.CB_SLA_0x25_L
	CPU._ops.CB_SLA_0x26_HL = CB_SLA_0x26_HL
	CPU._cbmap[0x26] = CPU._ops.CB_SLA_0x26_HL
	CPU._ops.CB_SLA_0x27_A = CB_SLA_0x27_A
	CPU._cbmap[0x27] = CPU._ops.CB_SLA_0x27_A
	CPU._ops.CB_SRA_0x28_B = CB_SRA_0x28_B
	CPU._cbmap[0x28] = CPU._ops.CB_SRA_0x28_B
	CPU._ops.CB_SRA_0x29_C = CB_SRA_0x29_C
	CPU._cbmap[0x29] = CPU._ops.CB_SRA_0x29_C
	CPU._ops.CB_SRA_0x2A_D = CB_SRA_0x2A_D
	CPU._cbmap[0x2A] = CPU._ops.CB_SRA_0x2A_D
	CPU._ops.CB_SRA_0x2B_E = CB_SRA_0x2B_E
	CPU._cbmap[0x2B] = CPU._ops.CB_SRA_0x2B_E
	CPU._ops.CB_SRA_0x2C_H = CB_SRA_0x2C_H
	CPU._cbmap[0x2C] = CPU._ops.CB_SRA_0x2C_H
	CPU._ops.CB_SRA_0x2D_L = CB_SRA_0x2D_L
	CPU._cbmap[0x2D] = CPU._ops.CB_SRA_0x2D_L
	CPU._ops.CB_SRA_0x2E_HL = CB_SRA_0x2E_HL
	CPU._cbmap[0x2E] = CPU._ops.CB_SRA_0x2E_HL
	CPU._ops.CB_SRA_0x2F_A = CB_SRA_0x2F_A
	CPU._cbmap[0x2F] = CPU._ops.CB_SRA_0x2F_A
	CPU._ops.CB_SWAP_0x30_B = CB_SWAP_0x30_B
	CPU._cbmap[0x30] = CPU._ops.CB_SWAP_0x30_B
	CPU._ops.CB_SWAP_0x31_C = CB_SWAP_0x31_C
	CPU._cbmap[0x31] = CPU._ops.CB_SWAP_0x31_C
	CPU._ops.CB_SWAP_0x32_D = CB_SWAP_0x32_D
	CPU._cbmap[0x32] = CPU._ops.CB_SWAP_0x32_D
	CPU._ops.CB_SWAP_0x33_E = CB_SWAP_0x33_E
	CPU._cbmap[0x33] = CPU._ops.CB_SWAP_0x33_E
	CPU._ops.CB_SWAP_0x34_H = CB_SWAP_0x34_H
	CPU._cbmap[0x34] = CPU._ops.CB_SWAP_0x34_H
	CPU._ops.CB_SWAP_0x35_L = CB_SWAP_0x35_L
	CPU._cbmap[0x35] = CPU._ops.CB_SWAP_0x35_L
	CPU._ops.CB_SWAP_0x36_HL = CB_SWAP_0x36_HL
	CPU._cbmap[0x36] = CPU._ops.CB_SWAP_0x36_HL
	CPU._ops.CB_SWAP_0x37_A = CB_SWAP_0x37_A
	CPU._cbmap[0x37] = CPU._ops.CB_SWAP_0x37_A
	CPU._ops.CB_SRL_0x38_B = CB_SRL_0x38_B
	CPU._cbmap[0x38] = CPU._ops.CB_SRL_0x38_B
	CPU._ops.CB_SRL_0x39_C = CB_SRL_0x39_C
	CPU._cbmap[0x39] = CPU._ops.CB_SRL_0x39_C
	CPU._ops.CB_SRL_0x3A_D = CB_SRL_0x3A_D
	CPU._cbmap[0x3A] = CPU._ops.CB_SRL_0x3A_D
	CPU._ops.CB_SRL_0x3B_E = CB_SRL_0x3B_E
	CPU._cbmap[0x3B] = CPU._ops.CB_SRL_0x3B_E
	CPU._ops.CB_SRL_0x3C_H = CB_SRL_0x3C_H
	CPU._cbmap[0x3C] = CPU._ops.CB_SRL_0x3C_H
	CPU._ops.CB_SRL_0x3D_L = CB_SRL_0x3D_L
	CPU._cbmap[0x3D] = CPU._ops.CB_SRL_0x3D_L
	CPU._ops.CB_SRL_0x3E_HL = CB_SRL_0x3E_HL
	CPU._cbmap[0x3E] = CPU._ops.CB_SRL_0x3E_HL
	CPU._ops.CB_SRL_0x3F_A = CB_SRL_0x3F_A
	CPU._cbmap[0x3F] = CPU._ops.CB_SRL_0x3F_A
	CPU._ops.CB_BIT_0x40_0_B = CB_BIT_0x40_0_B
	CPU._cbmap[0x40] = CPU._ops.CB_BIT_0x40_0_B
	CPU._ops.CB_BIT_0x41_0_C = CB_BIT_0x41_0_C
	CPU._cbmap[0x41] = CPU._ops.CB_BIT_0x41_0_C
	CPU._ops.CB_BIT_0x42_0_D = CB_BIT_0x42_0_D
	CPU._cbmap[0x42] = CPU._ops.CB_BIT_0x42_0_D
	CPU._ops.CB_BIT_0x43_0_E = CB_BIT_0x43_0_E
	CPU._cbmap[0x43] = CPU._ops.CB_BIT_0x43_0_E
	CPU._ops.CB_BIT_0x44_0_H = CB_BIT_0x44_0_H
	CPU._cbmap[0x44] = CPU._ops.CB_BIT_0x44_0_H
	CPU._ops.CB_BIT_0x45_0_L = CB_BIT_0x45_0_L
	CPU._cbmap[0x45] = CPU._ops.CB_BIT_0x45_0_L
	CPU._ops.CB_BIT_0x46_0_HL = CB_BIT_0x46_0_HL
	CPU._cbmap[0x46] = CPU._ops.CB_BIT_0x46_0_HL
	CPU._ops.CB_BIT_0x47_0_A = CB_BIT_0x47_0_A
	CPU._cbmap[0x47] = CPU._ops.CB_BIT_0x47_0_A
	CPU._ops.CB_BIT_0x48_1_B = CB_BIT_0x48_1_B
	CPU._cbmap[0x48] = CPU._ops.CB_BIT_0x48_1_B
	CPU._ops.CB_BIT_0x49_1_C = CB_BIT_0x49_1_C
	CPU._cbmap[0x49] = CPU._ops.CB_BIT_0x49_1_C
	CPU._ops.CB_BIT_0x4A_1_D = CB_BIT_0x4A_1_D
	CPU._cbmap[0x4A] = CPU._ops.CB_BIT_0x4A_1_D
	CPU._ops.CB_BIT_0x4B_1_E = CB_BIT_0x4B_1_E
	CPU._cbmap[0x4B] = CPU._ops.CB_BIT_0x4B_1_E
	CPU._ops.CB_BIT_0x4C_1_H = CB_BIT_0x4C_1_H
	CPU._cbmap[0x4C] = CPU._ops.CB_BIT_0x4C_1_H
	CPU._ops.CB_BIT_0x4D_1_L = CB_BIT_0x4D_1_L
	CPU._cbmap[0x4D] = CPU._ops.CB_BIT_0x4D_1_L
	CPU._ops.CB_BIT_0x4E_1_HL = CB_BIT_0x4E_1_HL
	CPU._cbmap[0x4E] = CPU._ops.CB_BIT_0x4E_1_HL
	CPU._ops.CB_BIT_0x4F_1_A = CB_BIT_0x4F_1_A
	CPU._cbmap[0x4F] = CPU._ops.CB_BIT_0x4F_1_A
	CPU._ops.CB_BIT_0x50_2_B = CB_BIT_0x50_2_B
	CPU._cbmap[0x50] = CPU._ops.CB_BIT_0x50_2_B
	CPU._ops.CB_BIT_0x51_2_C = CB_BIT_0x51_2_C
	CPU._cbmap[0x51] = CPU._ops.CB_BIT_0x51_2_C
	CPU._ops.CB_BIT_0x52_2_D = CB_BIT_0x52_2_D
	CPU._cbmap[0x52] = CPU._ops.CB_BIT_0x52_2_D
	CPU._ops.CB_BIT_0x53_2_E = CB_BIT_0x53_2_E
	CPU._cbmap[0x53] = CPU._ops.CB_BIT_0x53_2_E
	CPU._ops.CB_BIT_0x54_2_H = CB_BIT_0x54_2_H
	CPU._cbmap[0x54] = CPU._ops.CB_BIT_0x54_2_H
	CPU._ops.CB_BIT_0x55_2_L = CB_BIT_0x55_2_L
	CPU._cbmap[0x55] = CPU._ops.CB_BIT_0x55_2_L
	CPU._ops.CB_BIT_0x56_2_HL = CB_BIT_0x56_2_HL
	CPU._cbmap[0x56] = CPU._ops.CB_BIT_0x56_2_HL
	CPU._ops.CB_BIT_0x57_2_A = CB_BIT_0x57_2_A
	CPU._cbmap[0x57] = CPU._ops.CB_BIT_0x57_2_A
	CPU._ops.CB_BIT_0x58_3_B = CB_BIT_0x58_3_B
	CPU._cbmap[0x58] = CPU._ops.CB_BIT_0x58_3_B
	CPU._ops.CB_BIT_0x59_3_C = CB_BIT_0x59_3_C
	CPU._cbmap[0x59] = CPU._ops.CB_BIT_0x59_3_C
	CPU._ops.CB_BIT_0x5A_3_D = CB_BIT_0x5A_3_D
	CPU._cbmap[0x5A] = CPU._ops.CB_BIT_0x5A_3_D
	CPU._ops.CB_BIT_0x5B_3_E = CB_BIT_0x5B_3_E
	CPU._cbmap[0x5B] = CPU._ops.CB_BIT_0x5B_3_E
	CPU._ops.CB_BIT_0x5C_3_H = CB_BIT_0x5C_3_H
	CPU._cbmap[0x5C] = CPU._ops.CB_BIT_0x5C_3_H
	CPU._ops.CB_BIT_0x5D_3_L = CB_BIT_0x5D_3_L
	CPU._cbmap[0x5D] = CPU._ops.CB_BIT_0x5D_3_L
	CPU._ops.CB_BIT_0x5E_3_HL = CB_BIT_0x5E_3_HL
	CPU._cbmap[0x5E] = CPU._ops.CB_BIT_0x5E_3_HL
	CPU._ops.CB_BIT_0x5F_3_A = CB_BIT_0x5F_3_A
	CPU._cbmap[0x5F] = CPU._ops.CB_BIT_0x5F_3_A
	CPU._ops.CB_BIT_0x60_4_B = CB_BIT_0x60_4_B
	CPU._cbmap[0x60] = CPU._ops.CB_BIT_0x60_4_B
	CPU._ops.CB_BIT_0x61_4_C = CB_BIT_0x61_4_C
	CPU._cbmap[0x61] = CPU._ops.CB_BIT_0x61_4_C
	CPU._ops.CB_BIT_0x62_4_D = CB_BIT_0x62_4_D
	CPU._cbmap[0x62] = CPU._ops.CB_BIT_0x62_4_D
	CPU._ops.CB_BIT_0x63_4_E = CB_BIT_0x63_4_E
	CPU._cbmap[0x63] = CPU._ops.CB_BIT_0x63_4_E
	CPU._ops.CB_BIT_0x64_4_H = CB_BIT_0x64_4_H
	CPU._cbmap[0x64] = CPU._ops.CB_BIT_0x64_4_H
	CPU._ops.CB_BIT_0x65_4_L = CB_BIT_0x65_4_L
	CPU._cbmap[0x65] = CPU._ops.CB_BIT_0x65_4_L
	CPU._ops.CB_BIT_0x66_4_HL = CB_BIT_0x66_4_HL
	CPU._cbmap[0x66] = CPU._ops.CB_BIT_0x66_4_HL
	CPU._ops.CB_BIT_0x67_4_A = CB_BIT_0x67_4_A
	CPU._cbmap[0x67] = CPU._ops.CB_BIT_0x67_4_A
	CPU._ops.CB_BIT_0x68_5_B = CB_BIT_0x68_5_B
	CPU._cbmap[0x68] = CPU._ops.CB_BIT_0x68_5_B
	CPU._ops.CB_BIT_0x69_5_C = CB_BIT_0x69_5_C
	CPU._cbmap[0x69] = CPU._ops.CB_BIT_0x69_5_C
	CPU._ops.CB_BIT_0x6A_5_D = CB_BIT_0x6A_5_D
	CPU._cbmap[0x6A] = CPU._ops.CB_BIT_0x6A_5_D
	CPU._ops.CB_BIT_0x6B_5_E = CB_BIT_0x6B_5_E
	CPU._cbmap[0x6B] = CPU._ops.CB_BIT_0x6B_5_E
	CPU._ops.CB_BIT_0x6C_5_H = CB_BIT_0x6C_5_H
	CPU._cbmap[0x6C] = CPU._ops.CB_BIT_0x6C_5_H
	CPU._ops.CB_BIT_0x6D_5_L = CB_BIT_0x6D_5_L
	CPU._cbmap[0x6D] = CPU._ops.CB_BIT_0x6D_5_L
	CPU._ops.CB_BIT_0x6E_5_HL = CB_BIT_0x6E_5_HL
	CPU._cbmap[0x6E] = CPU._ops.CB_BIT_0x6E_5_HL
	CPU._ops.CB_BIT_0x6F_5_A = CB_BIT_0x6F_5_A
	CPU._cbmap[0x6F] = CPU._ops.CB_BIT_0x6F_5_A
	CPU._ops.CB_BIT_0x70_6_B = CB_BIT_0x70_6_B
	CPU._cbmap[0x70] = CPU._ops.CB_BIT_0x70_6_B
	CPU._ops.CB_BIT_0x71_6_C = CB_BIT_0x71_6_C
	CPU._cbmap[0x71] = CPU._ops.CB_BIT_0x71_6_C
	CPU._ops.CB_BIT_0x72_6_D = CB_BIT_0x72_6_D
	CPU._cbmap[0x72] = CPU._ops.CB_BIT_0x72_6_D
	CPU._ops.CB_BIT_0x73_6_E = CB_BIT_0x73_6_E
	CPU._cbmap[0x73] = CPU._ops.CB_BIT_0x73_6_E
	CPU._ops.CB_BIT_0x74_6_H = CB_BIT_0x74_6_H
	CPU._cbmap[0x74] = CPU._ops.CB_BIT_0x74_6_H
	CPU._ops.CB_BIT_0x75_6_L = CB_BIT_0x75_6_L
	CPU._cbmap[0x75] = CPU._ops.CB_BIT_0x75_6_L
	CPU._ops.CB_BIT_0x76_6_HL = CB_BIT_0x76_6_HL
	CPU._cbmap[0x76] = CPU._ops.CB_BIT_0x76_6_HL
	CPU._ops.CB_BIT_0x77_6_A = CB_BIT_0x77_6_A
	CPU._cbmap[0x77] = CPU._ops.CB_BIT_0x77_6_A
	CPU._ops.CB_BIT_0x78_7_B = CB_BIT_0x78_7_B
	CPU._cbmap[0x78] = CPU._ops.CB_BIT_0x78_7_B
	CPU._ops.CB_BIT_0x79_7_C = CB_BIT_0x79_7_C
	CPU._cbmap[0x79] = CPU._ops.CB_BIT_0x79_7_C
	CPU._ops.CB_BIT_0x7A_7_D = CB_BIT_0x7A_7_D
	CPU._cbmap[0x7A] = CPU._ops.CB_BIT_0x7A_7_D
	CPU._ops.CB_BIT_0x7B_7_E = CB_BIT_0x7B_7_E
	CPU._cbmap[0x7B] = CPU._ops.CB_BIT_0x7B_7_E
	CPU._ops.CB_BIT_0x7C_7_H = CB_BIT_0x7C_7_H
	CPU._cbmap[0x7C] = CPU._ops.CB_BIT_0x7C_7_H
	CPU._ops.CB_BIT_0x7D_7_L = CB_BIT_0x7D_7_L
	CPU._cbmap[0x7D] = CPU._ops.CB_BIT_0x7D_7_L
	CPU._ops.CB_BIT_0x7E_7_HL = CB_BIT_0x7E_7_HL
	CPU._cbmap[0x7E] = CPU._ops.CB_BIT_0x7E_7_HL
	CPU._ops.CB_BIT_0x7F_7_A = CB_BIT_0x7F_7_A
	CPU._cbmap[0x7F] = CPU._ops.CB_BIT_0x7F_7_A
	CPU._ops.CB_RES_0x80_0_B = CB_RES_0x80_0_B
	CPU._cbmap[0x80] = CPU._ops.CB_RES_0x80_0_B
	CPU._ops.CB_RES_0x81_0_C = CB_RES_0x81_0_C
	CPU._cbmap[0x81] = CPU._ops.CB_RES_0x81_0_C
	CPU._ops.CB_RES_0x82_0_D = CB_RES_0x82_0_D
	CPU._cbmap[0x82] = CPU._ops.CB_RES_0x82_0_D
	CPU._ops.CB_RES_0x83_0_E = CB_RES_0x83_0_E
	CPU._cbmap[0x83] = CPU._ops.CB_RES_0x83_0_E
	CPU._ops.CB_RES_0x84_0_H = CB_RES_0x84_0_H
	CPU._cbmap[0x84] = CPU._ops.CB_RES_0x84_0_H
	CPU._ops.CB_RES_0x85_0_L = CB_RES_0x85_0_L
	CPU._cbmap[0x85] = CPU._ops.CB_RES_0x85_0_L
	CPU._ops.CB_RES_0x86_0_HL = CB_RES_0x86_0_HL
	CPU._cbmap[0x86] = CPU._ops.CB_RES_0x86_0_HL
	CPU._ops.CB_RES_0x87_0_A = CB_RES_0x87_0_A
	CPU._cbmap[0x87] = CPU._ops.CB_RES_0x87_0_A
	CPU._ops.CB_RES_0x88_1_B = CB_RES_0x88_1_B
	CPU._cbmap[0x88] = CPU._ops.CB_RES_0x88_1_B
	CPU._ops.CB_RES_0x89_1_C = CB_RES_0x89_1_C
	CPU._cbmap[0x89] = CPU._ops.CB_RES_0x89_1_C
	CPU._ops.CB_RES_0x8A_1_D = CB_RES_0x8A_1_D
	CPU._cbmap[0x8A] = CPU._ops.CB_RES_0x8A_1_D
	CPU._ops.CB_RES_0x8B_1_E = CB_RES_0x8B_1_E
	CPU._cbmap[0x8B] = CPU._ops.CB_RES_0x8B_1_E
	CPU._ops.CB_RES_0x8C_1_H = CB_RES_0x8C_1_H
	CPU._cbmap[0x8C] = CPU._ops.CB_RES_0x8C_1_H
	CPU._ops.CB_RES_0x8D_1_L = CB_RES_0x8D_1_L
	CPU._cbmap[0x8D] = CPU._ops.CB_RES_0x8D_1_L
	CPU._ops.CB_RES_0x8E_1_HL = CB_RES_0x8E_1_HL
	CPU._cbmap[0x8E] = CPU._ops.CB_RES_0x8E_1_HL
	CPU._ops.CB_RES_0x8F_1_A = CB_RES_0x8F_1_A
	CPU._cbmap[0x8F] = CPU._ops.CB_RES_0x8F_1_A
	CPU._ops.CB_RES_0x90_2_B = CB_RES_0x90_2_B
	CPU._cbmap[0x90] = CPU._ops.CB_RES_0x90_2_B
	CPU._ops.CB_RES_0x91_2_C = CB_RES_0x91_2_C
	CPU._cbmap[0x91] = CPU._ops.CB_RES_0x91_2_C
	CPU._ops.CB_RES_0x92_2_D = CB_RES_0x92_2_D
	CPU._cbmap[0x92] = CPU._ops.CB_RES_0x92_2_D
	CPU._ops.CB_RES_0x93_2_E = CB_RES_0x93_2_E
	CPU._cbmap[0x93] = CPU._ops.CB_RES_0x93_2_E
	CPU._ops.CB_RES_0x94_2_H = CB_RES_0x94_2_H
	CPU._cbmap[0x94] = CPU._ops.CB_RES_0x94_2_H
	CPU._ops.CB_RES_0x95_2_L = CB_RES_0x95_2_L
	CPU._cbmap[0x95] = CPU._ops.CB_RES_0x95_2_L
	CPU._ops.CB_RES_0x96_2_HL = CB_RES_0x96_2_HL
	CPU._cbmap[0x96] = CPU._ops.CB_RES_0x96_2_HL
	CPU._ops.CB_RES_0x97_2_A = CB_RES_0x97_2_A
	CPU._cbmap[0x97] = CPU._ops.CB_RES_0x97_2_A
	CPU._ops.CB_RES_0x98_3_B = CB_RES_0x98_3_B
	CPU._cbmap[0x98] = CPU._ops.CB_RES_0x98_3_B
	CPU._ops.CB_RES_0x99_3_C = CB_RES_0x99_3_C
	CPU._cbmap[0x99] = CPU._ops.CB_RES_0x99_3_C
	CPU._ops.CB_RES_0x9A_3_D = CB_RES_0x9A_3_D
	CPU._cbmap[0x9A] = CPU._ops.CB_RES_0x9A_3_D
	CPU._ops.CB_RES_0x9B_3_E = CB_RES_0x9B_3_E
	CPU._cbmap[0x9B] = CPU._ops.CB_RES_0x9B_3_E
	CPU._ops.CB_RES_0x9C_3_H = CB_RES_0x9C_3_H
	CPU._cbmap[0x9C] = CPU._ops.CB_RES_0x9C_3_H
	CPU._ops.CB_RES_0x9D_3_L = CB_RES_0x9D_3_L
	CPU._cbmap[0x9D] = CPU._ops.CB_RES_0x9D_3_L
	CPU._ops.CB_RES_0x9E_3_HL = CB_RES_0x9E_3_HL
	CPU._cbmap[0x9E] = CPU._ops.CB_RES_0x9E_3_HL
	CPU._ops.CB_RES_0x9F_3_A = CB_RES_0x9F_3_A
	CPU._cbmap[0x9F] = CPU._ops.CB_RES_0x9F_3_A
	CPU._ops.CB_RES_0xA0_4_B = CB_RES_0xA0_4_B
	CPU._cbmap[0xA0] = CPU._ops.CB_RES_0xA0_4_B
	CPU._ops.CB_RES_0xA1_4_C = CB_RES_0xA1_4_C
	CPU._cbmap[0xA1] = CPU._ops.CB_RES_0xA1_4_C
	CPU._ops.CB_RES_0xA2_4_D = CB_RES_0xA2_4_D
	CPU._cbmap[0xA2] = CPU._ops.CB_RES_0xA2_4_D
	CPU._ops.CB_RES_0xA3_4_E = CB_RES_0xA3_4_E
	CPU._cbmap[0xA3] = CPU._ops.CB_RES_0xA3_4_E
	CPU._ops.CB_RES_0xA4_4_H = CB_RES_0xA4_4_H
	CPU._cbmap[0xA4] = CPU._ops.CB_RES_0xA4_4_H
	CPU._ops.CB_RES_0xA5_4_L = CB_RES_0xA5_4_L
	CPU._cbmap[0xA5] = CPU._ops.CB_RES_0xA5_4_L
	CPU._ops.CB_RES_0xA6_4_HL = CB_RES_0xA6_4_HL
	CPU._cbmap[0xA6] = CPU._ops.CB_RES_0xA6_4_HL
	CPU._ops.CB_RES_0xA7_4_A = CB_RES_0xA7_4_A
	CPU._cbmap[0xA7] = CPU._ops.CB_RES_0xA7_4_A
	CPU._ops.CB_RES_0xA8_5_B = CB_RES_0xA8_5_B
	CPU._cbmap[0xA8] = CPU._ops.CB_RES_0xA8_5_B
	CPU._ops.CB_RES_0xA9_5_C = CB_RES_0xA9_5_C
	CPU._cbmap[0xA9] = CPU._ops.CB_RES_0xA9_5_C
	CPU._ops.CB_RES_0xAA_5_D = CB_RES_0xAA_5_D
	CPU._cbmap[0xAA] = CPU._ops.CB_RES_0xAA_5_D
	CPU._ops.CB_RES_0xAB_5_E = CB_RES_0xAB_5_E
	CPU._cbmap[0xAB] = CPU._ops.CB_RES_0xAB_5_E
	CPU._ops.CB_RES_0xAC_5_H = CB_RES_0xAC_5_H
	CPU._cbmap[0xAC] = CPU._ops.CB_RES_0xAC_5_H
	CPU._ops.CB_RES_0xAD_5_L = CB_RES_0xAD_5_L
	CPU._cbmap[0xAD] = CPU._ops.CB_RES_0xAD_5_L
	CPU._ops.CB_RES_0xAE_5_HL = CB_RES_0xAE_5_HL
	CPU._cbmap[0xAE] = CPU._ops.CB_RES_0xAE_5_HL
	CPU._ops.CB_RES_0xAF_5_A = CB_RES_0xAF_5_A
	CPU._cbmap[0xAF] = CPU._ops.CB_RES_0xAF_5_A
	CPU._ops.CB_RES_0xB0_6_B = CB_RES_0xB0_6_B
	CPU._cbmap[0xB0] = CPU._ops.CB_RES_0xB0_6_B
	CPU._ops.CB_RES_0xB1_6_C = CB_RES_0xB1_6_C
	CPU._cbmap[0xB1] = CPU._ops.CB_RES_0xB1_6_C
	CPU._ops.CB_RES_0xB2_6_D = CB_RES_0xB2_6_D
	CPU._cbmap[0xB2] = CPU._ops.CB_RES_0xB2_6_D
	CPU._ops.CB_RES_0xB3_6_E = CB_RES_0xB3_6_E
	CPU._cbmap[0xB3] = CPU._ops.CB_RES_0xB3_6_E
	CPU._ops.CB_RES_0xB4_6_H = CB_RES_0xB4_6_H
	CPU._cbmap[0xB4] = CPU._ops.CB_RES_0xB4_6_H
	CPU._ops.CB_RES_0xB5_6_L = CB_RES_0xB5_6_L
	CPU._cbmap[0xB5] = CPU._ops.CB_RES_0xB5_6_L
	CPU._ops.CB_RES_0xB6_6_HL = CB_RES_0xB6_6_HL
	CPU._cbmap[0xB6] = CPU._ops.CB_RES_0xB6_6_HL
	CPU._ops.CB_RES_0xB7_6_A = CB_RES_0xB7_6_A
	CPU._cbmap[0xB7] = CPU._ops.CB_RES_0xB7_6_A
	CPU._ops.CB_RES_0xB8_7_B = CB_RES_0xB8_7_B
	CPU._cbmap[0xB8] = CPU._ops.CB_RES_0xB8_7_B
	CPU._ops.CB_RES_0xB9_7_C = CB_RES_0xB9_7_C
	CPU._cbmap[0xB9] = CPU._ops.CB_RES_0xB9_7_C
	CPU._ops.CB_RES_0xBA_7_D = CB_RES_0xBA_7_D
	CPU._cbmap[0xBA] = CPU._ops.CB_RES_0xBA_7_D
	CPU._ops.CB_RES_0xBB_7_E = CB_RES_0xBB_7_E
	CPU._cbmap[0xBB] = CPU._ops.CB_RES_0xBB_7_E
	CPU._ops.CB_RES_0xBC_7_H = CB_RES_0xBC_7_H
	CPU._cbmap[0xBC] = CPU._ops.CB_RES_0xBC_7_H
	CPU._ops.CB_RES_0xBD_7_L = CB_RES_0xBD_7_L
	CPU._cbmap[0xBD] = CPU._ops.CB_RES_0xBD_7_L
	CPU._ops.CB_RES_0xBE_7_HL = CB_RES_0xBE_7_HL
	CPU._cbmap[0xBE] = CPU._ops.CB_RES_0xBE_7_HL
	CPU._ops.CB_RES_0xBF_7_A = CB_RES_0xBF_7_A
	CPU._cbmap[0xBF] = CPU._ops.CB_RES_0xBF_7_A
	CPU._ops.CB_SET_0xC0_0_B = CB_SET_0xC0_0_B
	CPU._cbmap[0xC0] = CPU._ops.CB_SET_0xC0_0_B
	CPU._ops.CB_SET_0xC1_0_C = CB_SET_0xC1_0_C
	CPU._cbmap[0xC1] = CPU._ops.CB_SET_0xC1_0_C
	CPU._ops.CB_SET_0xC2_0_D = CB_SET_0xC2_0_D
	CPU._cbmap[0xC2] = CPU._ops.CB_SET_0xC2_0_D
	CPU._ops.CB_SET_0xC3_0_E = CB_SET_0xC3_0_E
	CPU._cbmap[0xC3] = CPU._ops.CB_SET_0xC3_0_E
	CPU._ops.CB_SET_0xC4_0_H = CB_SET_0xC4_0_H
	CPU._cbmap[0xC4] = CPU._ops.CB_SET_0xC4_0_H
	CPU._ops.CB_SET_0xC5_0_L = CB_SET_0xC5_0_L
	CPU._cbmap[0xC5] = CPU._ops.CB_SET_0xC5_0_L
	CPU._ops.CB_SET_0xC6_0_HL = CB_SET_0xC6_0_HL
	CPU._cbmap[0xC6] = CPU._ops.CB_SET_0xC6_0_HL
	CPU._ops.CB_SET_0xC7_0_A = CB_SET_0xC7_0_A
	CPU._cbmap[0xC7] = CPU._ops.CB_SET_0xC7_0_A
	CPU._ops.CB_SET_0xC8_1_B = CB_SET_0xC8_1_B
	CPU._cbmap[0xC8] = CPU._ops.CB_SET_0xC8_1_B
	CPU._ops.CB_SET_0xC9_1_C = CB_SET_0xC9_1_C
	CPU._cbmap[0xC9] = CPU._ops.CB_SET_0xC9_1_C
	CPU._ops.CB_SET_0xCA_1_D = CB_SET_0xCA_1_D
	CPU._cbmap[0xCA] = CPU._ops.CB_SET_0xCA_1_D
	CPU._ops.CB_SET_0xCB_1_E = CB_SET_0xCB_1_E
	CPU._cbmap[0xCB] = CPU._ops.CB_SET_0xCB_1_E
	CPU._ops.CB_SET_0xCC_1_H = CB_SET_0xCC_1_H
	CPU._cbmap[0xCC] = CPU._ops.CB_SET_0xCC_1_H
	CPU._ops.CB_SET_0xCD_1_L = CB_SET_0xCD_1_L
	CPU._cbmap[0xCD] = CPU._ops.CB_SET_0xCD_1_L
	CPU._ops.CB_SET_0xCE_1_HL = CB_SET_0xCE_1_HL
	CPU._cbmap[0xCE] = CPU._ops.CB_SET_0xCE_1_HL
	CPU._ops.CB_SET_0xCF_1_A = CB_SET_0xCF_1_A
	CPU._cbmap[0xCF] = CPU._ops.CB_SET_0xCF_1_A
	CPU._ops.CB_SET_0xD0_2_B = CB_SET_0xD0_2_B
	CPU._cbmap[0xD0] = CPU._ops.CB_SET_0xD0_2_B
	CPU._ops.CB_SET_0xD1_2_C = CB_SET_0xD1_2_C
	CPU._cbmap[0xD1] = CPU._ops.CB_SET_0xD1_2_C
	CPU._ops.CB_SET_0xD2_2_D = CB_SET_0xD2_2_D
	CPU._cbmap[0xD2] = CPU._ops.CB_SET_0xD2_2_D
	CPU._ops.CB_SET_0xD3_2_E = CB_SET_0xD3_2_E
	CPU._cbmap[0xD3] = CPU._ops.CB_SET_0xD3_2_E
	CPU._ops.CB_SET_0xD4_2_H = CB_SET_0xD4_2_H
	CPU._cbmap[0xD4] = CPU._ops.CB_SET_0xD4_2_H
	CPU._ops.CB_SET_0xD5_2_L = CB_SET_0xD5_2_L
	CPU._cbmap[0xD5] = CPU._ops.CB_SET_0xD5_2_L
	CPU._ops.CB_SET_0xD6_2_HL = CB_SET_0xD6_2_HL
	CPU._cbmap[0xD6] = CPU._ops.CB_SET_0xD6_2_HL
	CPU._ops.CB_SET_0xD7_2_A = CB_SET_0xD7_2_A
	CPU._cbmap[0xD7] = CPU._ops.CB_SET_0xD7_2_A
	CPU._ops.CB_SET_0xD8_3_B = CB_SET_0xD8_3_B
	CPU._cbmap[0xD8] = CPU._ops.CB_SET_0xD8_3_B
	CPU._ops.CB_SET_0xD9_3_C = CB_SET_0xD9_3_C
	CPU._cbmap[0xD9] = CPU._ops.CB_SET_0xD9_3_C
	CPU._ops.CB_SET_0xDA_3_D = CB_SET_0xDA_3_D
	CPU._cbmap[0xDA] = CPU._ops.CB_SET_0xDA_3_D
	CPU._ops.CB_SET_0xDB_3_E = CB_SET_0xDB_3_E
	CPU._cbmap[0xDB] = CPU._ops.CB_SET_0xDB_3_E
	CPU._ops.CB_SET_0xDC_3_H = CB_SET_0xDC_3_H
	CPU._cbmap[0xDC] = CPU._ops.CB_SET_0xDC_3_H
	CPU._ops.CB_SET_0xDD_3_L = CB_SET_0xDD_3_L
	CPU._cbmap[0xDD] = CPU._ops.CB_SET_0xDD_3_L
	CPU._ops.CB_SET_0xDE_3_HL = CB_SET_0xDE_3_HL
	CPU._cbmap[0xDE] = CPU._ops.CB_SET_0xDE_3_HL
	CPU._ops.CB_SET_0xDF_3_A = CB_SET_0xDF_3_A
	CPU._cbmap[0xDF] = CPU._ops.CB_SET_0xDF_3_A
	CPU._ops.CB_SET_0xE0_4_B = CB_SET_0xE0_4_B
	CPU._cbmap[0xE0] = CPU._ops.CB_SET_0xE0_4_B
	CPU._ops.CB_SET_0xE1_4_C = CB_SET_0xE1_4_C
	CPU._cbmap[0xE1] = CPU._ops.CB_SET_0xE1_4_C
	CPU._ops.CB_SET_0xE2_4_D = CB_SET_0xE2_4_D
	CPU._cbmap[0xE2] = CPU._ops.CB_SET_0xE2_4_D
	CPU._ops.CB_SET_0xE3_4_E = CB_SET_0xE3_4_E
	CPU._cbmap[0xE3] = CPU._ops.CB_SET_0xE3_4_E
	CPU._ops.CB_SET_0xE4_4_H = CB_SET_0xE4_4_H
	CPU._cbmap[0xE4] = CPU._ops.CB_SET_0xE4_4_H
	CPU._ops.CB_SET_0xE5_4_L = CB_SET_0xE5_4_L
	CPU._cbmap[0xE5] = CPU._ops.CB_SET_0xE5_4_L
	CPU._ops.CB_SET_0xE6_4_HL = CB_SET_0xE6_4_HL
	CPU._cbmap[0xE6] = CPU._ops.CB_SET_0xE6_4_HL
	CPU._ops.CB_SET_0xE7_4_A = CB_SET_0xE7_4_A
	CPU._cbmap[0xE7] = CPU._ops.CB_SET_0xE7_4_A
	CPU._ops.CB_SET_0xE8_5_B = CB_SET_0xE8_5_B
	CPU._cbmap[0xE8] = CPU._ops.CB_SET_0xE8_5_B
	CPU._ops.CB_SET_0xE9_5_C = CB_SET_0xE9_5_C
	CPU._cbmap[0xE9] = CPU._ops.CB_SET_0xE9_5_C
	CPU._ops.CB_SET_0xEA_5_D = CB_SET_0xEA_5_D
	CPU._cbmap[0xEA] = CPU._ops.CB_SET_0xEA_5_D
	CPU._ops.CB_SET_0xEB_5_E = CB_SET_0xEB_5_E
	CPU._cbmap[0xEB] = CPU._ops.CB_SET_0xEB_5_E
	CPU._ops.CB_SET_0xEC_5_H = CB_SET_0xEC_5_H
	CPU._cbmap[0xEC] = CPU._ops.CB_SET_0xEC_5_H
	CPU._ops.CB_SET_0xED_5_L = CB_SET_0xED_5_L
	CPU._cbmap[0xED] = CPU._ops.CB_SET_0xED_5_L
	CPU._ops.CB_SET_0xEE_5_HL = CB_SET_0xEE_5_HL
	CPU._cbmap[0xEE] = CPU._ops.CB_SET_0xEE_5_HL
	CPU._ops.CB_SET_0xEF_5_A = CB_SET_0xEF_5_A
	CPU._cbmap[0xEF] = CPU._ops.CB_SET_0xEF_5_A
	CPU._ops.CB_SET_0xF0_6_B = CB_SET_0xF0_6_B
	CPU._cbmap[0xF0] = CPU._ops.CB_SET_0xF0_6_B
	CPU._ops.CB_SET_0xF1_6_C = CB_SET_0xF1_6_C
	CPU._cbmap[0xF1] = CPU._ops.CB_SET_0xF1_6_C
	CPU._ops.CB_SET_0xF2_6_D = CB_SET_0xF2_6_D
	CPU._cbmap[0xF2] = CPU._ops.CB_SET_0xF2_6_D
	CPU._ops.CB_SET_0xF3_6_E = CB_SET_0xF3_6_E
	CPU._cbmap[0xF3] = CPU._ops.CB_SET_0xF3_6_E
	CPU._ops.CB_SET_0xF4_6_H = CB_SET_0xF4_6_H
	CPU._cbmap[0xF4] = CPU._ops.CB_SET_0xF4_6_H
	CPU._ops.CB_SET_0xF5_6_L = CB_SET_0xF5_6_L
	CPU._cbmap[0xF5] = CPU._ops.CB_SET_0xF5_6_L
	CPU._ops.CB_SET_0xF6_6_HL = CB_SET_0xF6_6_HL
	CPU._cbmap[0xF6] = CPU._ops.CB_SET_0xF6_6_HL
	CPU._ops.CB_SET_0xF7_6_A = CB_SET_0xF7_6_A
	CPU._cbmap[0xF7] = CPU._ops.CB_SET_0xF7_6_A
	CPU._ops.CB_SET_0xF8_7_B = CB_SET_0xF8_7_B
	CPU._cbmap[0xF8] = CPU._ops.CB_SET_0xF8_7_B
	CPU._ops.CB_SET_0xF9_7_C = CB_SET_0xF9_7_C
	CPU._cbmap[0xF9] = CPU._ops.CB_SET_0xF9_7_C
	CPU._ops.CB_SET_0xFA_7_D = CB_SET_0xFA_7_D
	CPU._cbmap[0xFA] = CPU._ops.CB_SET_0xFA_7_D
	CPU._ops.CB_SET_0xFB_7_E = CB_SET_0xFB_7_E
	CPU._cbmap[0xFB] = CPU._ops.CB_SET_0xFB_7_E
	CPU._ops.CB_SET_0xFC_7_H = CB_SET_0xFC_7_H
	CPU._cbmap[0xFC] = CPU._ops.CB_SET_0xFC_7_H
	CPU._ops.CB_SET_0xFD_7_L = CB_SET_0xFD_7_L
	CPU._cbmap[0xFD] = CPU._ops.CB_SET_0xFD_7_L
	CPU._ops.CB_SET_0xFE_7_HL = CB_SET_0xFE_7_HL
	CPU._cbmap[0xFE] = CPU._ops.CB_SET_0xFE_7_HL
	CPU._ops.CB_SET_0xFF_7_A = CB_SET_0xFF_7_A
	CPU._cbmap[0xFF] = CPU._ops.CB_SET_0xFF_7_A

	return CPU
}
