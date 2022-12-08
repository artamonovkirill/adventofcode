package rucksack

import (
	"testing"

	"gotest.tools/assert"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, duplicates("example.txt"), 157)
	assert.Equal(t, identifiers("example.txt"), 70)
}

func TestCalculatesPriority(t *testing.T) {
	for _, tc := range []struct {
		letter   rune
		expected int
	}{
		{letter: 'a', expected: 1},
		{letter: 'p', expected: 16},
		{letter: 'z', expected: 26},
		{letter: 'A', expected: 27},
		{letter: 'L', expected: 38},
		{letter: 'Z', expected: 52},
	} {
		t.Run(string(tc.letter), func(t *testing.T) {
			// expect:
			assert.Equal(t, priority(tc.letter), tc.expected)
		})
	}
}
