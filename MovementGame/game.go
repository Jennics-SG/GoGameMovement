package MovementGame

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 600
	SCREEN_HEIGHT = 600
)

var (
	bgColor = color.RGBA{50, 100, 50, 50}
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Layout(outsideWidth, insideWidth int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)
}
