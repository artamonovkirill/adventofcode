package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	// expect:
	assert.Equal(t, LongestDistance("example.txt"), 4)
	assert.Equal(t, LongestDistance("example2.txt"), 8)
}

func TestSolve2(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int
	}{
		{"example.txt", 1},
		{"example2.txt", 1},
		{"example3.txt", 4},
		{"example4.txt", 4},
		{"example5.txt", 8},
		{"example6.txt", 10},
	} {
		t.Run(tc.input, func(t *testing.T) {
			// expect:
			assert.Equal(t, Insides(tc.input), tc.expected)
		})
	}
}
