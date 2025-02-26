package main

type GPU struct {
	_vram []byte
	_oam  []byte
}

func newGPU() *GPU {
	gpu := &GPU{}

	return gpu
}
