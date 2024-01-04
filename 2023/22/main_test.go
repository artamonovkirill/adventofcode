package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt"), 5)
	assert.Equal(t, Solve("example2.txt"), 3)
}

func TestSolve2(t *testing.T) {
	// expect:
	assert.Equal(t, Solve2("example.txt"), 7)
}
