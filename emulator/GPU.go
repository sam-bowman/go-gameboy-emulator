package main

type GPU struct {
	_vram []byte
	_oam  []byte
}

func newGPU() *GPU {
	gpu := &GPU{}

	gpu._vram = make([]byte, 0xFFFF)
	gpu._oam = make([]byte, 0xFFFF)

	return gpu
}
