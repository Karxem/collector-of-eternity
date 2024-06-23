package main

import (
	"embed"
	"image"
	_ "image/jpeg" // Ensure JPEG format is recognized
	_ "image/png"  // Ensure PNG format is recognized
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 360

	frameOX     = 0   // X-coordinate of the first frame
	frameOY     = 0   // Y-coordinate of the first frame (updated to start from the top)
	frameWidth  = 100 // Width of each frame
	frameHeight = 100 // Height of each frame
	frameCount  = 6   // Number of frames in the spritesheet
	scaleFactor = 4   // Factor to scale up the sprite
)

type Game struct {
	count int
}

//go:embed assets/*
var assets embed.FS
var playerImage *ebiten.Image

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY

	// Scale up the Sprite
	op.GeoM.Scale(scaleFactor, scaleFactor)

	// Poistion the Sprite
	op.GeoM.Translate(screenWidth-(frameWidth*scaleFactor)/2, screenHeight-(frameHeight*scaleFactor)/2)

	screen.DrawImage(playerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	playerImage = loadImage("Characters/Lancer/Lancer-Idle.png")

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Pick it up")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func loadImage(name string) *ebiten.Image {
	file, err := assets.Open("assets/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
