package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	buttonImage     *ebiten.Image
	backgroundColor color.Color
	IsPressed       bool
}

func NewGame() *Game {
	// Create a button image
	buttonImage := ebiten.NewImage(200, 100)
	buttonImage.Fill(color.RGBA{0, 128, 0, 255})
	return &Game{buttonImage: buttonImage,
		backgroundColor: color.Black,
		IsPressed:       false,
	}
}

func (g *Game) Update() error {
	IsCurPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	screenWidth, screenHeight := ebiten.WindowSize()

	buttonWidth := g.buttonImage.Bounds().Dx()
	buttonHeight := g.buttonImage.Bounds().Dy()

	X := screenWidth / 2
	Y := screenHeight / 2

	buttonLeft := X - buttonWidth/2
	buttonRight := X + buttonWidth/2
	buttonTop := Y - buttonHeight/2
	buttonBottom := Y + buttonHeight/2

	if IsCurPressed && !g.IsPressed {
		x, y := ebiten.CursorPosition()
		// Check if the mouse is within the button's area
		if x >= buttonLeft && x <= buttonRight && y >= buttonTop && y <= buttonBottom {
			fmt.Println("Button Test")
			switch g.backgroundColor {
			case color.Black:
				g.backgroundColor = color.RGBA{100, 100, 100, 100}
			case color.RGBA{100, 100, 100, 100}:
				g.backgroundColor = color.RGBA{255, 255, 255, 255}
			case color.RGBA{255, 255, 255, 255}:
				g.backgroundColor = color.RGBA{255, 0, 132, 100}
			case color.RGBA{255, 0, 132, 100}:
				g.backgroundColor = color.RGBA{4, 0, 255, 100}
			default:
				g.backgroundColor = color.Black
			}

		}
	}
	g.IsPressed = IsCurPressed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := ebiten.WindowSize()
	screen.Fill(g.backgroundColor)

	// Draw the button image at the specified position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth)/2-100, float64(screenHeight)/2-50)

	screen.DrawImage(g.buttonImage, op)

	ebiten.SetWindowTitle("Game Engine")
}
func (g *Game) Layout(int, int) (int, int) {
	return ebiten.WindowSize()
}

func main() {
	game := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
