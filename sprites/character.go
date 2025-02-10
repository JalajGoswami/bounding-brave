package sprites

import (
	"bounding-brave/animation"
	"bounding-brave/utils"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type CharacterState uint8

const (
	Idle CharacterState = iota
	Attacking
	Running
	JumpPrep
	JumpAscent
	JumpDescent
	JumpLanding
	JumpReload
)

func (s CharacterState) Tiles() (first, last, repeat int) {
	repeat = -1
	switch s {
	case Idle:
		first = 0
		last = 5
	case Running:
		first = 15
		last = 23
	case JumpPrep:
		first = 24
		last = 25
		repeat = 0
	case JumpAscent:
		first = 26
		last = 29
	case JumpDescent:
		first = 33
		last = 36
	case JumpLanding:
		first = 37
		last = 39
		repeat = 0
	case JumpReload:
		first = 30
		last = 32
		repeat = 0
	}
	return
}

func (s CharacterState) Next() CharacterState {
	next := s
	switch s {
	case JumpPrep:
		next = JumpAscent
	case JumpLanding:
		next = Idle
	case JumpReload:
		next = JumpAscent
	}
	return next
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
	isSecondJump    bool
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

	// handling inputs
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.dx -= 0.1
		if c.state < JumpPrep || c.state > JumpReload {
			c.state = Running
		}
		c.flipped = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.dx += 0.1
		if c.state < JumpPrep || c.state > JumpReload {
			c.state = Running
		}
		c.flipped = false
	} else {
		c.dx = 0
	}
	if (inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsKeyJustPressed(ebiten.KeyW) ||
		inpututil.IsKeyJustPressed(ebiten.KeyUp)) && !c.isSecondJump {
		c.dy = -2
		if c.state == JumpAscent || c.state == JumpDescent {
			c.state = JumpReload
			c.isSecondJump = true
		} else {
			c.state = JumpPrep
		}
	}

	// change intermediate state
	if c.state == JumpAscent && c.dy >= 0 {
		c.state = JumpDescent
	}

	// pseudo gravity & speed constraints
	c.dy += 0.1
	if math.Abs(c.dx) > 1 {
		c.dx = c.dx / math.Abs(c.dx) // set 1 if dx is greater (only magnitude)
	}
	if c.dy > 5 {
		c.dy = 5
	}

	// applying current speed
	c.x += c.dx
	c.y += c.dy

	if c.state != prevState {
		c.animat.ChangeBounds(c.state.Tiles())
	} else {
		done := c.animat.Update()
		if done {
			c.state = c.state.Next()
			c.animat.ChangeBounds(c.state.Tiles())
		}
	}

}

func (c *Character) Collides(collidable Sprite) {
	bounds := collidable.Bounds()
	dx, dy := utils.RectIntersectFloat(bounds, c.x, c.y, c.characterWidth, c.characterHeight)
	if dx < dy {
		// translating rect by dx in x-axis
		direction := (c.x - float64(bounds.Min.X)) / math.Abs(c.x-float64(bounds.Min.X))
		c.x += direction * dx
		c.dx = 0
		if c.state == Running {
			c.state = Idle
			c.animat.ChangeBounds(c.state.Tiles())
		}
	} else {
		// translating rect by dy in y-axis
		direction := c.y - float64(bounds.Min.Y)
		c.y += math.Copysign(dy, direction)
		c.dy = 0
		if c.state == JumpDescent && math.Signbit(direction) {
			c.state = JumpLanding
			c.animat.ChangeBounds(c.state.Tiles())
			c.isSecondJump = false
		}
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
	if c.state == Running {
		c.state = Idle
	}
}

func NewCharacter(spriteSheet *SpriteSheet, x, y float64, posInSrcImg image.Point) *Character {
	first, last, repeat := Idle.Tiles()
	characterHeight := spriteSheet.tileHeight - posInSrcImg.Y
	characterWidth := spriteSheet.tileWidth - (2 * posInSrcImg.X)
	return &Character{
		spriteSheet:     spriteSheet,
		x:               x,
		y:               y,
		state:           Idle,
		animat:          animation.NewAnimation(first, last, repeat, 100),
		posInSrcImg:     posInSrcImg,
		characterWidth:  characterWidth,
		characterHeight: characterHeight,
	}
}
