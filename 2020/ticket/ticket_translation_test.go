package ticket

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesPart1Example(t *testing.T) {
	// given:
	input := "example/part1.txt"

	// expect:
	assert.Equal(t, invalidTickets(input), 71)
}

func TestSolvesPart2Example(t *testing.T) {
	// given:
	input := "example/part2.txt"

	// expect:
	assert.DeepEqual(t, solve(input), map[string]int{
		"row":   11,
		"class": 12,
		"seat":  13,
	})
}
