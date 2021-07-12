package token

type (
	Type string

	Token struct {
		Type    Type
		Literal string
	}
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals

	IDENT = "IDENT" // add, foobar, x, z, ...
	INT   = "INT"   // 123456

	// Operators

	ASSIGN = "="
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
