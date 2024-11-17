package objects

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SpriteObject interface {
	GetBitmapObject() BitmapObject
	GetAnimatedObject() AnimatedObject
	LoadBitmaps(folderPath string, bmNum int) error
	GetName(num int) (string, error)
	SetBitmap(name int, bmNum int) error
	MoveObject(x, y, num int) error
}

type spriteObject struct {
	bitmapObject   BitmapObject
	animatedObject AnimatedObject
	dictionary     map[int]string
	name           string
}

func NewSpriteObject(bitmapObject BitmapObject, name string) SpriteObject {
	return &spriteObject{
		bitmapObject:   bitmapObject,
		animatedObject: nil,
		dictionary:     make(map[int]string),
		name:           name,
	}
}

func (spriteObject *spriteObject) GetBitmapObject() BitmapObject {
	return spriteObject.bitmapObject
}

func (spriteObject *spriteObject) GetAnimatedObject() AnimatedObject {
	return spriteObject.animatedObject
}

func (spriteObject *spriteObject) LoadBitmaps(folderPath string, bmNum int) error {
	files, err := os.ReadDir(folderPath)

	if err != nil {
		return err
	}
	for i, file := range files {
		// Пропускаем директории
		if file.IsDir() {
			continue
		}

		// Извлекаем имя файла без расширения
		fileName := file.Name()
		nameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		// Записываем в карту, где ключ — это название файла без расширения, а значение — номер
		spriteObject.dictionary[i] = nameWithoutExt
	}
	amOb := NewAnimatedObject(spriteObject.bitmapObject.GetDrawableObject().GetGameObject(), len(spriteObject.dictionary), spriteObject.name)
	spriteObject.animatedObject = amOb

	for key := range spriteObject.dictionary {
		width, height, err := getImageSize(folderPath + "/" + spriteObject.dictionary[key] + ".png")
		if err != nil {
			fmt.Println("Ошибка при получении размеров изображения:", err)
			return err
		}
		spriteObject.GetBitmapObject().GetBitmapHandler(bmNum).Create(spriteObject.dictionary[key], width, height, spriteObject.animatedObject.GetGameObject().GetBackgroundColor())
		err = spriteObject.GetBitmapObject().GetBitmapHandler(bmNum).Load(spriteObject.dictionary[key], folderPath+"/"+spriteObject.dictionary[key]+".png")
		if err != nil {
			return err
		}
	}

	return nil
}

func (spriteObject *spriteObject) SetBitmap(name int, bmNum int) error {
	err := spriteObject.animatedObject.Animate(name)
	if err != nil {
		return err
	}
	val, err := spriteObject.GetName(name)
	if err != nil {
		return err
	}
	err = spriteObject.bitmapObject.Draw(val, bmNum)
	if err != nil {
		return err
	}
	return nil

}

func (spriteObject *spriteObject) GetName(num int) (string, error) {
	val, exists := spriteObject.dictionary[num]
	if !exists {
		return "", errors.New("value not in list")
	}
	return val, nil
}

func (spriteObject *spriteObject) MoveObject(x, y, num int) error {
	err := spriteObject.bitmapObject.GetBitmapHandler(num).SetCoords(x, y)
	if err != nil {
		return err
	}
	return nil
}
