package systems

import (
	"log"
	"pick-it-up/internal/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type CollsionHandler struct {
	Entities []objects.Entity
}

func NewCollisionHandler() *CollsionHandler {
	return &CollsionHandler{}
}

func (ch *CollsionHandler) Update(player objects.Player) error {
	for i, m := range ch.Entities {
		if player.BoundingBox.Intersects(m.BoundingBox) {
			log.Println("Intersected with entity:", i)
		}
	}

	return nil
}

func (ch *CollsionHandler) Draw(screen *ebiten.Image) {

}

func (ch *CollsionHandler) AddEntity(e objects.Entity) {
	ch.Entities = append(ch.Entities, e)
}
