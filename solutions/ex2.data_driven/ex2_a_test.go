package solutions_ex2

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

//  1. Add data-driven / table tests

type CalculatorDisplaySuite struct {
    suite.Suite
    cd *CalculatorDisplay
}

func (suite *CalculatorDisplaySuite) SetupTest() {
    suite.cd = &CalculatorDisplay{}
}

func (suite *CalculatorDisplaySuite) TestCheckWithDataTable() {
    cases := []struct {
        sequence string
        expected string
    }{
        {"1", "1"},
        {"12", "12"},
        {"1+", "1"},
    }

    for _, tt := range cases {
        suite.Run(tt.sequence, func() {
            suite.PressSequence(tt.sequence)
            suite.ShouldDisplay(tt.expected)
        })
    }
}

func (suite *CalculatorDisplaySuite) ShouldDisplay(expected string) {
    suite.Equal( expected, suite.cd.GetDisplay())
}

func (suite *CalculatorDisplaySuite) PressSequence(sequence string) {
    suite.cd = &CalculatorDisplay{}
	for _, c := range sequence {
        suite.cd.Press(string(c))
    }
}


func TestCalculatorDisplaySuite(t *testing.T) {
    suite.Run(t, new(CalculatorDisplaySuite))
}
