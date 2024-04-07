package ast

import "monkey-interpreter/token"

type ExpressionStatement struct {
	Token      token.Token // 该表达式中第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
