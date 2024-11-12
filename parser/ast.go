package parser

import (
	"bytes"
	"fmt"
	"strings"
	"wavy/lexer"
)

type Node interface {
	TokenValue() string
	String() string
	Tree(indentation string, base bool) string
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
func (program *Program) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	out.WriteString(indentation + "PROGRAM\n")
	for _, s := range program.Statements {
		out.WriteString(s.Tree(indentation, true))
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
func (i *Identifier) Tree(indentation string, base bool) string {
	if base {
		return fmt.Sprintf("%sIDENTIFIER (%s)\n", indentation, i.Value)
	}
	return fmt.Sprintf("IDENTIFIER (%s)\n", i.Value)
}

type AssignmentStatement struct {
	Token    lexer.Token
	Name     *Identifier
	Operator string
	Value    Expression
}

func (as *AssignmentStatement) expressionNode()    {}
func (as *AssignmentStatement) statementNode()     {}
func (as *AssignmentStatement) TokenValue() string { return as.Token.Value }

func (as *AssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(as.Name.String())
	out.WriteString(" " + as.Operator + " ")
	out.WriteString(as.Value.String())
	out.WriteString(";")
	return out.String()
}
func (as *AssignmentStatement) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "ASSIGNMENT\n")
	} else {
		out.WriteString("ASSIGNMENT\n")
	}
	out.WriteString(indentation + "├── IDENTIFIER (" + as.Name.Value + ")\n")
	out.WriteString(indentation + "├── " + as.Operator + "\n")
	out.WriteString(indentation + "└── " + as.Value.Tree(indentation+"    ", false))
	return out.String()
}

type LoadStatement struct {
	Token lexer.Token
	Value Expression
}

func (ls *LoadStatement) statementNode()     {}
func (ls *LoadStatement) TokenValue() string { return ls.Token.Value }
func (ls *LoadStatement) String() string {
	return fmt.Sprintf("load('%s');", ls.Value)
}
func (ls *LoadStatement) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "LOAD\n")
	} else {
		out.WriteString("LOAD\n")
	}
	out.WriteString(indentation + "  └── " + ls.Value.Tree(indentation, false) + "\n")
	return out.String()
}

// export(audio1, 'looped_audio1.wav')
type ExportStatement struct {
	Token lexer.Token
	Name  string
	Value string
}

func (es *ExportStatement) expressionNode()    {}
func (es *ExportStatement) TokenValue() string { return es.Token.Value }
func (es *ExportStatement) String() string {
	return fmt.Sprintf("export %s, '%s';", es.Name, es.Value)
}
func (es *ExportStatement) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	out.WriteString(indentation + "EXPORT\n")
	out.WriteString(indentation + "  ├── " + es.Name + "\n")
	out.WriteString(indentation + "  └── " + es.Value + "\n")
	return out.String()
}

type ReturnStatement struct {
	Token       lexer.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()     {}
func (rs *ReturnStatement) TokenValue() string { return rs.Token.Value }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenValue() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "RETURN\n")
	} else {
		out.WriteString("RETURN\n")
	}
	if rs.ReturnValue != nil {
		out.WriteString(indentation + "└── " + rs.ReturnValue.Tree(indentation+"   ", false))
	}
	return out.String()
}

type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()     {}
func (es *ExpressionStatement) TokenValue() string { return es.Token.Value }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) Tree(indentation string, base bool) string {
	if es.Expression != nil {
		return es.Expression.Tree(indentation, false)
	}
	return indentation + "EXPRESSION_STATEMENT\n"
}

type IntegerValue struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerValue) expressionNode()    {}
func (il *IntegerValue) TokenValue() string { return il.Token.Value }
func (il *IntegerValue) String() string     { return il.Token.Value }
func (il *IntegerValue) Tree(indentation string, base bool) string {
	if base {
		return fmt.Sprintf("%sINTEGER (%d)\n", indentation, il.Value)
	}
	return fmt.Sprintf("INTEGER (%d)\n", il.Value)
}

type FloatValue struct {
	Token lexer.Token
	Value float64
}

