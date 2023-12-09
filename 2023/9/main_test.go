package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, PredictFuture("example.txt"), 114)
	assert.Equal(t, ExtrapolateHistory("example.txt"), 2)
}
