package solutions_ex1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. Write tests
// 2. Incremental coding

func TestCalculatorDisplay_AtStartDisplays0(t *testing.T) {
    cd := &CalculatorDisplay{}
    assert.Equal(t, "0", cd.GetDisplay())
}

func TestCalculatorDisplay_Pressing1Displays1(t *testing.T) {
    cd := &CalculatorDisplay{}
    cd.Press("1")
    assert.Equal(t, "1", cd.GetDisplay())
}

func TestCalculatorDisplay_Pressing12Displays12(t *testing.T) {
    cd := &CalculatorDisplay{}
    cd.Press("1")
    cd.Press("2")
    assert.Equal(t, "12", cd.GetDisplay())
}
