package main

//go:generate goyacc -o ./parser/parser.go -p YY ./parser/parser.go.y

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/negasus/golangconf2023/parser"
	"github.com/negasus/golangconf2023/runner"
)

//go:embed demo.sl
var scriptContent string

func main() {
	lx := parser.NewLexer([]byte(scriptContent))

	p := parser.YYNewParser()

	p.Parse(lx)

	if len(lx.Errors()) > 0 {
		fmt.Printf("Обнаружены ошибки\n")
		fmt.Printf("%s", strings.Join(lx.Errors(), "\n"))
		os.Exit(1)
	}

	err := runner.Run(lx.Statements())
	if err != nil {
		fmt.Printf("Ошибка выполнения: %s\n", err)
		os.Exit(1)
	}
}
