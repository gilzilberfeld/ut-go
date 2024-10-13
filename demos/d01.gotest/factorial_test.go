package demos

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	factorial := Factorial{}
	result := factorial.Calculate(1)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
	result = factorial.Calculate(2)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
	result = factorial.Calculate(3)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestNegativeFactorial(t *testing.T) {
	t.Skip("Not implemented yet")
	factorial := Factorial{}
	result := factorial.Calculate(-3)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestBigFactorial(t *testing.T) {
	factorial := Factorial{}
	result := factorial.Calculate(10)
	expected := 3628800
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
