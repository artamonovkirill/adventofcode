package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestHash(t *testing.T) {
	// expect:
	assert.Equal(t, hash("HASH"), 52)
	assert.Equal(t, Hash("example.txt"), 1320)
}

func TestLenses(t *testing.T) {
	assert.Equal(t, Lenses("example.txt"), 145)
}
