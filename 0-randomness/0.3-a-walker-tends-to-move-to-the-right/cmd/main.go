package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "A traditional Random Walk")
	rl.SetTargetFPS(60)

	walker := NewWalker()

	target := rl.LoadRenderTexture(640, 240)
	rl.BeginTextureMode(target)
	rl.ClearBackground(rl.RayWhite)
	rl.EndTextureMode()
	for !rl.WindowShouldClose() {
		walker.Step()
		rl.BeginTextureMode(target)
		walker.Show()
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height)), rl.Vector2Zero(), rl.White)
		rl.EndDrawing()
	}
	rl.UnloadRenderTexture(target)

	rl.CloseWindow()
}

type Walker struct {
	X, Y float32
}

func NewWalker() *Walker {
	return &Walker{X: 320, Y: 120}
}

func (w *Walker) Show() {
	rl.DrawCircleV(rl.NewVector2(w.X, w.Y), 1, rl.Black)
}

func (w *Walker) Step() {
	offset := rand.Float32()
	if offset < 0.4 {
		w.X++
		return
	}
	if offset < 0.6 {
		w.X--
		return
	}
	if offset < 0.8 {
		w.Y++
		return
	}
	w.Y--

}
