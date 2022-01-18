package combo

import (
	"gotest.tools/assert"
	"testing"
)

func TestFindsLoopSize(t *testing.T) {
	assert.Equal(t, loopSize(7, 5764801), 8)
	assert.Equal(t, loopSize(7, 17807724), 11)
	assert.Equal(t, key(17807724, 8), 14897079)
}
