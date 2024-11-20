package objects

import (
	"errors"
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// absolute calculates the absolute value of a float64 number.
// @param num float64: The number to compute the absolute value for.
// @return float64: The absolute value.
func absolute(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

// PrimitiveRendererClass defines the interface for rendering various shapes and primitives.
type PrimitiveRendererСlass interface {
	// Draws a single pixel on the screen.
	// @param x, y int: Coordinates of the pixel.
	// @param col color.Color: The color of the pixel.
	plotPixel(int, int, color.Color)

	// Draws a line segment using Bresenham's algorithm.
	// @param startX, startY int: Starting coordinates of the line.
	// @param finalX, finalY int: Ending coordinates of the line.
	// @param col color.Color: The color of the line.
	// @return error: Returns an error if the line cannot be drawn.
	segment(int, int, int, int, color.Color) error

	// Draws a square with optional rotation.
	// @param X, Y int: Top-left corner of the square.
	// @param S int: Side length of the square.
	// @param angle int: Rotation angle in degrees.
	// @param col color.Color: The color of the square.
	// @return error: Returns an error if the square is invalid.
	DrawSquare(int, int, int, int, color.Color) error

	// Draws a polyline connecting multiple points.
	// @param points []Point2D: List of points to connect.
	// @param lineColor color.Color: The color of the polyline.
	DrawPolyline([]Point2D, color.Color)

	// Draws a circle using a midpoint algorithm.
	// @param x_, y_ int: Center of the circle.
	// @param radius int: Radius of the circle.
	// @param col color.Color: The color of the circle.
	DrawCircle(int, int, int, color.Color)

	// Fills a square area with the specified color.
	// @param x, y int: Top-left corner of the square.
	// @param s int: Side length of the square.
	// @param col color.Color: The fill color.
	FillSquare(int, int, int, color.Color)

	// Draws a polygon defined by a set of points.
	// @param points []Point2D: List of polygon vertices.
	// @param lineColor color.Color: The color of the polygon edges.
	// @return error: Returns an error if the polygon is invalid.
	DrawPolygon([]Point2D, color.Color) error

	// Fills an area using the flood-fill algorithm.
	// @param x, y int: Starting coordinates for the fill.
	// @param fillColor color.Color: The fill color.
	// @param boundaryColor color.Color: The color marking the boundaries.
	FloodFill(int, int, color.Color, color.Color)

	// Fills an area up to the boundary color using the border-fill algorithm.
	// @param x, y int: Starting coordinates for the fill.
	// @param fillColor color.Color: The fill color.
	// @param borderColor color.Color: The boundary color.
	BorderFill(int, int, color.Color, color.Color)

	// Draws an ellipse using a midpoint algorithm.
	// @param center Point2D: Center of the ellipse.
	// @param a, b int: Semi-major and semi-minor axes of the ellipse.
	// @param col color.Color: The color of the ellipse.
	DrawEllipse(Point2D, int, int, color.Color)
}

// primitiveRendererСlass is a concrete implementation of the PrimitiveRendererСlass interface.
type primitiveRendererСlass struct {
	screen *ebiten.Image
	startX int
	startY int

	S               int
	col             color.Color
	backgroundColor color.Color
	lines           []LineSegment
}

// NewPrimitiveRendererClass creates a new instance of the PrimitiveRendererClass.
// @param screen *ebiten.Image: The screen to draw on.
// @param backgroundColor color.Color: The background color of the screen.
// @return PrimitiveRendererClass: The created instance.
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

// Draws a single pixel on the screen.
// @param x, y int: Coordinates of the pixel.
// @param col color.Color: The color of the pixel.
func (primitive *primitiveRendererСlass) plotPixel(x int, y int, col color.Color) {

	primitive.screen.Set(x, y, col)
}

// Draws a line segment using Bresenham's algorithm.
// @param startX, startY int: Starting coordinates of the line.
// @param finalX, finalY int: Ending coordinates of the line.
// @param col color.Color: The color of the line.
// @return error: Returns an error if the line cannot be drawn.
func (primitive *primitiveRendererСlass) segment(startX int, startY int, finalX int, finalY int, col color.Color) error {
	dx := math.Abs(float64(finalX) - float64(startX))
	dy := math.Abs(float64(finalY) - float64(startY))
	sx := -1
	if startX < finalX {
		sx = 1
	}
	sy := -1
	if startY < finalY {
		sy = 1
	}
	err := dx - dy

	for {
		primitive.plotPixel(startX, startY, col)
		if startX == finalX && startY == finalY {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			startX += sx
		}
		if e2 < dx {
			err += dx
			startY += sy
		}
	}

	return nil
}

// Draws a square with optional rotation.
// @param X, Y int: Top-left corner of the square.
// @param S int: Side length of the square.
// @param angle int: Rotation angle in degrees.
// @param col color.Color: The color of the square.
// @return error: Returns an error if the square is invalid.
func (primitive *primitiveRendererСlass) DrawSquare(X int, Y int, S int, angle int, col color.Color) error {
	if X <= 0 && Y <= 0 && S < 1 {
		return fmt.Errorf("Square should be on the screen and not smaller than 1 px")
	}

	radAngle := float64(angle) * math.Pi / 180.0

	// Вычисляем смещение от центра для каждой вершины
	centrX := X + S/2
	centrY := Y + S/2

	x1, y1 := X, Y
	x2, y2 := X+S, Y
	x3, y3 := X+S, Y+S
	x4, y4 := X, Y+S

	x1, y1 = rotatePoint(x1, y1, centrX, centrY, radAngle)
	x2, y2 = rotatePoint(x2, y2, centrX, centrY, radAngle)
	x3, y3 = rotatePoint(x3, y3, centrX, centrY, radAngle)

	x4, y4 = rotatePoint(x4, y4, centrX, centrY, radAngle)

	// Вычисляем координаты вершин с учетом угла поворота

	// Рисуем стороны квадрата с использованием обновленных координат
	primitive.segment(x1, y1, x2, y2, col)
	primitive.segment(x2, y2, x3, y3, col)
	primitive.segment(x3, y3, x4, y4, col)
	primitive.segment(x4, y4, x1, y1, col)

	// Сохраняем параметры
	primitive.startX = X
	primitive.startY = Y
	primitive.S = S

	return nil
}

// Draws a polyline connecting multiple points.
// @param points []Point2D: List of points to connect.
// @param lineColor color.Color: The color of the polyline.
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

// Draws a circle using a midpoint algorithm.
// @param x_, y_ int: Center of the circle.
// @param radius int: Radius of the circle.
// @param col color.Color: The color of the circle.
func (primitive *primitiveRendererСlass) DrawCircle(x_, y_ int, radius int, col color.Color) {
	centerX, centerY := x_, y_

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

// Draws an ellipse using a midpoint algorithm.
// @param center Point2D: Center of the ellipse.
// @param a, b int: Semi-major and semi-minor axes of the ellipse.
// @param col color.Color: The color of the ellipse.
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

// Draws a polygon defined by a set of points.
// @param points []Point2D: List of polygon vertices.
// @param lineColor color.Color: The color of the polygon edges.
// @return error: Returns an error if the polygon is invalid.
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

// Fills a square area with the specified color.
// @param x, y int: Top-left corner of the square.
// @param s int: Side length of the square.
// @param col color.Color: The fill color.
func (primitive *primitiveRendererСlass) FillSquare(x int, y int, s int, col color.Color) {
	for i := x; i <= x+s; i++ {
		for j := y; j <= y+s; j++ {
			primitive.plotPixel(i, j, col)
		}
	}
}

// Fills an area up to the boundary color using the border-fill algorithm.
// @param x, y int: Starting coordinates for the fill.
// @param fillColor color.Color: The fill color.
// @param borderColor color.Color: The boundary color.
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

// Fills an area using the flood-fill algorithm.
// @param x, y int: Starting coordinates for the fill.
// @param fillColor color.Color: The fill color.
// @param boundaryColor color.Color: The color marking the boundaries.
func (primitive *primitiveRendererСlass) FloodFill(x, y int, fillColor color.Color, boundaryColor color.Color) {
	width, height := primitive.screen.Size()
	originalColor := primitive.screen.At(x, y)

	if originalColor == fillColor || originalColor == boundaryColor {
		return
	}

	var floodFillRecursive func(x, y int)
	floodFillRecursive = func(x, y int) {
		if x < 0 || x >= width || y < 0 || y >= height {
			return
		}

		currentColor := primitive.screen.At(x, y)
		if currentColor != originalColor {
			return // Not the original color, stop recursion
		}

		primitive.screen.Set(x, y, fillColor)

		floodFillRecursive(x+1, y) // Right
		floodFillRecursive(x-1, y) // Left
		floodFillRecursive(x, y+1) // Down
		floodFillRecursive(x, y-1) // Up
	}

	floodFillRecursive(x, y)
}
