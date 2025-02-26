package main

import (
	"log"
	"os"
	"strconv"

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
	GPU *GPU

	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

// newGameboy creates a new GAMEBOY and initializes the CPU and MMU.
func newGameboy() *GAMEBOY {
	GAMEBOY := &GAMEBOY{}

	GAMEBOY.CPU = newCPU()
	GAMEBOY.MMU = newMMU()
	GAMEBOY.GPU = newGPU()

	// Log configuration
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	GAMEBOY.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	GAMEBOY.WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	GAMEBOY.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

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
	ebitenutil.DebugPrint(screen, "In BIOS: "+strconv.FormatBool(g.GAMEBOY.MMU._inbios))
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

	// Start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
