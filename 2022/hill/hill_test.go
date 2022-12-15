package hill

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, path("example.txt"), 29)
}

func TestAllAroundAreOnGrid(t *testing.T) {
	// given:
	grid, _ := parse("example.txt")

	// expect:
	for _, row := range grid {
		for _, cell := range row {
			around(cell, grid)
		}
	}
}
