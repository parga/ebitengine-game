package entities

import (
	"image"
	"videogame/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type Potion struct {
	*Sprite
	AmtHeal uint
}

func UpdatePotions(potions []*Potion, player *Player) []*Potion{
	for index, potion := range potions {
		haveCollide := RectsCollide(
			image.Rect(int(player.X), int(player.Y), int(player.X+16), int(player.Y+16)),
			image.Rect(int(potion.X), int(potion.Y), int(potion.X+8), int(potion.Y+8)),
		)

		if haveCollide {
			player.Health += potion.AmtHeal
			potions = append(potions[:index], potions[index+1:]...)
		}
	}
	return potions
}

func DrawPotions(potions []*Potion, screen *ebiten.Image, cam *camera.Camera) {
	for _, potion := range potions {
		potion.Sprite.Draw(screen, cam)
	}
}
