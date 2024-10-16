package lexer

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	KEYWORD    = "KEYWORD"
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"
	FLOAT      = "FLOAT"
	STRING     = "STRING"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"
	BANG     = "BANG"
	LT       = "LT"
	GT       = "GT"

	EQUALS     = "EQUALS"
	NOT_EQUALS = "NOT_EQUALS"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	COLON     = "COLON"

	LPR      = "LPR"
	RPR      = "RPR"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"

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
