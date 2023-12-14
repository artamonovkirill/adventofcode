package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt"), 64)
}

func TestTilt(t *testing.T) {
	// given:
	m := parse("example.txt")

	// when:
	tiltUp(m)

	// then:
	assert.Equal(t, toString(m), "OOOO.#.O..\nOO..#....#\nOO..O##..O\nO..#.OO...\n........#.\n..#....#.#\n..O..#.O.O\n..O.......\n#....###..\n#....#....")
}

func TestCycle(t *testing.T) {
	// given:
	m := parse("example.txt")

	// when:
	cycle(m)

	// then:
	assert.Equal(t, toString(m), ".....#....\n....#...O#\n...OO##...\n.OO#......\n.....OOO#.\n.O#...O#.#\n....O#....\n......OOOO\n#...O###..\n#..OO#....")

	// when:
	cycle(m)

	// then:
	assert.Equal(t, toString(m), ".....#....\n....#...O#\n.....##...\n..O#......\n.....OOO#.\n.O#...O#.#\n....O#...O\n.......OOO\n#..OO###..\n#.OOO#...O")

	// when:
	cycle(m)

	// then:
	assert.Equal(t, toString(m), ".....#....\n....#...O#\n.....##...\n..O#......\n.....OOO#.\n.O#...O#.#\n....O#...O\n.......OOO\n#...O###.O\n#.OOO#...O")
}
