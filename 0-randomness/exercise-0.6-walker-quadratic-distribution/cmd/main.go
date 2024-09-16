package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Walker Quadratic Distribution")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()
	target := rl.LoadRenderTexture(640, 240)
	defer rl.UnloadRenderTexture(target)
	rl.BeginTextureMode(target)
	rl.ClearBackground(rl.RayWhite)
	rl.EndTextureMode()
	walker := NewWalker()
	for !rl.WindowShouldClose() {
		walker.Step()
		rl.BeginTextureMode(target)
		walker.Draw()
		rl.EndTextureMode()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height)), rl.Vector2Zero(), rl.White)
		rl.EndDrawing()
	}
}

func randomQuadratic() float64 {
	step := rand.Float64()*2 - 1

	quadraticStep := math.Pow(step, 2) * math.Copysign(1, step)
	return quadraticStep
}

type Walker struct {
	x, y float32
}

func NewWalker() *Walker {
	return &Walker{
		x: 320,
		y: 120,
	}
}

func (w *Walker) Step() {
	w.x += float32(randomQuadratic()) * 5
	w.y += float32(randomQuadratic()) * 5

	if w.x < 0 {
		w.x = 0
	}
	if w.x > 640 {
		w.x = 640
	}
}

func (w *Walker) Draw() {
	rl.DrawCircleV(rl.NewVector2(w.x, w.y), 5, rl.Black)
}
