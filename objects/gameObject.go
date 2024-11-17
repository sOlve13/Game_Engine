package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	GetScreen() *ebiten.Image
	GetBackgroundColor() color.Color
	SetScreen(screen *ebiten.Image)
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

func NewWScreenGameObject(backgroundColor color.Color) GameObject {
	return &gameObject{
		screen:          nil,
		backgroundColor: backgroundColor,
	}
}
func (gameObject *gameObject) SetScreen(screen *ebiten.Image) {
	gameObject.screen = screen
}
func (gameObject *gameObject) GetScreen() *ebiten.Image {
	return gameObject.screen
}

func (gameObject *gameObject) GetBackgroundColor() color.Color {
	return gameObject.backgroundColor
}
