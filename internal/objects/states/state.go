package states

type State interface {
	Enter()
	Update()
	Exit()
}

type StateMachine struct {
	CurrentState State
}

func (sm *StateMachine) ChangeState(ns State) {
	if sm.CurrentState != nil {
		sm.CurrentState.Exit()
	}

	sm.CurrentState = ns
	sm.CurrentState.Enter()
}

func (sm *StateMachine) Update() {
	if sm.CurrentState == nil {
		panic("No current state found in state machine")
	}

	sm.CurrentState.Update()
}
