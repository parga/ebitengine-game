package entities

import (
	"image"
	"videogame/animations"
	"videogame/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerState uint8

const (
	Down PlayerState = iota
	Up
	Left
	Right
)

type Sprite struct {
	Id           int
	Image        *ebiten.Image
	X, Y, Dx, Dy float64
	Animations   map[PlayerState]*animations.Animation
}

func (s *Sprite) Draw(screen *ebiten.Image, cam *camera.Camera) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.X, s.Y)
	opts.GeoM.Translate(cam.X, cam.Y)
	screen.DrawImage(s.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
	opts.GeoM.Reset()
}
