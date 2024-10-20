package solutions_ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDisplay_b struct {
    mock.Mock
}

func (md *MockDisplay_b) IsOn() bool {
    args := md.Called()
    return args.Bool(0)
}

func (md *MockDisplay_b) Show(display string) {
    md.Called(display)
}

func TestWhenExternalDisplayIsOff_DisplayNotConnected_b(t *testing.T) {
    mockDisplay := new(MockDisplay_b)
    mockDisplay.On("IsOn").Return(false)

    cd := NewCalculatorDisplay(mockDisplay)
    assert.False(t, cd.IsDisplayConnected)
}

func TestWhenDisplayIsOn_DisplayIsCorrect_b(t *testing.T) {
    mockDisplay := new(MockDisplay_b)
    mockDisplay.On("IsOn").Return(true)
    mockDisplay.On("Show", "1").Return()

    cd := NewCalculatorDisplay(mockDisplay)
    cd.Press("1")

    mockDisplay.AssertExpectations(t)
}

func TestWhenDisplayIsOn_DisplayIsCorrectWithCaptor(t *testing.T) {
    mockDisplay := new(MockDisplay_b)
    theArgs := mock.Arguments{}
    mockDisplay.On("IsOn").Return(true)
    mockDisplay.On("Show", mock.Anything).Run(func(args mock.Arguments) {
		theArgs =append(theArgs, args[0])
    }).Return()

    cd := NewCalculatorDisplay(mockDisplay)
    cd.Press("1")

    assert.Equal(t, "1", theArgs[0])
}
