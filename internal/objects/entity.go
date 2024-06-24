package objects

import (
	"image"
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	GameObject
	libs.FrameData
	CurrentAnimation libs.Animation
	AnimationFrame   int

	// TODO: Use AnimationSet
	IdleAnimation  libs.Animation
	WalkAnimation  libs.Animation
	HurtAnimation  libs.Animation
	DeathAnimation libs.Animation
}

func NewEntity(position libs.Vector, idleAnimation libs.Animation) *Entity {
	return &Entity{
		GameObject:       *NewGameObject(position),
		IdleAnimation:    idleAnimation,
		CurrentAnimation: idleAnimation,
		AnimationFrame:   0,
	}
}

func (e *Entity) Update() {
	e.AnimationFrame++
}

func (e *Entity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(float64(e.CurrentAnimation.ScaleFactor), float64(e.CurrentAnimation.ScaleFactor))
	op.GeoM.Translate(e.Position.X, e.Position.Y)

	// Calculate position
	i := (e.AnimationFrame / 5) % e.CurrentAnimation.FrameCount
	sx, sy := e.CurrentAnimation.FrameOX+i*e.CurrentAnimation.FrameWidth, e.CurrentAnimation.FrameOY
	screen.DrawImage(e.CurrentAnimation.Image.SubImage(image.Rect(sx, sy, sx+e.CurrentAnimation.FrameWidth, sy+e.CurrentAnimation.FrameHeight)).(*ebiten.Image), op)
}