func (fl *FloatValue) expressionNode()    {}
func (fl *FloatValue) TokenValue() string { return fl.Token.Value }
func (fl *FloatValue) String() string     { return fl.Token.Value }
func (fl *FloatValue) Tree(indentation string, base bool) string {
	if base {
		return fmt.Sprintf("%sFLOAT (%f)\n", indentation, fl.Value)
	}
	return fmt.Sprintf("FLOAT (%f)\n", fl.Value)
}

type PrefixExpression struct {
	Token    lexer.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()    {}
func (pe *PrefixExpression) TokenValue() string { return pe.Token.Value }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}
func (pe *PrefixExpression) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	out.WriteString(indentation + "PREFIX\n")
	out.WriteString(indentation + "  ├── " + pe.Operator + "\n")
	out.WriteString(indentation + "  └── " + pe.Right.Tree(indentation+"    ", false))
	return out.String()
}

type InfixExpression struct {
	Token    lexer.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()    {}
func (ie *InfixExpression) TokenValue() string { return ie.Token.Value }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
func (ie *InfixExpression) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "INFIX\n")
	} else {
		out.WriteString("INFIX\n")
	}
	out.WriteString(indentation + "  ├── " + ie.Left.Tree(indentation+"  │  ", false))
	out.WriteString(indentation + "  ├── " + ie.Operator + "\n")
	out.WriteString(indentation + "  └── " + ie.Right.Tree(indentation+"     ", false))
	return out.String()
}

type Boolean struct {
	Token lexer.Token
	Value bool
}

func (b *Boolean) expressionNode()    {}
func (b *Boolean) TokenValue() string { return b.Token.Value }
func (b *Boolean) String() string     { return b.Token.Value }
func (b *Boolean) Tree(indentation string, base bool) string {
	return fmt.Sprintf("%sBOOLEAN (%t)\n", indentation, b.Value)
}

type IfExpression struct {
	Token       lexer.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()    {}
func (ie *IfExpression) TokenValue() string { return ie.Token.Value }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
func (ie *IfExpression) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "IF_STATEMENT\n")
	} else {
		out.WriteString("IF_STATEMENT\n")
	}
	out.WriteString(indentation + "├── CONDITION\n")
	out.WriteString(indentation + ie.Condition.Tree("│   ", false))
	out.WriteString(indentation + "├── CONSEQUENCE\n")
	out.WriteString(indentation + "│   " + ie.Consequence.Tree(indentation+"│   ", false))
	if ie.Alternative != nil {
		out.WriteString(indentation + "└── ALTERNATIVE\n")
		out.WriteString(indentation + "    " + ie.Alternative.Tree(indentation+"      ", false))
	}
	return out.String()
}

type BlockStatement struct {
	Token      lexer.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()     {}
func (bs *BlockStatement) TokenValue() string { return bs.Token.Value }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
func (bs *BlockStatement) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "BLOCK\n")
	} else {
		out.WriteString("BLOCK\n")
	}
	for i, s := range bs.Statements {
		if i == len(bs.Statements)-1 {
			out.WriteString(indentation + "└── " + s.Tree(indentation+"    ", false))
		} else {
			out.WriteString(indentation + "├── " + s.Tree(indentation+"│   ", false))
		}
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      lexer.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()    {}
func (fl *FunctionLiteral) TokenValue() string { return fl.Token.Value }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenValue())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
func (fl *FunctionLiteral) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "FUNCTION\n")
	} else {
		out.WriteString("FUNCTION\n")
	}
	out.WriteString(indentation + "  ├── PARAMETERS\n")
	for i, param := range fl.Parameters {
		prefix := "  │   ├── "
		if i == len(fl.Parameters)-1 {
			prefix = "  │   └── "
		}
		out.WriteString(indentation + prefix + param.Tree(indentation+"      ", false))
	}
	out.WriteString(indentation + "  └── BODY\n")
	out.WriteString(indentation + "      " + fl.Body.Tree(indentation+"      ", false))
	return out.String()
}

type CallExpression struct {
	Token     lexer.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()    {}
func (ce *CallExpression) TokenValue() string { return ce.Token.Value }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (ce *CallExpression) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "CALL_EXPRESSION\n")
	} else {
		out.WriteString("CALL_EXPRESSION\n")
	}
	out.WriteString(indentation + "├── FUNCTION\n")
	out.WriteString(indentation + ce.Function.Tree("│   ", true))
	out.WriteString(indentation + "└── ARGUMENTS\n")
	for i, arg := range ce.Arguments {
		prefix := "    ├── "
		if i == len(ce.Arguments)-1 {
			prefix = "    └── "
		}
		out.WriteString(indentation + prefix + arg.Tree(indentation+"      ", false))
	}
	return out.String()
}

