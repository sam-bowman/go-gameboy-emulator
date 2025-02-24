package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var gbResWidth = 160
var gbResHeight = 144
var gbRenderScale = 4

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
	count   int
	GAMEBOY GAMEBOY
}

// Update function, called every frame.
func (g *Game) Update() error {
	if g.GAMEBOY.CPU._r.T == 0 && g.GAMEBOY.CPU._r.M == 0 {
		g.GAMEBOY.CPU.exec(&g.GAMEBOY)
	} else {
		//Reduce Register Clock
		g.GAMEBOY.CPU._r.T -= 1
		if g.GAMEBOY.CPU._r.T%4 == 0 {
			g.GAMEBOY.CPU._r.M -= 1
		}
		//Increase Clock
		g.GAMEBOY.CPU._c.T += 1
		if g.GAMEBOY.CPU._c.T%4 == 0 {
			g.GAMEBOY.CPU._c.M += 1
		}
	}
	return nil
}

// Draw function, called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World")
}

// Layout function, called when the window is resized.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gbResWidth, gbResHeight
}

// Main function, called when the program starts.
func main() {
	// Initialize Ebiten
	ebiten.SetWindowSize(gbResWidth*gbRenderScale, gbResHeight*gbRenderScale)
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
