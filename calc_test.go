package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	teste := sum(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	teste := sum(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubtractionCorrect(t *testing.T) {
	teste := subtraction(3, 2)
	resultado := 1

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubtractionIncorrect(t *testing.T) {
	teste := subtraction(3, 2)
	resultado := 0

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultiplicationCorrect(t *testing.T) {
	teste := multiplication(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultiplicationIncorrect(t *testing.T) {
	teste := multiplication(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivisionCorrect(t *testing.T) {
	teste := division(3, 2)
	resultado := 1.5

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivisionIncorrect(t *testing.T) {
	teste := division(3, 2)
	resultado := 2.0

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
