package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 240, "Perlin Noise Pixel")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()
	target := rl.LoadRenderTexture(640, 240)
	defer rl.UnloadRenderTexture(target)
	rl.BeginTextureMode(target)
	rl.ClearBackground(rl.RayWhite)
	rl.EndTextureMode()
	noiseDetail(8, 0.65)
	for !rl.WindowShouldClose() {
		rl.BeginTextureMode(target)
		xoff := float64(0)
		for x := 0; x < 640; x++ {
			yoff := float64(0)
			for y := 0; y < 240; y++ {
				bright := uint8(rl.Remap(float32(noise(xoff, yoff)), 0, 1, 0, 255))
				rl.DrawPixel(int32(x), int32(y), rl.NewColor(bright, bright/2, 255-bright, 255))
				yoff += 0.01
			}
			xoff += 0.01
		}
		rl.EndTextureMode()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height)), rl.Vector2Zero(), rl.White)
		rl.EndDrawing()
	}
}

type Walker struct {
	x, y   float32
	tx, ty float32
}

func NewWalker() *Walker {
	return &Walker{
		x:  320,
		y:  120,
		ty: 10000,
	}
}

const (
	PERLIN_YWRAPB = 4
	PERLIN_YWRAP  = 1 << PERLIN_YWRAPB
	PERLIN_ZWRAPB = 8
	PERLIN_ZWRAP  = 1 << PERLIN_ZWRAPB
	PERLIN_SIZE   = 4095
)

var (
	perlinOctaves    = 4
	perlinAmpFalloff = 0.5
	perlin           []float64
)

func scaledCosine(i float64) float64 {
	return 0.5 * (1.0 - math.Cos(i*math.Pi))
}

func noise(x float64, y float64) float64 {
	z := float64(0)
	if perlin == nil {
		perlin = make([]float64, PERLIN_SIZE+1)
		for i := range perlin {
			perlin[i] = rand.Float64()
		}
	}
	if x < 0 {
		x = -x
	}
	xi := int(math.Floor(x))
	yi := int(math.Floor(y))
	zi := int(math.Floor(z))
	xf := x - float64(xi)
	yf := y - float64(yi)
	zf := z - float64(zi)
	var rxf, ryf float64

	r := 0.0
	ampl := 0.5

	n1, n2, n3 := 0.0, 0.0, 0.0

	for o := 0; o < perlinOctaves; o++ {
		of := xi + (yi << PERLIN_YWRAPB) + (zi << PERLIN_ZWRAPB)

		rxf = scaledCosine(xf)
		ryf = scaledCosine(yf)

		n1 = perlin[of&PERLIN_SIZE]
		n1 += rxf * (perlin[(of+1)&PERLIN_SIZE] - n1)
		n2 = perlin[(of+PERLIN_YWRAP)&PERLIN_SIZE]
		n2 += rxf * (perlin[(of+PERLIN_YWRAP+1)&PERLIN_SIZE] - n2)
		n1 += ryf * (n2 - n1)

		of += PERLIN_ZWRAP
		n2 = perlin[of&PERLIN_SIZE]
		n2 += rxf * (perlin[(of+1)&PERLIN_SIZE] - n2)
		n3 = perlin[(of+PERLIN_YWRAP)&PERLIN_SIZE]
		n3 += rxf * (perlin[(of+PERLIN_YWRAP+1)&PERLIN_SIZE] - n3)
		n2 += ryf * (n3 - n2)

		n1 += scaledCosine(zf) * (n2 - n1)
		r += n1 * ampl
		ampl *= perlinAmpFalloff
		xi <<= 1
		yi <<= 1
		xf *= 2
		yf *= 2
		zi <<= 1
		zf *= 2

		if xf >= 1.0 {
			xi++
			xf--
		}

		if yf >= 1.0 {
			yi++
			yf--
		}

		if zf >= 1.0 {
			zi++
			zf--
		}
	}
	return r
}

func noiseDetail(lod int, falloff float64) {
	if lod > 0 {
		perlinOctaves = lod
	}
	if falloff > 0 {
		perlinAmpFalloff = falloff
	}
}
