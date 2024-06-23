package main

import (
	"embed"
	"image"
	_ "image/png" // Ensure PNG format is recognized
	"log"
	"pick-it-up/gameobjects/player"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 360
)

type Game struct {
	player *player.Player
}

//go:embed assets/*
var assets embed.FS

var idleImage *ebiten.Image
var walkImage *ebiten.Image

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	idleImage = loadImage("Characters/Archer/Archer-Idle.png")
	walkImage = loadImage("Characters/Archer/Archer-Walk.png")

	g := &Game{
		player: player.NewPlayer(idleImage, walkImage, 6, 8, ScreenWidth, ScreenHeight),
	}

	// Set Window size and title
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Pick it up!")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// Used to load image assets from the embeded assets folder
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
