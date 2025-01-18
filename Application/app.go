// Application logic

package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// VARIABLES ==================================================================
const (
	SCREEN_WIDTH  = 600
	SCREEN_HEIGHT = 600
)

var (
	bgColor = color.RGBA{50, 100, 50, 50}
)

// ============================================================================
// Game Class =================================================================
type Game struct {
	cache  AssetCache
	player Player
}

// Constructor
func NewGame() *Game {
	return &Game{}
}

// Initialise game
func (g *Game) Init() {
	// Load player sprite
	g.cache.addAsset("assets/player.png", "player")

	// Create player
	g.player = *NewPlayer(
		*g.cache.find("player"),
		SCREEN_WIDTH/2, SCREEN_HEIGHT/2,
	)
}

// Get screen dims
func (g *Game) Layout(outsideWidth, insideWidth int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// Update logic of game
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Draw to game screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	g.player.Draw(screen)
}

// ============================================================================
