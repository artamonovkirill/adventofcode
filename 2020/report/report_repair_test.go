package report

import (
	"gotest.tools/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	// given:
	report := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	//when:
	result := solve(report)

	// then:
	assert.Equal(t, result, 241861950)
}
