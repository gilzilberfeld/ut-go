package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialTable(t *testing.T) {
    theFactorial := new(Factorial)
    cases := []struct {
        input    int
        expected int
    }{
        {1, 1},
        {2, 2},
        {3, 6},
    }

    for _, c := range cases {
            assert.Equal(t, c.expected, theFactorial.Calculate(c.input))
    }
}
