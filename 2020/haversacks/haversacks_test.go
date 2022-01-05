package haversacks

import (
	"gotest.tools/assert"
	"testing"
)

func TestExampleRules(t *testing.T) {
	// given:
	input := "example.txt"

	// expect:
	assert.Equal(t, solve(input), 126)
}
