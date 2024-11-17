package objects

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type BitmapObject interface {
	GetDrawableObject() DrawableObject
	GetBitmapHandler(int) BitmapHandler
	Draw(names string, num int) error
}

type bitmapObject struct {
	bitmapHandlers []BitmapHandler
	drawableObject DrawableObject
}

func NewBitmapObject(bitmapHandlers []BitmapHandler, drawableObject DrawableObject) BitmapObject {
	return &bitmapObject{
		drawableObject: drawableObject,
		bitmapHandlers: bitmapHandlers,
	}
}

func (bitmapObject *bitmapObject) GetDrawableObject() DrawableObject {
	return bitmapObject.drawableObject
}

func (bitmapObject *bitmapObject) GetBitmapHandler(num int) BitmapHandler {
	return bitmapObject.bitmapHandlers[num]
}
func (bitmapObject *bitmapObject) Draw(name string, num int) error {

	screen := bitmapObject.GetDrawableObject().GetGameObject().GetScreen()
	handler := bitmapObject.bitmapHandlers[num]
	img, exists := handler.Get(name)

	if !exists {
		return errors.New("name not exist")
	}
	op := &ebiten.DrawImageOptions{}
	x, y := handler.GetCords()
	x_, y_ := float64(x), float64(y)
	scaleX, scaleY := 3.0, 3.0
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(x_, y_) // Установка позиции
	// Отрисовка битмапа
	screen.DrawImage(img, op)
	bitmapObject.GetDrawableObject().Draw()
	return nil

}
