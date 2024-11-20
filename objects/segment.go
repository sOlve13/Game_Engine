package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// LineSegment defines the interface for working with line segments.
// Includes methods for drawing, updating points, and retrieving coordinates.
type LineSegment interface {
	// Draws a line segment using the Bresenham algorithm.
	// @param startPoint Point2D: Starting point of the line.
	// @param finalPoint Point2D: Ending point of the line.
	// @param col color.Color: The color of the line.
	// @return error: Returns an error if drawing fails.
	Segment(startPoint Point2D, finalPoint Point2D, col color.Color) error

	// Draws a line segment using the default stroke rendering.
	// @param startPoint Point2D: Starting point of the line.
	// @param finalPoint Point2D: Ending point of the line.
	// @param col color.Color: The color of the line.
	SegmentDefault(startPoint Point2D, finalPoint Point2D, col color.Color)

	// Updates the ending point of the line segment.
	// @param newPoint Point2D: The new ending point.
	ChangeFinal(newPoint Point2D)

	// Updates the starting point of the line segment.
	// @param newPoint Point2D: The new starting point.
	ChangeStart(newPoint Point2D)

	// Retrieves the coordinates of the ending point.
	// @return (int, int): X and Y coordinates of the ending point.
	GetFinal() (int, int)

	// Retrieves the coordinates of the starting point.
	// @return (int, int): X and Y coordinates of the starting point.
	GetStart() (int, int)
}

// lineSegment is a concrete implementation of the LineSegment interface.
type lineSegment struct {
	screen          *ebiten.Image
	startPoint      Point2D
	finalPoint      Point2D
	col             color.Color
	backgroundColor color.Color
}

// NewLineSegment creates a new instance of a line segment.
// @param screen *ebiten.Image: The screen to draw the line on.
// @param backgroundColor color.Color: The background color of the screen.
// @return LineSegment: A new LineSegment instance.
func NewLineSegment(screen *ebiten.Image, backgroundColor color.Color) LineSegment {
	return &lineSegment{
		screen:          screen,
		startPoint:      nil,
		finalPoint:      nil,
		col:             nil, // Нулевое значение для интерфейса color.Color
		backgroundColor: backgroundColor,
	}
}

// plotPixel draws a single pixel at the specified coordinates.
// @param x, y int: Coordinates of the pixel.
// @param col color.Color: The color of the pixel.
func (primitive *lineSegment) plotPixel(x int, y int, col color.Color) {
	primitive.screen.Set(x, y, col)

}

// Segment draws a line segment using the Bresenham algorithm.
// @param startPoint Point2D: The starting point of the line.
// @param finalPoint Point2D: The ending point of the line.
// @param col color.Color: The color of the line.
func (primitive *lineSegment) Segment(startPoint Point2D, finalPoint Point2D, col color.Color) error {
	startX, startY := startPoint.GetCoords()
	finalX, finalY := finalPoint.GetCoords()
	primitive.col = col
	primitive.startPoint = startPoint
	primitive.finalPoint = finalPoint

	// Разница координат
	deltX := abs(finalX - startX)
	deltY := abs(finalY - startY)

	// Определяем направление шага
	stepX := 1
	if startX > finalX {
		stepX = -1
	}
	stepY := 1
	if startY > finalY {
		stepY = -1
	}

	err := deltX - deltY

	x, y := startX, startY
	for {
		primitive.plotPixel(x, y, col)
		if x == finalX && y == finalY {
			break
		}
		e2 := err * 2
		if e2 > -deltY {
			err -= deltY
			x += stepX
		}
		if e2 < deltX {
			err += deltX
			y += stepY
		}
	}

	return nil
}

// SegmentDefault draws a line segment using the default vector stroke rendering.
// @param startPoint Point2D: The starting point of the line.
// @param finalPoint Point2D: The ending point of the line.
// @param col color.Color: The color of the line.
func (primitive *lineSegment) SegmentDefault(startPoint Point2D, finalPoint Point2D, col color.Color) {
	primitive.startPoint = startPoint
	primitive.finalPoint = finalPoint
	x1, y1 := startPoint.GetCoords()
	x2, y2 := finalPoint.GetCoords()
	primitive.col = col
	x1_ := float32(x1)
	x2_ := float32(x2)
	y1_ := float32(y1)
	y2_ := float32(y2)
	vector.StrokeLine(primitive.screen, x1_, y1_, x2_, y2_, 1, col, false)
}

// ChangeStart updates the starting point of the line segment.
// @param newPoint Point2D: The new starting point of the line.
func (primitive *lineSegment) ChangeStart(newPoint Point2D) {
	color := primitive.col
	primitive.Segment(primitive.startPoint, primitive.finalPoint, primitive.backgroundColor)
	primitive.Segment(newPoint, primitive.finalPoint, color)
}

// ChangeFinal updates the ending point of the line segment.
// @param newPoint Point2D: The new ending point of the line.
func (primitive *lineSegment) ChangeFinal(newPoint Point2D) {
	color := primitive.col
	primitive.Segment(primitive.startPoint, primitive.finalPoint, primitive.backgroundColor)
	primitive.Segment(primitive.startPoint, newPoint, color)
}

// GetFinal retrieves the coordinates of the ending point.
// @return (int, int): X and Y coordinates of the ending point.
func (primitive *lineSegment) GetFinal() (int, int) {
	return primitive.finalPoint.GetCoords()
}

// GetStart retrieves the coordinates of the starting point.
// @return (int, int): X and Y coordinates of the starting point.
func (primitive *lineSegment) GetStart() (int, int) {
	return primitive.startPoint.GetCoords()
}
