package objects

// DrawableObject represents an object that can be drawn and undrawn.
// Also this object inherit GameObject(the basic class of hierarchy)
// It also provides a method to retrieve the associated game object.
type DrawableObject interface {
	// GetGameObject returns the associated game object of the drawable object.
	// @return GameObject: The associated game object.
	GetGameObject() GameObject

	// Draw draws the object, marking it as drawn.
	// @return error: Returns nil if the drawing operation is successful.
	Draw() error

	// UnDraw undraws the object, marking it as not drawn.
	// @return error: Returns nil if the undrawing operation is successful.
	UnDraw() error
}

// drawableObject is an internal implementation of the DrawableObject interface.
// It contains a reference to a game object and a flag indicating whether the object is drawn.
type drawableObject struct {
	gameObject GameObject // The associated game object.
	isDrawn    bool       // A flag indicating whether the object is drawn or not.
}

// NewDrawableObject creates a new instance of a drawable object with the specified game object.
// The object is initially set to not be drawn (isDrawn is false).
// @param gameObject GameObject: The game object to be associated with the drawable object.
// @return DrawableObject: A new instance of drawable object with the provided game object.
func NewDrawableObject(gameObject GameObject) DrawableObject {
	return &drawableObject{
		gameObject: gameObject,
		isDrawn:    false,
	}
}

// GetGameObject returns the associated game object of the drawable object.
// @return GameObject: The associated game object.
func (drawableObject *drawableObject) GetGameObject() GameObject {
	return drawableObject.gameObject
}

// Draw marks the object as drawn by setting the isDrawn flag to true.
// @return error: Returns nil if the drawing operation is successful.
func (drawableObject *drawableObject) Draw() error {
	drawableObject.isDrawn = true
	return nil
}

// UnDraw marks the object as not drawn by setting the isDrawn flag to false.
// @return error: Returns nil if the undrawing operation is successful.
func (drawableObject *drawableObject) UnDraw() error {
	drawableObject.isDrawn = false
	return nil
}

// GetIsDrawn returns the current state of the drawable object (whether it is drawn or not).
// @return bool: Returns true if the object is drawn, false otherwise.
func (drawableObject *drawableObject) GetIsDrawn() bool {
	return drawableObject.isDrawn
}
