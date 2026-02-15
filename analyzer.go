package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var allowedMethods = map[string]bool{
	"Info":    true,
	"Error":   true,
	"Warn":    true,
	"Debug":   true,
	"Print":   true,
	"Printf":  true,
	"Println": true,
}

var allowedLibs = map[string]bool{
	"log":  true,
	"slog": true,
	"zap":  true,
}

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
	call, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}
	funcName := call.Fun
	selector, ok := funcName.(*ast.SelectorExpr)
	if !ok {
		return true
	}
	var needX any
	switch x := selector.X.(type) {
	case *ast.Ident:
		needX = x
	case *ast.CallExpr:
		subFunName := x.Fun
		subSelector, ok := subFunName.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		subX := subSelector.X
		subXIdent, ok := subX.(*ast.Ident)
		if !ok {
			return true
		}
		needX = subXIdent
	default:
		return true
	}
	x, ok := needX.(*ast.Ident)
	if !ok {
		return true
	}
	sel := selector.Sel
	if !allowedLibs[x.Name] {
		return true
	}
	if !allowedMethods[sel.Name] {
		return true
	}
	// дальнейшая логика
	return true
}
