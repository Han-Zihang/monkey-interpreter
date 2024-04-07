package ast

import "monkey-interpreter/token"

type ReturnStatement struct {
	Token       token.Token // Token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
