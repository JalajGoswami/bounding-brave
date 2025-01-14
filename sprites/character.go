package sprites

import (
	"bounding-brave/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type CharacterState uint8

const (
	Idle CharacterState = iota
	Attacking
	Running
	Jumping
)

func (s CharacterState) Tiles() (first, last int) {
	switch s {
	case Idle:
		first = 0
		last = 5
	case Running:
		first = 15
		last = 23
	}
	return
}

type Character struct {
	spriteSheet *SpriteSheet
	x, y        float64
	state       CharacterState
	animat      *animation.Animation
}

func (c *Character) Draw(screen *ebiten.Image) {
	indx := c.animat.Frame()
	img := c.spriteSheet.Tile(indx)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(c.x, c.y)
	screen.DrawImage(img, opts)
}

func (c *Character) Update() {
	prevState := c.state
	c.state = Idle
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.x -= 1
		c.state = Running
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.x += 1
		c.state = Running
	}

	if c.state != prevState {
		c.animat.ChangeBounds(c.state.Tiles())
	}
	c.animat.Update()
}

func NewCharacter(spriteSheet *SpriteSheet, x, y float64) *Character {
	first, last := Idle.Tiles()
	return &Character{
		spriteSheet: spriteSheet,
		x:           x,
		y:           y,
		state:       Idle,
		animat:      animation.NewAnimation(first, last, 100),
	}
}
