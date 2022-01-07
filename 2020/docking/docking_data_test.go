package docking

import (
	"gotest.tools/assert"
	"testing"
)

func TestExampleCommand(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, solve(input), 208)
}
