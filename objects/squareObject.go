package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SquareObject interface {
	GetShapeObject() ShapeObject
	Draw() error
}

type squareObject struct {
	shapeObject  ShapeObject
	squareTop    Point2D
	squareLenght int
	primitive    PrimitiveRendererСlass
	color        color.Color
}

func NewSquareObject(shapeObject ShapeObject, x, y int, squareLenght int, color color.Color) SquareObject {
	return &squareObject{
		shapeObject:  shapeObject,
		squareTop:    NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		squareLenght: squareLenght,
		color:        color,
		primitive:    NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

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
func (squareObject *squareObject) GetShapeObject() ShapeObject {
	return squareObject.shapeObject
}

func (squareObject *squareObject) Draw() error {
	x, y := squareObject.squareTop.GetCoords()
	squareObject.primitive.DrawSquare(x+squareObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+squareObject.GetShapeObject().GetTransformableObject().GetTranslationY(), squareObject.squareLenght+squareObject.GetShapeObject().GetTransformableObject().GetScale(), squareObject.GetShapeObject().GetTransformableObject().GetAngle(), squareObject.color)
	squareObject.shapeObject.GetDrawableObject().Draw()
	return nil
}
