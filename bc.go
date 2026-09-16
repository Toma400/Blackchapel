package main

// https://github.com/tducasse/ebiten-collisions
// https://github.com/MelonFunction/ebiten-collider

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"

	"bc/core"
)

type Game struct{}

var objects = core.RegisterObjects()
var test_mp = core.ReadInteriorMap("room", objects)
// variables
var varStr map[string]string
var varInt map[string]int
var varFlt map[string]bool

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	ebitenutil.DebugPrint(screen, test_mp.Name)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Baedoor Travels (v.0.1.0)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
