package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tk := l.NextToken()
		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token_type wrong. expected:%q, got:%q", i, tt.expectedType, tk.Type)
		}
		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token_literal wrong. expected:%q, got:%q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}
