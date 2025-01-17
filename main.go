package main

import (
	"GoGameMovement/MovementGame"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := MovementGame.NewGame()

	ebiten.SetWindowSize(MovementGame.SCREEN_WIDTH, MovementGame.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
