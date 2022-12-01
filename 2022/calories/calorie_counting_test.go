package calories

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, max("example.txt"), 24000)
	assert.Equal(t, topThree("example.txt"), 45000)
}
