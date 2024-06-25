package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneManager struct {
	currentScene Scene
}

func NewSceneManager(initialScene Scene) *SceneManager {
	return &SceneManager{
		currentScene: initialScene,
	}
}

func (sm *SceneManager) SetScene(scene Scene) {
	sm.currentScene = scene
}

func (sm *SceneManager) Update() error {
	if sm.currentScene != nil {
		return sm.currentScene.Update()
	}
	return nil
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	if sm.currentScene != nil {
		sm.currentScene.Draw(screen)
	}
}
