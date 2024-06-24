package game

import (
	assets "pick-it-up"
	"pick-it-up/internal/libs"
	"pick-it-up/internal/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Entity   objects.Entity
	Skeleton objects.Entity
}

func NewGame() *Game {
	entity := objects.NewEntity(libs.Vector{X: float64(400 - (100*4)/2), Y: float64(300 - (100*4)/2)}, assets.PlayerIdleAnimation)
	skeleton := objects.NewEntity(libs.Vector{X: float64(600 - (100*4)/2), Y: float64(300 - (100*4)/2)}, assets.SkeletonIdleAnimation)
	return &Game{
		Entity:   *entity,
		Skeleton: *skeleton,
	}

}

func (g *Game) Update() error {
	g.Entity.Update()
	g.Skeleton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Entity.Draw(screen)
	g.Skeleton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
