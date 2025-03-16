package engine

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	Screen *ebiten.Image
	Camera *Camera
}
