package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(640, 480, "3D Bouncing Sphere")
	rl.SetTargetFPS(60)

	// Create 3D vectors for position and speed
	position := rl.NewVector3(0, 0, 0)
	speed := rl.NewVector3(2.5, 2, 0)

	// Camera setup
	camera := rl.Camera{
		Position:   rl.NewVector3(0, 10, 20),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		// Move the sphere
		position = rl.Vector3Add(position, speed)

		// Check for boundaries in 3D and reverse speed if necessary
		if position.X > 10 || position.X < -10 {
			speed.X *= -1
		}
		if position.Y > 10 || position.Y < -10 {
			speed.Y *= -1
		}
		if position.Z > 10 || position.Z < -10 {
			speed.Z *= -1
		}

		// Drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)
		rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 20, 20, 20, rl.LightGray) // Draw a box for reference
		rl.DrawSphere(position, 1, rl.Maroon)                              // Draw the sphere
		rl.EndMode3D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
