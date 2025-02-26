package main

func readByte(GB *GAMEBOY, addr uint16) uint8 {
	switch addr {

	case 0x0000: //BIOS & ROM0
		if GB.MMU._inbios {
			if addr < 0x0100 {
				return GB.MMU._bios[addr]
			} else if GB.CPU._r.PC == 0x0100 {
				GB.MMU._inbios = false
			}
			return GB.MMU._rom[addr]
		}

	case 0x1000, 0x2000, 0x3000: //ROM0
		return GB.MMU._rom[addr]

	case 0x4000, 0x5000, 0x6000, 0x7000: //ROM1
		return GB.MMU._rom[addr]

	case 0x8000, 0x9000: //Graphics-VRAM
		return GB.GPU._vram[addr&0x1FFF]

	case 0xA000, 0xB000: //ERAM
		return GB.MMU._eram[addr&0x1FFF]

	case 0xC000, 0xD000: //WRAM
		return GB.MMU._wram[addr&0x1FFF]

	case 0xE000: //WRAM-Shadow
		return GB.MMU._wram[addr&0x1FFF]

	case 0xF000: //WRAM-Shadow, I/O, Zero-Page RAM
		switch addr & 0x0F00 {

		case 0x000, 0x100, 0x200, 0x300, 0x400, 0x500, 0x600: //WRAM-Shadow
			return GB.MMU._wram[addr&0x1FFF]

		case 0x700, 0x800, 0x900, 0xA00, 0xB00, 0xC00, 0xD00: //WRAM-Shadow
			return GB.MMU._wram[addr&0x1FFF]

		case 0xE00: //Graphics-OAM
			if addr < 0xFAE0 {
				return GB.GPU._oam[addr&0xFF]
			} else {
				return 0
			}

		case 0xF00: //Zero-Page
			if addr >= 0xFF80 {
				return GB.MMU._zram[addr&0x7F]
			} else {
				//I/O Handling
				//NEEDS CODE
				return 0
			}
		}
		return GB.MMU._wram[addr&0x1FFF]

	}

	return 0
}

func read2Bytes(GB *GAMEBOY, addr uint16) uint16 {
	return uint16(GB.MMU.readByte(GB, addr)) | uint16(GB.MMU.readByte(GB, addr+1))<<8
}

func writeByte(GB *GAMEBOY, addr uint16, val uint8) {

}

func write2Bytes(GB *GAMEBOY, addr uint16, val uint16) {

}

type MMU struct {
	readByte    func(GB *GAMEBOY, addr uint16) uint8
	read2Bytes  func(GB *GAMEBOY, addr uint16) uint16
	writeByte   func(GB *GAMEBOY, addr uint16, val uint8)
	write2Bytes func(GB *GAMEBOY, addr uint16, val uint16)

	_inbios bool

	_bios []byte
	_rom  []byte
	_wram []byte
	_eram []byte
	_zram []byte
}

func newMMU() *MMU {
	mmu := &MMU{}

	mmu.readByte = readByte
	mmu.read2Bytes = read2Bytes
	mmu.writeByte = writeByte
	mmu.write2Bytes = write2Bytes

	mmu._inbios = true

	return mmu
}
