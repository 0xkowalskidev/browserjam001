package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "BrowserJam 001")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(color.RGBA{0, 0, 0, 1})
		rl.DrawText("Hello World", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
