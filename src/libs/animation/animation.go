package animation

import "github.com/hajimehoshi/ebiten/v2"

// Type that holds the sprite image and the animation frame count
type Animation struct {
	Image      *ebiten.Image
	FrameCount int
}
