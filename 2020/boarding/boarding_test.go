package boarding

import (
	"gotest.tools/assert"
	"testing"
)

var passes = []struct {
	input    string
	expected Pass
}{
	{
		input:    "FBFBBFFRLR",
		expected: Pass{Row: 44, Column: 5, ID: 357},
	},
	{
		input:    "BFFFBBFRRR",
		expected: Pass{Row: 70, Column: 7, ID: 567},
	},
	{
		input:    "FFFBBBFRRR",
		expected: Pass{Row: 14, Column: 7, ID: 119},
	},
	{
		input:    "BBFFBBFRLL",
		expected: Pass{Row: 102, Column: 4, ID: 820},
	},
}

func TestExamplePasses(t *testing.T) {
	for _, tc := range passes {
		t.Run(tc.input, func(t *testing.T) {
			// expect:
			assert.Equal(t, parse(tc.input), tc.expected)
		})
	}
}
