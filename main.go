package main

import (
	"image/color"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render(node *Node) {
	switch node.Type {
	case DocumentNode:
		rl.ClearBackground(color.RGBA{25, 25, 25, 1})

	case ElementNode:
		switch node.TagName {
		case "h1":
			text := getTextContent(node)
			rl.DrawText(text, 0, 0, 40, rl.Red)
		}
	case TextNode:
		rl.DrawText(node.Data, 0, 0, 20, rl.White)
	}

	for _, child := range node.Children {
		Render(child)
	}
}

func getTextContent(node *Node) string {
	var text string
	if node.Type == TextNode {
		text += node.Data
	}
	for _, child := range node.Children {
		text += getTextContent(child)
	}
	return text
}

func main() {
	rl.InitWindow(800, 450, "BrowserJam 001")
	defer rl.CloseWindow()

	html, err := os.ReadFile("input.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	tokens := Tokenize(string(html))

	document := Parse(tokens)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		Render(document)

		rl.EndDrawing()
	}
}
