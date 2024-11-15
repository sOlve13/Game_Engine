package objects

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type LineObject interface {
	GetShapeObject() ShapeObject
	Draw() error
	UnDraw() error
	Translate(x, y int) error
	Scale(S int) error
	Rotate(angle int) error
}

type lineObject struct {
	shapeObject ShapeObject
	start       Point2D
	finish      Point2D

	segment LineSegment
	color   color.Color
}

func NewLineObject(shapeObject ShapeObject, x1, y1, x2, y2 int, color color.Color) LineObject {
	return &lineObject{
		shapeObject: shapeObject,
		start:       NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x1, y1, color),
		finish:      NewPoint2D(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x2, y2, color),
		color:       color,
		segment:     NewLineSegment(shapeObject.GetDrawableObject().GetGameObject().GetScreen(), shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor()),
	}
}

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
func (lineObject *lineObject) GetShapeObject() ShapeObject {
	return lineObject.shapeObject
}

func (lineObject *lineObject) Draw() error {
	x1, y1 := lineObject.start.GetCoords()
	x2, y2 := lineObject.finish.GetCoords()

	// Применяем трансляцию
	translationX := lineObject.GetShapeObject().GetTransformableObject().GetTranslationX()
	translationY := lineObject.GetShapeObject().GetTransformableObject().GetTranslationY()
	x1, y1 = x1+translationX, y1+translationY
	x2, y2 = x2+translationX, y2+translationY

	// Применяем масштабирование к обеим точкам
	scale := lineObject.shapeObject.GetTransformableObject().GetScale()
	dx := (x2 - x1) * (scale - 1)
	dy := (y2 - y1) * (scale - 1)
	x2, y2 = x2+dx, y2+dy

	// Применяем поворот относительно центра линии
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
func (lineObject *lineObject) UnDraw() error {
	x1, y1 := lineObject.start.GetCoords()
	x2, y2 := lineObject.finish.GetCoords()

	// Применяем трансляцию
	translationX := lineObject.GetShapeObject().GetTransformableObject().GetTranslationX()
	translationY := lineObject.GetShapeObject().GetTransformableObject().GetTranslationY()
	x1, y1 = x1+translationX, y1+translationY
	x2, y2 = x2+translationX, y2+translationY

	// Применяем масштабирование к обеим точкам
	scale := lineObject.shapeObject.GetTransformableObject().GetScale()
	dx := (x2 - x1) * (scale - 1)
	dy := (y2 - y1) * (scale - 1)
	x2, y2 = x2+dx, y2+dy

	// Применяем поворот относительно центра линии
	radAngle := float64(lineObject.shapeObject.GetTransformableObject().GetAngle()) * math.Pi / 180.0
	centrX, centrY := (x1+x2)/2, (y1+y2)/2
	x1, y1 = rotatePoint(x1, y1, centrX, centrY, radAngle)
	x2, y2 = rotatePoint(x2, y2, centrX, centrY, radAngle)

	lineObject.segment.Segment(NewPoint2D(lineObject.shapeObject.GetDrawableObject().GetGameObject().GetScreen(), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x1, y1, lineObject.color), NewPoint2D(lineObject.shapeObject.GetDrawableObject().GetGameObject().GetScreen(), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor(), x2, y2, lineObject.color), lineObject.shapeObject.GetDrawableObject().GetGameObject().GetBackgroundColor())
	lineObject.shapeObject.GetDrawableObject().Draw()
	return nil

}

func (lineObject *lineObject) Translate(x, y int) error {
	lineObject.UnDraw()
	lineObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	lineObject.Draw()
	return nil
}

func (lineObject *lineObject) Scale(S int) error {
	lineObject.UnDraw()
	lineObject.GetShapeObject().GetTransformableObject().Scale(S)
	lineObject.Draw()
	return nil
}

func (lineObject *lineObject) Rotate(angle int) error {
	lineObject.UnDraw()
	lineObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	lineObject.Draw()
	return nil
}
