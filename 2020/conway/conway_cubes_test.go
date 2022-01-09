package conway

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// given:
	input := "example.txt"
	init := parse(input)

	// when:
	updated := advance(init, 6)

	// then:
	assert.Equal(t, lit(updated), 848)
}
