package rope

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, process("example.txt", 2), 13)
	assert.Equal(t, process("example.txt", 10), 1)
	assert.Equal(t, process("example2.txt", 10), 36)
}
