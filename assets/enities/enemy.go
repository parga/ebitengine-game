package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"videogame/camera"
)

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

func UpdateEnemies(enemies []*Enemy, player *Player, colliders []image.Rectangle) {
	for _, sprite := range enemies {
		sprite.Dy = 0
		sprite.Dx = 0
		if sprite.FollowsPlayer {
			if sprite.X < player.X {
				sprite.Dx = 1
			} else if sprite.X > player.X {
				sprite.Dx = -1
			}
			if sprite.Y < player.Y {
				sprite.Dy = 1
			} else if sprite.Y > player.Y {
				sprite.Dy = -1
			}
		}
		sprite.X += sprite.Dx
		CheckCollisionHorizontal(sprite.Sprite, colliders)

		sprite.Y += sprite.Dy
		CheckCollisionVertical(sprite.Sprite, colliders)
	}
}

func DrawEnemies(enemies []*Enemy, screen *ebiten.Image, cam *camera.Camera) {
	for _, enemy := range enemies {
		enemy.Sprite.Draw(screen, cam)
	}
}
