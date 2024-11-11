package parser

import (
	"wavy/lexer"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
	errors       []string
}

func Init(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}
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
	case lexer.IDENTIFIER:
		return p.parseAssignmentStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssignmentStatement() *AssignmentStatement {
	stmt := &AssignmentStatement{}

	stmt.Name = &Identifier{Token: p.currentToken, Value: p.currentToken.Value}

	if !p.peek(lexer.ASSIGN) {
		return nil
	}

	for !p.isCurrentToken(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{Token: p.currentToken}

	p.nextToken()

	for !p.isCurrentToken(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
