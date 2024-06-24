package libs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Image *ebiten.Image
	FrameData
}

type AnimationSet struct {
	Animations map[string]Animation
}

func NewAnimation(image *ebiten.Image, frameData FrameData) *Animation {
	return &Animation{
		Image:     image,
		FrameData: frameData,
	}
}

func NewAnimationSet() *AnimationSet {
	return &AnimationSet{
		Animations: make(map[string]Animation),
	}
}

func (as *AnimationSet) AddAnimation(name string, anim Animation) {
	as.Animations[name] = anim
}

func (as *AnimationSet) GetAnimation(name string) (Animation, bool) {
	anim, found := as.Animations[name]
	return anim, found
}
