package objects

import (
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	Position libs.Vector
}

func NewGameObject(position libs.Vector) *GameObject {
	return &GameObject{
		Position: position,
	}
}

func (g *GameObject) Update() {

}

func (g *GameObject) Draw(screen *ebiten.Image) {

}
