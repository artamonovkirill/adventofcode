package password

import (
	"gotest.tools/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	// given:
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	// when:
	result := valid(input)

	// then:
	assert.Equal(t, result, 1)
}

func TestPartialInput(t *testing.T) {
	// given:
	input := []string{
		"2-7 p: pbhhzpmppb",
	}

	// when:
	result := valid(input)

	// then:
	assert.Equal(t, result, 0)
}
