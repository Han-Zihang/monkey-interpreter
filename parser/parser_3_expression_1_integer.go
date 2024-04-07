package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"strconv"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	v, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("parse int %q err", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = v
	return lit
}
