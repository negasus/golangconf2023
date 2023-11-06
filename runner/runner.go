package runner

import (
	"fmt"
	"strconv"

	"github.com/negasus/golangconf2023/ast"
)

var (
	vars = map[string]int{}
)

func Run(stmts []ast.Stmt) error {
	for _, s := range stmts {
		switch ss := s.(type) {
		case *ast.PrintStmt:
			if err := printStatement(ss); err != nil {
				return err
			}
		case *ast.AssignStmt:
			if err := assignStatement(ss); err != nil {
				return err
			}
		default:
			return fmt.Errorf("неожиданная инструкция %T", s)
		}
	}

	return nil
}

func assignStatement(s *ast.AssignStmt) error {
	varName := s.Ident.Value
	value, err := calcExpr(s.Expr)
	if err != nil {
		return err
	}
	vars[varName] = value
	return nil
}

func printStatement(s *ast.PrintStmt) error {
	switch ss := s.Expr.(type) {
	case *ast.StringExpr:
		v := ss.Value[1 : len(ss.Value)-1] // trim quotes
		fmt.Printf("%s\n", v)

	default:
		v, err := calcExpr(s.Expr)
		if err != nil {
			return err
		}
		fmt.Printf("%d\n", v)
	}

	return nil
}

func calcExpr(e ast.Expr) (int, error) {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		return strconv.Atoi(ex.Value)

	case *ast.IdentExpr:
		v, ok := vars[ex.Value]
		if !ok {
			return 0, fmt.Errorf("переменная %s не найдена", ex.Value)
		}
		return v, nil

	case *ast.BinaryExpr:
		vLeft, errLeft := calcExpr(ex.Lhs)
		if errLeft != nil {
			return 0, errLeft
		}
		vRight, errRight := calcExpr(ex.Rhs)
		if errRight != nil {
			return 0, errRight
		}

		switch ex.Operator {
		case "+":
			return vLeft + vRight, nil
		case "-":
			return vLeft - vRight, nil
		case "*":
			return vLeft * vRight, nil
		}
	}

	return 0, fmt.Errorf("неожиданное выражение %#v", e)
}
