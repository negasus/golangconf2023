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
	var lastVarName string

	for _, s := range stmts {
		st, ok := s.(*ast.AssignStmt)
		if !ok {
			return fmt.Errorf("неожиданная инструкция %T", s)
		}

		varName := st.Ident.Value
		value, err := calcExpr(st.Expr)
		if err != nil {
			return err
		}
		vars[varName] = value
		fmt.Printf("%s = %d\n", varName, value)

		lastVarName = varName
	}

	v := vars[lastVarName]

	fmt.Printf("Ответ: %s = %d\n", lastVarName, v)

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
	case *ast.ArithmeticOpExpr:
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
		case "*":
			return vLeft * vRight, nil
		}
	}

	panic(fmt.Sprintf("ошибка в выражении, %#v", e))
}
