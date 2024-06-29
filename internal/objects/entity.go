package objects

import (
	"image"
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	GameObject
	libs.FrameData
	BoundingBox      libs.Rect
	CurrentAnimation libs.Animation
	AnimationFrame   int
	Animations       libs.AnimationSet
	Direction        int
	IsDamaged        bool
}

func NewEntity(position libs.Vector, boundingBoxWidth, boundingBoxHeight int, animations libs.AnimationSet) *Entity {
	return &Entity{
		GameObject:       *NewGameObject(position),
		BoundingBox:      libs.NewRect(position.X, position.Y, float64(boundingBoxWidth), float64(boundingBoxHeight)),
		Animations:       animations,
		CurrentAnimation: animations.Animations["idle"],
		AnimationFrame:   0,
		Direction:        1,
		IsDamaged:        false,
	}
}

func (e *Entity) Update() {
	e.AnimationFrame++
	e.AnimationFrame %= e.CurrentAnimation.FrameCount * 5

	if e.Position.X != e.BoundingBox.X || e.Position.Y != e.BoundingBox.Y {
		e.UpdateBoundingBox()
	}

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

func (e *Entity) UpdateBoundingBox() {
	e.BoundingBox.X = e.Position.X
	e.BoundingBox.Y = e.Position.Y
}

func (e *Entity) TakeDamage() {
	if e.IsDamaged {
		for i := 0; i <= e.CurrentAnimation.FrameCount*2; i++ {
			e.IsDamaged = false
			e.CurrentAnimation = e.Animations.Animations["idle"]
		}
		return
	}

	if val, ok := e.Animations.Animations["hurt"]; ok {
		e.CurrentAnimation = val
	}
	e.IsDamaged = true
}
