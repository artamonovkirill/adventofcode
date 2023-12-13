package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt", 1), 21)
	assert.Equal(t, Solve("example.txt", 5), 525152)
}

func TestSolveOne(t *testing.T) {
	for _, tc := range []struct {
		input    string
		factor   int
		expected int
	}{
		{"???.### 1,1,3", 1, 1},
		{".??..??...?##. 1,1,3", 1, 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1, 1},
		{"????.#...#... 4,1,1", 1, 1},
		{"????.######..#####. 1,6,5", 1, 4},
		{"?###???????? 3,2,1", 1, 10},

		{"???.### 1,1,3", 5, 1},
		{".??..??...?##. 1,1,3", 5, 16384},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 5, 1},
		{"????.#...#... 4,1,1", 5, 16},
		{"????.######..#####. 1,6,5", 5, 2500},
		{"?###???????? 3,2,1", 5, 506250},
	} {
		t.Run(tc.input, func(t *testing.T) {
			// expect:
			assert.Equal(t, Matches(tc.input, tc.factor), tc.expected)
		})
	}
}
