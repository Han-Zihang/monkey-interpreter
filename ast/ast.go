package ast

import "monkey-interpreter/token"

type Node interface {
	TokenLiteral() string // 用来调试
}

type Statement interface {
	Node
	statementNode() // 占位, 编译检查
}

type Expression interface {
	Node
	expressionNode() // 占位, 编译检查
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
