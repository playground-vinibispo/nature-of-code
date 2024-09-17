package main

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Accelerating Toward the Mouse")
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
	position     rl.Vector2
	velocity     rl.Vector2
	acceleration rl.Vector2
	topSpeed     float32
}

func random(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

func NewMover() *Mover {
	return &Mover{
		position:     rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		velocity:     rl.Vector2Zero(),
		acceleration: rl.Vector2Zero(),
		topSpeed:     5,
	}
}

func Vector2Limit(v rl.Vector2, limit float32) rl.Vector2 {
	if rl.Vector2Length(v) > limit {
		v = rl.Vector2Normalize(v)
		v = rl.Vector2Scale(v, limit)
	}
	return v
}

func Vector2Random() rl.Vector2 {
	angle := rand.Float64() * 2 * math.Pi
	return rl.NewVector2(float32(math.Cos(angle)), float32(math.Sin(angle)))
}

func (m *Mover) Update() {
	width := float32(rl.GetScreenWidth())
	mouse := rl.GetMousePosition()
	dir := rl.Vector2Subtract(mouse, m.position)
	distance := rl.Vector2Length(dir)
	dir = rl.Vector2Normalize(dir)
	if distance == 0 {
		distance = 0.0001
	}
	strength := 1/distance + distance/width
	dir = rl.Vector2Scale(dir, strength)
	fmt.Println(dir)
	m.acceleration = dir
	m.velocity = rl.Vector2Add(m.velocity, m.acceleration)
	m.velocity = Vector2Limit(m.velocity, m.topSpeed)
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
