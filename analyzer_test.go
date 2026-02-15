package loglint

import (
	"fmt"
	"testing"
)

// Тест checkLowerCase - Успех
func TestCheckLowerCaseSuccess(t *testing.T) {
	// preparing
	msg := "the bay harbour butcher"
	expectedResult := true

	// test
	result := checkLowerCase(msg)
	fmt.Println(result)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}

func TestCheckLowerCaseFailure(t *testing.T) {
	// preparing
	msg := "The Bay Harbour Butcher"
	expectedResult := false

	// test
	result := checkLowerCase(msg)
	fmt.Println(result)

	// assert
	if result != expectedResult {
		t.Errorf("expected result - %v", expectedResult)
	}
}
