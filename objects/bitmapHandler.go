package objects

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type bitmapHandler struct {
	bitmaps map[string]*ebiten.Image // Хранилище для битмапов с ключами
	x, y    int
}

type BitmapHandler interface {
	Create(name string, width, height int, clr color.Color)
	Delete(name string)
	Load(name, filePath string) error
	Save(name, filePath string) error
	Copy(srcName, destName string) error
	Get(name string) (*ebiten.Image, bool)
	GetCords() (x, y int)
	SetCoords(x, y int) error
}

// NewBitmapHandler создаёт новый экземпляр BitmapHandler
func NewBitmapHandler(x, y int) BitmapHandler {
	return &bitmapHandler{
		bitmaps: make(map[string]*ebiten.Image),
		x:       x,
		y:       y,
	}
}

func (bh *bitmapHandler) Create(name string, width, height int, clr color.Color) {
	img := ebiten.NewImage(width, height)
	img.Fill(clr) // Заливка цветом
	bh.bitmaps[name] = img
}

func (bh *bitmapHandler) Delete(name string) {
	delete(bh.bitmaps, name)
}

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

func (bh *bitmapHandler) Copy(srcName, destName string) error {
	src, exists := bh.bitmaps[srcName]
	if !exists {
		return os.ErrNotExist
	}

	dst := ebiten.NewImageFromImage(src)
	bh.bitmaps[destName] = dst
	return nil
}

func (bh *bitmapHandler) Get(name string) (*ebiten.Image, bool) {
	img, exists := bh.bitmaps[name]
	return img, exists
}
func (bh *bitmapHandler) GetCords() (x, y int) {
	return bh.x, bh.y
}

func (bh *bitmapHandler) SetCoords(x, y int) error {
	bh.x, bh.y = x, y
	return nil
}
