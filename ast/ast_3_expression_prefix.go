package ast

import "monkey-interpreter/token"

type PrefixExpression struct {
	Token    token.Token //	 前缀词法单元, 比如 !
	Operator string
	Right    Expression
}
