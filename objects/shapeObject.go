package objects

type ShapeObject interface {
	GetDrawableObject() DrawableObject
	GetTransformableObject() TransformableObject
}

type shapeObject struct {
	drawableObject      DrawableObject
	transformableObject TransformableObject
}

func NewShapeObject(drawableObject_ DrawableObject, transforableObject_ TransformableObject) ShapeObject {
	return &shapeObject{
		drawableObject:      drawableObject_,
		transformableObject: transforableObject_,
	}
}
func (shapeObject *shapeObject) GetDrawableObject() DrawableObject {
	return shapeObject.drawableObject
}

func (shapeObject *shapeObject) GetTransformableObject() TransformableObject {
	return shapeObject.transformableObject
}
