package ast

type Token struct {
	Type   int
	Value  string
	Line   int
	Column int
}
