package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point2D interface {
	GetCoords() (int, int)
	ChangeCoords(int, int)
	PlotPixel()
}

type point2D struct {
	screen          *ebiten.Image
	X               int
	Y               int
	col             color.Color
	backgroundColor color.Color
}

func NewPoint2D(screen *ebiten.Image, backgroundCol color.Color, x, y int, col color.Color) Point2D {
	return &point2D{
		screen:          screen,
		X:               x,
		Y:               y,
		col:             col, // Нулевое значение для интерфейса color.Color
		backgroundColor: backgroundCol,
	}
}

func (primitive *point2D) PlotPixel() {
	primitive.screen.Set(primitive.X, primitive.Y, primitive.col)
}
func (primitive *point2D) GetCoords() (int, int) {
	return primitive.X, primitive.Y
}

func (primitive *point2D) ChangeCoords(x int, y int) {
	primitive.X = x
	primitive.Y = y
}
