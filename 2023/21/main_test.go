package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt", 1), 2)
	assert.Equal(t, Solve("example.txt", 2), 4)
	assert.Equal(t, Solve("example.txt", 3), 6)
	assert.Equal(t, Solve("example.txt", 6), 16)
}
