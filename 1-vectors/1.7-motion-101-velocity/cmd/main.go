package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Motion 101 (Velocity)")
	mover := NewMover()

	rl.SetTargetFPS(60)

	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		mover.Update()
		mover.CheckEdges()
		mover.Show()

		rl.EndDrawing()
	}
}

type Mover struct {
	position rl.Vector2
	velocity rl.Vector2
}

func random(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

func NewMover() *Mover {
	return &Mover{
		position: rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		velocity: rl.NewVector2(random(-2, 2), random(-2, 2)),
	}
}

func (m *Mover) Update() {
	m.position = rl.Vector2Add(m.position, m.velocity)
}

func (m *Mover) Show() {
	rl.DrawCircleV(m.position, 48, rl.NewColor(175, 175, 177, 255))
	rl.DrawCircleLines(int32(m.position.X), int32(m.position.Y), 48, rl.NewColor(0, 0, 0, 255))
}

func (m *Mover) CheckEdges() {
	if m.position.X > float32(rl.GetScreenWidth()) {
		m.position.X = 0
	} else if m.position.X < 0 {
		m.position.X = float32(rl.GetScreenWidth())
	}

	if m.position.Y > float32(rl.GetScreenHeight()) {
		m.position.Y = 0
	} else if m.position.Y < 0 {
		m.position.Y = float32(rl.GetScreenHeight())
	}
}
