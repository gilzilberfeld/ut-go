package demos

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// The Suite type
// Embeds also the asserts package
type FactorialSuite struct {
    suite.Suite
    theFactorial Factorial
}

// Runs before each test
func (suite *FactorialSuite) SetupTest() {
    suite.theFactorial = Factorial{}
}

func (suite *FactorialSuite) TestFactorialTests() {
    suite.Equal(1, suite.theFactorial.Calculate(1))
    suite.Equal(2, suite.theFactorial.Calculate(2))
    suite.Equal(6, suite.theFactorial.Calculate(3))
}

func (suite *FactorialSuite) TestNegativeFactorial() {
	suite.T().Skip("Not implemented yet")
    suite.Equal(0, suite.theFactorial.Calculate(-3))
}

func (suite *FactorialSuite) TestBigFactorial() {
    suite.Equal(3628800, suite.theFactorial.Calculate(10))
}

func TestFactorialSuite(t *testing.T) {
    suite.Run(t, new(FactorialSuite))
}
