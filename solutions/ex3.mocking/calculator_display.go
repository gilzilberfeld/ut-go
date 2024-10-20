package solutions_ex3

import "strconv"

type OperationType int

const (
	Plus OperationType = iota
	Div
)

 // 1. Modify CalculatorDisplay to accept external display
   
type CalculatorDisplay struct {
	externalDisplay    IExternalDisplay
	IsDisplayConnected bool
	display            string
	lastArgument       int
	newArgument        bool
	shouldReset        bool
	lastOperation      OperationType
}

func NewCalculatorDisplay(external IExternalDisplay) *CalculatorDisplay {
	return &CalculatorDisplay{
		externalDisplay:    external,
		IsDisplayConnected: false,
		display:            "",
		lastArgument:       0,
		newArgument:        false,
		shouldReset:        true,
	}
}

func (cd *CalculatorDisplay) Press(key string) {
	switch key {
	case "+":
		cd.lastOperation = Plus
		cd.lastArgument = cd.parseDisplay()
		cd.newArgument = true
	case "/":
		cd.lastOperation = Div
		cd.lastArgument = cd.parseDisplay()
		cd.newArgument = true
	case "=":
		currentArgument := cd.parseDisplay()
		switch cd.lastOperation {
		case Plus:
			cd.display = formatInt(cd.lastArgument + currentArgument)
		case Div:
			if currentArgument == 0 {
				cd.display = "Division By Zero Error"
			} else {
				cd.display = formatInt(cd.lastArgument / currentArgument)
			}
		}
		cd.shouldReset = true
	default:
		if cd.shouldReset {
			cd.display = ""
			cd.shouldReset = false
		}
		if cd.newArgument {
			cd.display = ""
			cd.newArgument = false
		}
		cd.display += key
	}

	if cd.externalDisplay.IsOn() {
		cd.externalDisplay.Show(cd.GetDisplay())
	} else {
		cd.IsDisplayConnected = false
	}
}

func (cd *CalculatorDisplay) GetDisplay() string {
	if cd.display == "" {
		return "0"
	}
	if len(cd.display) > 5 {
		return "E"
	}
	return cd.display
}

func (cd *CalculatorDisplay) parseDisplay() int {
	value, _ := strconv.Atoi(cd.display)
	return value
}

func formatInt(value int) string {
	return strconv.Itoa(value)
}
