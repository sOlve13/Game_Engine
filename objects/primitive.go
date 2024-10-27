package objects

import (
	"errors"
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
	DrawCircle(Point2D, int, color.Color)
	FillSquare(int, int, int, color.Color)
	DrawPolygon([]Point2D, color.Color) error
	FloodFill(int, int, color.Color, color.Color)
	BorderFill(int, int, color.Color, color.Color)
	DrawEllipse(Point2D, int, int, color.Color)
}

type primitiveRendererСlass struct {
	screen *ebiten.Image
	startX int
	startY int

	S               int
	col             color.Color
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

func (primitive *primitiveRendererСlass) DrawCircle(center Point2D, radius int, col color.Color) {
	centerX, centerY := center.GetCoords()

	x := 0
	y := radius
	decison := 1 - radius

	plotSymmetricPoints := func(x, y int) {
		primitive.plotPixel(centerX+x, centerY+y, col) // Octant 1
		primitive.plotPixel(centerX-x, centerY+y, col) // Octant 2
		primitive.plotPixel(centerX+x, centerY-y, col) // Octant 8
		primitive.plotPixel(centerX-x, centerY-y, col) // Octant 7
		primitive.plotPixel(centerX+y, centerY+x, col) // Octant 3
		primitive.plotPixel(centerX-y, centerY+x, col) // Octant 4
		primitive.plotPixel(centerX+y, centerY-x, col) // Octant 6
		primitive.plotPixel(centerX-y, centerY-x, col) // Octant 5
	}

	plotSymmetricPoints(x, y)

	for x < y {
		// Update the decision parameter based on the current point
		if decison < 0 {
			// Move horizontally
			decison += 2*x + 1 // Corrected from +3 to +1
		} else {
			// Move diagonally
			decison += 2*(x-y) + 1
			y-- // Decrease y when moving diagonally
		}
		x++

		// Plot the symmetric points for the new (x, y)
		plotSymmetricPoints(x, y)
	}
}

func (primitive *primitiveRendererСlass) DrawEllipse(center Point2D, a int, b int, col color.Color) {
	centerX, centerY := center.GetCoords()
	x := 0
	y := b
	a2 := a * a
	b2 := b * b
	decision := b2 - (a2 * b) + (a2 / 4)

	// Function to plot symmetric points
	plotSymmetricPoints := func(x, y int) {
		primitive.plotPixel(centerX+x, centerY+y, col) // Quadrant I
		primitive.plotPixel(centerX-x, centerY+y, col) // Quadrant II
		primitive.plotPixel(centerX+x, centerY-y, col) // Quadrant IV
		primitive.plotPixel(centerX-x, centerY-y, col) // Quadrant III
	}

	// Draw the ellipse in the first half
	plotSymmetricPoints(x, y)
	for (b2 * x) <= (a2 * y) {
		x++
		if decision < 0 {
			decision += b2 * (2*x + 1)
		} else {
			y--
			decision += b2*(2*x+1) - 2*a2*y
		}
		plotSymmetricPoints(x, y)
	}

	// Now handle the lower half of the ellipse
	x = a
	y = 0
	decision = a2 - (b2 * a) + (b2 / 4)

	// Draw the ellipse in the second half
	plotSymmetricPoints(x, y)
	for (a2 * y) <= (b2 * x) {
		y++
		if decision < 0 {
			decision += a2 * (2*y + 1)
		} else {
			x--
			decision += a2*(2*y+1) - 2*b2*x
		}
		plotSymmetricPoints(x, y)
	}
}

func (pr *primitiveRendererСlass) DrawPolygon(points []Point2D, lineColor color.Color) error {
	if len(points) < 2 {
		return errors.New("Polygon can't consist of < 3 points")
	}
	st_x, st_y := points[0].GetCoords()
	fn_x, fn_y := points[len(points)-1].GetCoords()
	if st_x != fn_x && st_y != fn_y {
		return errors.New("First and last points should be same")
	}

	for i := 0; i < len(points)-1; i++ {
		startPoint := points[i]
		endPoint := points[i+1]
		line := NewLineSegment(pr.screen, pr.backgroundColor) // Use transparent as background
		line.Segment(startPoint, endPoint, lineColor)
		pr.lines = append(pr.lines, line)
	}
	pr.col = lineColor
	minX := 100000
	minY := 100000
	maxX := 0
	maxY := 0
	for _, p := range points {

		x, y := p.GetCoords()
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}

		if y < minY {
			minY = y
		}
	}
	medX := (minX + maxX) / 2
	medY := (minY + maxY) / 2
	for i := minX; i < maxX; i++ {

		if isPointInPolygon(NewPoint2D(pr.screen, pr.backgroundColor, i, medY, pr.col), points, pr.screen, pr.backgroundColor) {
			pr.FloodFill(i+1, medY+1, pr.col, pr.backgroundColor)

			return nil
		}
	}
	for i := minY; i < maxY; i++ {
		if isPointInPolygon(NewPoint2D(pr.screen, pr.backgroundColor, medX, i, pr.col), points, pr.screen, pr.backgroundColor) {
			pr.FloodFill(medX+1, i+1, pr.col, pr.backgroundColor)
			return nil
		}
	}

	return errors.New("Point was not found")
}

func (primitive *primitiveRendererСlass) FillSquare(x int, y int, s int, col color.Color) {
	for i := x; i <= x+s; i++ {
		for j := y; j <= y+s; j++ {
			primitive.plotPixel(i, j, col)
		}
	}
}

func (primitive *primitiveRendererСlass) BorderFill(x int, y int, fillColor color.Color, borderColor color.Color) {
	if primitive.screen.At(x, y) == borderColor || primitive.screen.At(x, y) == fillColor {
		return
	}
	primitive.plotPixel(x, y, fillColor)
	primitive.BorderFill(x+1, y, fillColor, borderColor)
	primitive.BorderFill(x-1, y, fillColor, borderColor)
	primitive.BorderFill(x, y+1, fillColor, borderColor)
	primitive.BorderFill(x, y-1, fillColor, borderColor)
}

func (primitive *primitiveRendererСlass) FloodFill(x, y int, fillColor color.Color, boundaryColor color.Color) {
	width, height := primitive.screen.Size()
	originalColor := primitive.screen.At(x, y)

	// Check if the initial pixel is already the fill color or is the boundary color
	if originalColor == fillColor || originalColor == boundaryColor {
		return // No need to fill
	}

	// Define the recursive function inside FloodFill
	var floodFillRecursive func(x, y int)
	floodFillRecursive = func(x, y int) {
		// Check bounds
		if x < 0 || x >= width || y < 0 || y >= height {
			return // Out of bounds
		}

		currentColor := primitive.screen.At(x, y)
		if currentColor != originalColor {
			return // Not the original color, stop recursion
		}

		// Set the pixel to the fill color
		primitive.screen.Set(x, y, fillColor)

		// Recursive calls for neighboring pixels
		floodFillRecursive(x+1, y) // Right
		floodFillRecursive(x-1, y) // Left
		floodFillRecursive(x, y+1) // Down
		floodFillRecursive(x, y-1) // Up
	}

	// Call the recursive function
	floodFillRecursive(x, y)
}
