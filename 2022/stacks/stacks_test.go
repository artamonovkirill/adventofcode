package stacks

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, process("example.txt"), "MCD")
}
