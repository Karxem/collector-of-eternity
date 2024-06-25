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
	Animations       libs.AnimationSet
	Direction        int
}

func NewEntity(position libs.Vector, animations libs.AnimationSet) *Entity {
	return &Entity{
		GameObject:       *NewGameObject(position),
		Animations:       animations,
		CurrentAnimation: animations.Animations["idle"],
		AnimationFrame:   0,
		Direction:        1,
	}
}

func (e *Entity) Update() {
	e.AnimationFrame++
	e.AnimationFrame %= e.CurrentAnimation.FrameCount * 5

	// log.Printf("Position: (%f, %f), Frame: %d", e.Position.X, e.Position.Y, e.AnimationFrame)
}

func (e *Entity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Set scale based on direction
	if e.Direction == -1 {
		op.GeoM.Scale(-float64(e.CurrentAnimation.ScaleFactor), float64(e.CurrentAnimation.ScaleFactor)) // Mirror horizontally
		op.GeoM.Translate(float64(e.CurrentAnimation.FrameWidth*e.CurrentAnimation.ScaleFactor), 0)      // Adjust translation
	} else {
		op.GeoM.Scale(float64(e.CurrentAnimation.ScaleFactor), float64(e.CurrentAnimation.ScaleFactor))
	}

	// Calculate position
	i := (e.AnimationFrame / 5) % e.CurrentAnimation.FrameCount
	sx, sy := e.CurrentAnimation.FrameOX+i*e.CurrentAnimation.FrameWidth, e.CurrentAnimation.FrameOY

	op.GeoM.Translate(e.Position.X, e.Position.Y)
	screen.DrawImage(e.CurrentAnimation.Image.SubImage(image.Rect(sx, sy, sx+e.CurrentAnimation.FrameWidth, sy+e.CurrentAnimation.FrameHeight)).(*ebiten.Image), op)

	// log.Printf("Drawing frame: %d at position: (%f, %f)", e.AnimationFrame, e.Position.X, e.Position.Y)
}
