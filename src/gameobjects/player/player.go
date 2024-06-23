package player

import (
	"image"
	"math"
	"pick-it-up/geometries/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	frameOX     = 0   // X-coordinate of the first frame
	frameOY     = 0   // Y-coordinate of the first frame (updated to start from the top)
	frameWidth  = 100 // Width of each frame
	frameHeight = 100 // Height of each frame
	scaleFactor = 4   // Factor to scale up the sprite
)

type Player struct {
	Image          *ebiten.Image
	Count          int
	FrameCount     int
	PlayerPosition vector.Vector
	ScreenWidth    int
	ScreenHeight   int
}

func NewPlayer(image *ebiten.Image, frameCount int, screenWidth, screenHeight int) *Player {
	return &Player{
		Image:          image,
		Count:          0,
		FrameCount:     frameCount,
		PlayerPosition: vector.Vector{X: float64(screenWidth - (frameWidth*scaleFactor)/2), Y: float64(screenHeight - (frameHeight*scaleFactor)/2)},
		ScreenWidth:    screenWidth,
		ScreenHeight:   screenHeight,
	}
}

func (p *Player) Update() {
	p.Count++

	speed := 3.0
	var delta vector.Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.PlayerPosition.X += delta.X
	p.PlayerPosition.Y += delta.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Set scale
	op.GeoM.Scale(scaleFactor, scaleFactor)

	// Calculate position
	i := (p.Count / 5) % p.FrameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	op.GeoM.Translate(p.PlayerPosition.X, p.PlayerPosition.Y)

	// Draw the scaled image
	screen.DrawImage(p.Image.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
