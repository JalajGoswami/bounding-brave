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
	terrains    []*sprites.Terrain
}

func (g *Game) Update() error {
	g.hero.Update()
	for _, terrain := range g.terrains {
		if terrain.Bounds().Overlaps(g.hero.Bounds()) {
			g.hero.Collides(terrain)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, background := range g.backgrounds {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2, 2)
		screen.DrawImage(background, opts)
	}

	// solid objects
	for _, terrain := range g.terrains {
		terrain.Draw(screen)
	}

	g.hero.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func InitGame() *Game {
	terrainTileSet := LoadImage("assets/terrains.png")
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
		terrains: []*sprites.Terrain{
			sprites.NewTerrain(terrainTileSet, 100, 300, 120, 168, 70, 23),
		},
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
