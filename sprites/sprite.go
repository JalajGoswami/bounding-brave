package sprites

import (
	"bounding-brave/engine"
	"image"
)

type Sprite interface {
	Draw(scene *engine.Scene)
	Update()
	Bounds() image.Rectangle
}

type Box interface {
	Bounds() image.Rectangle
}
