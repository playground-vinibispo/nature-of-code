package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Normalizing a Vector")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		mouse := rl.GetMousePosition()
		center := rl.NewVector2(320, 120)
		sub := rl.Vector2Subtract(mouse, center)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.PushMatrix()
		rl.Translatef(center.X, center.Y, 0)
		rl.DrawLineEx(rl.Vector2Zero(), sub, 2, rl.DarkGray)
		rl.DrawLineEx(rl.Vector2Zero(), rl.Vector2Scale(rl.Vector2Normalize(sub), 50), 8, rl.Black)
		rl.PopMatrix()

		rl.EndDrawing()
	}
}
