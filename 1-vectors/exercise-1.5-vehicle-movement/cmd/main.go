package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Vehicle Movement")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()
	vehicle := NewVehicle(10, 200)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		vehicle.Update()
		vehicle.CheckEdges()
		vehicle.Show()
		rl.EndDrawing()
	}
}

type Vehicle struct {
	position     rl.Vector2
	velocity     rl.Vector2
	acceleration rl.Vector2
	maxSpeed     float32
}

func NewVehicle(x, y float32) *Vehicle {
	return &Vehicle{
		position:     rl.NewVector2(x, y),
		velocity:     rl.NewVector2(0, 0),
		acceleration: rl.NewVector2(0, 0),
		maxSpeed:     10,
	}
}

func Vector2Limit(v rl.Vector2, max float32) rl.Vector2 {
	if rl.Vector2Length(v) > max {
		return rl.Vector2Normalize(v)
	}
	return v
}

func (v *Vehicle) Update() {
	v.velocity = rl.Vector2Add(v.velocity, v.acceleration)
	v.velocity = Vector2Limit(v.velocity, v.maxSpeed)
	fmt.Println(v.velocity)
	v.position = rl.Vector2Add(v.position, v.velocity)

	if rl.IsKeyDown(rl.KeyRight) {
		v.acceleration = rl.Vector2Add(v.acceleration, rl.NewVector2(0.1, 0))
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		v.acceleration = rl.Vector2Zero()
	}
}

func (v *Vehicle) Show() {
	rl.DrawRectangleV(v.position, rl.NewVector2(20, 10), rl.Black)
}

func (v *Vehicle) CheckEdges() {
	screenW := float32(rl.GetScreenWidth())
	if v.position.X > screenW {
		v.position.X = 0
	} else if v.position.X < 0 {
		v.position.X = screenW
	}
}
