package treehouse

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, Visible("example.txt"), 21)
	assert.Equal(t, BestScore("example.txt"), 8)
}
