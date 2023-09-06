%{
package parser

import (
  "github.com/negasus/golangconf2023/ast"
)
%}

%type<nil> program
%type<stmt> stmt
%type<expr> expr

%union {
	token 		*ast.Token
	stmt 		ast.Stmt
	expr 		ast.Expr
}


%token<token> IDENT NUMBER EQ PRINT LPAREN RPAREN STRING

%left PLUS MINUS
%left MUL

%%

program: {}
	|
	program stmt {
		l := YYlex.(*Lexer)
		l.stmts = append(l.stmts, $2)
	}

stmt:
	IDENT EQ expr  {
		$$ = &ast.AssignStmt{Ident: &ast.IdentExpr{Value:$1.Value}, Expr: $3}
	}
	|
	PRINT LPAREN expr RPAREN {
		$$ = &ast.PrintStmt{Expr: $3}
	}

expr:
	STRING {
		$$ = &ast.StringExpr{Value: $1.Value}
	}
	|
        NUMBER {
        	$$ = &ast.NumberExpr{Value: $1.Value}
        }
        |
        IDENT {
        	$$ = &ast.IdentExpr{Value:$1.Value}
        }
        |
        expr PLUS expr {
        	$$ = &ast.ArithmeticOpExpr{Lhs: $1, Operator: "+", Rhs: $3}
        }
        |
        expr MINUS expr {
        	$$ = &ast.ArithmeticOpExpr{Lhs: $1, Operator: "-", Rhs: $3}
        }
        |
        expr MUL expr {
        	$$ = &ast.ArithmeticOpExpr{Lhs: $1, Operator: "*", Rhs: $3}
        }
%%
