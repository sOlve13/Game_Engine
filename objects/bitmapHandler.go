package objects

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// bitmapHandler is an internal implementation of the BitmapHandler interface.
// It manages a collection of bitmaps and provides methods to create, load, save, copy, and manipulate them.
type bitmapHandler struct {
	bitmaps map[string]*ebiten.Image // A map for storing bitmaps by their names.
	x, y    int                      // Coordinates for positioning the bitmaps.
}

// BitmapHandler defines the interface for handling bitmaps.
// It provides methods to create, delete, load, save, copy, and retrieve bitmaps.
type BitmapHandler interface {
	// Create creates a new bitmap with the specified name, width, height, and color.
	// @param name string: The name of the bitmap.
	// @param width int: The width of the bitmap.
	// @param height int: The height of the bitmap.
	// @param clr color.Color: The color to fill the bitmap with.
	Create(name string, width, height int, clr color.Color)

	// Delete removes a bitmap by its name.
	// @param name string: The name of the bitmap to delete.
	Delete(name string)

	// Load loads a bitmap from a file at the specified file path.
	// @param name string: The name to assign to the loaded bitmap.
	// @param filePath string: The file path of the image to load.
	// @return error: Returns nil if the loading operation is successful, or an error if there is a failure.
	Load(name, filePath string) error

	// Save saves the bitmap to a file at the specified file path.
	// @param name string: The name of the bitmap to save.
	// @param filePath string: The file path where the bitmap will be saved.
	// @return error: Returns nil if the save operation is successful, or an error if there is a failure.
	Save(name, filePath string) error

	// Copy copies a bitmap from one name to another.
	// @param srcName string: The name of the source bitmap.
	// @param destName string: The name of the destination bitmap.
	// @return error: Returns nil if the copy operation is successful, or an error if there is a failure.
	Copy(srcName, destName string) error

	// Get retrieves a bitmap by its name.
	// @param name string: The name of the bitmap to retrieve.
	// @return *ebiten.Image: The bitmap associated with the specified name, or nil if not found.
	// @return bool: True if the bitmap was found, false otherwise.
	Get(name string) (*ebiten.Image, bool)

	// GetCords retrieves the current coordinates of the bitmap handler.
	// @return x, y int: The current x and y coordinates of the bitmap handler.
	GetCords() (x, y int)

	// SetCoords sets the coordinates for the bitmap handler.
	// @param x int: The new x-coordinate.
	// @param y int: The new y-coordinate.
	// @return error: Returns nil if the coordinates are successfully updated.
	SetCoords(x, y int) error
}

// NewBitmapHandler creates a new instance of BitmapHandler with specified coordinates.
// @param x int: The initial x-coordinate.
// @param y int: The initial y-coordinate.
// @return BitmapHandler: The newly created BitmapHandler instance.
func NewBitmapHandler(x, y int) BitmapHandler {
	return &bitmapHandler{
		bitmaps: make(map[string]*ebiten.Image),
		x:       x,
		y:       y,
	}
}

// Create creates a new bitmap with the specified name, width, height, and color.
// @param name string: The name of the bitmap.
// @param width int: The width of the bitmap.
// @param height int: The height of the bitmap.
// @param clr color.Color: The color to fill the bitmap with.
func (bh *bitmapHandler) Create(name string, width, height int, clr color.Color) {
	img := ebiten.NewImage(width, height)
	img.Fill(clr) // Fill the image with the specified color.
	bh.bitmaps[name] = img
}

// Delete removes a bitmap by its name.
// @param name string: The name of the bitmap to delete.
func (bh *bitmapHandler) Delete(name string) {
	delete(bh.bitmaps, name)
}

// Load loads a bitmap from a file and stores it with the given name.
// @param name string: The name to assign to the loaded bitmap.
// @param filePath string: The file path of the image to load.
// @return error: Returns nil if the loading operation is successful, or an error if there is a failure.
func (bh *bitmapHandler) Load(name, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	bitmap := ebiten.NewImageFromImage(img)
	bh.bitmaps[name] = bitmap
	return nil
}

// Save saves the bitmap to a file at the specified file path.
// @param name string: The name of the bitmap to save.
// @param filePath string: The file path where the bitmap will be saved.
// @return error: Returns nil if the save operation is successful, or an error if there is a failure.
func (bh *bitmapHandler) Save(name, filePath string) error {
	bitmap, exists := bh.bitmaps[name]
	if !exists {
		return os.ErrNotExist
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img := bitmap.SubImage(bitmap.Bounds()).(*image.RGBA)

	return png.Encode(file, img)
}

// Copy copies a bitmap from one name to another.
// @param srcName string: The name of the source bitmap.
// @param destName string: The name of the destination bitmap.
// @return error: Returns nil if the copy operation is successful, or an error if there is a failure.
func (bh *bitmapHandler) Copy(srcName, destName string) error {
	src, exists := bh.bitmaps[srcName]
	if !exists {
		return os.ErrNotExist
	}

	dst := ebiten.NewImageFromImage(src)
	bh.bitmaps[destName] = dst
	return nil
}

// Get retrieves a bitmap by its name.
// @param name string: The name of the bitmap to retrieve.
// @return *ebiten.Image: The bitmap associated with the specified name, or nil if not found.
// @return bool: True if the bitmap was found, false otherwise.
func (bh *bitmapHandler) Get(name string) (*ebiten.Image, bool) {
	img, exists := bh.bitmaps[name]
	return img, exists
}

// GetCords retrieves the current coordinates of the bitmap handler.
// @return x, y int: The current x and y coordinates of the bitmap handler.
func (bh *bitmapHandler) GetCords() (x, y int) {
	return bh.x, bh.y
}

// SetCoords sets the coordinates for the bitmap handler.
// @param x int: The new x-coordinate.
// @param y int: The new y-coordinate.
// @return error: Returns nil if the coordinates are successfully updated.
func (bh *bitmapHandler) SetCoords(x, y int) error {
	bh.x, bh.y = x, y
	return nil
}
