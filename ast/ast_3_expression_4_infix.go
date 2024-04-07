package ast

import (
	"monkey-interpreter/token"
	"strings"
)

type InfixExpression struct {
	Token    token.Token // 运算符词法单元, 如 +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	var out strings.Builder
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Operator)
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
