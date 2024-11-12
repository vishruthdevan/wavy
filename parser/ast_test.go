package parser

import (
	"testing"
	"wavy/lexer"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AssignmentStatement{
				Name: &Identifier{
					Token: lexer.Token{Type: lexer.IDENTIFIER, Value: "myVar"},
					Value: "myVar",
				},
				Operator: "=",
				Value: &Identifier{
					Token: lexer.Token{Type: lexer.IDENTIFIER, Value: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
