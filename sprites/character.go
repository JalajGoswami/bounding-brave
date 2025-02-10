package sprites

import (
	"bounding-brave/animation"
	"bounding-brave/utils"
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
	spriteSheet     *SpriteSheet
	x, y            float64
	dx, dy          float64
	state           CharacterState
	flipped         bool
	animat          *animation.Animation
	posInSrcImg     image.Point
	characterWidth  int
	characterHeight int
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
	opts.GeoM.Translate(float64(-c.posInSrcImg.X), float64(-c.posInSrcImg.Y))
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
		c.dy = -2
	}
	c.dy += 0.1
	if math.Abs(c.dx) > 1 {
		c.dx = c.dx / math.Abs(c.dx) // set 1 if dx is greater (only magnitude)
	}
	if c.dy > 5 {
		c.dy = 5
	}
	c.x += c.dx
	c.y += c.dy

	if c.state != prevState {
		c.animat.ChangeBounds(c.state.Tiles())
	}
	c.animat.Update()
}

func (c *Character) Collides(collidable Sprite) {
	bounds := collidable.Bounds()
	dx, dy := utils.RectIntersectFloat(bounds, c.x, c.y, c.characterWidth, c.characterHeight)
	if dx < dy {
		// if c.x > float64(bounds.Min.X) {
		// 	c.x = float64(bounds.Max.X)
		// } else {
		// 	c.x = float64(bounds.Min.X) - float64(c.characterWidth)
		// }
		direction := (c.x - float64(bounds.Min.X)) / math.Abs(c.x-float64(bounds.Min.X))
		c.x += direction * dx
		c.dx = 0
	} else {
		// if c.y > float64(bounds.Min.Y) {
		// 	c.y = float64(bounds.Max.Y)
		// } else {
		// 	c.y = float64(bounds.Min.Y) - float64(c.characterHeight)
		// }
		direction := (c.y - float64(bounds.Min.Y)) / math.Abs(c.y-float64(bounds.Min.Y))
		c.y += direction * dy
		c.dy = 0
	}
}

func (c *Character) Bounds() image.Rectangle {
	return image.Rect(
		int(c.x), int(c.y),
		int(math.Ceil(c.x))+c.characterWidth,
		int(math.Ceil(c.y))+c.characterHeight,
	)
}

func (c *Character) reset() {
	c.state = Idle
}

func NewCharacter(spriteSheet *SpriteSheet, x, y float64, posInSrcImg image.Point) *Character {
	first, last := Idle.Tiles()
	characterHeight := spriteSheet.tileHeight - posInSrcImg.Y
	characterWidth := spriteSheet.tileWidth - (2 * posInSrcImg.X)
	return &Character{
		spriteSheet:     spriteSheet,
		x:               x,
		y:               y,
		state:           Idle,
		animat:          animation.NewAnimation(first, last, 100),
		posInSrcImg:     posInSrcImg,
		characterWidth:  characterWidth,
		characterHeight: characterHeight,
	}
}
