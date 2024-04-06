package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken2(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`
	tests := []*expectedToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
	}
	checkExpectedToken(t, input, tests)
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []*expectedToken{
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
	checkExpectedToken(t, input, tests)
}

func checkExpectedToken(t *testing.T, input string, tests []*expectedToken) {
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

type expectedToken struct {
	expectedType    token.Type
	expectedLiteral string
}
