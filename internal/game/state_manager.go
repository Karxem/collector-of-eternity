package game

type GameState int

const (
	StatePlaying GameState = iota
	StatePaused
	StateGameOver
)

type StateManager struct {
	currentState GameState
}

func (sm *StateManager) SetState(state GameState) {
	sm.currentState = state
}

func (sm *StateManager) GetState() GameState {
	return sm.currentState
}
