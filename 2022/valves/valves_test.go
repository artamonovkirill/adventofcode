package valves

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesPart1(t *testing.T) {
	// expect:
	assert.Equal(t, max("example.txt", 1), 1651)
}

func TestSolvesPart2(t *testing.T) {
	// expect:
	assert.Equal(t, max("example.txt", 2), 1707)
}
