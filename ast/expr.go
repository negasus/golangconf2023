package ast

type Expr interface{}

type IdentExpr struct {
	Value string
}

type NumberExpr struct {
	Value string
}

type StringExpr struct {
	Value string
}

type BinaryExpr struct {
	Operator string
	Lhs      Expr
	Rhs      Expr
}
