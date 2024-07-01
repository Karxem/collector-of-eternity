package states

type State int

const (
	Idle State = iota
	Walk
	Attack
	Hurt
	Dead
)
