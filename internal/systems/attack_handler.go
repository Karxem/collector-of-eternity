package systems

import (
	"log"
	"pick-it-up/internal/objects"
)

type AttackHandler struct {
	Entities []*objects.Entity
}

func NewAttackHandler() *AttackHandler {
	return &AttackHandler{}
}

func (ah *AttackHandler) Update(player objects.Player) error {
	for i, e := range ah.Entities {
		if player.HitBox.Intersects(e.BoundingBox) {
			if player.IsAttacking {
				log.Println("Attack frame hit on entity:", i)
				e.TakeDamage()
			}
		}
	}

	return nil
}

func (ah *AttackHandler) AddEntity(e *objects.Entity) {
	ah.Entities = append(ah.Entities, e)
}
