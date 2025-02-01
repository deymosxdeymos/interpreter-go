package token

// We are using string for more easy time
// when debugging (e.g when you use int, it will output it as numbers
// instead of the token types values)
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// We are using const because for this Monkey language,
// we have a predefined and limited amount of token types
// so this can be defined with const
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, etc...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGn = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
