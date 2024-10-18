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
	PlotPixel(int, int, color.Color, ...bool)
	Segment(int, int, int, int, color.Color, ...bool) error
	DrawSquare(int, int, int, color.Color) error
}

type primitiveRendererСlass struct {
	screen        *ebiten.Image
	startX        int
	startY        int
	finalSegX     int
	finalSegY     int
	S             int
	col           color.Color
	primitiveType string
}

func NewPrimitiveRendererclass(screen *ebiten.Image) PrimitiveRendererСlass {
	return &primitiveRendererСlass{
		screen:        screen,
		startX:        0,
		startY:        0,
		finalSegX:     0,
		finalSegY:     0,
		S:             0,
		col:           nil, // Нулевое значение для интерфейса color.Color
		primitiveType: "",
	}
}
func (primitive *primitiveRendererСlass) PlotPixel(x int, y int, col color.Color, flags ...bool) {
	var flag bool
	if len(flags) > 0 {
		flag = flags[0]
	} else {
		flag = true
	}
	primitive.screen.Set(x, y, col)

	if flag {
		primitive.startX = x
		primitive.startY = y
		primitive.col = col
		primitive.primitiveType = "dot"
	}
}

func (primitive *primitiveRendererСlass) Segment(startX int, startY int, finalX int, finalY int, col color.Color, flags ...bool) error {
	var flag bool
	var err error
	if len(flags) > 0 {
		flag = flags[0]
	} else {
		flag = true
	}
	if flag {
		primitive.startX = startX
		primitive.startY = startY
		primitive.finalSegX = finalX
		primitive.finalSegY = finalY
		primitive.col = col
		primitive.primitiveType = "segment"
	}
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
			primitive.PlotPixel(startX, y, col, false)
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
			primitive.PlotPixel(x, int(y), col, false)
			y += slope // Increment y
		}
	} else { // Case where |slope| > 1, swap roles of x and y
		x := float64(startX)
		step := 1
		if deltY < 0 {
			step = -1
		}
		for y := startY; y != finalY+step; y += step {
			primitive.PlotPixel(int(x), y, col, false) // Plot at rounded (x, y)
			x += 1 / slope                             // Increment x
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

	primitive.Segment(X, Y, X+S, Y, col, false)
	primitive.Segment(X, Y, X, Y+S, col, false)
	primitive.Segment(X+S, Y, X+S, Y+S, col, false)
	primitive.Segment(X, Y+S, X+S, Y+S, col, false)

	primitive.startX = X
	primitive.startY = Y
	primitive.S = S
	primitive.primitiveType = "square"

	return nil
}
