package tube

import (
	"github.com/advendofcode/util"
	"gotest.tools/assert"
	"testing"
)

func TestSolvesStrengthExample(t *testing.T) {
	// expect:
	assert.Equal(t, strength("small_example.txt"), 0)
	assert.Equal(t, strength("example.txt"), 13140)
}

func TestSolvesCRTExample(t *testing.T) {
	// expect:
	assert.Equal(t, crt("example.txt"), util.Text("expected.txt"))
}
