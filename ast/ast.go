package ast

import (
	"strings"
)

type Node interface {
	TokenLiteral() string // 用来调试
	String() string       // 用来调试, 打印比较 AST 节点
}

// Statement 语句
type Statement interface {
	Node
	statementNode() // 占位, 编译检查
}

// Expression 表达式
type Expression interface {
	Node
	expressionNode() // 占位, 编译检查
}

// Program 程序
// 由一系列 Statement 语句组成
// Statement 和 Expression 的区别是
// Statement 语句不会产生值
// Expression 表达式会产生值
// 比如 let x = 5;
// let x = 5; 不会产生值
// 5 会产生值, 就是 5
// return 5; 不会产生值
// add(5, 5) 会产生值
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out strings.Builder // strings.Builder 性能比原文中 bytes.Buffer 略快约 10%
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
