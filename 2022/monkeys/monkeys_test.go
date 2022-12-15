package monkeys

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, business("example.txt", 1), 4*6)
	assert.Equal(t, business("example.txt", 20), 103*99)
	assert.Equal(t, business("example.txt", 1000), 5204*5192)
	assert.Equal(t, business("example.txt", 10000), 2713310158)
}
