package lexer

type Lexer struct {
	input        string
	position     int
	nextPosition int
	current      rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) advance() {
	if l.nextPosition < len(l.input) {
		l.current = rune(l.input[l.nextPosition])
	} else {
		l.current = 0
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

