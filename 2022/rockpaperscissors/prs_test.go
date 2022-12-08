package rockpaperscissors

import (
	"testing"

	"gotest.tools/assert"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, best("example.txt"), 12)
}
