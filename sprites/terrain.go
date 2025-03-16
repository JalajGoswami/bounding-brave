package sprites

import (
	"bounding-brave/engine"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Terrain struct {
	tileSet       *ebiten.Image
	x, y          float64
	tileX, tileY  int
	width, height int
}

func (t *Terrain) Draw(scene *engine.Scene) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(t.x, t.y)
	scene.Camera.ApplyCam(&opts.GeoM)
	img := t.tileSet.SubImage(
		image.Rect(t.tileX, t.tileY, t.tileX+t.width, t.tileY+t.height),
	).(*ebiten.Image)
	scene.Screen.DrawImage(img, opts)
}

func (t *Terrain) Update() {}

func (t *Terrain) Bounds() image.Rectangle {
	return image.Rect(int(t.x), int(t.y), int(t.x)+t.width, int(t.y)+t.height)
}

func NewTerrain(sourceImage *ebiten.Image, x, y float64, tileX, tileY, width, height int) *Terrain {
	return &Terrain{
		tileSet: sourceImage,
		x:       x, y: y,
		tileX: tileX, tileY: tileY,
		width: width, height: height,
	}
}
