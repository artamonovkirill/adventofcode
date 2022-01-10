package monster

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, solve(input), 12)
}
