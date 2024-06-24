package main

import (
	"embed"
	"image"
	_ "image/png" // Ensure PNG format is recognized
	"log"
	"pick-it-up/gameobjects/player"
	"pick-it-up/libs/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 500
	ScreenHeight = 300
)

type Game struct {
	player *player.Player
}

//go:embed assets/*
var assets embed.FS

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
	idleAnimation := animation.Animation{Image: loadImage("Characters/Knight/Knight-Idle.png"), FrameCount: 6}
	walkAnimation := animation.Animation{Image: loadImage("Characters/Knight/Knight-Walk.png"), FrameCount: 8}
	attackAnimation := animation.Animation{Image: loadImage("Characters/Knight/Knight-Attack02.png"), FrameCount: 10}

	g := &Game{
		player: player.NewPlayer(idleAnimation, walkAnimation, attackAnimation, ScreenWidth, ScreenHeight),
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
