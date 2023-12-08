package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt"), 2)
	assert.Equal(t, Solve("example2.txt"), 6)
}

func TestSolve2(t *testing.T) {
	// expect:
	assert.Equal(t, Solve2("example3.txt"), 6)
}
