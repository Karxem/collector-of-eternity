package gamestate

import "github.com/hajimehoshi/ebiten/v2"

type GameState interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type StateManager struct {
	currentState GameState
}

func (s *StateManager) SetState(state GameState) {
	s.currentState = state
}

func (s *StateManager) Update() error {
	return s.currentState.Update()
}

func (s *StateManager) Draw(screen *ebiten.Image) {
	s.currentState.Draw(screen)
}
