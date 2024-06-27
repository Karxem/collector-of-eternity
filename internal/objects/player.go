package objects

import (
	"math"
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Entity
	IsAttacking bool
	AttackFrame int
}

func NewPlayer(position libs.Vector, boundingBoxWidth, boundingBoxHeight int, animations libs.AnimationSet) *Player {
	return &Player{
		Entity:      *NewEntity(position, boundingBoxWidth, boundingBoxHeight, animations),
		IsAttacking: false,
		AttackFrame: 0,
	}
}

func (p *Player) Update() {
	moveInputListener(p)
	useAttackMoves(p)

	p.Entity.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}

func moveInputListener(p *Player) {
	/*
		TODO: This prevents the current animation to get set to idle when not moving.
			   The current state of the player should be handled by something like state machine that prevents inputs on certain states.
	*/
	if p.IsAttacking {
		return
	}

	speed := 3.0
	var delta libs.Vector

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X = -speed
		p.Direction = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X = speed
		p.Direction = 1
	}

	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	// Sets either walk or idle animation wether the player is moving or not
	if delta.X != 0 || delta.Y != 0 {
		p.CurrentAnimation = p.Animations.Animations["walk"]
	} else {
		p.CurrentAnimation = p.Animations.Animations["idle"]
	}

	p.Position.X += delta.X
	p.Position.Y += delta.Y
}

func useAttackMoves(p *Player) {
	if p.IsAttacking {
		p.AttackFrame++

		p.AnimationFrame++ // Increment animation frame

		// Check if the attack animation has completed
		var animFrames = p.CurrentAnimation.FrameCount
		if p.AttackFrame >= animFrames*2 {
			p.IsAttacking = false
			p.CurrentAnimation = p.Animations.Animations["idle"]
		}
		return
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		p.IsAttacking = true
		p.CurrentAnimation = p.Animations.Animations["primary-attack"]
		p.AttackFrame = 0 // Reset the attack frame counter
		p.Entity.AnimationFrame = 0
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		p.IsAttacking = true
		p.CurrentAnimation = p.Animations.Animations["secondary-attack"]
		p.AttackFrame = 0 // Reset the attack frame counter
		p.Entity.AnimationFrame = 0
	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		p.IsAttacking = true
		p.CurrentAnimation = p.Animations.Animations["heavy-attack"]
		p.AttackFrame = 0 // Reset the attack frame counter
		p.Entity.AnimationFrame = 0
	}
}
