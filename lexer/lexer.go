package lexer

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
	default:
	}
	lexer.advance()
	return t
}