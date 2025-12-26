package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // 입력에서 현재위치(현재문자)
	readPosition int  // 입력에서 현재 읽는 위치(현재 문자의 다음)
	ch           byte // 현재 조사하고 있는 문자
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.LBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL in ascii
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
