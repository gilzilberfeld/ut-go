package demos

import (
	"strconv"
)

// CalculatorDisplay struct
type CalculatorDisplay struct {
    display      string
    lastArgument int
    result       int
    newArgument  bool
    shouldReset  bool
    lastOperation OperationType
}

type OperationType int

const (
    Plus OperationType = iota
    Div
)

func (calcDisplay *CalculatorDisplay) Press(key string) {
    if key == "+" {
        calcDisplay.lastOperation = Plus
        calcDisplay.lastArgument, _ = strconv.Atoi(calcDisplay.display)
        calcDisplay.newArgument = true
    } else if key == "/" {
        calcDisplay.lastOperation = Div
        calcDisplay.lastArgument, _ = strconv.Atoi(calcDisplay.display)
        calcDisplay.newArgument = true
    } else if key == "=" {
        currentArgument, _ := strconv.Atoi(calcDisplay.display)
        if calcDisplay.lastOperation == Plus {
            calcDisplay.display = strconv.Itoa(calcDisplay.lastArgument + currentArgument)
        }
        if calcDisplay.lastOperation == Div && currentArgument == 0 {
            calcDisplay.display = "Division By Zero Error"
        }
        calcDisplay.shouldReset = true
    } else {
        if calcDisplay.shouldReset {
            calcDisplay.display = ""
            calcDisplay.shouldReset = false
        }
        if calcDisplay.newArgument {
            calcDisplay.display = ""
            calcDisplay.newArgument = false
        }
        calcDisplay.display += key
    }
}

func (c *CalculatorDisplay) GetDisplay() string {
    if c.display == "" {
        return "0"
    }
    if len(c.display) > 5 {
        return "E"
    }
    return c.display
}