package systems

import (
	"log"
	"pick-it-up/internal/objects"
)

type CollisionHandler struct {
	Entities []*objects.Entity
}

func NewCollisionHandler() *CollisionHandler {
	return &CollisionHandler{}
}

func (ch *CollisionHandler) Update(player objects.Player) error {
	for i, m := range ch.Entities {
		if player.BoundingBox.Intersects(m.BoundingBox) {
			log.Println("Intersected with entity:", i)
		}
	}

	return nil
}

func (ch *CollisionHandler) AddEntity(e *objects.Entity) {
	ch.Entities = append(ch.Entities, e)
}
