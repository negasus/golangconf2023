package ast

type Expr interface{}

type IdentExpr struct {
	Value string
}

type NumberExpr struct {
	Value string
}

type ArithmeticOpExpr struct {
	Operator string
	Lhs      Expr
	Rhs      Expr
}
