package objects

// TransformableObject represents an object that can be transformed.
// It provides methods for rotating, scaling, and translating the object,
// as well as retrieving the current transformation values (scale, angle, and translation).
// Also this object inherit GameObject(the basic class of hierarchy)
type TransformableObject interface {
	// GetGameObject returns the associated game object of the transformable object.
	// @return GameObject: The associated game object.
	GetGameObject() GameObject

	// Rotate rotates the object by the specified angle.
	// @param angle int: The angle by which to rotate the object.
	// @return error: Returns nil if the rotation is applied successfully.
	Rotate(angle int) error

	// Scale scales the object by the specified scale factor.
	// @param scale int: The scaling factor for the object.
	// @return error: Returns nil if the scaling is applied successfully.
	Scale(scale int) error

	// Translate moves the object by the specified x and y values.
	// @param x int: The x translation value.
	// @param y int: The y translation value.
	// @return error: Returns nil if the translation is applied successfully.
	Translate(x, y int) error

	// GetScale returns the current scale factor of the object.
	// @return int: The current scale of the object.
	GetScale() int

	// GetAngle returns the current rotation angle of the object.
	// @return int: The current angle of the object.
	GetAngle() int

	// GetTranslationX returns the current x translation of the object.
	// @return int: The current x translation value of the object.
	GetTranslationX() int

	// GetTranslationY returns the current y translation of the object.
	// @return int: The current y translation value of the object.
	GetTranslationY() int
}

// transformableObject is an internal implementation of the TransformableObject interface.
// It contains a reference to a game object, scale, angle, and translation values.
type transformableObject struct {
	gameObject   GameObject // The associated game object.
	scale        int        // The current scale factor of the object.
	angle        int        // The current rotation angle of the object.
	translationX int        // The current x translation value of the object.
	translationY int        // The current y translation value of the object.
}

// NewTransformableObject creates a new instance of a transformable object with the specified game object.
// The object is initially set with default transformations: scale=1, angle=0, and translationX=0, translationY=0.
// @param gameObject GameObject: The game object to be associated with the transformable object.
// @return TransformableObject: A new instance of the transformable object with default transformations.
func NewTransformableObject(gameObject GameObject) TransformableObject {
	return &transformableObject{
		gameObject:   gameObject,
		scale:        1,
		angle:        0,
		translationX: 0,
		translationY: 0,
	}
}

// GetGameObject returns the associated game object of the transformable object.
// @return GameObject: The associated game object.
func (transformableObject *transformableObject) GetGameObject() GameObject {
	return transformableObject.gameObject
}

// Rotate rotates the object by the specified angle and updates the angle.
// @param angle int: The angle to rotate the object.
// @return error: Returns nil if the rotation is successfully applied.
func (transformableObject *transformableObject) Rotate(angle int) error {
	transformableObject.angle = angle
	return nil
}

// Scale scales the object by the specified scale factor and updates the scale.
// @param scale int: The scaling factor for the object.
// @return error: Returns nil if the scaling is successfully applied.
func (transformableObject *transformableObject) Scale(scale int) error {
	transformableObject.scale = scale
	return nil
}

// Translate translates the object by the specified x and y values and updates the translation.
// @param x int: The x translation value.
// @param y int: The y translation value.
// @return error: Returns nil if the translation is successfully applied.
func (transformableObject *transformableObject) Translate(x, y int) error {
	transformableObject.translationX = x
	transformableObject.translationY = y
	return nil
}

// GetScale returns the current scale factor of the transformable object.
// @return int: The current scale of the object.
func (t *transformableObject) GetScale() int {
	return t.scale
}

// GetAngle returns the current angle of rotation for the transformable object.
// @return int: The current angle of the object.
func (t *transformableObject) GetAngle() int {
	return t.angle
}

// GetTranslationX returns the current x translation value of the transformable object.
// @return int: The current x translation value of the object.
func (t *transformableObject) GetTranslationX() int {
	return t.translationX
}

// GetTranslationY returns the current y translation value of the transformable object.
// @return int: The current y translation value of the object.
func (t *transformableObject) GetTranslationY() int {
	return t.translationY
}