type ForLoopExpression struct {
	Token       lexer.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (fle *ForLoopExpression) expressionNode()    {}
func (fle *ForLoopExpression) TokenValue() string { return fle.Token.Value }
func (fle *ForLoopExpression) String() string {
	var out bytes.Buffer
	out.WriteString("for (")
	out.WriteString(fle.Condition.String())
	out.WriteString(" ) {")
	out.WriteString(fle.Consequence.String())
	out.WriteString("}")
	return out.String()
}
func (fle *ForLoopExpression) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	out.WriteString(indentation + "FOR_LOOP\n")
	out.WriteString(indentation + " ├── CONDITION\n")
	out.WriteString(indentation + " │   " + fle.Condition.Tree(indentation+" │   ", false))
	out.WriteString(indentation + " └── CONSEQUENCE\n")
	out.WriteString(indentation + "      " + fle.Consequence.Tree(indentation+"      ", false))
	return out.String()
}

type StringValue struct {
	Token lexer.Token
	Value string
}

func (sl *StringValue) expressionNode()    {}
func (sl *StringValue) TokenValue() string { return sl.Token.Value }
func (sl *StringValue) String() string     { return sl.Token.Value }
func (sl *StringValue) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "STRING\n")
	} else {
		out.WriteString("STRING\n")
	}
	out.WriteString(indentation + "  └── " + sl.Value + "\n")
	return out.String()
}

type ArrayLiteral struct {
	Token    lexer.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()    {}
func (al *ArrayLiteral) TokenValue() string { return al.Token.Value }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
func (al *ArrayLiteral) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "ARRAY_LITERAL\n")
	} else {
		out.WriteString("ARRAY_LITERAL\n")
	}
	for i, elem := range al.Elements {
		prefix := "  ├── "
		if i == len(al.Elements)-1 {
			prefix = "  └── "
		}
		out.WriteString(indentation + prefix + elem.Tree(indentation, false))
	}
	return out.String()
}

type IndexExpression struct {
	Token lexer.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()    {}
func (ie *IndexExpression) TokenValue() string { return ie.Token.Value }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
func (ie *IndexExpression) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	if base {
		out.WriteString(indentation + "INDEX_EXPRESSION\n")
	} else {
		out.WriteString("INDEX_EXPRESSION\n")
	}
	out.WriteString(indentation + "     ├── LEFT\n")
	out.WriteString(indentation + "     │   " + ie.Left.Tree(indentation+"  │   ", false))
	out.WriteString(indentation + "     └── INDEX\n")
	out.WriteString(indentation + "         " + ie.Index.Tree(indentation+"      ", false))
	return out.String()
}

type ForeachStatement struct {
	Token lexer.Token
	Index string
	Ident string
	Value Expression
	Body  *BlockStatement
}

func (fes *ForeachStatement) expressionNode() {}

func (fes *ForeachStatement) TokenValue() string { return fes.Token.Value }

func (fes *ForeachStatement) String() string {
	var out bytes.Buffer
	out.WriteString("foreach ")
	out.WriteString(fes.Ident)
	out.WriteString(" ")
	out.WriteString(fes.Value.String())
	out.WriteString(fes.Body.String())
	return out.String()
}
func (fes *ForeachStatement) Tree(indentation string, base bool) string {
	var out bytes.Buffer
	out.WriteString(indentation + "FOREACH_STATEMENT\n")
	out.WriteString(indentation + "  ├── IDENTIFIER (" + fes.Ident + ")\n")
	out.WriteString(indentation + "  ├── INDEX\n")
	out.WriteString(indentation + "  │   └── " + fes.Index + "\n")
	out.WriteString(indentation + "  ├── VALUE\n")
	out.WriteString(indentation + "  │   └── " + fes.Value.Tree(indentation+"  │   ", false))
	out.WriteString(indentation + "  └── BODY\n")
	out.WriteString(indentation + "      └── " + fes.Body.Tree(indentation+"        ", false))
	return out.String()
}
