package objects

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SpriteObject represents an interface for managing and manipulating sprite objects.
// Also this object inherit BitmapObject and AnimatedObject
type SpriteObject interface {
	// GetBitmapObject retrieves the associated BitmapObject.
	// @return BitmapObject: The BitmapObject linked to the SpriteObject.
	GetBitmapObject() BitmapObject

	// GetAnimatedObject retrieves the associated AnimatedObject.
	// @return AnimatedObject: The AnimatedObject linked to the SpriteObject.
	GetAnimatedObject() AnimatedObject

	// LoadBitmaps loads bitmap images from a folder and associates them with the SpriteObject.
	// @param folderPath string: The path to the folder containing bitmap images.
	// @param bmNum int: The index of the BitmapHandler to store the bitmaps in.
	// @return error: Returns nil if the operation succeeds or an error if a failure occurs.
	LoadBitmaps(folderPath string, bmNum int) error

	// GetName retrieves the name of a bitmap by its index.
	// @param num int: The index of the bitmap.
	// @return string: The name of the bitmap at the specified index.
	// @return error: Returns nil if the operation succeeds or an error if the index is invalid.
	GetName(num int) (string, error)

	// SetBitmap selects a bitmap for drawing by its index and renders it using the specified BitmapHandler.
	// @param name int: The index of the bitmap to render.
	// @param bmNum int: The index of the BitmapHandler to use for rendering.
	// @return error: Returns nil if the operation succeeds or an error if rendering fails.
	SetBitmap(name int, bmNum int) error

	// MoveObject moves the sprite to the specified coordinates.
	// @param x int: The x-coordinate for the new position.
	// @param y int: The y-coordinate for the new position.
	// @param num int: The index of the BitmapHandler to update.
	// @return error: Returns nil if the operation succeeds or an error if movement fails.
	MoveObject(x, y, num int) error
}

// spriteObject is an implementation of the SpriteObject interface.
// It manages bitmaps, animation, and positional updates for a sprite.
type spriteObject struct {
	bitmapObject   BitmapObject   // The associated BitmapObject.
	animatedObject AnimatedObject // The associated AnimatedObject.
	dictionary     map[int]string // A mapping of bitmap indices to names.
	name           string         // The name of the sprite.
}

// NewSpriteObject creates a new instance of a SpriteObject.
// @param bitmapObject BitmapObject: The BitmapObject to associate with the SpriteObject.
// @param name string: The name of the sprite.
// @return SpriteObject: A new instance of SpriteObject.
func NewSpriteObject(bitmapObject BitmapObject, name string) SpriteObject {
	return &spriteObject{
		bitmapObject:   bitmapObject,
		animatedObject: nil,
		dictionary:     make(map[int]string),
		name:           name,
	}
}

// GetBitmapObject retrieves the associated BitmapObject.
// @return BitmapObject: The BitmapObject linked to the SpriteObject.
func (spriteObject *spriteObject) GetBitmapObject() BitmapObject {
	return spriteObject.bitmapObject
}

// GetAnimatedObject retrieves the associated AnimatedObject.
// @return AnimatedObject: The AnimatedObject linked to the SpriteObject.
func (spriteObject *spriteObject) GetAnimatedObject() AnimatedObject {
	return spriteObject.animatedObject
}

// LoadBitmaps loads bitmap images from a folder and associates them with the SpriteObject.
// @param folderPath string: The path to the folder containing bitmap images.
// @param bmNum int: The index of the BitmapHandler to store the bitmaps in.
// @return error: Returns nil if the operation succeeds or an error if a failure occurs.
func (spriteObject *spriteObject) LoadBitmaps(folderPath string, bmNum int) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for i, file := range files {
		// Skip directories.
		if file.IsDir() {
			continue
		}

		// Extract the file name without extension.
		fileName := file.Name()
		nameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		// Map the index to the file name (without extension).
		spriteObject.dictionary[i] = nameWithoutExt
	}

	// Initialize an AnimatedObject.
	amOb := NewAnimatedObject(spriteObject.bitmapObject.GetDrawableObject().GetGameObject(), len(spriteObject.dictionary), spriteObject.name)
	spriteObject.animatedObject = amOb

	// Load each bitmap into the BitmapHandler.
	for key := range spriteObject.dictionary {
		width, height, err := getImageSize(filepath.Join(folderPath, spriteObject.dictionary[key]+".png"))
		if err != nil {
			fmt.Println("Error retrieving image dimensions:", err)
			return err
		}
		spriteObject.GetBitmapObject().GetBitmapHandler(bmNum).Create(spriteObject.dictionary[key], width, height, spriteObject.animatedObject.GetGameObject().GetBackgroundColor())
		err = spriteObject.GetBitmapObject().GetBitmapHandler(bmNum).Load(spriteObject.dictionary[key], filepath.Join(folderPath, spriteObject.dictionary[key]+".png"))
		if err != nil {
			return err
		}
	}

	return nil
}

// SetBitmap selects a bitmap for drawing by its index and renders it using the specified BitmapHandler.
// @param name int: The index of the bitmap to render.
// @param bmNum int: The index of the BitmapHandler to use for rendering.
// @return error: Returns nil if the operation succeeds or an error if rendering fails.
func (spriteObject *spriteObject) SetBitmap(name int, bmNum int) error {
	err := spriteObject.animatedObject.Animate(name)
	if err != nil {
		return err
	}
	val, err := spriteObject.GetName(name)
	if err != nil {
		return err
	}
	err = spriteObject.bitmapObject.Draw(val, bmNum)
	if err != nil {
		return err
	}
	return nil
}

// GetName retrieves the name of a bitmap by its index.
// @param num int: The index of the bitmap.
// @return string: The name of the bitmap at the specified index.
// @return error: Returns nil if the operation succeeds or an error if the index is invalid.
func (spriteObject *spriteObject) GetName(num int) (string, error) {
	val, exists := spriteObject.dictionary[num]
	if !exists {
		return "", errors.New("value not in list")
	}
	return val, nil
}

// MoveObject moves the sprite to the specified coordinates.
// @param x int: The x-coordinate for the new position.
// @param y int: The y-coordinate for the new position.
// @param num int: The index of the BitmapHandler to update.
// @return error: Returns nil if the operation succeeds or an error if movement fails.
func (spriteObject *spriteObject) MoveObject(x, y, num int) error {
	err := spriteObject.bitmapObject.GetBitmapHandler(num).SetCoords(x, y)
	if err != nil {
		return err
	}
	return nil
}
