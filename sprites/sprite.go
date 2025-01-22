package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	Draw(screen *ebiten.Image)
	Update()
	Bounds() image.Rectangle
}
