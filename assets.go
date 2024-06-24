package assets

import (
	"embed"
	"image/png"
	"log"
	"pick-it-up/internal/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var embeddedFiles embed.FS

var (
	PlayerIdleAnimation   libs.Animation
	SkeletonIdleAnimation libs.Animation
)

func init() {
	// Load player image
	PlayerIdleAnimation = *libs.NewAnimation(loadImage("Characters/Knight/Knight-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4})
	SkeletonIdleAnimation = *libs.NewAnimation(loadImage("Characters/Skeleton/Skeleton-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4})
}

func loadImage(path string) *ebiten.Image {
	file, err := embeddedFiles.Open("assets/" + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
