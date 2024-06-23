package player

import (
	"image"
	"math"
	"pick-it-up/libs/animation"
	"pick-it-up/libs/vector"

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
	IdleAnimation   animation.Animation
	WalkAnimation   animation.Animation
	CurrentAnim     *animation.Animation
	Count           int
	PlayerPosition  vector.Vector
	PlayerDirection int
	ScreenWidth     int
	ScreenHeight    int
}

func NewPlayer(idleAnimation, walkAnimation animation.Animation, screenWidth, screenHeight int) *Player {
	return &Player{
		IdleAnimation:  idleAnimation,
		WalkAnimation:  walkAnimation,
		CurrentAnim:    &idleAnimation, // Defaults to the idle animation
		Count:          0,              // Determines the current animation frame
		PlayerPosition: vector.Vector{X: float64(screenWidth - (frameWidth*scaleFactor)/2), Y: float64(screenHeight - (frameHeight*scaleFactor)/2)},
		ScreenWidth:    screenWidth,
		ScreenHeight:   screenHeight,
	}
}

func (p *Player) Update() {
	p.Count++

	speed := 3.0
	var delta vector.Vector

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X = -speed
		p.PlayerDirection = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X = speed
		p.PlayerDirection = 1
	}

	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	// Change animation based on movement
	if delta.X != 0 || delta.Y != 0 {
		p.CurrentAnim = &p.WalkAnimation
	} else {
		p.CurrentAnim = &p.IdleAnimation
	}

	p.PlayerPosition.X += delta.X
	p.PlayerPosition.Y += delta.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Set scale based on direction
	if p.PlayerDirection == -1 {
		op.GeoM.Scale(-scaleFactor, scaleFactor)              // Mirror horizontally
		op.GeoM.Translate(float64(frameWidth*scaleFactor), 0) // Adjust translation
	} else {
		op.GeoM.Scale(scaleFactor, scaleFactor)
	}

	// Calculate position
	i := (p.Count / 5) % p.CurrentAnim.FrameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	op.GeoM.Translate(p.PlayerPosition.X, p.PlayerPosition.Y)

	// Draw the scaled image
	screen.DrawImage(p.CurrentAnim.Image.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
