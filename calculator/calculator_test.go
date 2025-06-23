package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	testing := sum(1, 2, 5, 10)
	result := 18
	if testing != result {
		t.Error("Valor Esperado", result, "Valor Retornado", testing)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	testing := sum(1, 2, 5, 10)
	result := 75
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	testing := sub(100, 5)
	result := 95
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}

func TestShouldSubIncorrect(t *testing.T) {
	testing := sub(100, 5)
	result := 5
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	testing := mult(10, 10)
	result := 100
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}

func TestShouldMultIncorrect(t *testing.T) {
	testing := mult(10, 10)
	result := 1024
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	testing := div(50, 2)
	result := 25
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}

func TestShouldDivIncorrect(t *testing.T) {
	testing := div(50, 2)
	result := 15
	if testing != result {
		t.Error("Valor Esperado:", result, "Valor Retornado:", testing)
	}
}
