package engine

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	firstFrame   int
	lastFrame    int
	currentFrame int
	speedInTicks int
	frameCounter int
	repeat       int // negative indicates infinite
}

func (a *Animation) Frame() int {
	return a.currentFrame
}

func (a *Animation) Update() bool {
	if a.repeat == 0 && a.currentFrame == a.lastFrame {
		return true
	}
	if a.repeat > 0 {
		a.repeat--
	}
	a.frameCounter--
	if a.frameCounter < 0 {
		a.currentFrame++
		a.frameCounter = a.speedInTicks
		if a.currentFrame > a.lastFrame {
			a.currentFrame = a.firstFrame
		}
	}
	return false
}

func (a *Animation) ChangeBounds(firstFrame, lastFrame, repeat int) {
	a.firstFrame = firstFrame
	a.lastFrame = lastFrame
	a.repeat = repeat
	a.currentFrame = firstFrame
	a.frameCounter = a.speedInTicks
}

// frameTime will be rounded off to a nearest multiple of rendering frame time
func NewAnimation(firstFrame, lastFrame, repeat int, frameTime float64) *Animation {
	renderingFrameTime := 1000 / float64(ebiten.TPS())
	if frameTime < renderingFrameTime {
		frameTime = renderingFrameTime
	}
	speedInTicks := int(math.Round(frameTime / renderingFrameTime))
	return &Animation{
		firstFrame:   firstFrame,
		lastFrame:    lastFrame,
		currentFrame: firstFrame,
		speedInTicks: speedInTicks,
		frameCounter: speedInTicks,
		repeat:       repeat,
	}
}
