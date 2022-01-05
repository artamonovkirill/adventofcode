package adapter

import (
	"gotest.tools/assert"
	"testing"
)

var examples = []struct {
	file     string
	expected int
}{
	{
		file:     "example/1.txt",
		expected: 8,
	},
	{
		file:     "example/2.txt",
		expected: 19208,
	},
}

func TestExamples(t *testing.T) {
	for _, tc := range examples {
		t.Run(tc.file, func(t *testing.T) {
			// given:
			input := tc.file

			// expect:
			assert.Equal(t, solve(input), tc.expected)
		})
	}
}
