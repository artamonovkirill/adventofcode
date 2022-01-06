package shuttle

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// given:
	timestamp := 939
	busses := "7,13,x,x,59,x,31,19"

	// expect:
	assert.Equal(t, find(timestamp, busses), 295)
}

func TestFindsEarliest(t *testing.T) {
	cases := []struct {
		busses   string
		expected int
	}{
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}

	for _, c := range cases {
		t.Run(c.busses, func(t *testing.T) {
			assert.Equal(t, earliest(c.busses, 0, 0), c.expected)
		})
	}
}
