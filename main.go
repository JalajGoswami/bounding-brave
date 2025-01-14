package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	backgrounds      []*ebiten.Image
	characterTileset *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, background := range g.backgrounds {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2, 2)
		screen.DrawImage(background, opts)
	}

	img := g.characterTileset.SubImage(image.Rect(0, 0, 56, 56)).(*ebiten.Image)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(100, 150)
	screen.DrawImage(img, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func InitGame() *Game {
	return &Game{
		backgrounds: []*ebiten.Image{
			LoadImage("assets/backgrounds/layer_1.png"),
			LoadImage("assets/backgrounds/layer_2.png"),
			LoadImage("assets/backgrounds/layer_3.png"),
		},
		characterTileset: LoadImage("assets/characters/hero.png"),
	}
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Bounding Brave")

	game := InitGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func LoadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
