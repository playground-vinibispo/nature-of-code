package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Bouncing Ball with No Vectors")
	rl.SetTargetFPS(60)
	position := rl.NewVector2(100, 100)
	velocity := rl.NewVector2(2.5, 2)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		position = rl.Vector2Add(position, velocity)
		if position.X > 620 || position.X < 20 {
			velocity.X = velocity.X * -1
		}
		if position.Y > 220 || position.Y < 20 {
			velocity.Y = velocity.Y * -1
		}
		rl.DrawCircleV(position, 20, rl.Maroon)
		rl.EndDrawing()
	}
}
