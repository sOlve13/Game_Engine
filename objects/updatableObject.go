package objects

// UpdatableObject represents an object that can be updated with a list of points and can retrieve that list.
// Also this object inherit GameObject(the basic class of hierarchy)
// It also has a reference to a game object associated with it.
type UpdatableObject interface {
	// GetPointsList returns the current list of points associated with the updatable object.
	// @return []Point2D: A slice of Point2D representing the points list.
	GetPointsList() []Point2D

	// Update functions updates the list of points associated with the updatable object.
	// @param pointsList []Point2D: A slice of Point2D to update the points list of the object.
	Update(pointsList []Point2D)
}

// updatableObject is an internal implementation of the UpdatableObject interface.
// It contains a reference to a GameObject and a list of points (Point2D).
type updatableObject struct {
	gameObject GameObject // The associated game object.
	poitsList  []Point2D  // The list of points associated with the object.
}

// NewUpdatableObject creates a new instance of an updatable object with the specified game object.
// The points list is initialized as an empty slice.
// @param gameObject GameObject: The game object to be associated with the updatable object.
// @return UpdatableObject: A new instance of updatable object with the provided game object and an empty points list.
func NewUpdatableObject(gameObject GameObject) UpdatableObject {
	return &updatableObject{
		gameObject: gameObject,
		poitsList:  make([]Point2D, 0),
	}
}

// Update updates the list of points associated with the updatable object.
// This method replaces the existing points list with the provided one.
// @param pointsList []Point2D: A slice of Point2D representing the new points list.
func (updatableObject *updatableObject) Update(pointsList []Point2D) {
	updatableObject.poitsList = pointsList
}

// GetPointsList returns the current list of points associated with the updatable object.
// @return []Point2D: A slice of Point2D representing the current points list.
func (updatableObject *updatableObject) GetPointsList() []Point2D {
	return updatableObject.poitsList
}

// GetGameObject returns the associated game object of the updatable object.
// @return GameObject: The associated game object.
func (updatableObject *updatableObject) GetGameObject() GameObject {
	return updatableObject.gameObject
}
