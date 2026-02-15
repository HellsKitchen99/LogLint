package loglint

import (
	"testing"
)

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
