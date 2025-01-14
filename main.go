package main

import (
	"bounding-brave/sprites"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var assets embed.FS

type Game struct {
	backgrounds []*ebiten.Image
	hero        *sprites.Character
}

func (g *Game) Update() error {
	g.hero.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, background := range g.backgrounds {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2, 2)
		screen.DrawImage(background, opts)
	}

	g.hero.Draw(screen)
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
		hero: sprites.NewCharacter(
			sprites.NewSpriteSheet(LoadImage("assets/characters/hero.png"), 448, 616, 56),
			100, 200,
		),
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
	img, _, err := ebitenutil.NewImageFromFileSystem(assets, path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
