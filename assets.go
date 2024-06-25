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
	PlayerAnimations   libs.AnimationSet
	SkeletonAnimations libs.AnimationSet
	LancerAnimations   libs.AnimationSet
	OrcAnimations      libs.AnimationSet
)

func init() {
	PlayerAnimations = *libs.NewAnimationSet()
	PlayerAnimations.AddAnimation("idle", *libs.NewAnimation(loadImage("Characters/Knight/Knight-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	PlayerAnimations.AddAnimation("walk", *libs.NewAnimation(loadImage("Characters/Knight/Knight-Walk.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SkeletonAnimations = *libs.NewAnimationSet()
	SkeletonAnimations.AddAnimation("idle", *libs.NewAnimation(loadImage("Characters/Skeleton/Skeleton-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	OrcAnimations = *libs.NewAnimationSet()
	OrcAnimations.AddAnimation("idle", *libs.NewAnimation(loadImage("Characters/Orc/Orc-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	LancerAnimations = *libs.NewAnimationSet()
	LancerAnimations.AddAnimation("idle", *libs.NewAnimation(loadImage("Characters/Lancer/Lancer-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
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
