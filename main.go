package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"Game_Engine/objects"

	"github.com/hajimehoshi/ebiten/v2"
)

var errorLogger *log.Logger

func init() {
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	errorLogger = log.New(f, "", log.LstdFlags)
}

type Game struct {
	buttonImage     *ebiten.Image
	backgroundColor color.Color
	IsPressed       bool
}

func NewGame() *Game {
	buttonImage := ebiten.NewImage(200, 100)
	buttonImage.Fill(color.RGBA{220, 220, 220, 255})
	return &Game{
		buttonImage:     buttonImage,
		backgroundColor: color.Black,
		IsPressed:       false,
	}
}

func (g *Game) setBackgroundColor(R int, G int, B int, A int) {
	if R < 0 || R > 255 || G < 0 || G > 255 || B < 0 || B > 255 || A < 0 || A > 255 {
		logError(fmt.Errorf("Colors must be in the range of 0 to 255"))
	}
	g.backgroundColor = color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
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
		if x >= buttonLeft && x <= buttonRight && y >= buttonTop && y <= buttonBottom {
			g.buttonImage.Fill(color.RGBA{128, 128, 128, 255})
			fmt.Println("Button Test")
			switch g.backgroundColor {
			case color.Black:
				g.setBackgroundColor(70, 70, 70, 100)
			case color.RGBA{70, 70, 70, 100}:
				g.setBackgroundColor(255, 255, 255, 255)
			case color.RGBA{255, 255, 255, 255}:
				g.setBackgroundColor(255, 0, 132, 100)
			case color.RGBA{255, 0, 132, 100}:
				g.setBackgroundColor(4, 0, 255, 100)
			default:
				g.backgroundColor = color.Black
			}
		}
	}
	if !IsCurPressed && g.IsPressed {
		g.buttonImage.Fill(color.RGBA{220, 220, 220, 255})
	}
	g.IsPressed = IsCurPressed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := ebiten.WindowSize()
	screen.Fill(g.backgroundColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth)/2-100, float64(screenHeight)/2-50)

	screen.DrawImage(g.buttonImage, op)
	ebiten.SetWindowTitle("Game Engine")

	col := color.RGBA{150, 100, 200, 255} // Setting the color of segment/square
	testSegment := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSegment.Segment(200, 100, 200, 50, col)

	testSquare1 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSquare2 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSquare1.DrawSquare(50, 200, 200, col)
	testSquare2.DrawSquare(650, 200, 100, col)

}

func (g *Game) Layout(int, int) (int, int) {
	return ebiten.WindowSize()
}

func logError(err error) {
	errorLogger.Println(err)
}

func (g *Game) drawsquare(screen *ebiten.Image, X int, Y int, S int, col color.Color) error {
	if X <= 0 && Y <= 0 && S < 1 {
		logError(fmt.Errorf("Square should be on the screen and not smaller than 1 px"))
		return nil
	}
	g.Segment(screen, X, Y, X+S, Y, col)
	g.Segment(screen, X, Y, X, Y+S, col)
	g.Segment(screen, X+S, Y, X+S, Y+S, col)
	g.Segment(screen, X, Y+S, X+S, Y+S, col)

	return nil
}

func (g *Game) Segment(screen *ebiten.Image, startX int, startY int, finalX int, finalY int, col color.Color) error {
	deltX := finalX - startX
	deltY := finalY - startY

	if deltX == 0 && deltY == 0 {
		logError(fmt.Errorf("Line can't be 0"))
		return nil // No line to draw
	}

	if deltX == 0 { // Vertical line case
		step := 1
		if deltY < 0 {
			step = -1
		}
		for y := startY; y != finalY+step; y += step {
			g.plotPixel(screen, startX, y, col)
		}
		return nil
	}

	var slope float64
	if deltX != 0 {
		slope = float64(deltY) / float64(deltX) // Calculate slope
	}

	if absolute(slope) <= 1 { // Case where |slope| <= 1
		y := float64(startY)
		step := 1
		if deltX < 0 {
			step = -1
		}
		for x := startX; x != finalX+step; x += step {
			g.plotPixel(screen, x, int(y), col)
			y += slope // Increment y
		}
	} else { // Case where |slope| > 1, swap roles of x and y
		x := float64(startX)
		step := 1
		if deltY < 0 {
			step = -1
		}
		for y := startY; y != finalY+step; y += step {
			g.plotPixel(screen, int(x), y, col) // Plot at rounded (x, y)
			x += 1 / slope                      // Increment x
		}
	}

	return nil
}

func (g *Game) plotPixel(screen *ebiten.Image, x int, y int, col color.Color) {
	// Draw a pixel by setting it on the screen image
	screen.Set(x, y, col)

}

func absolute(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	tps := flag.Int("tps", 60, "Number of ticks per second (TPS)")
	flag.Parse()
	game := NewGame()

	width, height := 800, 600
	if width <= 0 || height <= 0 {
		logError(fmt.Errorf("invalid window size: %d x %d", width, height))
	} else {
		ebiten.SetWindowSize(width, height)
	}
	ebiten.SetTPS(*tps)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		logError(err)
	}

	logError(fmt.Errorf("this is a test error to check logging"))
}
