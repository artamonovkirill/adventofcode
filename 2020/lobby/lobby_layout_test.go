package lobby

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	init := initialize("example.txt")
	assert.Equal(t, len(init), 10)

	cases := []struct {
		steps    int
		expected int
	}{
		{1, 15},
		{2, 12},
		{100, 2208},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.steps, c.expected), func(t *testing.T) {
			assert.Equal(t, len(advance(init, c.steps)), c.expected)
		})
	}
}

func TestCalculatesNeighbours(t *testing.T) {
	cases := []struct {
		p        Point
		l        Lobby
		expected int
	}{
		{Point{0, 0, 0}, Lobby{}, 0},
		{Point{0, 0, 0}, Lobby{
			Point{0, 1, 1}: true,
		}, 1},
		{Point{0, 1, 1}, Lobby{
			Point{0, 0, 0}: true,
		}, 1},
		{Point{0, 0, 0}, Lobby{
			Point{0, 1, 1}:   true,
			Point{1, 1, 0}:   true,
			Point{1, 0, -1}:  true,
			Point{0, -1, -1}: true,
			Point{-1, -1, 0}: true,
			Point{-1, 0, 1}:  true,
		}, 6},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.p, c.l), func(t *testing.T) {
			assert.Equal(t, black(neighbours(c.p), c.l), c.expected)
		})
	}
}
