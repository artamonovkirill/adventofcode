package cleanup

import (
	"testing"

	"gotest.tools/assert"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, count("example.txt"), 4)
}
