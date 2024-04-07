package parser

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}
