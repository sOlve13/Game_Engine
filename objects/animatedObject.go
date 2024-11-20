package objects

import (
	"errors"
	"io/ioutil"
	"strconv"
)

// AnimatedObject represents an object that can be animated by changing its frame.
// It provides methods to get the associated game object, animate to a specific frame,
// and get the current frame of the animation.
// Also this object inherit GameObject(the basic class of hierarchy)
type AnimatedObject interface {
	// GetGameObject returns the associated game object of the animated object.
	// @return GameObject: The associated game object.
	GetGameObject() GameObject

	// Animate animates the object by changing its frame to the specified frame number.
	// @param numerOfFrame int: The frame number to which the object should be animated.
	// @return error: Returns an error if the frame number is out of bounds.
	Animate(numerOfFrame int) error

	// GetCurrentFrame returns the current frame of the animation.
	// @return int: The current frame number of the animation.
	GetCurrentFrame() int
}

// animatedObject is an internal implementation of the AnimatedObject interface.
// It contains a reference to a game object, the number of frames in the animation,
// the current frame of the animation, and the name used to store animation data.
type animatedObject struct {
	gameObject     GameObject // The associated game object.
	numberOfFrames int        // The total number of frames in the animation.
	currentFrame   int        // The current frame of the animation.
	name           string     // The name used to store and load animation data.
}

// NewAnimatedObject creates a new instance of an animated object with the specified game object,
// number of frames in the animation, and the name for storing animation data.
// It also reads the current frame from a file based on the provided name.
// @param gameObject GameObject: The game object to be associated with the animated object.
// @param numberOfFrames int: The total number of frames in the animation.
// @param name string: The name used to store the animation state in a file.
// @return AnimatedObject: A new instance of the animated object with the provided parameters.
func NewAnimatedObject(gameObject GameObject, numberOfFrames int, name string) AnimatedObject {
	// Load the current frame from a file
	filename := name + ".txt"
	data, _ := ioutil.ReadFile(filename)
	current_f, _ := strconv.Atoi(string(data))
	return &animatedObject{
		gameObject:     gameObject,
		numberOfFrames: numberOfFrames,
		currentFrame:   current_f,
		name:           name,
	}
}

// GetGameObject returns the associated game object of the animated object.
// @return GameObject: The associated game object.
func (animatedObject *animatedObject) GetGameObject() GameObject {
	return animatedObject.gameObject
}

// Animate animates the object by changing its frame to the specified frame number.
// If the frame number is valid (within the range of frames), it writes the new frame number to a file
// and updates the current frame of the object.
// @param numerOfFrame int: The frame number to animate to.
// @return error: Returns an error if the frame number is invalid (greater than or equal to the maximum number of frames).
func (animatedObject *animatedObject) Animate(numerOfFrame int) error {
	if 0 <= numerOfFrame && numerOfFrame < animatedObject.numberOfFrames {
		// Write the new frame number to a file
		writeToFile(animatedObject.name+".txt", strconv.Itoa(numerOfFrame))
		animatedObject.currentFrame = numerOfFrame
	} else {
		// Return an error if the frame number is invalid
		return errors.New("animation error: frame>maximum")
	}
	writeToFile(animatedObject.name+".txt", strconv.Itoa(numerOfFrame))
	animatedObject.currentFrame = numerOfFrame
	return nil
}

// GetCurrentFrame returns the current frame number of the animation.
// @return int: The current frame number.
func (animatedObject *animatedObject) GetCurrentFrame() int {
	return animatedObject.currentFrame
}
