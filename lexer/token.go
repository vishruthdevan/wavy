package lexer

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	KEYWORD     = "KEYWORD"
	IDENTIFIER = "IDENTIFIER"
	INTEGER     = "INTEGER"
	FLOAT       = "FLOAT"
	STRING      = "STRING"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	LT       = "<"
	GT       = ">"

	EQUALS     = "=="
	NOT_EQUALS = "!="

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPR      = "("
	RPR      = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL     = "NULL"
	FOR      = "FOR"
	IN       = "IN"
	LOAD     = "LOAD"
	EXPORT   = "EXPORT"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"return":   RETURN,
	"if":       IF,
	"else":     ELSE,
	"true":     TRUE,
	"false":    FALSE,
	"null":     NULL,
	"for":      FOR,
	"in":       IN,
	"load":     LOAD,
	"export":   EXPORT,
}
