package entities

import (
	"fmt"
	"videogame/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type Potion struct {
	*Sprite
	AmtHeal uint
}

func UpdatePotions(potions []*Potion, player *Player) {
	for _, potion := range potions {
		if player.X > potion.X {
			player.Health += potion.AmtHeal
			fmt.Println("Player healed by", potion.AmtHeal)
		}
	}
}

func DrawPotions(potions []*Potion, screen *ebiten.Image, cam *camera.Camera) {
  for _, potion := range potions {
    potion.Sprite.Draw(screen, cam)
  }
}
