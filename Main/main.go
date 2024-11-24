// Package main store the realisation of game
package main_

import (
	"Game_Engine/objects"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Variable which stores error logger
var errorLogger *log.Logger

// Itialisation of error logger
func init() {
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	errorLogger = log.New(f, "", log.LstdFlags)
}

// Struct which holds crusial for engine objects
type Game struct {
	buttonImage                              *ebiten.Image
	backgroundColor                          color.Color
	IsPressed                                bool
	xTranslate                               int
	yTranslate                               int
	translationSpeed                         int
	angle                                    int
	isRight, isLeft, isTop, isDown, isAttack bool
}

// Initalisation of Game with
// @param screenWidth, screenHeight int: which supply information about size of screen
func NewGame(screenWidth, screenHeight int) *Game {
	buttonImage := ebiten.NewImage(200, 100)
	buttonImage.Fill(color.RGBA{220, 220, 220, 255})

	return &Game{
		buttonImage:      buttonImage,
		backgroundColor:  color.Black,
		IsPressed:        false,
		xTranslate:       0,
		yTranslate:       0,
		translationSpeed: 1,
		angle:            0,
		isRight:          false,
		isTop:            false,
		isAttack:         false,
		isLeft:           false,
		isDown:           false,
	}
}

// Function for setting background color
// @param R int, G int, B int, A int: colors
func (g *Game) setBackgroundColor(R int, G int, B int, A int) {
	if R < 0 || R > 255 || G < 0 || G > 255 || B < 0 || B > 255 || A < 0 || A > 255 {
		logError(fmt.Errorf("Colors must be in the range of 0 to 255"))
	}
	g.backgroundColor = color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
}

// Function which is beeing runned every tick to update information about game
func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.yTranslate = g.yTranslate - g.translationSpeed
		g.isTop = true
	} else {
		g.isTop = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.yTranslate = g.yTranslate + g.translationSpeed
		g.isDown = true
	} else {
		g.isDown = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.xTranslate = g.xTranslate - g.translationSpeed
		g.isLeft = true
	} else {
		g.isLeft = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.xTranslate = g.xTranslate + g.translationSpeed
		g.isRight = true
	} else {
		g.isRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		if g.translationSpeed <= 1 {
			g.translationSpeed = 1
		} else {
			g.translationSpeed = g.translationSpeed - 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		g.translationSpeed = g.translationSpeed + 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.angle >= 360 {
			g.angle = g.angle - 360
		}

		g.angle = g.angle + 1
	}
	/*
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
			g.IsPressed = IsCurPressed*/
	return nil
}

