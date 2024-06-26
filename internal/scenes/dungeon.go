package scenes

import (
	assets "pick-it-up"
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

type Dungeon struct {
	Tier       string
	CreepCount int
	TileMaps   libs.TileMapSet
}

func NewDungeon(tier string, creepCount int, tileMaps libs.TileMapSet) *Dungeon {
	return &Dungeon{
		Tier:       tier,
		CreepCount: creepCount,
		TileMaps:   tileMaps,
	}
}

func (d *Dungeon) Update() error {
	return nil
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	for _, v := range d.TileMaps {
		drawTilemap(screen, 32, v)
	}
}

func drawTilemap(screen *ebiten.Image, tileSize int, tilemap [][]int) {
	for y, row := range tilemap {
		for x, tile := range row {
			// Calculate tile position in the tileset
			tilesetX := tile % 2 // Atlas is 2 tiles wide
			tilesetY := tile / 2 // Atlas is 6 tiles long
			subImage := libs.GetTile(assets.LoadImage("Tiles/TileAtlas.png"), 32, tilesetX, tilesetY)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
			op.GeoM.Scale(float64(3.12), float64(3.12)) // FIXME: Magic number
			screen.DrawImage(subImage, op)
		}
	}
}
