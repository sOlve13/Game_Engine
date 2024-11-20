package objects

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

// BitmapObject represents an object that can manage multiple bitmaps, draw them, and access their handlers.
// Also this object inherit BitmapHandler and DrawableObject
type BitmapObject interface {
	// GetDrawableObject returns the associated DrawableObject.
	// @return DrawableObject: The drawable object associated with the BitmapObject.
	GetDrawableObject() DrawableObject

	// GetBitmapHandler returns a BitmapHandler by its index.
	// @param num int: The index of the BitmapHandler to retrieve.
	// @return BitmapHandler: The BitmapHandler at the specified index.
	GetBitmapHandler(num int) BitmapHandler

	// Draw renders the bitmap on the screen using the provided handler and bitmap name.
	// @param name string: The name of the bitmap to be drawn.
	// @param num int: The index of the BitmapHandler to use for drawing.
	// @return error: Returns nil if the drawing operation is successful, or an error if there is a failure.
	Draw(name string, num int) error
}

// bitmapObject is an implementation of the BitmapObject interface.
// It stores a list of BitmapHandlers and a DrawableObject, and provides methods to retrieve and draw bitmaps.
type bitmapObject struct {
	bitmapHandlers []BitmapHandler // A slice of BitmapHandlers.
	drawableObject DrawableObject  // The associated DrawableObject.
}

// NewBitmapObject creates a new instance of bitmapObject with the provided BitmapHandlers and DrawableObject.
// @param bitmapHandlers []BitmapHandler: The list of BitmapHandlers to be associated with the BitmapObject.
// @param drawableObject DrawableObject: The DrawableObject associated with the BitmapObject.
// @return BitmapObject: A new bitmapObject instance.
func NewBitmapObject(bitmapHandlers []BitmapHandler, drawableObject DrawableObject) BitmapObject {
	return &bitmapObject{
		drawableObject: drawableObject,
		bitmapHandlers: bitmapHandlers,
	}
}

// GetDrawableObject returns the associated DrawableObject.
// @return DrawableObject: The DrawableObject associated with the BitmapObject.
func (bitmapObject *bitmapObject) GetDrawableObject() DrawableObject {
	return bitmapObject.drawableObject
}

// GetBitmapHandler returns a BitmapHandler by its index.
// @param num int: The index of the BitmapHandler to retrieve.
// @return BitmapHandler: The BitmapHandler at the specified index.
func (bitmapObject *bitmapObject) GetBitmapHandler(num int) BitmapHandler {
	return bitmapObject.bitmapHandlers[num]
}

// Draw renders the bitmap with the specified name and handler index on the screen.
// @param name string: The name of the bitmap to be drawn.
// @param num int: The index of the BitmapHandler to use for drawing.
// @return error: Returns nil if the drawing operation is successful, or an error if the bitmap is not found.
func (bitmapObject *bitmapObject) Draw(name string, num int) error {
	screen := bitmapObject.GetDrawableObject().GetGameObject().GetScreen() // Retrieve the screen from the DrawableObject.
	handler := bitmapObject.bitmapHandlers[num]                            // Get the BitmapHandler at the specified index.
	img, exists := handler.Get(name)                                       // Retrieve the image by its name.

	if !exists {
		return errors.New("name not exist") // Return an error if the bitmap is not found.
	}

	// Set up the drawing options.
	op := &ebiten.DrawImageOptions{}
	x, y := handler.GetCords() // Get the coordinates of the BitmapHandler.
	x_, y_ := float64(x), float64(y)
	scaleX, scaleY := 3.0, 3.0    // Scale factors (adjustable).
	op.GeoM.Scale(scaleX, scaleY) // Apply scaling.
	op.GeoM.Translate(x_, y_)     // Apply translation (positioning).

	// Draw the bitmap on the screen.
	screen.DrawImage(img, op)

	// Draw the associated DrawableObject.
	bitmapObject.GetDrawableObject().Draw()
	return nil
}
