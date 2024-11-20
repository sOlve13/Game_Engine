package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// CircleObject represents a circle object that can be drawn, transformed (scaled, rotated, translated), and erased from the screen.
// It provides methods for manipulating the circle and applying transformations to it.
// Also this object inherit ShapeObject and use primitive for drawing circle
type CircleObject interface {
	// GetShapeObject returns the associated shape object for the circle.
	// @return ShapeObject: The associated shape object for the circle.
	GetShapeObject() ShapeObject

	// Draw draws the circle on the screen with its current transformations.
	// @return error: Returns nil if the drawing operation is successful.
	Draw() error

	// UnDraw erases the circle from the screen.
	// @return error: Returns nil if the erase operation is successful.
	UnDraw() error

	// Translate moves the circle by a given offset on the x and y axes.
	// @param x int: The offset on the x-axis.
	// @param y int: The offset on the y-axis.
	// @return error: Returns nil if the translation operation is successful.
	Translate(x, y int) error

	// Scale changes the scale of the circle by a given factor.
	// @param S int: The scale factor.
	// @return error: Returns nil if the scaling operation is successful.
	Scale(S int) error

	// Rotate rotates the circle by a given angle.
	// @param angle int: The angle to rotate the circle.
	// @return error: Returns nil if the rotation operation is successful.
	Rotate(angle int) error
}

// circleObject is the internal implementation of the CircleObject interface.
// It contains the shape object, the center of the circle, its radius, a primitive renderer, and its color.
type circleObject struct {
	shapeObject ShapeObject            // The associated shape object for the circle.
	center      Point2D                // The center of the circle (coordinates).
	radius      int                    // The radius of the circle.
	primitive   PrimitiveRendererСlass // The renderer for drawing the circle.
	color       color.Color            // The color of the circle.
}

// NewCircleObject creates a new circle object with the specified shape object, center coordinates, radius, and color.
// @param shapeObject ShapeObject: The associated shape object for the circle.
// @param x int: The x-coordinate of the circle's center.
// @param y int: The y-coordinate of the circle's center.
// @param r int: The radius of the circle.
// @param color color.Color: The color of the circle.
// @return CircleObject: The new circle object instance.
func NewCircleObject(shapeObject ShapeObject, x, y, r int, color color.Color) CircleObject {
	return &circleObject{
		shapeObject: shapeObject,
		center:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		color:       color,
		radius:      r,
		primitive:   NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

// EnhancedNewCircleObject creates a new circle object with the specified screen, background color, center coordinates, radius, and color.
// This method also initializes a new game object and shape object.
// @param screen *ebiten.Image: The screen where the circle will be drawn.
// @param backgroundColor color.Color: The background color for the circle.
// @param x int: The x-coordinate of the circle's center.
// @param y int: The y-coordinate of the circle's center.
// @param r int: The radius of the circle.
// @param color color.Color: The color of the circle.
// @return CircleObject: The new circle object instance.
func EnhancedNewCircleObject(screen *ebiten.Image, backgroundColor color.Color, x, y, r int, color color.Color) CircleObject {
	gmob := NewGameObject(screen, backgroundColor)
	shapeObject := NewShapeObject(NewDrawableObject(gmob), NewTransformableObject(gmob))
	return &circleObject{
		shapeObject: shapeObject,
		center:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		radius:      r,
		color:       color,
		primitive:   NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

// GetShapeObject returns the associated shape object for the circle.
// @return ShapeObject: The associated shape object for the circle.
func (circleObject *circleObject) GetShapeObject() ShapeObject {
	return circleObject.shapeObject
}

// Draw draws the circle with its current transformations (translation, scaling, rotation).
// @return error: Returns nil if the drawing operation is successful.
func (circleObject *circleObject) Draw() error {
	x, y := circleObject.center.GetCoords()
	circleObject.primitive.DrawCircle(x+circleObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+circleObject.GetShapeObject().GetTransformableObject().GetTranslationY(), circleObject.radius*circleObject.shapeObject.GetTransformableObject().GetScale(), circleObject.color)
	circleObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

// UnDraw erases the circle from the screen by effectively removing it.
// @return error: Returns nil if the erase operation is successful.
func (circleObject *circleObject) UnDraw() error {
	x, y := circleObject.center.GetCoords()
	circleObject.primitive.DrawCircle(x+circleObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+circleObject.GetShapeObject().GetTransformableObject().GetTranslationY(), circleObject.radius*circleObject.shapeObject.GetTransformableObject().GetScale(), circleObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor())
	circleObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

// Translate moves the circle by a given offset on the x and y axes.
// @param x int: The offset on the x-axis.
// @param y int: The offset on the y-axis.
// @return error: Returns nil if the translation operation is successful.
func (circleObject *circleObject) Translate(x, y int) error {
	circleObject.UnDraw()
	circleObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	circleObject.Draw()
	return nil
}

// Scale changes the scale of the circle by a given factor.
// @param S int: The scale factor for the circle.
// @return error: Returns nil if the scaling operation is successful.
func (circleObject *circleObject) Scale(S int) error {
	circleObject.UnDraw()
	circleObject.GetShapeObject().GetTransformableObject().Scale(S)
	circleObject.Draw()
	return nil
}

// Rotate rotates the circle by a given angle.
// @param angle int: The angle by which to rotate the circle.
// @return error: Returns nil if the rotation operation is successful.
func (circleObject *circleObject) Rotate(angle int) error {
	circleObject.UnDraw()
	circleObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	circleObject.Draw()
	return nil
}
