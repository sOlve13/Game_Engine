package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point2D interface {
	PlotPixel(int, int, color.Color)
}

type point2D struct {
	screen          *ebiten.Image
	X               int
	Y               int
	col             color.Color
	backgroundColor color.Color
}

func NewPoint2D(screen *ebiten.Image, backfroundCol color.Color) Point2D {
	return &point2D{
		screen:          screen,
		X:               0,
		Y:               0,
		col:             nil, // Нулевое значение для интерфейса color.Color
		backgroundColor: backfroundCol,
	}
}

func (primitive *point2D) PlotPixel(x int, y int, col color.Color) {
	primitive.screen.Set(x, y, col)

}

func (primitive *point2D) GetCoords() (int, int) {
	return primitive.X, primitive.Y
}

func (primitive *point2D) ChangeCoords(x int, y int) {
	primitive.screen.Set(primitive.X, primitive.Y, primitive.backgroundColor)
	primitive.screen.Set(x, y, primitive.col)
}
