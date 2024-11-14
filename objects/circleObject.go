package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type CircleObject interface {
	GetShapeObject() ShapeObject
	Draw() error
	UnDraw() error
	Translate(x, y int) error
	Scale(S int) error
	Rotate(angle int) error
}

type circleObject struct {
	shapeObject ShapeObject
	center      Point2D
	primitive   PrimitiveRendererСlass
	color       color.Color
}

func NewCircleObject(shapeObject ShapeObject, x, y int, color color.Color) CircleObject {
	return &circleObject{
		shapeObject: shapeObject,
		center:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		color:       color,
		primitive:   NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

func EnhancedNewSircleObject(screen *ebiten.Image, backgroundColor color.Color, squareLenght, x, y int, color color.Color) CircleObject {
	gmob := NewGameObject(screen, backgroundColor)
	shapeObject := NewShapeObject(NewDrawableObject(gmob), NewTransformableObject(gmob))
	return &circleObject{
		shapeObject: shapeObject,
		center:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		color:       color,
		primitive:   NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}
func (circleObject *circleObject) GetShapeObject() ShapeObject {
	return circleObject.shapeObject
}

func (circleObject *circleObject) Draw() error {
	x, y := circleObject.center.GetCoords()
	circleObject.primitive.DrawSquare(x+squareObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+squareObject.GetShapeObject().GetTransformableObject().GetTranslationY(), squareObject.squareLenght*squareObject.GetShapeObject().GetTransformableObject().GetScale(), squareObject.GetShapeObject().GetTransformableObject().GetAngle(), squareObject.color)
	circleObject.shapeObject.GetDrawableObject().Draw()
	return nil
}
func (circleObject *circleObject) UnDraw() error {
	x, y := squareObject.squareTop.GetCoords()
	squareObject.primitive.DrawSquare(x+squareObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+squareObject.GetShapeObject().GetTransformableObject().GetTranslationY(), squareObject.squareLenght*squareObject.GetShapeObject().GetTransformableObject().GetScale(), squareObject.GetShapeObject().GetTransformableObject().GetAngle(), squareObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor())
	squareObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

func (circleObject *circleObject) Translate(x, y int) error {
	squareObject.UnDraw()
	squareObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	squareObject.Draw()
	return nil
}

func (circleObject *circleObject) Scale(S int) error {
	squareObject.UnDraw()
	squareObject.GetShapeObject().GetTransformableObject().Scale(S)
	squareObject.Draw()
	return nil
}

func (circleObject *circleObject) Rotate(angle int) error {
	squareObject.UnDraw()
	squareObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	squareObject.Draw()
	return nil
}
