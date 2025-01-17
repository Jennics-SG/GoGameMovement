package game

import (
	app "GoGameMovement/Application"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := app.NewGame()

	ebiten.SetWindowSize(app.SCREEN_WIDTH, app.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
