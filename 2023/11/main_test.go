package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt", 2), 374)
	assert.Equal(t, Solve("example.txt", 10), 1030)
	assert.Equal(t, Solve("example.txt", 100), 8410)
}
