package libs

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetTile(tileSet *ebiten.Image, tileSize, x, y int) *ebiten.Image {
	sx := x * tileSize
	sy := y * tileSize

	return tileSet.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}
