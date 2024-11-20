package objects

// ShapeObject represents an object that has both drawable and transformable characteristics.
// It provides methods to get the associated drawable and transformable objects.
// Also this object inherit TransformableObject and DrawableObject
type ShapeObject interface {
	// GetDrawableObject returns the drawable object associated with the shape object.
	// @return DrawableObject: The associated drawable object of the shape object.
	GetDrawableObject() DrawableObject

	// GetTransformableObject returns the transformable object associated with the shape object.
	// @return TransformableObject: The associated transformable object of the shape object.
	GetTransformableObject() TransformableObject
}

// shapeObject is an internal implementation of the ShapeObject interface.
// It holds references to a drawable object and a transformable object.
type shapeObject struct {
	drawableObject      DrawableObject      // The associated drawable object.
	transformableObject TransformableObject // The associated transformable object.
}

// NewShapeObject creates a new instance of a shape object with the specified drawable object and transformable object.
// @param drawableObject_ DrawableObject: The drawable object to be associated with the shape object.
// @param transforableObject_ TransformableObject: The transformable object to be associated with the shape object.
// @return ShapeObject: A new instance of the shape object with the provided drawable and transformable objects.
func NewShapeObject(drawableObject_ DrawableObject, transforableObject_ TransformableObject) ShapeObject {
	return &shapeObject{
		drawableObject:      drawableObject_,
		transformableObject: transforableObject_,
	}
}

// GetDrawableObject returns the associated drawable object of the shape object.
// @return DrawableObject: The associated drawable object.
func (shapeObject *shapeObject) GetDrawableObject() DrawableObject {
	return shapeObject.drawableObject
}

// GetTransformableObject returns the associated transformable object of the shape object.
// @return TransformableObject: The associated transformable object.
func (shapeObject *shapeObject) GetTransformableObject() TransformableObject {
	return shapeObject.transformableObject
}
