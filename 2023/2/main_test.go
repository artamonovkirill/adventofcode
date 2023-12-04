package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestValid(t *testing.T) {
	// expect:
	assert.Equal(t, Valid(Game{12, 13, 14}, "example.txt"), 1+2+5)
}

func TestPower(t *testing.T) {
	// expect:
	assert.Equal(t, Power("example.txt"), 2286)
}
