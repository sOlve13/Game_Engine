package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	GetScreen() *ebiten.Image
	GetBackgroundColor() color.Color
}

type gameObject struct {
	screen          *ebiten.Image
	backgroundColor color.Color
}

func NewGameObject(screen *ebiten.Image, backgroundColor color.Color) GameObject {
	return &gameObject{
		screen:          screen,
		backgroundColor: backgroundColor,
	}
}

func (gameObject *gameObject) GetScreen() *ebiten.Image {
	return gameObject.screen
}

func (gameObject *gameObject) GetBackgroundColor() color.Color {
	return gameObject.backgroundColor
}
