package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt", 1), 8*4)
	assert.Equal(t, Solve("example.txt", 1000), 32000000)
	assert.Equal(t, Solve("example2.txt", 1000), 11687500)
	assert.Equal(t, Solve("input.txt", 1000), 821985143)
}
