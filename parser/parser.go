package parser

import (
	"fmt"
	"wavy/lexer"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	lexer          *lexer.Lexer
	currentToken   lexer.Token
	peekToken      lexer.Token
	errors         []string
	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn
}

type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (p *Parser) isCurrentToken(t lexer.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) isPeekToken(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peek(t lexer.TokenType) bool {
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.currentToken, Value: p.currentToken.Value}
}

func Init(l *lexer.Lexer) *Parser {
	parser := &Parser{lexer: l, errors: []string{}}

	parser.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)
	parser.registerPrefix(lexer.IDENTIFIER, parser.parseIdentifier)

	parser.nextToken()
	parser.nextToken()
	return parser
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}
	for p.currentToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	// case lexer.IDENTIFIER:
	// 	return p.parseAssignmentStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// func (p *Parser) parseAssignmentStatement() *AssignmentStatement {
// 	stmt := &AssignmentStatement{}

// 	stmt.Name = &Identifier{Token: p.currentToken, Value: p.currentToken.Value}

// 	if !p.peek(lexer.ASSIGN) {
// 		return nil
// 	}

// 	for !p.isCurrentToken(lexer.SEMICOLON) {
// 		p.nextToken()
// 	}

// 	return stmt
// }

func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{Token: p.currentToken}

	p.nextToken()

	for !p.isCurrentToken(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{Token: p.currentToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.isPeekToken(lexer.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}
