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
	KnightAnimations   libs.AnimationSet
	SkeletonAnimations libs.AnimationSet
)

func init() {
	KnightAnimations = *libs.NewAnimationSet()
	KnightAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Walk.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Attack01.png"), libs.FrameData{FrameCount: 7, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Attack02.png"), libs.FrameData{FrameCount: 10, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("heavy-attack", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Attack03.png"), libs.FrameData{FrameCount: 11, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

	SkeletonAnimations = *libs.NewAnimationSet()
	SkeletonAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SkeletonAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
}

func LoadImage(path string) *ebiten.Image {
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
