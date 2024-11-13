package objects

type UpdatableObject interface {
	GetPointsList() []Point2D
	Update(pointsList []Point2D)
}

type updatableObject struct {
	gameObject GameObject
	poitsList  []Point2D
}

func NewUpdatableObject(gameObject GameObject) UpdatableObject {
	return &updatableObject{
		gameObject: gameObject,
		poitsList:  make([]Point2D, 0),
	}
}
func (updatableObject *updatableObject) Update(pointsList []Point2D) {
	updatableObject.poitsList = pointsList
}

func (updatableObject *updatableObject) GetPointsList() []Point2D {
	return updatableObject.poitsList
}

func (updatableObject *updatableObject) GetGameObject() GameObject {
	return updatableObject.gameObject
}
