package loglint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var modeOneLenFunc = 1
var modeTwoLenFunc = 2

// Тест checkLowerCase - Успех
func TestCheckLowerCaseSuccess(t *testing.T) {
	// preparing
	msg := "the bay harbour butcher"
	expectedResult := true

	// test
	result := checkLowerCase(msg)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkLowerCase - Провал
func TestCheckLowerCaseFailure(t *testing.T) {
	// preparing
	msg := "The Bay Harbour Butcher"
	expectedResult := false

	// test
	result := checkLowerCase(msg)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkEnglish - Успех
func TestCheckEnglishSuccess(t *testing.T) {
	// preparing
	msg := "something on english"
	expectedResult := true

	// test
	result := checkEnglish(msg)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkEnglish - Провал
func TestCheckEnglishFailure(t *testing.T) {
	// preparing
	msg := "что то на русском"
	expectedResult := false

	// test
	result := checkEnglish(msg)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkNoSpecialChars - Успех
func TestCheckNoSpecialCharsSuccess(t *testing.T) {
	// preparing
	msg := "abc_-123 "
	expectedResult := true

	// test
	result := checkNoSpecialChars(msg)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkNoSpecialChars - Провал
func TestCheckNoSpecialCharsFailure(t *testing.T) {
	// preparing
	msg := "@%😶‍🌫️🥶"
	expectedResult := false

	// test
	result := checkNoSpecialChars(msg)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkSensitive - Успех
func TestCheckSensitiveSuccess(t *testing.T) {
	// preparing
	word1, word2 := "Smth", ""
	exprs, err := checkSensitivePreparing(word1, word2, modeOneLenFunc)
	fmt.Println(exprs)
	if err != nil {
		t.Errorf("error while trying to parse go code to ast.File: %v", err)
	}
	expectedResult := true

	// test
	result := checkSensitive(exprs)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkSensitive - Провал (BasicLit)
func TestCheckSensitiveFailureBasicLit(t *testing.T) {
	// preparing
	msg1, msg2 := "password", ""
	exprs, err := checkSensitivePreparing(msg1, msg2, modeOneLenFunc)
	if err != nil {
		t.Errorf("error while trying to parse go code to ast.File: %v", err)
	}
	expectedResult := false

	// test
	result := checkSensitive(exprs)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkSensitive - Провал (Ident)
func TestCheckSensitiveFailureIdent(t *testing.T) {
	// preparing
	msg1, msg2 := "", "password"
	exprs, err := checkSensitivePreparing(msg1, msg2, modeOneLenFunc)
	if err != nil {
		t.Errorf("error while trying to parse go code to ast.File: %v", err)
	}
	expectedResult := false

	// test
	result := checkSensitive(exprs)

	// assert
	if result != expectedResult {
		t.Errorf("edxpected result - %v", expectedResult)
	}
}

// Тест checkSensitive - Провал (SelectorExpr, password)
func TestCheckSensitiveFailureSelectorExpr1(t *testing.T) {
	// preparing
	msg1, msg2 := "password", ""
	exprs, err := checkSensitivePreparing(msg1, msg2, modeTwoLenFunc)
	if err != nil {
		t.Errorf("error while trying to parse go code to ast.File: %v", err)
	}
	expectedResult := false

	// test
	result := checkSensitive(exprs)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Тест checkSensitive - Провал (SelectorExpr, user.Password)
func TestCheckSensitiveFailureSelectorExpr2(t *testing.T) {
	// preparing
	msg1, msg2 := "", "user.Password"
	exprs, err := checkSensitivePreparing(msg1, msg2, modeTwoLenFunc)
	if err != nil {
		t.Errorf("error while trying to parse go code to ast.File: %v", err)
	}
	expectedResult := false

	// test
	result := checkSensitive(exprs)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

// Прочие функции
func checkSensitivePreparing(word1, word2 string, mode int) ([]ast.Expr, error) {
	msg := ""
	switch mode {
	case 1:
		msg = fmt.Sprintf(`package main

		func main() {
			fmt.Println("%v", %v)
		}`, word1, word2)
	case 2:
		msg = fmt.Sprintf(`package main

		func main() {
			zap.L().Info("%v", %v)
		}`, word1, word2)
	}
	set := token.NewFileSet()
	node, err := parser.ParseFile(set, "", msg, 0)
	if err != nil {
		return []ast.Expr{}, err
	}

	var exprs []ast.Expr
	ast.Inspect(node, func(n ast.Node) bool {
		exp, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		exprs = append(exprs, exp.Args...)
		return true
	})
	return exprs, nil
}
