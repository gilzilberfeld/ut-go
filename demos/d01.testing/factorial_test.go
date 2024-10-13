package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	factorial := Factorial{}
	result := factorial.Calculate(1)
	// testing package, manual check
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	result = factorial.Calculate(2)
	// testify, with additional error message
	assert.Equal(t, 2, result, "DAMMIT! Expected 2, got %d", result)
	result = factorial.Calculate(3)
	// testify, without error message
	assert.Equal(t, 6, result)
}

func TestNegativeFactorial(t *testing.T) {
	t.Skip("Not implemented yet")
	factorial := Factorial{}
	result := factorial.Calculate(-3)
	assert.Equal(t, 0, result)
}

func TestBigFactorial(t *testing.T) {
	factorial := Factorial{}
	result := factorial.Calculate(10)
	expected := 3628800
	assert.Equal(t, expected, result)
}
