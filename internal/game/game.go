package game

import (
	"pick-it-up/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SceneManager scenes.SceneManager
}

func NewGame() *Game {
	testScene := scenes.NewTestScene()
	sceneManager := scenes.NewSceneManager(testScene)
	return &Game{
		SceneManager: *sceneManager,
	}

}

func (g *Game) Update() error {
	return g.SceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
