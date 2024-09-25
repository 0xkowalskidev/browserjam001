package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render(node *Node) {
	switch node.Type {
	case DocumentNode:
		rl.ClearBackground(color.RGBA{0, 0, 50, 1})
	case ElementNode:

	case TextNode:
		rl.DrawText(node.Data, 0, 0, 20, rl.LightGray)
	}

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			Render(child)
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "BrowserJam 001")
	defer rl.CloseWindow()

	html, err := os.ReadFile("input.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	tokens := Tokenize(string(html))
	for _, token := range tokens {
		fmt.Printf("Type: %v, Data: %q, Attributes: %v\n", token.Type, token.Data, token.Attributes)
	}

	document := Parse(tokens)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		Render(document)

		rl.EndDrawing()
	}
}
