package demos

import (
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	cd = new(CalculatorDisplay)
	log = new( TestLogger)
)

func PressSequence(sequence string) {
	
    for _, key := range strings.Split(sequence, "") {
        Press(key)
    }
}

func Press(key string) {
    cd.Press(string(key))
    log.Append(key, cd.GetDisplay())
}

func  TestCheckDisplay(t *testing.T) {
	t.Skip("Waiting for demo")
	cd := new(CalculatorDisplay)
	cd.Press("1")
    approvals.VerifyString(t, cd.GetDisplay())
}

func TestComplexOperations(t *testing.T) {
	t.Skip("Waiting for demo")
    cd = new(CalculatorDisplay)
	log = new(TestLogger)	

    PressSequence("1+2=")
    approvals.VerifyString(t, log.GetAll())

}