package objects

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type LineSegment interface {
	Segment(Point2D, Point2D, color.Color) error
	SegmentDefault(Point2D, Point2D, color.Color)
	ChangeFinal(Point2D)
	ChangeStart(Point2D)
	GetFinal() (int, int)
	GetStart() (int, int)
}

type lineSegment struct {
	screen          *ebiten.Image
	startPoint      Point2D
	finalPoint      Point2D
	col             color.Color
	backgroundColor color.Color
}

func NewLineSegment(screen *ebiten.Image, backgroundColor color.Color) LineSegment {
	return &lineSegment{
		screen:          screen,
		startPoint:      nil,
		finalPoint:      nil,
		col:             nil, // Нулевое значение для интерфейса color.Color
		backgroundColor: backgroundColor,
	}
}
func (primitive *lineSegment) plotPixel(x int, y int, col color.Color) {
	primitive.screen.Set(x, y, col)

}

func (primitive *lineSegment) Segment(startPoint Point2D, finalPoint Point2D, col color.Color) error {
	var err error
	startX, startY := startPoint.GetCoords()
	finalX, finalY := finalPoint.GetCoords()
	primitive.col = col
	primitive.startPoint = startPoint
	primitive.finalPoint = finalPoint
	deltX := finalX - startX
	deltY := finalY - startY
	if deltX == 0 && deltY == 0 {
		err = fmt.Errorf("Line can't be 0")
		return err // No line to draw
	}

	if deltX == 0 { // Vertical line case
		step := 1
		if deltY < 0 {
			step = -1
		}
		for y := startY; y != finalY+step; y += step {
			primitive.plotPixel(startX, y, col)
		}
		return nil
	}

	var slope float64
	if deltX != 0 {
		slope = float64(deltY) / float64(deltX) // Calculate slope
	}

	if absolute(slope) <= 1 { // Case where |slope| <= 1
		y := float64(startY)
		step := 1
		if deltX < 0 {
			step = -1
		}
		for x := startX; x != finalX+step; x += step {
			primitive.plotPixel(x, int(y), col)
			y += slope // Increment y
		}
	} else { // Case where |slope| > 1, swap roles of x and y
		x := float64(startX)
		step := 1
		if deltY < 0 {
			step = -1
		}
		for y := startY; y != finalY+step; y += step {
			primitive.plotPixel(int(x), y, col) // Plot at rounded (x, y)
			x += 1 / slope                      // Increment x
		}
	}

	return nil
}
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

func (primitive *lineSegment) ChangeStart(newPoint Point2D) {
	color := primitive.col
	primitive.Segment(primitive.startPoint, primitive.finalPoint, primitive.backgroundColor)
	primitive.Segment(newPoint, primitive.finalPoint, color)
}

func (primitive *lineSegment) ChangeFinal(newPoint Point2D) {
	color := primitive.col
	primitive.Segment(primitive.startPoint, primitive.finalPoint, primitive.backgroundColor)
	primitive.Segment(primitive.startPoint, newPoint, color)
}

func (primitive *lineSegment) GetFinal() (int, int) {
	return primitive.finalPoint.GetCoords()
}

func (primitive *lineSegment) GetStart() (int, int) {
	return primitive.startPoint.GetCoords()
}
