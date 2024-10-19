package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point2D interface {
	GetCoords() (int, int)
	ChangeCoords(int, int)
}

type point2D struct {
	screen          *ebiten.Image
	X               int
	Y               int
	col             color.Color
	backgroundColor color.Color
}

func NewPoint2D(screen *ebiten.Image, backfroundCol color.Color, x, y int, col color.Color) Point2D {
	screen.Set(x, y, col)
	return &point2D{
		screen:          screen,
		X:               x,
		Y:               y,
		col:             nil, // Нулевое значение для интерфейса color.Color
		backgroundColor: backfroundCol,
	}
}

func (primitive *point2D) GetCoords() (int, int) {
	return primitive.X, primitive.Y
}

func (primitive *point2D) ChangeCoords(x int, y int) {
	primitive.screen.Set(primitive.X, primitive.Y, primitive.backgroundColor)
	primitive.screen.Set(x, y, primitive.col)
}