// Draw function draws everything on screen, which is given as paramiter
func (g *Game) Draw(screen *ebiten.Image) {
	ebiten.SetWindowTitle("Game Engine")
	screenWidth, screenHeight := ebiten.WindowSize()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth)/2-100, float64(screenHeight)/2-50)
	col := color.RGBA{150, 100, 200, 255}
	white := color.RGBA{255, 255, 255, 255} // White color with full opacity

	//Test full layer of constructors

	testSegment := objects.NewLineSegment(screen, g.backgroundColor)
	testSegment.Segment(objects.NewPoint2D(screen, g.backgroundColor, 0, 0, white), objects.NewPoint2D(screen, g.backgroundColor, 200, 300, white), white)
	//testSegment.ChangeStart(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, white))

	gmOb := objects.NewGameObject(screen, g.backgroundColor)
	drawOb := objects.NewDrawableObject(gmOb)
	tranOb := objects.NewTransformableObject(gmOb)
	shapOb := objects.NewShapeObject(drawOb, tranOb)

	squaOb1 := objects.NewSquareObject(shapOb, 100, 100, 100, col)
	squaOb1.Draw()
	squaOb1.Scale(2)
	squaOb1.Rotate(-30)
	squaOb1.Translate(200, 200)
	squaOb2 := objects.EnhancedNewSquareObject(screen, g.backgroundColor, 100, 100, 100, col)
	squaOb2.Draw()
	squaOb2.Rotate(g.angle)
	squaOb2.Scale(2)
	squaOb2.Translate(100, 100)

	squaOb3 := objects.EnhancedNewSquareObject(screen, g.backgroundColor, 500, 500, 100, white)
	squaOb3.Draw()
	squaOb3.Rotate(g.angle)
	squaOb3.Scale(2)
	squaOb3.Translate(100, 100)

	lineOb1 := objects.EnhancedNewLineObject(screen, g.backgroundColor, 500, 500, 600, 600, col)

	lineOb1.Draw()
	lineOb1.Translate(-300, -400)
	lineOb1.Scale(4)

	lineOb1.Rotate(g.angle)
	lineOb2 := objects.EnhancedNewLineObject(screen, g.backgroundColor, 500, 500, 600, 600, col)
	lineOb2.Draw()
	lineOb2.Translate(-400, -300)
	lineOb2.Scale(4)

	circOb1 := objects.EnhancedNewCircleObject(screen, g.backgroundColor, 100, 600, 40, col)
	circOb1.Draw()
	circOb1.Scale(2)
	circOb1.Translate(0, -200)

	x, y := 100, 100

	player := objects.NewPlayerObject(screen, g.backgroundColor, col, x+g.xTranslate, y+g.yTranslate)
	err := player.LoadHero("Movement")

	player.SetRightMovement(createRange(12, 17))
	player.SetLeftMovement(createRange(6, 11))
	player.SetTopMovement(createRange(0, 5))
	player.SetDownMovement(createRange(18, 23))
	player.SetCalm(18)
	err = player.Move(g.isRight, g.isLeft, g.isTop, g.isDown, g.isAttack, g.xTranslate, g.yTranslate)
	if err != nil {
		logError(err)
	}

	if err != nil {
		logError(err)
	}
	/*
		col = color.RGBA{150, 100, 200, 255} // Setting the color of segment/square
		col2 := color.RGBA{50, 100, 200, 255}
		testSegment := objects.NewLineSegment(screen, g.backgroundColor)
		testSegment.Segment(objects.NewPoint2D(screen, g.backgroundColor, 0, 0, white), objects.NewPoint2D(screen, g.backgroundColor, 200, 300, white), white)
		testSegment.ChangeStart(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, col2))

		testSegmentDefault := objects.NewLineSegment(screen, g.backgroundColor)
		testSegmentDefault.Segment(objects.NewPoint2D(screen, g.backgroundColor, 300, 100, col), objects.NewPoint2D(screen, g.backgroundColor, 300, 300, col), col)
		testSegmentDefault.ChangeFinal(objects.NewPoint2D(screen, g.backgroundColor, 200, 100, col))

		testSquare1 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		testSquare2 := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		testSquare1.DrawSquare(50, 200, 200, 0, col)
		testSquare2.DrawSquare(950, 200, 100, 0, col)
		testPolyline := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)

		points := []objects.Point2D{
			objects.NewPoint2D(screen, g.backgroundColor, 500, 200, col),
			objects.NewPoint2D(screen, g.backgroundColor, 600, 300, col),
			objects.NewPoint2D(screen, g.backgroundColor, 700, 200, col),
			objects.NewPoint2D(screen, g.backgroundColor, 500, 200, col),
		}

		testPolyline.DrawPolyline(points, col)
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

		testPolygon := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		err = testPolygon.DrawPolygon(points2, col2)
		logError(err)
		testCircle := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		centerCircle := objects.NewPoint2D(screen, g.backgroundColor, 100, 100, col)
		x_, y_ := centerCircle.GetCoords()
		testCircle.DrawCircle(x_, y_, 50, col)
		centerEllipse := objects.NewPoint2D(screen, g.backgroundColor, 700, 700, col)
		testEllipse := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		testEllipse.DrawEllipse(centerEllipse, 100, 50, col)
		testBorderFill := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		testBorderFill.BorderFill(101, 102, col2, col)

		testFillSquare := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		testFillSquare.FillSquare(50, 200, 200, col)

		testFloodFill := objects.NewPrimitiveRendererclass(screen, g.backgroundColor)
		testFloodFill.FloodFill(951, 201, col, g.backgroundColor)
	*/
}

// Function which return windowsize of game
func (g *Game) Layout(int, int) (int, int) {
	return ebiten.WindowSize()
}

// Function for logging errors
func logError(err error) {
	errorLogger.Println(err)
}

// Main function which create game and handle other functions so everything can work fine
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
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		logError(err)
	}

}

// Help function for creating lists in range of given numbers
func createRange(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}
