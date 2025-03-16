package main

import (
	"bounding-brave/config"
	"bounding-brave/engine"
	"bounding-brave/sprites"
	"bounding-brave/sprites/character"
	"bounding-brave/utils"
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var assets embed.FS

type Game struct {
	scene       *engine.Scene
	backgrounds []*sprites.InfiniteTile
	hero        *character.Character
	terrains    []*sprites.Terrain
}

func (g *Game) Update() error {
	config.Tick = (config.Tick + 1) % 1000
	g.hero.Update()
	for _, terrain := range g.terrains {
		if terrain.Bounds().Overlaps(g.hero.Bounds()) {
			g.hero.Collides(terrain)
		}
	}
	g.scene.Camera.UpdateCamera(g.hero.Bounds().Min.X, g.hero.Bounds().Min.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Screen.Clear()
	for i, background := range g.backgrounds {
		parallexFactor := [...]float64{0.6, 0.8, 1}[i]
		background.DrawOnOffset(g.scene.Screen, parallexFactor*g.scene.Camera.Pos())
	}

	// solid objects
	for _, terrain := range g.terrains {
		terrain.Draw(g.scene)
	}

	g.hero.Draw(g.scene)
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.scene.Screen, opts)
	if config.DebugPrintText != "" {
		ebitenutil.DebugPrint(screen, config.DebugPrintText+g.hero.Bounds().String())
		config.DebugPrintText = ""
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}

func InitGame() *Game {
	terrainTileSet := LoadImage("assets/terrains.png")
	centerY := config.SCREEN_HEIGHT / 2
	return &Game{
		scene: &engine.Scene{
			Screen: ebiten.NewImage(config.SCREEN_WIDTH, config.SCREEN_HEIGHT),
			Camera: engine.NewCamera(100),
		},
		backgrounds: []*sprites.InfiniteTile{
			sprites.NewInfiniteTile(
				utils.ScaleImage(LoadImage("assets/backgrounds/layer_1.png"), 2),
			),
			sprites.NewInfiniteTile(
				utils.ScaleImage(LoadImage("assets/backgrounds/layer_2.png"), 2),
			),
			sprites.NewInfiniteTile(
				utils.ScaleImage(LoadImage("assets/backgrounds/layer_3.png"), 2),
			),
		},
		hero: character.NewCharacter(
			sprites.NewSpriteSheet(LoadImage("assets/characters/hero.png"), 448, 616, 56),
			100, float64(centerY),
			image.Pt(15, 23),
		),
		terrains: []*sprites.Terrain{
			sprites.NewTerrain(terrainTileSet, 100, 250, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 0, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, -100, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, -200, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 170, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 240, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 310, 300, 120, 168, 70, 23),
			// sprites.NewTerrain(terrainTileSet, 170, 180, 216, 144, 48, 120),
			sprites.NewTerrain(terrainTileSet, 100, 100, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 380, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 450, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 520, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 590, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 660, 300, 120, 168, 70, 23),
			sprites.NewTerrain(terrainTileSet, 730, 300, 120, 168, 70, 23),
		},
	}
}

func main() {
	ebiten.SetWindowSize(
		config.PIXED_DENSITY*config.SCREEN_WIDTH,
		config.PIXED_DENSITY*config.SCREEN_HEIGHT,
	)
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
