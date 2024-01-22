package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literal
	IDENT = "IDENT" // add, foobar, x, y ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ";"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyword
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
