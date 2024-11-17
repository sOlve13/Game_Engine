package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerObject interface {
	GetSpriteObject() SpriteObject
	LoadHero(folderPath string) error
	SetRightMovement(rmv []int) error
	SetLeftMovement(lmv []int) error
	SetTopMovement(tmv []int) error
	SetDownMovement(dmv []int) error
	SetAttack(att []int) error
	SetCalm(cal int) error
	Move(isRight, isLeft, isTop, isDown, isAttack bool, x, y int) error
}

type playerObject struct {
	spriteObject  SpriteObject
	calm          int
	rightMovement []int
	leftMovement  []int
	topMovement   []int
	downMovement  []int
	attack        []int
}

func NewPlayerObject(screen *ebiten.Image, backgroundColor color.Color, color color.Color, x, y int) PlayerObject {
	gmob := NewGameObject(screen, backgroundColor)
	bmHd := NewBitmapHandler(x, y)
	var myList []BitmapHandler
	myList = append(myList, bmHd)

	bmOb := NewBitmapObject(myList, NewDrawableObject(gmob))
	spriteObject := NewSpriteObject(bmOb, "player")

	return &playerObject{
		spriteObject: spriteObject,
	}
}
func (playerObject *playerObject) GetSpriteObject() SpriteObject {
	return playerObject.spriteObject
}
func (playerObject *playerObject) LoadHero(folderPath string) error {
	err := playerObject.spriteObject.LoadBitmaps(folderPath, 0)
	if err != nil {
		return err
	}
	return nil
}
func (playerObject *playerObject) SetRightMovement(rmv []int) error {
	playerObject.rightMovement = rmv
	return nil
}

func (playerObject *playerObject) SetLeftMovement(lmv []int) error {
	playerObject.leftMovement = lmv
	return nil
}
func (playerObject *playerObject) SetTopMovement(tmv []int) error {
	playerObject.topMovement = tmv
	return nil
}
func (playerObject *playerObject) SetDownMovement(dmv []int) error {
	playerObject.downMovement = dmv
	return nil
}

func (playerObject *playerObject) SetAttack(att []int) error {
	playerObject.attack = att
	return nil
}
func (playerObject *playerObject) SetCalm(cal int) error {
	playerObject.calm = cal
	return nil
}

func (playerObject *playerObject) Move(isRight, isLeft, isTop, isDown, isAttack bool, x, y int) error {

	if !isAttack && !isDown && !isLeft && !isRight && !isTop {
		err := playerObject.spriteObject.SetBitmap(playerObject.calm, 0)
		playerObject.spriteObject.MoveObject(x, y, 0)

		if err != nil {
			return err
		}
		return nil
	}
	if isRight && !isDown && !isTop && !isLeft && !isAttack {
		if !contains(playerObject.rightMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) {
			err := playerObject.spriteObject.SetBitmap(playerObject.rightMovement[0], 0)
			playerObject.spriteObject.MoveObject(x, y, 0)

			if err != nil {
				return err
			}
		} else {
			if indexOf(playerObject.rightMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) < len(playerObject.rightMovement)-1 {
				err := playerObject.spriteObject.SetBitmap(playerObject.rightMovement[indexOf(playerObject.rightMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame())+1], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			} else {
				err := playerObject.spriteObject.SetBitmap(playerObject.rightMovement[0], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	if !isRight && isDown && !isTop && !isLeft && !isAttack {
		if !contains(playerObject.downMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) {
			err := playerObject.spriteObject.SetBitmap(playerObject.downMovement[0], 0)
			playerObject.spriteObject.MoveObject(x, y, 0)
			if err != nil {
				return err
			}
		} else {
			if indexOf(playerObject.downMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) < len(playerObject.downMovement)-1 {
				err := playerObject.spriteObject.SetBitmap(playerObject.downMovement[indexOf(playerObject.downMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame())+1], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			} else {
				err := playerObject.spriteObject.SetBitmap(playerObject.downMovement[0], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	if !isRight && !isDown && isTop && !isLeft && !isAttack {
		if !contains(playerObject.topMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) {
			err := playerObject.spriteObject.SetBitmap(playerObject.topMovement[0], 0)
			playerObject.spriteObject.MoveObject(x, y, 0)
			if err != nil {
				return err
			}
		} else {
			if indexOf(playerObject.topMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) < len(playerObject.topMovement)-1 {

				err := playerObject.spriteObject.SetBitmap(playerObject.topMovement[indexOf(playerObject.topMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame())+1], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			} else {
				err := playerObject.spriteObject.SetBitmap(playerObject.topMovement[0], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	if !isRight && !isDown && !isTop && isLeft && !isAttack {
		if !contains(playerObject.leftMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) {
			err := playerObject.spriteObject.SetBitmap(playerObject.leftMovement[0], 0)
			playerObject.spriteObject.MoveObject(x, y, 0)
			if err != nil {
				return err
			}
		} else {
			if indexOf(playerObject.leftMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame()) < len(playerObject.leftMovement)-1 {
				err := playerObject.spriteObject.SetBitmap(playerObject.leftMovement[indexOf(playerObject.leftMovement, playerObject.spriteObject.GetAnimatedObject().GetCurrentFrame())+1], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			} else {
				err := playerObject.spriteObject.SetBitmap(playerObject.leftMovement[0], 0)
				playerObject.spriteObject.MoveObject(x, y, 0)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	err := playerObject.spriteObject.SetBitmap(playerObject.calm, 0)
	playerObject.spriteObject.MoveObject(x, y, 0)
	if err != nil {
		return err
	}
	return nil

}
