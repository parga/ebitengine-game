package main

import (
	"image"
	"image/color"
	"log"
	animations "videogame/animations"
	entities "videogame/assets/enities"
	spriteSheet "videogame/spritesheet"
  cam "videogame/camera"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	player            *entities.Player
	playerSpritesheet *spriteSheet.SpriteSheet
	enemies           []*entities.Enemy
	potions           []*entities.Potion
	tilemapJSON       *TilemapJSON
	tilesets          []Tileset
	tilemapImage      *ebiten.Image
	cam               *cam.Camera
	colliders         []image.Rectangle
}

func setKeyBindings(player *entities.Player) {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.Dx = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.Dx = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.Dy = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.Dy = 2
	}
}

func (g *Game) Update() error {
	g.player.Dx = 0
	g.player.Dy = 0
  setKeyBindings(g.player)

  entities.UpdatePlayer(g.player, g.colliders)
	entities.UpdateEnemies(g.enemies, g.player, g.colliders)
	entities.UpdatePotions(g.potions, g.player)

	g.cam.FollowTarget(g.player.X+8, g.player.Y+8, 320, 240)
	g.cam.Constrain(
		float64(g.tilemapJSON.Layers[0].Width)*16,
		float64(g.tilemapJSON.Layers[0].Height)*16,
		320,
		240,
	)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	opts := ebiten.DrawImageOptions{}

	// loop over the layers
	for layerIndex, layer := range g.tilemapJSON.Layers {
		for index, id := range layer.Data {
			if id == 0 {
				continue
			}

			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			img := g.tilesets[layerIndex].Img(id)
      // if layerIndex == 1 {
      //   g.colliders = append(
      //     g.colliders,
      //     image.Rect(
      //       x, 
      //       y, 
      //       g.tilesets[layerIndex].Img(id).Bounds().Size().X, 
      //       g.tilesets[layerIndex].Img(id).Bounds().Dy(),
      //     ),
      //   )
      // }
			opts.GeoM.Translate(float64(x), float64(y))
			opts.GeoM.Translate(g.cam.X, g.cam.Y)
			opts.GeoM.Translate(0.0, -(float64(img.Bounds().Dy()) + 16))

			screen.DrawImage(img, &opts)

			// srcX := (id - 1) % 22
			// srcY := (id - 1) / 22
			//
			// srcX *= 16
			// srcY *= 16
			//
			// opts.GeoM.Translate(float64(x), float64(y))
			// opts.GeoM.Translate(g.cam.X, g.cam.Y)
			//
			// screen.DrawImage(
			// 	g.tilemapImage.SubImage(
			// 		image.Rect(srcX, srcY, srcX+16, srcY+16),
			// 	).(*ebiten.Image),
			// 	&opts,
			// )
			opts.GeoM.Reset()
		}
	}

	opts.GeoM.Translate(g.player.X, g.player.Y)
	opts.GeoM.Translate(g.cam.X, g.cam.Y)

	playerFrame := 0
	activeAnimation := g.player.ActiveAnimation(int(g.player.Dx), int(g.player.Dy))
	if activeAnimation != nil {
		playerFrame = activeAnimation.Frame()
	}

	screen.DrawImage(g.player.Image.SubImage(g.playerSpritesheet.Rect(playerFrame)).(*ebiten.Image), &opts)
	opts.GeoM.Reset()

  entities.DrawEnemies(g.enemies, screen, g.cam)
  entities.DrawPotions(g.potions, screen, g.cam)


	for _, collider := range g.colliders {
		vector.StrokeRect(
			screen,
			float32(collider.Min.X)+float32(g.cam.X),
			float32(collider.Min.Y)+float32(g.cam.Y),
			float32(collider.Dx()),
			float32(collider.Dy()),
			1.0,
			color.RGBA{255, 0, 0, 255},
			true,
		)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeigh int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/pig.png")
	if err != nil {
		log.Fatal(err)
	}
	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	potionImg, _, err := ebitenutil.NewImageFromFile("assets/images/lifePot.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImage, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	tilesets, err := tilemapJSON.GenTilesets()
	if err != nil {
		log.Fatal(err)
	}

	playerSpriteSheet := spriteSheet.NewSpriteSheet(4, 7, 16)

	game := Game{
		player: &entities.Player{
			Sprite: &entities.Sprite{
				Image: playerImg,
				Y:     50,
				X:     50,
				Animations: map[entities.PlayerState]*animations.Animation{
					entities.Down:  animations.NewAnimation(4, 12, 4, 5.0),
					entities.Up:    animations.NewAnimation(5, 13, 4, 5.0),
					entities.Left:  animations.NewAnimation(6, 14, 4, 5.0),
					entities.Right: animations.NewAnimation(7, 15, 4, 5.0),
				},
			},
			Health: 100,
		},
		playerSpritesheet: playerSpriteSheet,
		enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Image: skeletonImg,
					Y:     100,
					X:     100,
					// Animations: map[entities.PlayerState]*animations.Animation{
					//   entities.Down:  animations.NewAnimation(4, 12, 4, 5.0),
					//   entities.Up:    animations.NewAnimation(5, 13, 4, 5.0),
					//   entities.Left:  animations.NewAnimation(6, 14, 4, 5.0),
					//   entities.Right: animations.NewAnimation(7, 15, 4, 5.0),
					// },
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &entities.Sprite{
					Image: skeletonImg,
					Y:     200,
					X:     200,
				},
				FollowsPlayer: true,
			},
		},
		potions: []*entities.Potion{
			{
				Sprite: &entities.Sprite{
					Image: potionImg,
					Y:     210,
					X:     100,
				},
				AmtHeal: 1,
			},
		},
		tilemapJSON:  tilemapJSON,
		tilemapImage: tilemapImage,
		tilesets:     tilesets,
		cam:          cam.NewCamera(0, 0),
		colliders: []image.Rectangle{
			image.Rect(100, 100, 116, 116),
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
