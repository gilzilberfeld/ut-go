package solutions_ex1

import (
	"strconv"
)

type CalculatorDisplay struct {
    display       string
    lastArgument  int
    newArgument   bool
    shouldReset   bool
    lastOperation OperationType
}

type OperationType int

const (
    Plus OperationType = iota
    Div
)

func (calc *CalculatorDisplay) Press(key string) {
    if key == "+" {
        calc.lastOperation = Plus
        calc.lastArgument, _ = strconv.Atoi(calc.display)
        calc.newArgument = true
    } else if key == "/" {
        calc.lastOperation = Div
        calc.lastArgument, _ = strconv.Atoi(calc.display)
        calc.newArgument = true
    } else if key == "=" {
        currentArgument, _ := strconv.Atoi(calc.display)
        if calc.lastOperation == Plus {
            calc.display = strconv.Itoa(calc.lastArgument + currentArgument)
        }
        if calc.lastOperation == Div && currentArgument == 0 {
            calc.display = "Division By Zero Error"
        }
        calc.shouldReset = true
    } else {
        if calc.shouldReset {
            calc.display = ""
            calc.shouldReset = false
        }
        if calc.newArgument {
            calc.display = ""
            calc.newArgument = false
        }
        calc.display += key
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
