package objects

import (
	"math"
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Entity
}

func NewPlayer(position libs.Vector, animations libs.AnimationSet) *Player {
	return &Player{
		Entity: *NewEntity(position, animations),
	}
}

func (p *Player) Update() {
	move(&p.Entity)

	p.Entity.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}

func move(e *Entity) {
	speed := 3.0
	var delta libs.Vector

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X = -speed
		e.Direction = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X = speed
		e.Direction = 1
	}

	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	if delta.X != 0 || delta.Y != 0 {
		e.CurrentAnimation = e.Animations.Animations["walk"]
	} else {
		e.CurrentAnimation = e.Animations.Animations["idle"]
	}

	e.Position.X += delta.X
	e.Position.Y += delta.Y
}
