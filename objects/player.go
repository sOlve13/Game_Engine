package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerObject interface {
	GetShapeObject() ShapeObject
	Draw() error
	UnDraw() error
	Translate(x, y int) error
	Scale(S int) error
	Rotate(angle int) error
}

type playerObject struct {
	shapeObject ShapeObject
	body        SquareObject
	head        CircleObject
	leg1        LineObject
	leg2        LineObject
	hand1       LineObject
	hand2       LineObject
	color       color.Color
}

func NewPlayerObject(screen *ebiten.Image, backgroundColor color.Color, color color.Color, x, y int) PlayerObject {
	gmob := NewGameObject(screen, backgroundColor)
	shapeObject := NewShapeObject(NewDrawableObject(gmob), NewTransformableObject(gmob))
	return &playerObject{
		shapeObject: shapeObject,
		color:       color,
		body:        NewSquareObject(shapeObject, x, y, 100, color),
		head:        NewCircleObject(shapeObject, x+50, y-25, 25, color),
		hand1:       NewLineObject(shapeObject, x, y, x-50, y+50, color),
		hand2:       NewLineObject(shapeObject, x+100, y, x+150, y+50, color),
		leg1:        NewLineObject(shapeObject, x, y+100, x-50, y+250, color),
		leg2:        NewLineObject(shapeObject, x+100, y+100, x+150, y+250, color),
	}
}
func (playerObject *playerObject) GetShapeObject() ShapeObject {
	return playerObject.shapeObject
}

func (playerObject *playerObject) Draw() error {
	playerObject.body.Draw()
	playerObject.head.Draw()
	playerObject.leg1.Draw()
	playerObject.leg2.Draw()
	playerObject.hand1.Draw()
	playerObject.hand2.Draw()
	playerObject.shapeObject.GetDrawableObject().Draw()
	return nil
}
func (playerObject *playerObject) UnDraw() error {
	playerObject.body.UnDraw()
	playerObject.head.UnDraw()
	playerObject.leg1.UnDraw()
	playerObject.leg2.UnDraw()
	playerObject.hand1.UnDraw()
	playerObject.hand2.UnDraw()
	playerObject.shapeObject.GetDrawableObject().UnDraw()
	return nil
}

func (playerObject *playerObject) Translate(x, y int) error {
	playerObject.UnDraw()
	playerObject.GetShapeObject().GetTransformableObject().Translate(x, y)
	playerObject.Draw()
	return nil
}

func (playerObject *playerObject) Scale(S int) error {
	//playerObject.UnDraw()
	//playerObject.GetShapeObject().GetTransformableObject().Scale(S)
	//playerObject.Draw()
	return nil
}

func (playerObject *playerObject) Rotate(angle int) error {
	//playerObject.UnDraw()
	//playerObject.GetShapeObject().GetTransformableObject().Rotate(angle)
	//playerObject.Draw()
	return nil
}
