package objects

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// LineObject represents a line object that can be drawn, transformed (scaled, rotated, translated), and erased from the screen.
// It provides methods for manipulating the line and applying transformations to it.
// Also this object inherit ShapeObject and use segment for drawing line
type LineObject interface {
	// GetShapeObject returns the associated shape object for the line.
	// @return ShapeObject: The associated shape object for the line.
	GetShapeObject() ShapeObject

	// Draw draws the line on the screen with its current transformations.
	// @return error: Returns nil if the drawing operation is successful.
	Draw() error

	// UnDraw erases the line from the screen.
	// @return error: Returns nil if the erase operation is successful.
	UnDraw() error

	// Translate moves the line by a given offset on the x and y axes.
	// @param x int: The offset on the x-axis.
	// @param y int: The offset on the y-axis.
	// @return error: Returns nil if the translation operation is successful.
	Translate(x, y int) error

	// Scale changes the scale of the line by a given factor.
	// @param S int: The scale factor.
	// @return error: Returns nil if the scaling operation is successful.
	Scale(S int) error

	// Rotate rotates the line by a given angle.
	// @param angle int: The angle to rotate the line.
	// @return error: Returns nil if the rotation operation is successful.
	Rotate(angle int) error
}

// lineObject is the internal implementation of the LineObject interface.
// It contains the shape object, the start and finish points of the line, the line segment, and its color.
type lineObject struct {
	shapeObject ShapeObject // The associated shape object for the line.
	start       Point2D     // The starting point of the line.
	finish      Point2D     // The finishing point of the line.
	segment     LineSegment // The line segment that uses for drawing the line.
	color       color.Color // The color of the line.
}

