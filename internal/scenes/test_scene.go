package scenes

import (
	assets "pick-it-up"
	"pick-it-up/internal/libs"
	"pick-it-up/internal/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type TestScene struct {
	Player   objects.Player
	Skeleton objects.Entity
}

func NewTestScene() *TestScene {
	player := objects.NewPlayer(libs.Vector{X: float64(400 - (100*4)/2), Y: float64(300 - (100*4)/2)}, assets.PlayerAnimations)
	skeleton := objects.NewEntity(libs.Vector{X: float64(600 - (100*4)/2), Y: float64(300 - (100*4)/2)}, assets.SkeletonAnimations)
	return &TestScene{
		Player:   *player,
		Skeleton: *skeleton,
	}
}

func (s *TestScene) Update() error {
	s.Player.Update()
	s.Skeleton.Update()
	return nil
}

func (s *TestScene) Draw(screen *ebiten.Image) {
	s.Player.Draw(screen)
	s.Skeleton.Draw(screen)
}
