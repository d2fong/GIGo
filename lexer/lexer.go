package lexer

import "github.com/d2fong/GIGO/token"

// Lexer reads input
type Lexer struct {
	input        string
	currPosition int
	nextPosition int
	currChar     byte
}

// NewLexer creates a new lexer
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.nextPosition >= len(lexer.input) {
		lexer.currChar = 0
	} else {
		lexer.currChar = lexer.input[lexer.nextPosition]
	}
	lexer.currPosition = lexer.nextPosition
	lexer.nextPosition++
}

func (lexer *Lexer) nextToken() token.Token {
	lexer.skipWhitespace()

	var t token.Token

	switch lexer.currChar {
	case '=':
		t = newToken(token.ASSIGN, lexer.currChar)
	case '+':
		t = newToken(token.PLUS, lexer.currChar)
	case ';':
		t = newToken(token.SEMICOLON, lexer.currChar)
	case '{':
		t = newToken(token.LBRACE, lexer.currChar)
	case '}':
		t = newToken(token.RBRACE, lexer.currChar)
	case '(':
		t = newToken(token.LPAREN, lexer.currChar)
	case ')':
		t = newToken(token.RPAREN, lexer.currChar)
	case ',':
		t = newToken(token.COMMA, lexer.currChar)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(lexer.currChar) {
			t.Literal = lexer.readIdentifier()
			return t
		} else if isDigit(lexer.currChar) {
			t.Type = token.INT
			t.Literal = lexer.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, lexer.currChar)
		}
	}

	lexer.readChar()
	return t
}

func (lexer *Lexer) readIdentifier() string {
	currPosition := lexer.currPosition
	for isLetter(lexer.currChar) {
		lexer.readChar()
	}

	return lexer.input[currPosition:lexer.currPosition]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currChar == ' ' || lexer.currChar == '\t' || lexer.currChar == '\n' || lexer.currChar == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readNumber() string {
	currPosition := lexer.currPosition
	for isDigit(lexer.currChar) {
		lexer.readChar()
	}

	return lexer.input[currPosition:lexer.currPosition]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
