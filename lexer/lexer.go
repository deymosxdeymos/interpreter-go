package lexer

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
