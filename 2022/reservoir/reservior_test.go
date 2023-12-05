//go:build flaky
// +build flaky

package reservoir

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, simulate("example.txt", false), 24)
	assert.Equal(t, simulate("example.txt", true), 93)
}
