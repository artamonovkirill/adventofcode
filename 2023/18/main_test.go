package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt", parse), 62)
	assert.Equal(t, Solve("input.txt", parse), 44436)
	assert.Equal(t, Solve("example.txt", parse2), 952408144115)
}
