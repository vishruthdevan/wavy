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
	lexer.advance()
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

	for isWhitespace(lexer.current) {
		lexer.advance()
	}

	switch lexer.current {
	case '=':
		if lexer.peek() == '=' {
			current := lexer.current
			lexer.advance()
			t.Type = EQUALS
			t.Value = string(current) + string(lexer.current)
		} else {
			t = initToken(ASSIGN, lexer.current)
		}
	case '!':
		if lexer.peek() == '=' {
			current := lexer.current
			lexer.advance()
			t = initToken(NOT_EQUALS, current+lexer.current)
		} else {
			t = initToken(BANG, lexer.current)
		}
	case '+':
		t = initToken(PLUS, lexer.current)
	case '-':
		t = initToken(MINUS, lexer.current)
	case '*':
		t = initToken(ASTERISK, lexer.current)
	case '/':
		t = initToken(SLASH, lexer.current)
	case '<':
		t = initToken(LT, lexer.current)
	case '>':
		t = initToken(GT, lexer.current)
	case ',':
		t = initToken(COMMA, lexer.current)
	case ';':
		t = initToken(SEMICOLON, lexer.current)
	case ':':
		t = initToken(COLON, lexer.current)
	case '"':
		t.Type = STRING
		t.Value = lexer.readString()
	case '\'':
		t.Type = STRING
		t.Value = lexer.readString()
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
		t.Type = EOF
		t.Value = ""
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

func isWhitespace(ch rune) bool {
	if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		return true
	}
	return false
}

func (lexer *Lexer) readWord() string {
	start := lexer.position
	if isValidChar(lexer.current) {
		lexer.advance()
		for isValidChar(lexer.current) || isDigit(lexer.current) {
			lexer.advance()
		}
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

func (lexer *Lexer) readString() string {
	start := lexer.position
	if lexer.current == '"' {
		lexer.advance()
		for lexer.current != '"' {
			lexer.advance()
		}
	}
	if lexer.current == '\'' {
		lexer.advance()
		for lexer.current != '\'' {
			lexer.advance()
		}
	}
	return lexer.input[start+1 : lexer.position]
}

func lookupKeyword(word string) TokenType {
	t, exists := keywords[word]
	if exists {
		return t
	}
	return IDENTIFIER
}

func (lexer *Lexer) peek() rune {
	if lexer.nextPosition < len(lexer.input) {
		return rune(lexer.input[lexer.nextPosition])
	}
	return 0
}
