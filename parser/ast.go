package parser

import (
	"bytes"
	"wavy/lexer"
)

type Node interface {
	TokenValue() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (program *Program) TokenValue() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenValue()
	} else {
		return ""
	}
}
func (program *Program) String() string {
	var out bytes.Buffer

	for _, s := range program.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Identifier struct {
	Token lexer.Token
	Value string
}

func (i *Identifier) expressionNode()    {}
func (i *Identifier) TokenValue() string { return i.Token.Value }
func (i *Identifier) String() string     { return i.Value }

type AssignmentStatement struct {
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode()     {}
func (as *AssignmentStatement) TokenValue() string { return "" }
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Name.String())
	out.WriteString(" = ")
	if as.Value != nil {
		out.WriteString(as.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