// NewLineObject creates a new line object with the specified shape object, start and finish points, and color.
// @param shapeObject ShapeObject: The associated shape object for the line.
// @param x1 int: The x-coordinate of the line's start point.
// @param y1 int: The y-coordinate of the line's start point.
// @param x2 int: The x-coordinate of the line's finish point.
// @param y2 int: The y-coordinate of the line's finish point.
// @param color color.Color: The color of the line.
// @return LineObject: The new line object instance.
func NewLineObject(shapeObject ShapeObject, x1, y1, x2, y2 int, color color.Color) LineObject {
	return &lineObject{
		shapeObject: shapeObject,
		start:       NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x1, y1, color),
		finish:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x2, y2, color),
		color:       color,
		segment:     NewLineSegment(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

// EnhancedNewLineObject creates a new line object with the specified screen, background color, start and finish points, and color.
// This method also initializes a new game object and shape object.
// @param screen *ebiten.Image: The screen where the line will be drawn.
// @param backgroundColor color.Color: The background color for the line.
// @param x1 int: The x-coordinate of the line's start point.
// @param y1 int: The y-coordinate of the line's start point.
// @param x2 int: The x-coordinate of the line's finish point.
// @param y2 int: The y-coordinate of the line's finish point.
// @param color color.Color: The color of the line.
// @return LineObject: The new line object instance.
func EnhancedNewLineObject(screen *ebiten.Image, backgroundColor color.Color, x1, y1, x2, y2 int, color color.Color) LineObject {
	gmob := NewGameObject(screen, backgroundColor)
	shapeObject := NewShapeObject(NewDrawableObject(gmob), NewTransformableObject(gmob))
	return &lineObject{
		shapeObject: shapeObject,
		start:       NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x1, y1, color),
		finish:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x2, y2, color),
		color:       color,
		segment:     NewLineSegment(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

// GetShapeObject returns the associated shape object for the line.
// @return ShapeObject: The associated shape object for the line.
func (lineObject *lineObject) GetShapeObject() ShapeObject {
	return lineObject.shapeObject
}

// Draw draws the line with its current transformations (translation, scaling, rotation).
// @return error: Returns nil if the drawing operation is successful.
func (lineObject *lineObject) Draw() error {
	x1, y1 := lineObject.start.GetCoords()
	x2, y2 := lineObject.finish.GetCoords()

	// Apply translation
	translationX := lineObject.GetShapeObject().GetTransformableObject().GetTranslationX()
	translationY := lineObject.GetShapeObject().GetTransformableObject().GetTranslationY()
	x1, y1 = x1+translationX, y1+translationY
	x2, y2 = x2+translationX, y2+translationY

	// Apply scaling to both points
	scale := lineObject.shapeObject.GetTransformableObject().GetScale()
	dx := (x2 - x1) * (scale - 1)
	dy := (y2 - y1) * (scale - 1)
	x2, y2 = x2+dx, y2+dy

	// Apply rotation around the center of the line
	radAngle := float64(lineObject.shapeObject.GetTransformableObject().GetAngle()) * math.Pi / 180.0
	centrX, centrY := (x1+x2)/2, (y1+y2)/2
	fmt.Println(lineObject.shapeObject.GetTransformableObject().GetAngle(), centrX, centrY, x1, y1, x2, y2)
	x1, y1 = rotatePoint(x1, y1, centrX, centrY, radAngle)
	x2, y2 = rotatePoint(x2, y2, centrX, centrY, radAngle)
	fmt.Println(lineObject.shapeObject.GetTransformableObject().GetAngle(), centrX, centrY, x1, y1, x2, y2)
	lineObject.segment.Segment(NewPoint2D(lineObject.shapeObject.GetDrawableObject().GetGameObject().GetScreen(), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x1, y1, lineObject.color), NewPoint2D(lineObject.shapeObject.GetDrawableObject().GetGameObject().GetScreen(), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x2, y2, lineObject.color), lineObject.color)
	lineObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

// UnDraw erases the line from the screen by effectively removing it.
// @return error: Returns nil if the erase operation is successful.
func (lineObject *lineObject) UnDraw() error {
	x1, y1 := lineObject.start.GetCoords()
	x2, y2 := lineObject.finish.GetCoords()

	// Apply translation
	translationX := lineObject.GetShapeObject().GetTransformableObject().GetTranslationX()
	translationY := lineObject.GetShapeObject().GetTransformableObject().GetTranslationY()
	x1, y1 = x1+translationX, y1+translationY
	x2, y2 = x2+translationX, y2+translationY

	// Apply scaling to both points
	scale := lineObject.shapeObject.GetTransformableObject().GetScale()
	dx := (x2 - x1) * (scale - 1)
	dy := (y2 - y1) * (scale - 1)
	x2, y2 = x2+dx, y2+dy

	// Apply rotation around the center of the line
	radAngle := float64(lineObject.shapeObject.GetTransformableObject().GetAngle()) * math.Pi / 180.0
	centrX, centrY := (x1+x2)/2, (y1+y2)/2
	x1, y1 = rotatePoint(x1, y1, centrX, centrY, radAngle)
	x2, y2 = rotatePoint(x2, y2, centrX, centrY, radAngle)

	lineObject.segment.Segment(NewPoint2D(lineObject.shapeObject.GetDrawableObject().GetGameObject().GetScreen(), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x1, y1, lineObject.color), NewPoint2D(lineObject.shapeObject.GetDrawableObject().GetGameObject().GetScreen(), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x2, y2, lineObject.color), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor())
	lineObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

// Translate moves the line by the specified x and y offsets.
// @param x int: The offset on the x-axis.
// @param y int: The offset on the y-axis.
// @return error: Returns nil if the translation operation is successful.
func (lineObject *lineObject) Translate(x, y int) error {
	lineObject.UnDraw()
	lineObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	lineObject.Draw()
	return nil
}

// Scale changes the scale of the line by the specified factor.
// @param S int: The scale factor.
// @return error: Returns nil if the scaling operation is successful.
func (lineObject *lineObject) Scale(S int) error {
	lineObject.UnDraw()
	lineObject.GetShapeObject().GetTransformableObject().Scale(S)
	lineObject.Draw()
	return nil
}

// Rotate rotates the line by the specified angle.
// @param angle int: The angle to rotate the line.
// @return error: Returns nil if the rotation operation is successful.
func (lineObject *lineObject) Rotate(angle int) error {
	lineObject.UnDraw()
	lineObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	lineObject.Draw()
	return nil
}
