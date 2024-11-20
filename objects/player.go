package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// PlayerObject defines an interface for managing the player, its movement, and animations.
// This object inherits SpriteObject
type PlayerObject interface {
	// GetSpriteObject returns the associated SpriteObject.
	// @return SpriteObject: The SpriteObject associated with this PlayerObject.
	GetSpriteObject() SpriteObject

	// LoadHero loads the player's assets from the specified folder.
	// @param folderPath string: Path to the folder containing the player's assets.
	// @return error: Returns nil if successful, otherwise returns an error.
	LoadHero(folderPath string) error

	// SetRightMovement sets the frame sequence for moving to the right.
	// @param rmv []int: Frame indices for moving to the right.
	// @return error: Returns nil if successful, otherwise returns an error.
	SetRightMovement(rmv []int) error

	// SetLeftMovement sets the frame sequence for moving to the left.
	// @param lmv []int: Frame indices for moving to the left.
	// @return error: Returns nil if successful, otherwise returns an error.
	SetLeftMovement(lmv []int) error

	// SetTopMovement sets the frame sequence for moving upward.
	// @param tmv []int: Frame indices for moving upward.
	// @return error: Returns nil if successful, otherwise returns an error.
	SetTopMovement(tmv []int) error

	// SetDownMovement sets the frame sequence for moving downward.
	// @param dmv []int: Frame indices for moving downward.
	// @return error: Returns nil if successful, otherwise returns an error.
	SetDownMovement(dmv []int) error

	// SetAttack sets the frame sequence for attacking.
	// @param att []int: Frame indices for attacking.
	// @return error: Returns nil if successful, otherwise returns an error.
	SetAttack(att []int) error

	// SetCalm sets the frame index for the idle state.
	// @param cal int: The frame index for the idle state.
	// @return error: Returns nil if successful, otherwise returns an error.
	SetCalm(cal int) error
	// Move functioun recives information about buttons pressed and then transformate it in the movement
	// @param isRight bool: indicates that right arrow is pressed or not.
	// @param isLeft bool: indicates that left arrow is pressed or not.
	// @param isTop bool: indicates that top arrow is pressed or not.
	// @param isDown bool: indicates that down arrow is pressed or not.
	// @param isAttac bool: not yet implemented
	// @param x, y int: represent current position of player
	// @return error: Returns nil if successful, otherwise returns an error.
	Move(isRight, isLeft, isTop, isDown, isAttack bool, x, y int) error
}

// playerObject implements the PlayerObject interface.
// It manages the player's animations, movement, and actions.
type playerObject struct {
	spriteObject  SpriteObject // The SpriteObject associated with the player.
	calm          int          // Frame index for the idle state.
	rightMovement []int        // Frame sequence for moving to the right.
	leftMovement  []int        // Frame sequence for moving to the left.
	topMovement   []int        // Frame sequence for moving upward.
	downMovement  []int        // Frame sequence for moving downward.
	attack        []int        // Frame sequence for attacking.
}

// NewPlayerObject creates a new player object.
// @param screen *ebiten.Image: The game screen for rendering.
// @param backgroundColor color.Color: The background color.
// @param color color.Color: The player's primary color.
// @param x int: Initial x-coordinate of the player.
// @param y int: Initial y-coordinate of the player.
// @return PlayerObject: A new PlayerObject instance.
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

// GetSpriteObject returns the associated SpriteObject.
// @return SpriteObject: The SpriteObject associated with this PlayerObject.
func (playerObject *playerObject) GetSpriteObject() SpriteObject {
	return playerObject.spriteObject
}

// LoadHero loads the player's assets from the specified folder.
// @param folderPath string: Path to the folder containing the player's assets.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) LoadHero(folderPath string) error {
	err := playerObject.spriteObject.LoadBitmaps(folderPath, 0)
	if err != nil {
		return err
	}
	return nil
}

// SetRightMovement sets the frame sequence for moving to the right.
// @param rmv []int: Frame indices for moving to the right.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) SetRightMovement(rmv []int) error {
	playerObject.rightMovement = rmv
	return nil
}

// SetLeftMovement sets the frame sequence for moving to the left.
// @param lmv []int: Frame indices for moving to the left.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) SetLeftMovement(lmv []int) error {
	playerObject.leftMovement = lmv
	return nil
}

// SetTopMovement sets the frame sequence for moving upward.
// @param tmv []int: Frame indices for moving upward.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) SetTopMovement(tmv []int) error {
	playerObject.topMovement = tmv
	return nil
}

// SetDownMovement sets the frame sequence for moving downward.
// @param dmv []int: Frame indices for moving downward.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) SetDownMovement(dmv []int) error {
	playerObject.downMovement = dmv
	return nil
}

// SetAttack sets the frame sequence for attacking.
// @param att []int: Frame indices for attacking.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) SetAttack(att []int) error {
	playerObject.attack = att
	return nil
}

// SetCalm sets the frame index for the idle state.
// @param cal int: The frame index for the idle state.
// @return error: Returns nil if successful, otherwise returns an error.
func (playerObject *playerObject) SetCalm(cal int) error {
	playerObject.calm = cal
	return nil
}

// Move functioun recives information about buttons pressed and then transformate it in the movement
// @param isRight bool: indicates that right arrow is pressed or not.
// @param isLeft bool: indicates that left arrow is pressed or not.
// @param isTop bool: indicates that top arrow is pressed or not.
// @param isDown bool: indicates that down arrow is pressed or not.
// @param isAttac bool: not yet implemented
// @param x, y int: represent current position of player
// @return error: Returns nil if successful, otherwise returns an error.
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
