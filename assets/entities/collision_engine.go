package entities

import (
	"image"
)

func CheckCollisionHorizontal(sprite *Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(int(sprite.X), int(sprite.Y), int(sprite.X)+16, int(sprite.Y)+16)) {
			if sprite.Dx > 0 {
				sprite.X = float64(collider.Min.X) - 16
			} else if sprite.Dx < 0 {
				sprite.X = float64(collider.Max.X)
			}
		}
	}
}

func CheckCollisionVertical(sprite *Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(int(sprite.X), int(sprite.Y), int(sprite.X)+16, int(sprite.Y)+16)) {
			if sprite.Dy > 0 {
				sprite.Y = float64(collider.Min.Y) - 16
			} else if sprite.Dy < 0 {
				sprite.Y = float64(collider.Max.Y)
			}
		}
	}

}

func RectsCollide(r1, r2 image.Rectangle) bool {
	return r1.Min.X < r2.Max.X && r1.Max.X > r2.Min.X && r1.Min.Y < r2.Max.Y && r1.Max.Y > r2.Min.Y
}
