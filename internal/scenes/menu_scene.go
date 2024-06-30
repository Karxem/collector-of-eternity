package scenes

import (
	"fmt"

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
	ebitenutil.DebugPrintAt(screen, "       Collector of Eternity\nMenu Scene - Press ENTER to start", 400, 280)
	fps := fmt.Sprintf("%f", ebiten.ActualFPS())
	tps := fmt.Sprintf("%f", ebiten.ActualTPS())
	ebitenutil.DebugPrintAt(screen, "FPS: "+fps, 10, 10)
	ebitenutil.DebugPrintAt(screen, "TPS: "+tps, 10, 30)
}
