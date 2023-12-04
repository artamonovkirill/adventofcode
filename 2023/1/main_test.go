package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, CalibrationValue("example.txt"), 142)
	assert.Equal(t, CalibrationValue("example2.txt"), 281)
	assert.Equal(t, CalibrationValue("example3.txt"), 78)
}
