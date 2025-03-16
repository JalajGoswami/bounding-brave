package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	tileSet     *ebiten.Image
	SheetWidth  int
	SheetHeight int
	TileWidth   int
	TileHeight  int
}

func (s *SpriteSheet) Tile(index int) *ebiten.Image {
	pos := index * s.TileWidth
	x := pos % s.SheetWidth
	y := (pos / s.SheetWidth) * s.TileHeight
	rect := image.Rect(x, y, x+s.TileWidth, y+s.TileHeight)
	subImg := s.tileSet.SubImage(rect).(*ebiten.Image)
	return subImg
}

func NewSpriteSheet(tileSet *ebiten.Image, width, height, tileSize int) *SpriteSheet {
	return &SpriteSheet{
		tileSet:     tileSet,
		SheetWidth:  width,
		SheetHeight: height,
		TileWidth:   tileSize,
		TileHeight:  tileSize,
	}
}
