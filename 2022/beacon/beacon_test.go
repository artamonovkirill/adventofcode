package beacon

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, impossible("example.txt", 10), 26)
	assert.Equal(t, frequency("example.txt", 20), 56000011)
}
