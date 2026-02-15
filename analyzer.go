package loglint

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"

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

// получения содержимого лога
func extractMessage(call *ast.CallExpr) (string, bool) {
	args := call.Args
	if len(args) == 0 {
		return "", false
	}
	basicLit, ok := args[0].(*ast.BasicLit)
	if !ok {
		return "", false
	}
	if basicLit.Kind != token.STRING {
		return "", false
	}
	basicLitValue, err := strconv.Unquote(basicLit.Value)
	if err != nil {
		return "", false
	}
	return basicLitValue, true
}

// проверка на регистр
func checkLowerCase(msg string, call *ast.CallExpr, pass *analysis.Pass) bool {
	for _, letter := range msg {
		if unicode.IsUpper(letter) {
			pass.Reportf(call.Pos(), "log message must not contain upper case letters")
			return false
		}
	}
	return true
}

// проверка на язык
func checkEnglish(msg string, call *ast.CallExpr, pass *analysis.Pass) bool {
	isEnglishLetter := func(letter rune) bool {
		if letter >= 'a' && letter <= 'z' {
			return true
		}
		return false
	}
	for _, letter := range msg {
		if unicode.IsLetter(letter) {
			if !isEnglishLetter(letter) {
				pass.Reportf(call.Pos(), "log message must consist only of English letters")
				return false
			}
		}
	}
	return true
}

// проверка на спец символы
func checkNoSpecialChars(msg string, call *ast.CallExpr, pass *analysis.Pass) bool {

}

// проверка на важные данные
func checkSensitive(msg string, call *ast.CallExpr, pass *analysis.Pass) bool {

}
