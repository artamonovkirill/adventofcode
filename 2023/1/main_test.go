package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt"), 142)
	assert.Equal(t, Solve("example2.txt"), 281)
	assert.Equal(t, Solve("example3.txt"), 78)
}
