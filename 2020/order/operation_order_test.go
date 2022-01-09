package order

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExamples(t *testing.T) {
	cases := []struct {
		expression string
		result     int
	}{
		{"1", 1},
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, c := range cases {
		t.Run(c.expression, func(t *testing.T) {
			assert.Equal(t, solve(c.expression), c.result)
		})
	}
}
