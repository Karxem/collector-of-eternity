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
)

func init() {
	PlayerAnimations = *libs.NewAnimationSet()
	PlayerAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	PlayerAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Walk.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	PlayerAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Attack01.png"), libs.FrameData{FrameCount: 9, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	PlayerAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Attack02.png"), libs.FrameData{FrameCount: 9, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	PlayerAnimations.AddAnimation("heavy-attack", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Attack03.png"), libs.FrameData{FrameCount: 12, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

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
