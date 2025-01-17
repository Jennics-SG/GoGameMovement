// Initialise app and add it to ebiten window

package main

import (
	app "GoGameMovement/Application"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Create game and initialise it
	game := app.NewGame()
	game.Init()

	// Ebiten window
	ebiten.SetWindowSize(app.SCREEN_WIDTH, app.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
