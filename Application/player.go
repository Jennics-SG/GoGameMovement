// Player logic n that

package app

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quartercastle/vector"
)

// Get dimensions of sprite
func GetSpriteDims(sprite ebiten.Image) (w, h int) {
	width := sprite.Bounds().Dx()
	height := sprite.Bounds().Dy()
	return width, height
}

// Degree to radian
func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// Player Class ===============================================================
type Player struct {
	sprite       ebiten.Image         // Player sprite
	pos          vector.MutableVector // Player position
	dims         vector.Vector        // Player dimensions
	rotSpeed     int                  // Speed of rotation in degrees
	acceleration float64              // Player acceleration
	maxSpeed     float64              // Maximum speed of player
	rotation     float64              // Current rotation of player in radians
	velocity     float64              // Current speed of player
	scale        float64              // Scale of player
}

// Constructor
func NewPlayer(sprite ebiten.Image, x, y float64) *Player {
	w, h := GetSpriteDims(sprite)

	return &Player{
		sprite:       sprite,
		pos:          vector.MutableVector{x, y},
		dims:         vector.Vector{float64(w), float64(h)},
		rotSpeed:     5,
		maxSpeed:     15,
		acceleration: 1,
		rotation:     0,
		velocity:     0,
		scale:        0.5,
	}
}

// Draw the player to the screen
func (p *Player) Draw(screen *ebiten.Image) {
	hw, hh := p.OffsetToCenter()

	// Get center of sprite
	xPos := p.pos.X() - hw
	yPos := p.pos.Y() - hh

	op := &ebiten.DrawImageOptions{}

	// Rotate player
	p.Rotate(op)

	op.GeoM.Scale(p.scale, p.scale)

	// Now the rotation is done we can move the player to where we want it to be
	op.GeoM.Translate(xPos, yPos)

	screen.DrawImage(&p.sprite, op)
}

func (p *Player) Update() {
	// Roatation logic
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= degToRad(float64(p.rotSpeed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += degToRad(float64(p.rotSpeed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.velocity <= p.maxSpeed {
		p.velocity += p.acceleration
	} else if p.velocity != 0 {
		p.velocity -= p.acceleration / 10
	}
	if p.velocity <= 0 {
		p.velocity = 0
	}

	// Move player in direction being faced
	deltaX := float64(p.velocity) * math.Sin(p.rotation)
	deltaY := float64(p.velocity) * math.Cos(p.rotation)

	// Make player teleport to other side of screen if out of bounds
	newX := p.pos.X() + deltaX
	newY := p.pos.Y() - deltaY

	if newX >= SCREEN_WIDTH+p.dims.X() {
		newX = 0
	} else if newX <= 0 {
		newX = SCREEN_WIDTH + p.dims.X()
	}
	if newY >= SCREEN_HEIGHT+p.dims.Y() {
		newY = 0
	} else if newY <= 0 {
		newY = SCREEN_HEIGHT
	}

	p.pos = vector.MutableVector{
		newX,
		newY,
	}

	log.Printf("CURRENT SPEED: %v\n", p.velocity)
}

func (p *Player) OffsetToCenter() (x, y float64) {
	// Get half width / height
	hw := p.dims.X() / 2
	hh := p.dims.Y() / 2
	return hw, hh
}

func (p *Player) Rotate(op *ebiten.DrawImageOptions) {
	// Get half width / height
	hw, hh := p.OffsetToCenter()

	// Move player sprite back by half w/h for pivot
	op.GeoM.Translate(-hw, -hh)

	// Now you can rotate as the pivot is in the center
	op.GeoM.Rotate(p.rotation)

	// And move it back
	op.GeoM.Translate(hw, hh)
}

// ============================================================================
