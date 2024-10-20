package solutions_ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDisplay struct {
    isOn  bool
    text  string
}

func (md *MockDisplay) IsOn() bool {
    return md.isOn
}

func (md *MockDisplay) Show(display string) {
    md.text = display
}

func (md *MockDisplay) GetText() string {
    return md.text
}

func TestWhenExternalDisplayIsOff_DisplayNotConnected(t *testing.T) {
    display := &MockDisplay{isOn: false}
    cd := NewCalculatorDisplay(display)
    assert.False(t, cd.IsDisplayConnected)
}

func TestWhenDisplayIsOn_DisplayIsCorrect(t *testing.T) {
    display := &MockDisplay{isOn: true}
    cd := NewCalculatorDisplay(display)
    cd.Press("1")
    assert.Equal(t, "1", display.GetText())
}
