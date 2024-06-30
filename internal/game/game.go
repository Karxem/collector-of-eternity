package game

import (
	"pick-it-up/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SceneManager scenes.SceneManager
}

func NewGame() *Game {
	menuScene := scenes.NewMenuScene()
	sceneManager := scenes.NewSceneManager(menuScene)
	return &Game{
		SceneManager: *sceneManager,
	}

}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.SceneManager.SetScene(scenes.NewTestScene())
	}
	return g.SceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
