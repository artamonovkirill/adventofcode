package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestPoints(t *testing.T) {
	// expect:
	assert.Equal(t, Points("example.txt"), 13)
}

func TestCards(t *testing.T) {
	// expect:
	assert.Equal(t, Cards("example.txt"), 30)
}
