package main

import (
	"log"
	"pick-it-up/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 500
	ScreenHeight = 300
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Pick it up!")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
