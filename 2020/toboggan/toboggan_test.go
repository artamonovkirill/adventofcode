package toboggan

import (
	"gotest.tools/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	// given:
	input := "example.txt"

	// when:
	result := solve(input)

	// then:
	assert.Equal(t, result, 336)
}
