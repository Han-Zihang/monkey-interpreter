package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/token"
)

/*
	The Pratt parser
	普拉特语法分析器主要思想是
	将解析函数 (语义代码) 与词法单元类型关联

	每个词法单元最多关联 2 个解析函数
	这取决于词法单元的位置是前缀还是后缀
*/

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// Monkey 语言优先级
const (
	_ int = iota
	LOWEST
	EQUALS  // ==
	COMPARE // > or <
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // f(x)
)

func (p *Parser) noPrefixFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
