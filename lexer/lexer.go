package lexer

import "strings"

type Lexer struct {
	input        string
	position     int
	nextPosition int
	current      rune
}

func Init(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}

func (lexer *Lexer) advance() {
	if lexer.nextPosition < len(lexer.input) {
		lexer.current = rune(lexer.input[lexer.nextPosition])
	} else {
		lexer.current = 0
	}
	lexer.position = lexer.nextPosition
	lexer.nextPosition += 1
}

func initToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Value: string(ch)}
}

func (lexer *Lexer) NextToken() Token {
	var t Token

	switch lexer.current {
	case '=':
		t = initToken(ASSIGN, lexer.current)
	case '+':
		t = initToken(PLUS, lexer.current)
	case '-':
		t = initToken(MINUS, lexer.current)
	case '*':
		t = initToken(ASTERISK, lexer.current)
	case '/':
		t = initToken(SLASH, lexer.current)
	case ',':
		t = initToken(COMMA, lexer.current)
	case ';':
		t = initToken(SEMICOLON, lexer.current)
	case '(':
		t = initToken(LPR, lexer.current)
	case ')':
		t = initToken(RPR, lexer.current)
	case '{':
		t = initToken(LBRACE, lexer.current)
	case '}':
		t = initToken(RBRACE, lexer.current)
	case '[':
		t = initToken(LBRACKET, lexer.current)
	case ']':
		t = initToken(RBRACKET, lexer.current)
	case 0:
		t = initToken(EOF, lexer.current)
	default:
		if isValidChar(lexer.current) {
			t.Value = lexer.readWord()
			t.Type = lookupKeyword(t.Value)
			return t
		} else if isDigit(lexer.current) {
			t.Value = lexer.readNumber()
			if strings.Contains(t.Value, ".") {
				t.Type = FLOAT
			} else {
				t.Type = INTEGER
			}
			return t
		} else {
			t.Type = ILLEGAL
			t.Value = string(lexer.current)
		}
	}
	lexer.advance()
	return t
}

func isValidChar(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readWord() string {
	start := lexer.position
	for isValidChar(lexer.current) {
		lexer.advance()
	}
	return lexer.input[start:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	start := lexer.position
	for isDigit(lexer.current) {
		lexer.advance()
	}
	if lexer.current == '.' {
		lexer.advance()
		for isDigit(lexer.current) {
			lexer.advance()
		}
	}

	return lexer.input[start:lexer.position]
}

func lookupKeyword(word string) TokenType {
	t, exists := keywords[word]
	if exists {
		return t
	}
	return INDENTIFIER
}
