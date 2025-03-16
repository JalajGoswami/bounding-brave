package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// simple horziontal camera for this game
type Camera struct {
	x     int
	initX int
}

func (c *Camera) UpdateCamera(x, y int) {
	// c.x = min(c.initX-x, 0)
	c.x = c.initX - x
}

func (c *Camera) Pos() float64 {
	return float64(c.x)
}

func (c *Camera) ApplyCam(geom *ebiten.GeoM) {
	geom.Translate(c.Pos(), 0)
}

func NewCamera(initX int) *Camera {
	return &Camera{0, initX}
}
