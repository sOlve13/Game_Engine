package objects

import (
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
