package main

func MMURB(addr uint16) uint8 {
	return 0
}

func MMURW(addr uint16) uint16 {
	return 0
}

func MMUWB(addr uint16, val uint8) {

}

func MMUWW(addr uint16, val uint16) {

}

type MMU struct {
	rb      func(uint16) uint8
	rw      func(uint16) uint16
	wb      func(uint16, uint8)
	ww      func(uint16, uint16)
	_inbios bool
	_rom    []byte
}

func newMMU() *MMU {
	mmu := &MMU{}
	mmu.rb = MMURB
	mmu.rw = MMURW
	mmu.wb = MMUWB
	mmu.ww = MMUWW
	mmu._inbios = true
	return mmu
}
