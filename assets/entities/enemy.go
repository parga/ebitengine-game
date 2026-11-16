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
	blockers := append([]image.Rectangle{}, colliders...)
	blockers = append(blockers, spriteBounds(player.Sprite))

	for index, sprite := range enemies {
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
		previousX := sprite.X
		sprite.X += sprite.Dx
		CheckCollisionHorizontal(sprite.Sprite, blockers)
		if overlapsOtherEnemy(sprite, index, enemies) {
			sprite.X = previousX
		}

		previousY := sprite.Y
		sprite.Y += sprite.Dy
		CheckCollisionVertical(sprite.Sprite, blockers)
		if overlapsOtherEnemy(sprite, index, enemies) {
			sprite.Y = previousY
		}
	}
}

func overlapsOtherEnemy(enemy *Enemy, enemyIndex int, enemies []*Enemy) bool {
	for index, other := range enemies {
		if index == enemyIndex {
			continue
		}
		if spriteBounds(enemy.Sprite).Overlaps(spriteBounds(other.Sprite)) {
			return true
		}
	}
	return false
}

func spriteBounds(sprite *Sprite) image.Rectangle {
	return image.Rect(int(sprite.X), int(sprite.Y), int(sprite.X)+16, int(sprite.Y)+16)
}

func DrawEnemies(enemies []*Enemy, screen *ebiten.Image, cam *camera.Camera) {
	for _, enemy := range enemies {
		enemy.Sprite.Draw(screen, cam)
	}
}
