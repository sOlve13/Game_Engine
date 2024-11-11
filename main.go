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
	buttonImage      *ebiten.Image
	backgroundColor  color.Color
	IsPressed        bool
	squareX          float64
	squareY          float64
	translationX     float64
	translationY     float64
	translationSpeed float64 // Speed at which the square moves
	angle            float64
	angleSpeed       float64
	lineStartX       float64
	lineStartY       float64
	lineEndX         float64
	lineEndY         float64
	bol              bool
	plus             float64
	playertransX     float64
	playertransY     float64
	playerX          float64
	playerY          float64
	playerSpeed      float64
}

func NewGame(screenWidth, screenHeight int) *Game {
	buttonImage := ebiten.NewImage(200, 100)
	buttonImage.Fill(color.RGBA{220, 220, 220, 255})

	return &Game{
		buttonImage:      buttonImage,
		backgroundColor:  color.Black,
		IsPressed:        false,
		squareX:          0,   // Initial position of the square
		squareY:          0,   // Initial position of the square
		translationX:     0,   // Start at the same position
		translationY:     0,   // Start at the same position
		translationSpeed: 0.2, // Speed of movement per frame
		angle:            0,
		angleSpeed:       4,
		lineStartX:       300, // Initial starting point of the line
		lineStartY:       200,
		lineEndX:         500, // Initial end point of the line
		lineEndY:         400,
		bol:              false,
		plus:             1,
		playertransX:     0,
		playertransY:     0,
		playerX:          0,
		playerY:          0,
		playerSpeed:      1,
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

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.translationY -= g.translationSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.translationX -= g.translationSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.translationY += g.translationSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.translationX += g.translationSpeed
	} else {
		g.translationX = 0
		g.translationY = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyN) {
		g.bol = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyM) {
		g.plus += 0.2
	}
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		g.plus -= 0.2
	}

	// Update line's start and end points by translating their positions
	g.lineStartX += g.translationX
	g.lineStartY += g.translationY
	g.lineEndX += g.translationX
	g.lineEndY += g.translationY

	// Continuous rotation
	g.angle += 1 // Change the increment to control rotation speed

	g.squareX += g.translationX
	g.squareY += g.translationY

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
	g.angle += g.angleSpeed
	if g.angle > 360 {
		g.angle -= 360
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.playertransY -= g.playerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.playertransX -= g.playerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.playertransY += g.playerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.playertransX += g.playerSpeed
	} else {
		g.playertransX = 0
		g.playertransY = 0
	}

	g.playerX += g.playertransX
	g.playerY += g.playertransY

	g.IsPressed = IsCurPressed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebiten.SetWindowTitle("Game Engine")
	//screenWidth, screenHeight := ebiten.WindowSize()
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(float64(screenWidth)/2-100, float64(screenHeight)/2-50)

	col := color.RGBA{150, 100, 200, 255} // Setting the color of segment/square
	col2 := color.RGBA{50, 100, 200, 255}
	testSegment := objects.NewLineSegment(screen, g.backgroundColor)
	testSegment.Segment(objects.NewPoint2D(screen, g.backgroundColor, 200, 100, col), objects.NewPoint2D(screen, g.backgroundColor, 200, 300, col), col)
	testSegment.ChangeStart(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, col))

	testSquare2 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSquare2.DrawSquare(950, 200, 100, col)
	//testSquare2.RotateSquare(950, 200, 100, 50, col)
	testSquare2.TranslateSquare(950, 200, 100, int(g.squareX), int(g.squareY), color.RGBA{255, 0, 0, 255})

	testSegmentDefault := objects.NewLineSegment(screen, g.backgroundColor)
	testSegmentDefault.Segment(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, col), objects.NewPoint2D(screen, g.backgroundColor, 300, 300, col), col)
	testSegmentDefault.ChangeFinal(objects.NewPoint2D(screen, g.backgroundColor, 200, 100, col))

	testSquare1 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testSquare1.DrawSquare(50, 200, 200, col)
	testPolyline := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	pointsPol := []objects.Point2D{
		objects.NewPoint2D(screen, g.backgroundColor, 500, 200, col),
		objects.NewPoint2D(screen, g.backgroundColor, 600, 300, col),
		objects.NewPoint2D(screen, g.backgroundColor, 700, 200, col),
		objects.NewPoint2D(screen, g.backgroundColor, 500, 200, col),
	}
	//testPolyline.TranslatePolyline(pointsPol, int(g.squareX), int(g.squareY), col)
	testPolyline.DrawPolyline(pointsPol, col)
	test_dot := objects.NewPoint2D(screen, g.backgroundColor, 500, 500, col)
	test_dot.PlotPixel()
	test_dot1 := objects.NewPoint2D(screen, g.backgroundColor, 501, 500, col)
	test_dot1.PlotPixel()
	test_dot2 := objects.NewPoint2D(screen, g.backgroundColor, 502, 500, col)
	test_dot2.PlotPixel()

	points2 := []objects.Point2D{
		objects.NewPoint2D(screen, g.backgroundColor, 750, 250, col),
		objects.NewPoint2D(screen, g.backgroundColor, 800, 350, col),
		objects.NewPoint2D(screen, g.backgroundColor, 1000, 500, col),
		objects.NewPoint2D(screen, g.backgroundColor, 600, 500, col),
		objects.NewPoint2D(screen, g.backgroundColor, 600, 350, col),
		objects.NewPoint2D(screen, g.backgroundColor, 750, 250, col),
	}
	testPolygon1 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	err := testPolygon1.DrawPolygon(points2, col2)
	testPolygon1.TranslatePolygon(points2, int(g.squareX), int(g.squareY), col)
	logError(err)
	testCircle := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	centerCircle := objects.NewPoint2D(screen, g.backgroundColor, 100, 100, col)
	//testCircle.TranslateCircle(centerCircle, 50, int(g.squareX), int(g.squareY), col)
	testCircle.DrawCircle(centerCircle, 50, col)
	centerEllipse := objects.NewPoint2D(screen, g.backgroundColor, 700, 700, col)
	testEllipse := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	//testEllipse.TranslateEllipse(centerEllipse, 100, 50, int(g.squareX), int(g.squareY), col)
	testEllipse.DrawEllipse(centerEllipse, 100, 50, col)

	testBorderFill := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testBorderFill.BorderFill(101, 102, col2, col)

	testFillSquare := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	testFillSquare.FillSquare(50, 200, 200, col)

	//testFloodFill := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
	//testFloodFill.FloodFill(951, 201, col, g.backgroundColor)
	if g.bol == true {
		testSquare2.ScaleSquare(950+int(g.squareX), 200+int(g.squareY), 100, float64(g.plus), col)
	}

	newplayer := objects.NewPlayer(screen, g.backgroundColor, col, "Renat", 200, 1000, 800, 50)
	newplayer.DrawPlayerSquare(color.RGBA{134, 234, 55, 34}, int(g.playerX), int(g.playerY))
	//newplayer.MovePlayer(int(g.playerX), int(g.playerY))

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
	width, height := 800, 600
	game := NewGame(800, 600)
	if width <= 0 || height <= 0 {
		logError(fmt.Errorf("invalid window size: %d x %d", width, height))
	} else {
		ebiten.SetWindowSize(width, height)
	}
	ebiten.SetTPS(*tps)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		logError(err)
	}

	logError(fmt.Errorf("this is a test error to check logging"))
}
