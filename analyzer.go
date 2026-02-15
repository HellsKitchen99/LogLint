package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",                                         // имя линтера
	Doc:  "finds invalid args in log funs based on 4 rules", // описание линтера
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			return ins(n, pass)
		})
	}
	return nil, nil
}

func ins(n ast.Node, pass *analysis.Pass) bool {

}
