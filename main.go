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

	testSegment := objects.NewLineSegment(screen, g.backgroundColor)
	testSegment.Segment(objects.NewPoint2D(screen, g.backgroundColor, 200, 100, col), objects.NewPoint2D(screen, g.backgroundColor, 200, 300, col), col)
	testSegment.ChangeStart(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, col))

	testSegmentDefault := objects.NewLineSegment(screen, g.backgroundColor)
	testSegmentDefault.Segment(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, col), objects.NewPoint2D(screen, g.backgroundColor, 300, 300, col), col)
	testSegmentDefault.ChangeFinal(objects.NewPoint2D(screen, g.backgroundColor, 200, 100, col))

	testSquare1 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSquare2 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSquare1.DrawSquare(50, 200, 200, col)
	testSquare2.DrawSquare(650, 200, 100, col)

	testPolyline := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	points := []objects.Point2D{
		objects.NewPoint2D(screen, g.backgroundColor, 500, 200, col),
		objects.NewPoint2D(screen, g.backgroundColor, 600, 300, col),
		objects.NewPoint2D(screen, g.backgroundColor, 700, 200, col),
	}
	testPolyline.DrawPolyline(points, col)
}

func (g *Game) Layout(int, int) (int, int) {
	return ebiten.WindowSize()
}

func logError(err error) {
	errorLogger.Println(err)
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
