package sprites

import (
	"bounding-brave/animation"
	"image"
	"math"

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
	dx, dy      float64
	state       CharacterState
	flipped     bool
	animat      *animation.Animation
}

func (c *Character) Draw(screen *ebiten.Image) {
	indx := c.animat.Frame()
	img := c.spriteSheet.Tile(indx)
	opts := &ebiten.DrawImageOptions{}
	if c.flipped {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(float64(img.Bounds().Dx()), 0)
	}
	opts.GeoM.Translate(c.x, c.y)
	screen.DrawImage(img, opts)
}

func (c *Character) Update() {
	prevState := c.state
	c.reset()
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.dx -= 0.1
		c.state = Running
		c.flipped = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.dx += 0.1
		c.state = Running
		c.flipped = false
	} else {
		c.dx = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) ||
		ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyUp) {
		c.dy = 5
	}
	if math.Abs(c.dx) > 1 {
		c.dx = c.dx / math.Abs(c.dx) // set 1 if dx is greater (only magnitude)
	}
	c.x += c.dx

	if c.state != prevState {
		c.animat.ChangeBounds(c.state.Tiles())
	}
	c.animat.Update()
}

func (c *Character) Collides(collidable Sprite) {
	bounds := collidable.Bounds()
	if c.dy > 0 {
		c.y = float64(bounds.Min.Y) - float64(c.spriteSheet.tileHeight)
	}
}

func (c *Character) Bounds() image.Rectangle {
	return image.Rect(
		int(c.x), int(c.y),
		int(c.x)+c.spriteSheet.tileWidth,
		int(c.y)+c.spriteSheet.sheetHeight,
	)
}

func (c *Character) reset() {
	c.state = Idle
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
