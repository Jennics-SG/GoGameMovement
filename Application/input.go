package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Control struct {
	key     ebiten.Key
	pressed bool
}

type Input struct {
	controls []Control
}

func NewInput() *Input {
	return &Input{}
}

var (
	controls = [...]Control{
		{ebiten.KeyW, false},
		{ebiten.KeyA, false},
		{ebiten.KeyS, false},
		{ebiten.KeyD, false},
	}
)

func (i *Input) Dir() {
	for i := range controls {
		controls[i].pressed = inpututil.IsKeyJustPressed(controls[i].key)
	}
}
