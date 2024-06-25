package game

import (
	assets "pick-it-up"
	"pick-it-up/internal/libs"
	"pick-it-up/internal/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player objects.Player
}

func NewGame() *Game {
	player := objects.NewPlayer(libs.Vector{X: float64(500 - (100*4)/2), Y: float64(300 - (100*4)/2)}, assets.PlayerAnimations)
	return &Game{
		Player: *player,
	}

}

func (g *Game) Update() error {
	g.Player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
