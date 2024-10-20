package solutions_ex2

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalculatorDisplaySuite_b struct {
    suite.Suite
    cd *CalculatorDisplay
}

func (suite *CalculatorDisplaySuite_b) SetupTest() {
    suite.cd = &CalculatorDisplay{}
}

func (suite *CalculatorDisplaySuite_b) TestCheckWithCSVSource() {
    cases := suite.GetCases()
    for _, c := range cases {
        suite.Run(c.sequence, func() {
            suite.PressSequence(c.sequence)
            suite.ShouldDisplay(c.expected)
        })
    }
}

func (suite *CalculatorDisplaySuite_b) ShouldDisplay(expected string) {
    assert.Equal(suite.T(), expected, suite.cd.GetDisplay())
}

func (suite *CalculatorDisplaySuite_b) PressSequence(sequence string) {
    suite.cd = &CalculatorDisplay{}
    for _, c := range sequence {
        suite.cd.Press(string(c))
    }
}

func (suite *CalculatorDisplaySuite_b) GetCases() []struct {
    sequence string
    expected string
} {
    file, err := os.Open("data.csv")
    if err != nil {
        suite.T().Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        suite.T().Fatal(err)
    }

    var cases []struct {
        sequence string
        expected string
    }

    for _, record := range records {
        cases = append(cases, struct {
            sequence string
            expected string
        }{
            sequence: record[0],
            expected: record[1],
        })
    }

    return cases
}

func TestCalculatorDisplaySuiteB(t *testing.T) {
    suite.Run(t, new(CalculatorDisplaySuite_b))
}
