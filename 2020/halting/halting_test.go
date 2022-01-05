package halting

import (
	"gotest.tools/assert"
	"testing"
)

func TestExampleInput(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, solve(input), 8)
}
