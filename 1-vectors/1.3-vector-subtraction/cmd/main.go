package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Vector Subtraction")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		mouse := rl.GetMousePosition()
		center := rl.NewVector2(320, 120)
		sub := rl.Vector2Subtract(mouse, center)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawLineEx(rl.Vector2Zero(), mouse, 2, rl.Red)
		rl.DrawLineEx(rl.Vector2Zero(), center, 2, rl.Blue)
		rl.PushMatrix()
		rl.Translatef(center.X, center.Y, 0)
		rl.DrawLineEx(rl.Vector2Zero(), sub, 2, rl.Green)
		rl.PopMatrix()
		rl.EndDrawing()
	}
}
