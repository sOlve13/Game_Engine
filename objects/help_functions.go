package objects

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Determines the orientation of the triplet (p, q, r).
// @return int: 0 if collinear, 1 if clockwise, 2 if counterclockwise.
func orientation(p, q, r Point2D) int {
	qX, qY := q.GetCoords()
	pX, pY := p.GetCoords()
	rX, rY := r.GetCoords()
	val := (qY-pY)*(rX-qX) - (qX-pX)*(rY-qY)
	if val == 0 {
		return 0 // Collinear
	} else if val > 0 {
		return 1 // Clockwise
	} else {
		return 2 // Counterclockwise
	}
}

// Checks if the point q lies on the line segment pr.
// @return bool: True if q is on segment pr, false otherwise.
func onSegment(p, q, r Point2D) bool {
	qX, qY := q.GetCoords()
	pX, pY := p.GetCoords()
	rX, rY := r.GetCoords()
	return qX <= max(pX, rX) && qX >= min(pX, rX) &&
		qY <= max(pY, rY) && qY >= min(pY, rY)
}

// Determines if two line segments (p1, q1) and (p2, q2) intersect.
// @return bool: True if the segments intersect, false otherwise.
func segmentsIntersect(p1, q1, p2, q2 Point2D) bool {
	// Find orientations for the four point pairs
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Special cases:
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}

	return false // No intersection
}

// Determines if a point is inside a polygon using the ray-casting algorithm.
// @return bool: True if the point is inside the polygon, false otherwise.
func isPointInPolygon(p Point2D, polygon []Point2D, screen *ebiten.Image, backgroundColor color.Color) bool {
	n := len(polygon)
	if n < 3 {
		return false
	}
	_, py := p.GetCoords()

	// Create a point far outside the polygon
	extreme := NewPoint2D(screen, backgroundColor, 1e10, py, backgroundColor)

	count := 0
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		// Check if the polygon edge intersects with the ray
		if segmentsIntersect(polygon[i], polygon[next], p, extreme) {
			// Check if the point lies on the edge of the polygon
			if orientation(polygon[i], p, polygon[next]) == 0 && onSegment(polygon[i], p, polygon[next]) {
				return false // Point lies on the edge
			}
			count++
		}
	}
	// Point is inside the polygon if the number of intersections is odd
	return count%2 == 1
}

// Rotates a point (x, y) around a center (cx, cy) by a specified angle.
// @return (int, int): The new coordinates of the rotated point.
func rotatePoint(x, y, cx, cy int, angle float64) (int, int) {
	xf, yf := float64(x-cx), float64(y-cy)
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	// Round cosine and sine values if close to 0
	if math.Abs(cosA) < 1e-10 {
		cosA = 0
	}
	if math.Abs(sinA) < 1e-10 {
		sinA = 0
	}

	newX := xf*cosA - yf*sinA
	newY := xf*sinA + yf*cosA

	return int(math.Round(newX)) + cx, int(math.Round(newY)) + cy
}

// Returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Retrieves the dimensions of an image from a file.
// @return (int, int, error): The width, height, and any error encountered.
func getImageSize(filePath string) (int, int, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	// Get image dimensions
	return img.Bounds().Dx(), img.Bounds().Dy(), nil
}

// Checks if a slice contains a specified integer.
// @return bool: True if the slice contains the item, false otherwise.
func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Returns the index of an item in a slice, or -1 if not found.
// @return int: The index of the item or -1 if not found.
func indexOf(slice []int, item int) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// Writes text to a specified file, overwriting any existing content.
// @return error: Returns nil if successful, otherwise returns an error.
func writeToFile(filename, text string) error {
	// Open the file for writing
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	return nil
}
