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
	Lancer   objects.Entity
	Orc      objects.Entity
}

func NewGame() *Game {
	entity := objects.NewEntity(libs.Vector{X: float64(400 - (100*4)/2), Y: float64(200 - (100*4)/2)}, assets.PlayerAnimations)
	skeleton := objects.NewEntity(libs.Vector{X: float64(600 - (100*4)/2), Y: float64(200 - (100*4)/2)}, assets.SkeletonAnimations)
	lancer := objects.NewEntity(libs.Vector{X: float64(400 - (100*4)/2), Y: float64(400 - (100*4)/2)}, assets.LancerAnimations)
	orc := objects.NewEntity(libs.Vector{X: float64(600 - (100*4)/2), Y: float64(400 - (100*4)/2)}, assets.OrcAnimations)
	return &Game{
		Entity:   *entity,
		Skeleton: *skeleton,
		Lancer:   *lancer,
		Orc:      *orc,
	}

}

func (g *Game) Update() error {
	g.Entity.Update()
	g.Skeleton.Update()
	g.Lancer.Update()
	g.Orc.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Entity.Draw(screen)
	g.Skeleton.Draw(screen)
	g.Lancer.Draw(screen)
	g.Orc.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
