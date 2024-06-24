package game

import (
	scene "pick-it-up/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	currentScene scene.Scene
}

func (sm *SceneManager) ChangeScene(scene scene.Scene) {
	sm.currentScene = scene
}

func (sm *SceneManager) Update() error {
	return sm.currentScene.Update()
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	sm.currentScene.Draw(screen)
}

func (sm *SceneManager) HandleInput() {
	sm.currentScene.HandleInput()
}
