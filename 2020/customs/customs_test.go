package customs

import (
	"gotest.tools/assert"
	"testing"
)

func TestExampleGroup(t *testing.T) {
	// expect:
	assert.Equal(t, one("abcx\nabcy\nabcz"), 3)
}

func TestExampleGroups(t *testing.T) {
	// given:
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"

	// expect:
	assert.Equal(t, many(input), 6)
}
