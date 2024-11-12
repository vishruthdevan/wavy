package parser

import (
	"fmt"
	"strconv"
	"wavy/lexer"
)

const (
	_ int = iota
	LOWEST
	ASSIGN
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[lexer.TokenType]int{
	lexer.ASSIGN:     ASSIGN,
	lexer.EQUALS:     EQUALS,
	lexer.NOT_EQUALS: EQUALS,
	lexer.LT:         LESSGREATER,
	lexer.GT:         LESSGREATER,
	lexer.PLUS:       SUM,
	lexer.MINUS:      SUM,
	lexer.SLASH:      PRODUCT,
	lexer.ASTERISK:   PRODUCT,
	lexer.LPR:        CALL,
}

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

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.currentToken, Value: p.currentToken.Value}
}

func Init(l *lexer.Lexer) *Parser {
	parser := &Parser{lexer: l, errors: []string{}}

	parser.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)
	parser.registerPrefix(lexer.IDENTIFIER, parser.parseIdentifier)
	parser.registerPrefix(lexer.INTEGER, parser.parseIntegerValue)
	parser.registerPrefix(lexer.BANG, parser.parsePrefixExpression)
	parser.registerPrefix(lexer.MINUS, parser.parsePrefixExpression)
	parser.registerPrefix(lexer.TRUE, parser.parseBoolean)
	parser.registerPrefix(lexer.FALSE, parser.parseBoolean)
	parser.registerPrefix(lexer.LPR, parser.parseGroupedExpression)
	parser.registerPrefix(lexer.IF, parser.parseIfExpression)
	parser.registerPrefix(lexer.FUNCTION, parser.parseFunctionLiteral)
	parser.registerPrefix(lexer.FOR, parser.parseForLoopExpression)
	parser.registerPrefix(lexer.FLOAT, parser.parseFloatValue)

	parser.infixParseFns = make(map[lexer.TokenType]infixParseFn)
	parser.registerInfix(lexer.ASSIGN, parser.parseAssignExpression)
	parser.registerInfix(lexer.PLUS, parser.parseInfixExpression)
	parser.registerInfix(lexer.MINUS, parser.parseInfixExpression)
	parser.registerInfix(lexer.SLASH, parser.parseInfixExpression)
	parser.registerInfix(lexer.ASTERISK, parser.parseInfixExpression)
	parser.registerInfix(lexer.EQUALS, parser.parseInfixExpression)
	parser.registerInfix(lexer.NOT_EQUALS, parser.parseInfixExpression)
	parser.registerInfix(lexer.LT, parser.parseInfixExpression)
	parser.registerInfix(lexer.GT, parser.parseInfixExpression)
	parser.registerInfix(lexer.LPR, parser.parseCallExpression)

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

	stmt.ReturnValue = p.parseExpression(LOWEST)

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
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.isPeekToken(lexer.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIntegerValue() Expression {
	lit := &IntegerValue{Token: p.currentToken}
	value, err := strconv.ParseInt(p.currentToken.Value, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseFloatValue() Expression {
	float := &FloatValue{Token: p.currentToken}
	value, err := strconv.ParseFloat(p.currentToken.Value, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.currentToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}
	float.Value = value
	return float
}

func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Value,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Value,
		Left:     left,
	}
	precedence := p.currentPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseBoolean() Expression {
	return &Boolean{Token: p.currentToken, Value: p.isCurrentToken(lexer.TRUE)}
}

func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.peek(lexer.RPR) {
		return nil
	}
	return exp
}

func (p *Parser) parseIfExpression() Expression {
	expression := &IfExpression{Token: p.currentToken}

	if !p.peek(lexer.LPR) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.peek(lexer.RPR) {
		return nil
	}

	if !p.peek(lexer.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.isPeekToken(lexer.ELSE) {
		p.nextToken()
		if !p.peek(lexer.LBRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{Token: p.currentToken}
	block.Statements = []Statement{}
	p.nextToken()
	for !p.isCurrentToken(lexer.RBRACE) && !p.isCurrentToken(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseFunctionLiteral() Expression {
	literal := &FunctionLiteral{Token: p.currentToken}

	if !p.peek(lexer.LPR) {
		return nil
	}

	literal.Parameters = p.parseFunctionParameters()

	if !p.peek(lexer.LBRACE) {
		return nil
	}

	literal.Body = p.parseBlockStatement()

	return literal
}

func (p *Parser) parseFunctionParameters() []*Identifier {
	identifiers := []*Identifier{}

	if p.isPeekToken(lexer.RPR) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &Identifier{Token: p.currentToken, Value: p.currentToken.Value}
	identifiers = append(identifiers, ident)

	for p.isPeekToken(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &Identifier{Token: p.currentToken, Value: p.currentToken.Value}
		identifiers = append(identifiers, ident)
	}

	if !p.peek(lexer.RPR) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function Expression) Expression {
	exp := &CallExpression{Token: p.currentToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []Expression {
	args := []Expression{}

	if p.isPeekToken(lexer.RPR) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.isPeekToken(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.peek(lexer.RPR) {
		return nil
	}

	return args
}

func (p *Parser) parseForLoopExpression() Expression {
	expression := &ForLoopExpression{Token: p.currentToken}
	if !p.peek(lexer.LPR) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.peek(lexer.RPR) {
		return nil
	}
	if !p.peek(lexer.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()
	return expression
}

func (p *Parser) parseAssignExpression(name Expression) Expression {
	stmt := &AssignmentStatement{Token: p.currentToken}
	if n, ok := name.(*Identifier); ok {
		stmt.Name = n
	} else {
		msg := "expected assign token to be IDENT, got null instead"

		if name != nil {
			msg = fmt.Sprintf("expected assign token to be IDENT, got %s instead", name.TokenValue())
		}
		p.errors = append(p.errors, msg)
	}

	p.nextToken()
	stmt.Operator = "="

	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}
