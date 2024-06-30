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
	KnightAnimations    libs.AnimationSet
	TemplarAnimations   libs.AnimationSet
	SwordsmanAnimations libs.AnimationSet
	AxemanAnimations    libs.AnimationSet
	SkeletonAnimations  libs.AnimationSet
	OrcAnimations       libs.AnimationSet
)

func init() {
	// Knight Animations
	KnightAnimations = *libs.NewAnimationSet()
	KnightAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Walk.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Attack01.png"), libs.FrameData{FrameCount: 7, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Attack02.png"), libs.FrameData{FrameCount: 10, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("heavy-attack", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Attack03.png"), libs.FrameData{FrameCount: 11, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	KnightAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Knight/Knight-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

	// Armored Axeman Animations
	AxemanAnimations = *libs.NewAnimationSet()
	AxemanAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	AxemanAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Walk.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	AxemanAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Attack01.png"), libs.FrameData{FrameCount: 9, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	AxemanAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Attack02.png"), libs.FrameData{FrameCount: 9, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	AxemanAnimations.AddAnimation("heavy-attack", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Attack03.png"), libs.FrameData{FrameCount: 12, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	AxemanAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	AxemanAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Armored Axeman/Armored Axeman-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

	// Knight Templar Animations
	TemplarAnimations = *libs.NewAnimationSet()
	TemplarAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	TemplarAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Walk02.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	TemplarAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Attack01.png"), libs.FrameData{FrameCount: 7, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	TemplarAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Attack02.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	TemplarAnimations.AddAnimation("heavy-attack", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Attack03.png"), libs.FrameData{FrameCount: 11, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	TemplarAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	TemplarAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Knight Templar/Knight Templar-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

	// Swordsman Animations
	SwordsmanAnimations = *libs.NewAnimationSet()
	SwordsmanAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SwordsmanAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Walk.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SwordsmanAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Attack01.png"), libs.FrameData{FrameCount: 7, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SwordsmanAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Attack02.png"), libs.FrameData{FrameCount: 15, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SwordsmanAnimations.AddAnimation("heavy-attack", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Attack03.png"), libs.FrameData{FrameCount: 12, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SwordsmanAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SwordsmanAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Swordsman/Swordsman-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

	// Skeleton Animations
	SkeletonAnimations = *libs.NewAnimationSet()
	SkeletonAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SkeletonAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Attack01.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SkeletonAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Attack02.png"), libs.FrameData{FrameCount: 7, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SkeletonAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	SkeletonAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Skeleton/Skeleton-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))

	// Orc Animations
	OrcAnimations = *libs.NewAnimationSet()
	OrcAnimations.AddAnimation("idle", *libs.NewAnimation(LoadImage("Characters/Orc/Orc-Idle.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	OrcAnimations.AddAnimation("walk", *libs.NewAnimation(LoadImage("Characters/Orc/Orc-Walk.png"), libs.FrameData{FrameCount: 8, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	OrcAnimations.AddAnimation("primary-attack", *libs.NewAnimation(LoadImage("Characters/Orc/Orc-Attack01.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	OrcAnimations.AddAnimation("secondary-attack", *libs.NewAnimation(LoadImage("Characters/Orc/Orc-Attack02.png"), libs.FrameData{FrameCount: 6, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	OrcAnimations.AddAnimation("hurt", *libs.NewAnimation(LoadImage("Characters/Orc/Orc-Hurt.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
	OrcAnimations.AddAnimation("death", *libs.NewAnimation(LoadImage("Characters/Orc/Orc-Death.png"), libs.FrameData{FrameCount: 4, FrameOX: 0, FrameOY: 0, FrameWidth: 100, FrameHeight: 100, ScaleFactor: 4}))
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
