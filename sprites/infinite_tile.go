package sprites

import (
	"bounding-brave/engine"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type InfiniteTile struct {
	img *ebiten.Image
}

func (i *InfiniteTile) Draw(scene *engine.Scene) {
	i.DrawOnOffset(scene.Screen, scene.Camera.Pos())
}

func (i *InfiniteTile) DrawOnOffset(screen *ebiten.Image, offset float64) {
	imageWidth := float64(i.img.Bounds().Dx())
	screenWidth := float64(screen.Bounds().Dx())

	startX := math.Mod(offset, imageWidth)
	factor := (math.Copysign(1, startX) + 1) / 2 // 0 if -ve 1 otherwise
	startX -= imageWidth * factor

	// Draw the images one after anoter to cover the screen
	for x := startX; x < screenWidth; x += imageWidth {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, 0)
		screen.DrawImage(i.img, op)
	}
}

func NewInfiniteTile(img *ebiten.Image) *InfiniteTile {
	return &InfiniteTile{img}
}
