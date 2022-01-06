package seating

import (
	"fmt"
	"github.com/advendofcode/util"
	"gotest.tools/assert"
	"strconv"
	"testing"
)

func TestAdvances(t *testing.T) {
	for _, steps := range []int{1, 2, 3, 4, 5, 6} {
		t.Run(strconv.Itoa(steps), func(t *testing.T) {
			// given:
			input := parse("example.txt")
			expected := fmt.Sprintf("steps/%d.txt", steps)

			// expect:
			assert.DeepEqual(t, toString(advance(input, steps)), util.Text(expected))
		})
	}
}

func TestSolves(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, solve(input), 26)
}

func TestFindsNeighbours(t *testing.T) {
	// given:
	input := parse("steps/1.txt")

	// expect:
	assert.DeepEqual(t, neighbours(input, 0, 2)["#"], 5)
}
