package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Point2D represents a 2D point interface with basic operations.
type Point2D interface {
	GetCoords() (int, int) // Retrieves the coordinates of the point.
	ChangeCoords(int, int) // Updates the coordinates of the point.
	PlotPixel()            // Draws the point on the screen.
}

// point2D is a concrete implementation of the Point2D interface.
type point2D struct {
	screen          *ebiten.Image // The screen where the point is drawn.
	X               int           // X-coordinate of the point.
	Y               int           // Y-coordinate of the point.
	col             color.Color   // The color of the point.
	backgroundColor color.Color   // Background color of the screen.
}

// NewPoint2D creates a new point2D instance.
// @param screen *ebiten.Image: The screen where the point will be drawn.
// @param backgroundCol color.Color: The background color of the screen.
// @param x, y int: Initial coordinates of the point.
// @param col color.Color: The color of the point.
// @return Point2D: A new point2D instance.
func NewPoint2D(screen *ebiten.Image, backgroundCol color.Color, x, y int, col color.Color) Point2D {
	return &point2D{
		screen:          screen,
		X:               x,
		Y:               y,
		col:             col,
		backgroundColor: backgroundCol,
	}
}

// PlotPixel draws the point on the screen using its defined color.
func (primitive *point2D) PlotPixel() {
	primitive.screen.Set(primitive.X, primitive.Y, primitive.col)
}

// GetCoords retrieves the current coordinates of the point.
// @return (int, int): X and Y coordinates of the point.
func (primitive *point2D) GetCoords() (int, int) {
	return primitive.X, primitive.Y
}

// ChangeCoords updates the coordinates of the point.
// @param x, y int: New coordinates for the point.
func (primitive *point2D) ChangeCoords(x int, y int) {
	primitive.X = x
	primitive.Y = y
}
