package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	tileSet     *ebiten.Image
	sheetWidth  int
	sheetHeight int
	tileWidth   int
	tileHeight  int
}

func (s *SpriteSheet) Tile(index int) *ebiten.Image {
	pos := index * s.tileWidth
	x := pos % s.sheetWidth
	y := (pos / s.sheetWidth) * s.tileHeight
	rect := image.Rect(x, y, x+s.tileWidth, y+s.tileHeight)
	subImg := s.tileSet.SubImage(rect).(*ebiten.Image)
	return subImg
}

func NewSpriteSheet(tileSet *ebiten.Image, width, height, tileSize int) *SpriteSheet {
	return &SpriteSheet{
		tileSet:     tileSet,
		sheetWidth:  width,
		sheetHeight: height,
		tileWidth:   tileSize,
		tileHeight:  tileSize,
	}
}
