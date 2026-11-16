package entities

import (
	"image"
	"videogame/animations"
)

type Player struct {
	*Sprite
	Health uint
}

func (p *Player) ActiveAnimation(dx, dy int) *animations.Animation {
	if dx > 0 {
		return p.Animations[Right]
	}
	if dx < 0 {
		return p.Animations[Left]
	}
	if dy > 0 {
		return p.Animations[Down]
	}
	if dy < 0 {
		return p.Animations[Up]
	}
	return nil
}

func UpdatePlayer(player *Player, colliders []image.Rectangle) {

  player.X += player.Dx
  CheckCollisionHorizontal(player.Sprite, colliders)

  player.Y += player.Dy
  CheckCollisionVertical(player.Sprite, colliders)

  activeAnimation := player.ActiveAnimation(int(player.Dx), int(player.Dy))
  if activeAnimation != nil {
    activeAnimation.Update()
  }
}
