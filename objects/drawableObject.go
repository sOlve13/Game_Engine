package objects

type DrawableObject interface {
	GetGameObject() GameObject
	Draw() error
	UnDraw() error
}

type drawableObject struct {
	gameObject GameObject
	isDrawn    bool
}

func NewDrawableObject(gameObject GameObject) DrawableObject {
	return &drawableObject{
		gameObject: gameObject,
		isDrawn:    false,
	}
}

func (drawableObject *drawableObject) GetGameObject() GameObject {
	return drawableObject.gameObject
}
func (drawableObject *drawableObject) Draw() error {
	drawableObject.isDrawn = true
	return nil
}
func (drawableObject *drawableObject) UnDraw() error {
	drawableObject.isDrawn = false
	return nil
}

func (drawableObject *drawableObject) GetIsDrawn() bool {
	return drawableObject.isDrawn
}
