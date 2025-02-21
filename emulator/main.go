package main

import (
	"bufio"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count int
	CPU   *CPU
	MMU   *MMU
}

func RetrieveROM(filename string) []byte {
	file, err := os.Open(filename)

	if err != nil {
		return nil
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes
}

func (g *Game) Update() error {
	var op = g.MMU._rom[g.CPU._r.PC]
	g.CPU._r.PC++
	g.CPU._r.PC &= 0xFFFF
	g.CPU._map[op](g.CPU, g.MMU)
	g.CPU._c.M += g.CPU._r.M
	g.CPU._c.T += g.CPU._r.T
	if g.MMU._inbios && g.CPU._r.PC == 0x0100 {
		g.MMU._inbios = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}

func main() {
	// Initialize Ebiten
	ebiten.SetWindowSize(160, 144)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{}

	// Initialize Gameboy Components
	game.CPU = newCPU()
	game.MMU = newMMU()

	// Load ROM
	gameboy_game := "../roms-for-testing/pokemon-fire-red.gb"
	rom := RetrieveROM(gameboy_game)
	game.MMU._rom = rom

	// Start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
