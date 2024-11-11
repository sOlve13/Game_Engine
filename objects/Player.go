package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player interface {
	DrawPlayerSquare(color.Color, int, int)
	//MovePlayer(dx, dy int)
}

type player struct {
	screen          *ebiten.Image
	health          int
	name            string
	backgroundColor color.Color
	x               int
	y               int
	size            int
}

func NewPlayer(screen *ebiten.Image, backgroundCol color.Color, col color.Color, name string, health int, x, y, s int) Player {
	return &player{
		screen:          screen,
		health:          health,
		name:            name,
		backgroundColor: backgroundCol,
		x:               x,
		y:               y,
		size:            s,
	}
}

func (player *player) DrawPlayerSquare(col color.Color, dx int, dy int) {
	player1 := NewPrimitiveRendererclass(player.screen, player.backgroundColor)
	player1.DrawSquare(player.x, player.y, player.size, col)
	player1.TranslateSquare(player.x, player.y, player.size, dx, dy, col)
}
