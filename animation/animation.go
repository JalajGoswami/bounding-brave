package animation

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
}

func (a *Animation) Frame() int {
	return a.currentFrame
}

func (a *Animation) Update() {
	a.frameCounter--
	if a.frameCounter < 0 {
		a.currentFrame++
		a.frameCounter = a.speedInTicks
		if a.currentFrame > a.lastFrame {
			a.currentFrame = a.firstFrame
		}
	}
}

func (a *Animation) ChangeBounds(firstFrame, lastFrame int) {
	a.firstFrame = firstFrame
	a.lastFrame = lastFrame
	a.currentFrame = firstFrame
	a.frameCounter = a.speedInTicks
}

// frameTime will be rounded off to a nearest multiple of rendering frame time
func NewAnimation(firstFrame, lastFrame int, frameTime float64) *Animation {
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
	}
}
