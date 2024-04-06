package lexer

import "monkey-interpreter/token"

// Lexer
// 目前只支持 Unicode 字符
type Lexer struct {
	input        string
	position     int  // 当前字符的位置
	nextPosition int  // 即将读取的字符的位置 (这门语言分析时只需要向后看一个字符)
	ch           byte // 当前字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token
	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, '=')
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	}
	l.readChar()
	return tk
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
