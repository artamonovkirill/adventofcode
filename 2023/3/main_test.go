package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestParts(t *testing.T) {
	// expect:
	assert.Equal(t, Parts("example.txt"), 4361)
}

func TestGears(t *testing.T) {
	// expect:
	assert.Equal(t, Gears("example.txt"), 467835)
}
