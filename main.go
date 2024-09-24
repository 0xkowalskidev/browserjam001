package main

import (
	"image/color"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "BrowserJam 001")
	defer rl.CloseWindow()

	html, err := os.ReadFile("input.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{0, 0, 50, 1})

		rl.DrawText(string(html), 0, 0, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
