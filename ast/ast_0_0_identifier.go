package ast

import "monkey-interpreter/token"

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// Identifier
// Monkey 程序中有些标识符会产生值
// 为了减少 AST 中各种类型节点的数量
// 在变量绑定中使用 Identifier 表示名称
// 之后在表达表达式中的标识符的时候
// 可以复用 Identifier 节点
func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }
