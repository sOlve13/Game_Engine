package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// SquareObject represents a square object that can be drawn, transformed (scaled, rotated, translated), and undrawn.
// It provides methods to manipulate the square's transformations and rendering.
// Also this object inherit ShapeObject and use primitive for drawing square
type SquareObject interface {
	// GetShapeObject returns the associated shape object of the square object.
	// @return ShapeObject: The associated shape object.
	GetShapeObject() ShapeObject

	// Draw draws the square object on the screen with its current transformations.
	// @return error: Returns nil if the drawing operation was successful.
	Draw() error

	// UnDraw removes the square object from the screen.
	// @return error: Returns nil if the undrawing operation was successful.
	UnDraw() error

	// Translate moves the square object by the specified x and y values.
	// @param x int: The x translation value.
	// @param y int: The y translation value.
	// @return error: Returns nil if the translation operation was successful.
	Translate(x, y int) error

	// Scale scales the square object by the specified scale factor.
	// @param S int: The scaling factor for the square.
	// @return error: Returns nil if the scaling operation was successful.
	Scale(S int) error

	// Rotate rotates the square object by the specified angle.
	// @param angle int: The angle to rotate the square object.
	// @return error: Returns nil if the rotation operation was successful.
	Rotate(angle int) error
}

// squareObject is an internal implementation of the SquareObject interface.
// It holds references to a shape object, square top position, square length,
// primitive renderer class, and the color of the square.
type squareObject struct {
	shapeObject  ShapeObject            // The associated shape object.
	squareTop    Point2D                // The top-left coordinates of the square.
	squareLenght int                    // The length of the sides of the square.
	primitive    PrimitiveRendererСlass // The primitive object - renderer for drawing the square.
	color        color.Color            // The color of the square.
}

// NewSquareObject creates a new square object with the specified shape object,
// top-left position, square length, and color.
// @param shapeObject ShapeObject: The shape object to associate with the square object.
// @param x int: The x-coordinate for the top-left corner of the square.
// @param y int: The y-coordinate for the top-left corner of the square.
// @param squareLenght int: The length of the sides of the square.
// @param color color.Color: The color of the square.
// @return SquareObject: A new instance of the square object.
func NewSquareObject(shapeObject ShapeObject, x, y int, squareLenght int, color color.Color) SquareObject {
	return &squareObject{
		shapeObject:  shapeObject,
		squareTop:    NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		squareLenght: squareLenght,
		color:        color,
		primitive:    NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

// EnhancedNewSquareObject creates a new square object with the specified screen, background color, square length,
// position, and color. This method initializes a new game object and shape object as well.
// @param screen *ebiten.Image: The screen where the square will be drawn.
// @param backgroundColor color.Color: The background color for the square.
// @param squareLenght int: The length of the sides of the square.
// @param x int: The x-coordinate for the top-left corner of the square.
// @param y int: The y-coordinate for the top-left corner of the square.
// @param color color.Color: The color of the square.
// @return SquareObject: A new instance of the square object.
func EnhancedNewSquareObject(screen *ebiten.Image, backgroundColor color.Color, squareLenght, x, y int, color color.Color) SquareObject {
	gmob := NewGameObject(screen, backgroundColor)
	shapeObject := NewShapeObject(NewDrawableObject(gmob), NewTransformableObject(gmob))
	return &squareObject{
		shapeObject:  shapeObject,
		squareTop:    NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		squareLenght: squareLenght,
		color:        color,
		primitive:    NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

// GetShapeObject returns the associated shape object of the square object.
// @return ShapeObject: The associated shape object.
func (squareObject *squareObject) GetShapeObject() ShapeObject {
	return squareObject.shapeObject
}

// Draw draws the square object on the screen with its current transformations (translation, scale, rotation).
// @return error: Returns nil if the drawing operation was successful.
func (squareObject *squareObject) Draw() error {
	x, y := squareObject.squareTop.GetCoords()
	squareObject.primitive.DrawSquare(x+squareObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+squareObject.GetShapeObject().GetTransformableObject().GetTranslationY(), squareObject.squareLenght*squareObject.GetShapeObject().GetTransformableObject().GetScale(), squareObject.GetShapeObject().GetTransformableObject().GetAngle(), squareObject.color)
	squareObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

// UnDraw removes the square object from the screen, effectively undrawing it.
// @return error: Returns nil if the undrawing operation was successful.
func (squareObject *squareObject) UnDraw() error {
	x, y := squareObject.squareTop.GetCoords()
	squareObject.primitive.DrawSquare(x+squareObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+squareObject.GetShapeObject().GetTransformableObject().GetTranslationY(), squareObject.squareLenght*squareObject.GetShapeObject().GetTransformableObject().GetScale(), squareObject.GetShapeObject().GetTransformableObject().GetAngle(), squareObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor())
	squareObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

// Translate moves the square object by the specified x and y values.
// @param x int: The x translation value.
// @param y int: The y translation value.
// @return error: Returns nil if the translation operation was successful.
func (squareObject *squareObject) Translate(x, y int) error {
	squareObject.UnDraw()
	squareObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	squareObject.Draw()
	return nil
}

// Scale scales the square object by the specified scale factor.
// @param S int: The scaling factor for the square.
// @return error: Returns nil if the scaling operation was successful.
func (squareObject *squareObject) Scale(S int) error {
	squareObject.UnDraw()
	squareObject.GetShapeObject().GetTransformableObject().Scale(S)
	squareObject.Draw()
	return nil
}

// Rotate rotates the square object by the specified angle.
// @param angle int: The angle to rotate the square object.
// @return error: Returns nil if the rotation operation was successful.
func (squareObject *squareObject) Rotate(angle int) error {
	squareObject.UnDraw()
	squareObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	squareObject.Draw()
	return nil
}
