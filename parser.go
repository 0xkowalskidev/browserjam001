package main

type NodeType int

const (
	DocumentNode NodeType = iota
	ElementNode
	TextNode
)

type Node struct {
	Type       NodeType
	TagName    string
	Attributes []Attribute
	Data       string
	Children   []*Node
}

func Parse(tokens []Token) *Node {
	var root *Node = &Node{Type: DocumentNode}
	var stack []*Node
	currentNode := root

	for _, token := range tokens {
		switch token.Type {
		case TokenStartTag:
			node := &Node{
				Type:       ElementNode,
				TagName:    token.Data,
				Attributes: token.Attributes,
			}
			currentNode.Children = append(currentNode.Children, node)
			if token.Type != TokenSelfClosingEndTag {
				stack = append(stack, currentNode)
				currentNode = node
			}
		case TokenEndTag:
			if len(stack) > 0 {
				currentNode = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case TokenText:
			node := &Node{
				Type: TextNode,
				Data: token.Data,
			}
			currentNode.Children = append(currentNode.Children, node)
		}
	}

	return root
}
