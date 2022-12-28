package flow

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, proceed("example.txt", 2022), 3068)
	assert.Equal(t, proceed("example.txt", 10000), 15148)
	//assert.Equal(t, proceed("example.txt", 1000000000000), 1514285714288)
}

func TestPredicts(t *testing.T) {
	// expect:
	assert.Equal(t, predict("example.txt", 2022), 3068)
	assert.Equal(t, predict("example.txt", 1000000000000), 1514285714288)
}

func TestParsesRocks(t *testing.T) {
	// given:
	rocks := parse()

	// expect:
	assert.Equal(t, len(rocks), 5)
	assert.Equal(t, rockToString(rocks[0]), "..@@@@.")
	assert.Equal(t, rockToString(rocks[1]), "...@...\n..@@@..\n...@...")
	assert.Equal(t, rockToString(rocks[2]), "....@..\n....@..\n..@@@..")
	assert.Equal(t, rockToString(rocks[3]), "..@....\n..@....\n..@....\n..@....")
	assert.Equal(t, rockToString(rocks[4]), "..@@...\n..@@...")
}
