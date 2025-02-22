package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// GAMEBOY struct, all emulation is handled by this.
type GAMEBOY struct {
	CPU *CPU
	MMU *MMU
}

// newGameboy creates a new GAMEBOY and initializes the CPU and MMU.
func newGameboy() *GAMEBOY {
	GAMEBOY := &GAMEBOY{}
	GAMEBOY.CPU = newCPU()
	GAMEBOY.MMU = newMMU()
	return GAMEBOY
}

// Game struct, used by ebiten to render the game.
type Game struct {
	count int
	GAMEBOY GAMEBOY
}

// Update function, called every frame.
func (g *Game) Update() error {
	// Game Logic.
	return nil
}

// Draw function, called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World")
}

// Layout function, called when the window is resized.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}

// Main function, called when the program starts.
func main() {
	// Initialize Ebiten
	ebiten.SetWindowSize(160, 144)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{}

	// Initialize Gameboy Components
	game.GAMEBOY = *newGameboy()

	// Load ROM
	game.GAMEBOY.MMU._rom = RetrieveROM("../roms-for-testing/pokemon-fire-red.gb")

	// Start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
