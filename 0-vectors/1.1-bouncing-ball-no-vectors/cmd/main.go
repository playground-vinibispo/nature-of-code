package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Bouncing Ball with No Vectors")
	rl.SetTargetFPS(60)
	xSpeed := 2.5
	ySpeed := 2.0
	x := 100.0
	y := 100.0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircle(int32(x), int32(y), 20, rl.Maroon)
		x += xSpeed
		y += ySpeed
		if x > 620 || x < 20 {
			xSpeed *= -1
		}
		if y > 220 || y < 20 {
			ySpeed *= -1
		}
		rl.EndDrawing()
	}
}
