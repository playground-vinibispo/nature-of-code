package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var randomCounts []float32

func main() {
	total := 20
	rl.InitWindow(640, 240, "An accept-reject distribution")
	rl.SetTargetFPS(60)

	randomCounts = make([]float32, total)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		index := int(acceptReject() * float32(len(randomCounts)))
		randomCounts[index]++
		w := float32(rl.GetScreenWidth()) / float32(len(randomCounts))

		for i := 0; i < len(randomCounts); i++ {
			rl.DrawRectangleV(rl.NewVector2(float32(i)*w, float32(rl.GetScreenHeight())-randomCounts[i]), rl.NewVector2(w-1, randomCounts[i]), rl.NewColor(127, 127, 127, 255))
			rl.DrawRectangleLinesEx(rl.NewRectangle(float32(i)*w, float32(rl.GetScreenHeight())-randomCounts[i], w-1, randomCounts[i]), 2, rl.Black)
		}
		rl.EndDrawing()
	}
}

func acceptReject() float32 {
	for {
		r1 := rand.Float32()
		probability := r1
		r2 := rand.Float32()
		if r2 < probability {
			return r1
		}
	}
}
