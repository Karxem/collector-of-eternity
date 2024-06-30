package scenes

import (
	assets "pick-it-up"
	"pick-it-up/internal/libs"
	"pick-it-up/internal/objects"
	"pick-it-up/internal/systems"

	"github.com/hajimehoshi/ebiten/v2"
)

var TileMap = [][]int{
	{4, 11, 11, 11, 11, 11, 11, 11, 11, 5},
	{8, 13, 13, 13, 13, 13, 13, 13, 13, 9},
	{8, 13, 13, 13, 13, 13, 13, 13, 13, 9},
	{8, 13, 13, 13, 13, 13, 13, 13, 13, 9},
	{8, 13, 13, 13, 13, 13, 13, 13, 13, 9},
	{6, 10, 10, 10, 10, 10, 10, 10, 10, 7},
}

var GroundMap = [][]int{
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
}

type TestScene struct {
	Dungeon
	Player           *objects.Player
	Skeleton         *objects.Entity
	Orc              *objects.Entity
	CollisionHandler *systems.CollisionHandler
	AttackHandler    *systems.AttackHandler
}

func NewTestScene() *TestScene {
	player := objects.NewPlayer(libs.Vector{X: float64(400 - (100*4)/2), Y: float64(300 - (100*4)/2)}, 50, 50, 80, 80, assets.AxemanAnimations)
	skeleton := objects.NewEntity(libs.Vector{X: float64(600 - (100*4)/2), Y: float64(300 - (100*4)/2)}, 50, 50, assets.SkeletonAnimations)
	orc := objects.NewEntity(libs.Vector{X: float64(400 - (100*4)/2), Y: float64(100 - (100*4)/2)}, 50, 50, assets.OrcAnimations)
	collisionHandler := systems.NewCollisionHandler()
	attackHandler := systems.NewAttackHandler()

	var tileMaps libs.TileMapSet
	tileMaps.AddTileMap(GroundMap)
	tileMaps.AddTileMap(TileMap)

	attackHandler.AddEntity(skeleton)
	attackHandler.AddEntity(orc)
	collisionHandler.AddEntity(orc)
	collisionHandler.AddEntity(orc)

	return &TestScene{
		Dungeon:          *NewDungeon("S", 100, tileMaps),
		Player:           player,
		Skeleton:         skeleton,
		Orc:              orc,
		CollisionHandler: collisionHandler,
		AttackHandler:    attackHandler,
	}
}

func (s *TestScene) Update() error {
	s.Player.Update()
	s.Skeleton.Update()
	s.Orc.Update()

	s.CollisionHandler.Update(*s.Player)
	s.AttackHandler.Update(*s.Player)

	return nil
}

func (s *TestScene) Draw(screen *ebiten.Image) {
	s.Dungeon.Draw(screen)

	s.Skeleton.Draw(screen)
	s.Orc.Draw(screen)
	s.Player.Draw(screen)
}
