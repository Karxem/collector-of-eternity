package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MenuScene struct {
	Scene
}

func NewMenuScene() *MenuScene {
	return &MenuScene{}
}

func (s *MenuScene) Update() error {
	return nil
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Menu Scene - Press ENTER to start")
}
