package solutions_ex1

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// 3. Refactor tests

type CalculatorDisplaySuite struct {
    suite.Suite
    cd *CalculatorDisplay
}


func (suite *CalculatorDisplaySuite) SetupTest() {
    suite.cd = &CalculatorDisplay{}
}

func (suite *CalculatorDisplaySuite) TestAtStartDisplays0() {
    suite.Equal( "0", suite.cd.GetDisplay())
}

func (suite *CalculatorDisplaySuite) TestPressing1Displays1() {
    suite.cd.Press("1")
    suite.Equal( "1", suite.cd.GetDisplay())
}

func (suite *CalculatorDisplaySuite) TestPressing12Displays12() {
    suite.cd.Press("1")
    suite.cd.Press("2")
    suite.Equal("12", suite.cd.GetDisplay())
}

func TestCalculatorDisplaySuite(t *testing.T) {
    suite.Run(t, new(CalculatorDisplaySuite))
}
