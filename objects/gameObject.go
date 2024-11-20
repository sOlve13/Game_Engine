package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// GameObject represents an object in the game that can have a screen (image) and a background color.
// This object is the basic object in class hierarchy.
// It provides methods to get and set the screen and retrieve the background color.
type GameObject interface {
	// GetScreen returns the current screen (image) associated with the game object.
	// @return *ebiten.Image: The image (screen) of the game object.
	GetScreen() *ebiten.Image

	// GetBackgroundColor returns the background color of the game object.
	// @return color.Color: The background color of the game object.
	GetBackgroundColor() color.Color

	// SetScreen sets the screen (image) for the game object.
	// @param screen *ebiten.Image: The screen (image) to be associated with the game object.
	SetScreen(screen *ebiten.Image)
}

// gameObject is an internal implementation of the GameObject interface.
// It contains a screen (image) and a background color.
type gameObject struct {
	screen          *ebiten.Image
	backgroundColor color.Color
}

// NewGameObject creates a new instance of a game object with the specified screen (image) and background color.
// @param screen *ebiten.Image: The screen (image) to be associated with the game object.
// @param backgroundColor color.Color: The background color of the game object.
// @return GameObject: A new instance of the game object with the given screen and background color.
func NewGameObject(screen *ebiten.Image, backgroundColor color.Color) GameObject {
	return &gameObject{
		screen:          screen,
		backgroundColor: backgroundColor,
	}
}

// NewWScreenGameObject creates a new instance of a game object without a screen, only with a background color.
// The screen is initially set to nil.
// @param backgroundColor color.Color: The background color of the game object.
// @return GameObject: A new instance of the game object with the given background color and no screen.
func NewWScreenGameObject(backgroundColor color.Color) GameObject {
	return &gameObject{
		screen:          nil,
		backgroundColor: backgroundColor,
	}
}

// SetScreen sets the screen (image) for the game object.
// This method allows you to change the screen associated with the game object.
// @param screen *ebiten.Image: The screen (image) to set for the game object.
func (gameObject *gameObject) SetScreen(screen *ebiten.Image) {
	gameObject.screen = screen
}

// GetScreen returns the current screen (image) associated with the game object.
// @return *ebiten.Image: The current screen (image) of the game object.
func (gameObject *gameObject) GetScreen() *ebiten.Image {
	return gameObject.screen
}

// GetBackgroundColor returns the background color of the game object.
// @return color.Color: The background color of the game object.
func (gameObject *gameObject) GetBackgroundColor() color.Color {
	return gameObject.backgroundColor
}
