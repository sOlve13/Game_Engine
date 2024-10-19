package objects

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func absolute(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

type PrimitiveRendererСlass interface {
	plotPixel(int, int, color.Color)
	segment(int, int, int, int, color.Color) error
	DrawSquare(int, int, int, color.Color) error
	DrawPolyline([]Point2D, color.Color)
}

type primitiveRendererСlass struct {
	screen          *ebiten.Image
	startX          int
	startY          int
	finalSegX       int
	finalSegY       int
	S               int
	col             color.Color
	primitiveType   string
	backgroundColor color.Color
	lines           []LineSegment
}

func NewPrimitiveRendererclass(screen *ebiten.Image, backgroundColor color.Color) PrimitiveRendererСlass {
	return &primitiveRendererСlass{
		screen:          screen,
		startX:          0,
		startY:          0,
		S:               0,
		col:             nil, // Нулевое значение для интерфейса color.Color
		primitiveType:   "",
		backgroundColor: backgroundColor,
		lines:           make([]LineSegment, 0),
	}
}
func (primitive *primitiveRendererСlass) plotPixel(x int, y int, col color.Color) {

	primitive.screen.Set(x, y, col)
}

func (primitive *primitiveRendererСlass) segment(startX int, startY int, finalX int, finalY int, col color.Color) error {
	var err error

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

func (primitive *primitiveRendererСlass) DrawSquare(X int, Y int, S int, col color.Color) error {
	var err error
	if X <= 0 && Y <= 0 && S < 1 {
		err = fmt.Errorf("Square should be on the screen and not smaller than 1 px")
		return err
	}

	primitive.segment(X, Y, X+S, Y, col)
	primitive.segment(X, Y, X, Y+S, col)
	primitive.segment(X+S, Y, X+S, Y+S, col)
	primitive.segment(X, Y+S, X+S, Y+S, col)

	primitive.startX = X
	primitive.startY = Y
	primitive.S = S
	primitive.primitiveType = "square"

	return nil
}

func (pr *primitiveRendererСlass) DrawPolyline(points []Point2D, lineColor color.Color) {
	if len(points) < 2 {
		return // Need at least two points to draw a polyline
	}

	for i := 0; i < len(points)-1; i++ {
		startPoint := points[i]
		endPoint := points[i+1]
		line := NewLineSegment(pr.screen, color.Transparent) // Use transparent as background
		line.Segment(startPoint, endPoint, lineColor)
		pr.lines = append(pr.lines, line)
	}
	pr.col = lineColor

}
