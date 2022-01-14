package crab

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	first := []int{9, 2, 6, 3, 1}
	second := []int{5, 8, 4, 7, 10}

	_, score := solve(first, second)

	assert.Equal(t, score, 291)
}
