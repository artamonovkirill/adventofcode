package main

import (
	"github.com/advendofcode/util"
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt", 2), 374)
	assert.Equal(t, Solve("example.txt", 10), 1030)
	assert.Equal(t, Solve("example.txt", 100), 8410)
}

func TestExpand(t *testing.T) {
	// given:
	input := util.Lines("example.txt")
	expected := util.Lines("expanded.txt")

	// expect:
	assert.DeepEqual(t, expand(input, 2), expected)
}
