package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func randomGaussian() float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	return math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
}

var (
	screenWidth  int32 = 640
	screenHeight int32 = 240
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Colored Dots")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()
	target := rl.LoadRenderTexture(640, 240)
	defer rl.UnloadRenderTexture(target)
	rl.BeginTextureMode(target)
	rl.ClearBackground(rl.RayWhite)
	rl.EndTextureMode()
	centerModeX := float32(screenWidth) / 2
	centerModeY := float32(screenHeight) / 2
	standardDeviation := float32(30)
	for !rl.WindowShouldClose() {
		rl.BeginTextureMode(target)
		for range 10 {
			x := centerModeX + float32(randomGaussian())*standardDeviation
			y := centerModeY + float32(randomGaussian())*standardDeviation
			r := rl.Remap(float32(randomGaussian()*50+150), 50, 150, 0, 255)
			g := rl.Remap(float32(randomGaussian()*50+150), 50, 150, 0, 255)
			b := rl.Remap(float32(randomGaussian()*50+150), 50, 150, 0, 255)
			rl.DrawCircleV(rl.NewVector2(x, y), 8, rl.NewColor(uint8(r), uint8(g), uint8(b), 255))
		}
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height)), rl.Vector2Zero(), rl.White)
		standardDeviation = DrawSlider(50, float32(screenHeight-50), 300, 1, 100, standardDeviation)
		rl.EndDrawing()
	}
}

func DrawSlider(x, y, width float32, minValue, maxValue, currentValue float32) float32 {
	rl.DrawRectangleV(rl.NewVector2(x, y), rl.NewVector2(width, 20), rl.LightGray)
	sliderPos := (currentValue - minValue) / (maxValue - minValue) * width
	rl.DrawRectangleV(rl.NewVector2(x+sliderPos-10, y-5), rl.NewVector2(20, 30), rl.DarkGray)

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mousePos := rl.GetMousePosition()
		if mousePos.X >= x && mousePos.X <= x+width {
			sliderPos = mousePos.X - x
			currentValue = sliderPos/width*(maxValue-minValue) + minValue
		}
	}
	return currentValue
}
