package objects

import (
	"errors"
	"io/ioutil"
	"strconv"
)

type AnimatedObject interface {
	GetGameObject() GameObject
	Animate(numerOfFrame int) error
	GetCurrentFrame() int
}

type animatedObject struct {
	gameObject     GameObject
	numberOfFrames int
	currentFrame   int
	name           string
}

func NewAnimatedObject(gameObject GameObject, numberOfFrames int, name string) AnimatedObject {
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

func (animatedObject *animatedObject) GetGameObject() GameObject {
	return animatedObject.gameObject
}
func (animatedObject *animatedObject) Animate(numerOfFrame int) error {
	if 0 <= numerOfFrame && numerOfFrame < animatedObject.numberOfFrames {
		writeToFile(animatedObject.name+".txt", strconv.Itoa(numerOfFrame))
		animatedObject.currentFrame = numerOfFrame
	} else {
		return errors.New("animation error: frame>maximum")
	}
	writeToFile(animatedObject.name+".txt", strconv.Itoa(numerOfFrame))
	animatedObject.currentFrame = numerOfFrame
	return nil
}
func (animatedObject *animatedObject) GetCurrentFrame() int {
	return animatedObject.currentFrame
}
