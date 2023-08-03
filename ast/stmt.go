package ast

type Stmt interface{}

type AssignStmt struct {
	Ident *IdentExpr
	Expr  Expr
}
