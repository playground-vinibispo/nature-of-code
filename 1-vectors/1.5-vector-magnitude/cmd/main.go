package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Vector Magnitude")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		mouse := rl.GetMousePosition()
		center := rl.NewVector2(320, 120)
		sub := rl.Vector2Subtract(mouse, center)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		mag := rl.Vector2Length(sub)
		rl.DrawRectangleV(rl.Vector2Zero(), rl.NewVector2(mag, 10), rl.Black)
		rl.PushMatrix()
		rl.Translatef(center.X, center.Y, 0)
		rl.DrawLineV(rl.Vector2Zero(), sub, rl.Red)
		rl.PopMatrix()

		rl.EndDrawing()
	}
}
