package lexer

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	KEYWORD = "KEYWORD"
	INDENTIFIER = "INDENTIFIER"
	INTEGER     = "INTEGER"
	FLOAT       = "FLOAT"
	STRING      = "STRING"
	
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	
	COMMA     = ","
	SEMICOLON = ";"
	
	LPR      = "("
	RPR      = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)
