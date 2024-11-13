package objects

type TransformableObject interface {
	GetGameObject() GameObject
	Rotate(angle int) error
	Scale(scale int) error
	Translate(x, y int) error
	GetScale() int
	GetAngle() int
	GetTranslationX() int
	GetTranslationY() int
}

type transformableObject struct {
	gameObject   GameObject
	scale        int
	angle        int
	translationX int
	translationY int
}

func NewTransformableObject(gameObject GameObject) TransformableObject {
	return &transformableObject{
		gameObject:   gameObject,
		scale:        1,
		angle:        0,
		translationX: 0,
		translationY: 0,
	}
}

func (transformableObject *transformableObject) GetGameObject() GameObject {
	return transformableObject.gameObject
}

func (transformableObject *transformableObject) Rotate(angle int) error {
	transformableObject.angle = angle
	return nil
}

func (transformableObject *transformableObject) Scale(scale int) error {
	transformableObject.angle = scale
	return nil
}

func (transformableObject *transformableObject) Translate(x, y int) error {
	transformableObject.translationX = x
	transformableObject.translationY = y
	return nil
}

func (t *transformableObject) GetScale() int {
	return t.scale
}

func (t *transformableObject) GetAngle() int {
	return t.angle
}

func (t *transformableObject) GetTranslationX() int {
	return t.translationX
}

func (t *transformableObject) GetTranslationY() int {
	return t.translationY
}
