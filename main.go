package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

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

	PrintNodeTree(document, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		Render(document)

		rl.EndDrawing()
	}
}

func PrintNodeTree(node *Node, depth int) {
	indent := strings.Repeat("  ", depth)

	if node.Type == ElementNode {
		fmt.Printf("%s<Element> %s - <Parent> %s\n", indent, node.TagName, node.Parent.TagName)
	} else if node.Type == TextNode {
		fmt.Printf("%s<Text> %s\n", indent, node.Data)
	}

	for _, child := range node.Children {
		PrintNodeTree(child, depth+1)
	}
}
