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
	radius      int
	primitive   PrimitiveRendererСlass
	color       color.Color
}

func NewCircleObject(shapeObject ShapeObject, x, y, r int, color color.Color) CircleObject {
	return &circleObject{
		shapeObject: shapeObject,
		center:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x, y, color),
		color:       color,
		radius:      r,
		primitive:   NewPrimitiveRendererclass(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

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
func (circleObject *circleObject) GetShapeObject() ShapeObject {
	return circleObject.shapeObject
}

func (circleObject *circleObject) Draw() error {
	x, y := circleObject.center.GetCoords()
	circleObject.primitive.DrawCircle(x+circleObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+circleObject.GetShapeObject().GetTransformableObject().GetTranslationY(), circleObject.radius*circleObject.shapeObject.GetTransformableObject().GetScale(), circleObject.color)
	circleObject.shapeObject.GetDrawableObject().Draw()
	return nil
}
func (circleObject *circleObject) UnDraw() error {
	x, y := circleObject.center.GetCoords()
	circleObject.primitive.DrawCircle(x+circleObject.GetShapeObject().GetTransformableObject().GetTranslationX(), y+circleObject.GetShapeObject().GetTransformableObject().GetTranslationY(), circleObject.radius*circleObject.shapeObject.GetTransformableObject().GetScale(), circleObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor())
	circleObject.shapeObject.GetDrawableObject().Draw()
	return nil
}

func (circleObject *circleObject) Translate(x, y int) error {
	circleObject.UnDraw()
	circleObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	circleObject.Draw()
	return nil
}

func (circleObject *circleObject) Scale(S int) error {
	circleObject.UnDraw()
	circleObject.GetShapeObject().GetTransformableObject().Scale(S)
	circleObject.Draw()
	return nil
}

func (circleObject *circleObject) Rotate(angle int) error {
	circleObject.UnDraw()
	circleObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	circleObject.Draw()
	return nil
}
