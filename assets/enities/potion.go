package entities

import (
	"fmt"
	"time"
	"videogame/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type Potion struct {
	*Sprite
	AmtHeal uint
}

func UpdatePotions(potions []*Potion, player *Player) {
	for _, potion := range potions {
		if player.X > potion.X &&
			player.X < potion.X + float64(potion.Image.Bounds().Dx()) &&
			player.Y > potion.Y &&
			player.Y < potion.Y - float64(potion.Image.Bounds().Dy()) {

			player.Health += potion.AmtHeal
			fmt.Println("Player healed by", potion.AmtHeal, time.Now())

		}
	}
}

func DrawPotions(potions []*Potion, screen *ebiten.Image, cam *camera.Camera) {
	for _, potion := range potions {
		potion.Sprite.Draw(screen, cam)
	}
}
