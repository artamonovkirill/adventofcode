package encoding

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesFirstPart(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, check(input, 5), 127)
}

func TestSolvesSecondPart(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, find(input, 127), 62)
}
