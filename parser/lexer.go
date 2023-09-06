package parser

import (
	"fmt"
	"regexp"

	"github.com/negasus/golangconf2023/ast"
)

type token struct {
	tokenType int
	re        *regexp.Regexp
}

var (
	tokens = []token{
		{tokenType: LPAREN, re: regexp.MustCompile("^\\(")},
		{tokenType: RPAREN, re: regexp.MustCompile("^\\)")},
		{tokenType: EQ, re: regexp.MustCompile("^равно")},
		{tokenType: PRINT, re: regexp.MustCompile("^показать")},
		{tokenType: MUL, re: regexp.MustCompile("^умножить на")},
		{tokenType: PLUS, re: regexp.MustCompile("^плюс")},
		{tokenType: MINUS, re: regexp.MustCompile("^минус")},
		{tokenType: NUMBER, re: regexp.MustCompile("^[0-9]+")},
		{tokenType: IDENT, re: regexp.MustCompile("^[А-Я]+")},
		{tokenType: STRING, re: regexp.MustCompile("^'[^']+'")},
	}
)

type Lexer struct {
	src []byte
	idx int

	lastToken *ast.Token

	line   int
	column int

	errors []string

	stmts []ast.Stmt
}

func NewLexer(src []byte) *Lexer {
	lx := &Lexer{
		src:    src,
		column: 1,
		line:   1,
	}

	return lx
}

func (lx *Lexer) Statements() []ast.Stmt {
	return lx.stmts
}

func (lx *Lexer) Errors() []string {
	return lx.errors
}

func (lx *Lexer) Error(s string) {
	lx.errors = append(lx.errors, fmt.Sprintf("ошибка '%s' на позиции %d:%d около '%s'",
		s,
		lx.lastToken.Line,
		lx.lastToken.Column,
		lx.lastToken.Value,
	))
}

func (lx *Lexer) Lex(lval *YYSymType) int {
	lx.skipSpaces()

	if lx.idx >= len(lx.src) {
		return 0
	}

	for _, re := range tokens {
		res := re.re.FindSubmatch(lx.src[lx.idx:])
		if len(res) == 0 {
			continue
		}

		lval.token = &ast.Token{
			Type:   re.tokenType,
			Value:  string(res[0]),
			Line:   lx.line,
			Column: lx.column,
		}

		lx.lastToken = lval.token

		lx.idx += len(res[0])
		lx.column += len(res[0])

		return re.tokenType
	}

	return 0
}

func (lx *Lexer) skipSpaces() {
	for {
		if lx.idx >= len(lx.src) {
			return
		}

		ch := lx.src[lx.idx]

		if ch == ' ' {
			lx.idx++
			lx.column++
			continue
		}

		if ch == '\n' {
			lx.idx++
			lx.column = 1
			lx.line++
			continue
		}

		break
	}
}
