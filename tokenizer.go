package main

type TokenType int

const (
	TokenText TokenType = iota
	TokenStartTag
	TokenEndTag
	TokenSelfClosingEndTag
)

type Token struct {
	Type       TokenType
	Data       string
	Attributes map[string]string
}

func Tokenize(html string) []Token {
	var tokens []Token
	length := len(html)
	i := 0

	for i < length {
		char := html[i]

		if char == '<' {
			// Check if it's an end tag
			if i+1 < length && html[i+1] == '/' {
				// End tag
				i += 2 // Skip '</'
				// Read the tag name
				tagNameStart := i
				for i < length && html[i] != '>' && !isWhitespace(html[i]) {
					i++
				}
				tagName := html[tagNameStart:i]
				// Skip until '>'
				for i < length && html[i] != '>' {
					i++
				}
				if i < length && html[i] == '>' {
					i++ // Skip '>'
				}
				tokens = append(tokens, Token{
					Type: TokenEndTag,
					Data: tagName,
				})
			} else {
				// Start tag or self-closing tag
				i++ // Skip '<'
				// Read the tag name
				tagNameStart := i
				for i < length && !isWhitespace(html[i]) && html[i] != '>' && html[i] != '/' {
					i++
				}
				tagName := html[tagNameStart:i]
				// Read attributes
				attrs := make(map[string]string)

				for i < length && html[i] != '>' && html[i] != '/' {
					// Skip whitespace
					for i < length && isWhitespace(html[i]) {
						i++
					}
					// Read attribute name
					attrNameStart := i
					for i < length && html[i] != '=' && !isWhitespace(html[i]) && html[i] != '>' && html[i] != '/' {
						i++
					}
					if attrNameStart == i {
						break // No attribute name found
					}
					attrName := html[attrNameStart:i]
					// Skip whitespace
					for i < length && isWhitespace(html[i]) {
						i++
					}
					// Skip '='
					if i < length && html[i] == '=' {
						i++
						// Skip whitespace
						for i < length && isWhitespace(html[i]) {
							i++
						}
						// Read attribute value
						attrValue := ""
						if i < length && (html[i] == '"' || html[i] == '\'') {
							quote := html[i]
							i++
							valueStart := i
							for i < length && html[i] != quote {
								i++
							}
							attrValue = html[valueStart:i]
							i++ // Skip closing quote
						} else {
							// Unquoted attribute value
							valueStart := i
							for i < length && !isWhitespace(html[i]) && html[i] != '>' && html[i] != '/' {
								i++
							}
							attrValue = html[valueStart:i]
						}
						attrs[attrName] = attrValue
					} else {
						// Attribute without value
						attrs[attrName] = ""
					}
				}
				// Check for self-closing tag
				selfClosing := false
				if i < length && html[i] == '/' {
					selfClosing = true
					i++ // Skip '/'
				}
				// Skip '>'
				if i < length && html[i] == '>' {
					i++
				}
				tokenType := TokenStartTag
				if selfClosing {
					tokenType = TokenSelfClosingEndTag
				}
				tokens = append(tokens, Token{
					Type:       tokenType,
					Data:       tagName,
					Attributes: attrs,
				})
			}
		} else {
			// Text node
			textStart := i
			for i < length && html[i] != '<' {
				i++
			}
			text := html[textStart:i]
			// Ignore empty or whitespace-only text nodes
			if len(text) > 0 {
				tokens = append(tokens, Token{
					Type: TokenText,
					Data: text,
				})
			}
		}
	}

	return tokens
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}
