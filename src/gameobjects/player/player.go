package player

import (
	"image"
	"math"
	"pick-it-up/libs/animation"
	"pick-it-up/libs/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	AttackAnimation animation.Animation
	CurrentAnim     *animation.Animation
	AnimationFrame  int
	PlayerPosition  vector.Vector
	PlayerDirection int
	ScreenWidth     int
	ScreenHeight    int
	IsAttacking     bool // New field to track attacking state
}

func NewPlayer(idleAnimation, walkAnimation, attackAnimation animation.Animation, screenWidth, screenHeight int) *Player {
	return &Player{
		IdleAnimation:   idleAnimation,
		WalkAnimation:   walkAnimation,
		AttackAnimation: attackAnimation,
		CurrentAnim:     &idleAnimation, // Defaults to the idle animation
		AnimationFrame:  0,              // Determines the current animation frame
		PlayerPosition:  vector.Vector{X: float64(screenWidth - (frameWidth*scaleFactor)/2), Y: float64(screenHeight - (frameHeight*scaleFactor)/2)},
		ScreenWidth:     screenWidth,
		ScreenHeight:    screenHeight,
		IsAttacking:     false,
	}
}

func (p *Player) Update() {
	if p.IsAttacking {
		p.AnimationFrame++
		if p.AnimationFrame >= p.AttackAnimation.FrameCount*5 {
			p.IsAttacking = false
			p.AnimationFrame = 0
		}

		return
	}

	speed := 3.0
	var delta vector.Vector

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.IsAttacking = true
		p.AnimationFrame = 0
		p.CurrentAnim = &p.AttackAnimation
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if p.PlayerPosition.Y <= -float64(frameHeight)-float64(frameHeight)/2 {
			// log.Println("W:", "Out of bounds!")
			return
		}
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if p.PlayerPosition.X <= -float64(frameWidth)-float64(frameWidth)/2 {
			// log.Println("A:", "Out of bounds!")
			return
		}
		delta.X = -speed
		p.PlayerDirection = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if p.PlayerPosition.Y >= float64(p.ScreenHeight)+float64(frameHeight)/2 {
			// log.Println("S:", "Out of bounds!")
			return
		}
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if p.PlayerPosition.X >= float64(p.ScreenWidth)+float64(p.ScreenWidth)/2 {
			// log.Println("D:", "Out of bounds!")
			return
		}
		delta.X = speed
		p.PlayerDirection = 1
	}

	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	if delta.X != 0 || delta.Y != 0 {
		p.CurrentAnim = &p.WalkAnimation
	} else {
		p.CurrentAnim = &p.IdleAnimation
	}

	p.PlayerPosition.X += delta.X
	p.PlayerPosition.Y += delta.Y
	p.AnimationFrame++
	p.AnimationFrame %= p.CurrentAnim.FrameCount * 5
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
	i := (p.AnimationFrame / 5) % p.CurrentAnim.FrameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	op.GeoM.Translate(p.PlayerPosition.X, p.PlayerPosition.Y)

	// Draw the scaled image
	screen.DrawImage(p.CurrentAnim.Image.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
