package parser

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/token"
)

func (p *Parser) ParseProgram() *ast.Program {
	/*
		// 伪代码逻辑
		program = newProgramASTNode()

		advanceTokens()

		for currentToken() != EOF_TOKEN {
			statement = null

			if currentToken() == LET_TOKEN {
				statement = parseLetStatement()
			} else if currentToken() == RETURN_TOKEN {
				statement = parseReturnStatement()
			} else if currentToken() == IF_TOKEN {
				statement = parseIFStatement()
			}

			if statement != null {
				program.Statements.push({statement})
			}

			advanceTokens()
		}

		return program

	*/
	program := &ast.Program{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}
