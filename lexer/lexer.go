package lexer

import "monkey/token"

// NextToken examines the current character and returns the corresponding token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Switch on the current character to determine what kind of token to create
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// NUL byte indicates end of file
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// For any other character, check if it's a letter
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			// If not a letter, mark as illegal token
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads in an identifier (variable name, keyword, etc)
// Continues reading until it hits a non-letter character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks if the given character is a letter or underscore
// Used to identify valid identifier characters
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// newToken creates a new Token with the given type and character value
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

type Lexer struct {
	input        string // Holds the entire input string we need to tokenize
	position     int    // Points to current character (the one in ch)
	readPosition int    // Points to next character to read
	ch           byte   // Current character under examination (like a window showing one char at a time)
}

// New creates a Lexer instance with the given input string
// This is our constructor - sets up initial state
func New(input string) *Lexer {
	l := &Lexer{input: input} // Store input string and return a pointer to new Lexer
	// Initialize lexer by reading the first character
	l.readChar()
	return l
}

// readChar moves through the input string, reading one character at a time
// Updates ch to current character and advances our positions
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL signifies either EOF or haven't read anything yet
	} else {
		l.ch = l.input[l.readPosition] // Get character at readPosition and store in ch
	}

	l.position = l.readPosition // Move current to where we just read
	l.readPosition += 1         // Set up for next character
}
