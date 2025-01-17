// Application logic

package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game constants
const (
	SCREEN_WIDTH  = 600
	SCREEN_HEIGHT = 600
)

// Globals
var (
	bgColor = color.RGBA{50, 100, 50, 50}
)

// Class variables
type Game struct {
	cache AssetCache
}

// Constructor
func NewGame() *Game {
	return &Game{}
}

// Initialise game
func (g *Game) Init() {
	// Load player sprite
	g.cache.addAsset("assets/player.png")
}

// Get screen dims
func (g *Game) Layout(outsideWidth, insideWidth int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// Update logic of game
func (g *Game) Update() error {
	return nil
}

// Draw to game screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)
}
