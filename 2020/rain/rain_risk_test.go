package rain

import (
	"gotest.tools/assert"
	"strconv"
	"testing"
)

func TestSolvesExampleInput(t *testing.T) {
	// given:
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	// expect:
	assert.Equal(t, solve(input), 286)
}

func TestProcessesCommands(t *testing.T) {
	cases := []struct {
		input    Ship
		command  string
		expected Ship
	}{
		{
			Ship{0, 0, Waypoint{1, 10}},
			"F10",
			Ship{10, 100, Waypoint{1, 10}},
		},
		{
			Ship{10, 100, Waypoint{1, 10}},
			"N3",
			Ship{10, 100, Waypoint{4, 10}},
		},
		{
			Ship{10, 100, Waypoint{4, 10}},
			"F7",
			Ship{38, 170, Waypoint{4, 10}},
		},
		{
			Ship{38, 170, Waypoint{4, 10}},
			"R90",
			Ship{38, 170, Waypoint{-10, 4}},
		},
		{
			Ship{38, 170, Waypoint{-10, 4}},
			"F11",
			Ship{-72, 214, Waypoint{-10, 4}},
		},
	}
	for _, c := range cases {
		t.Run(c.command, func(t *testing.T) {
			// when:
			result := process(c.input, c.command)

			// then:
			assert.Equal(t, result, c.expected)
		})
	}
}

func TestTurning(t *testing.T) {
	cases := []struct {
		input    Waypoint
		action   uint8
		angle    int
		expected Waypoint
	}{
		{
			Waypoint{1, 2},
			'R',
			90,
			Waypoint{-2, 1},
		},
		{
			Waypoint{-2, 1},
			'R',
			90,
			Waypoint{-1, -2},
		},
		{
			Waypoint{-1, -2},
			'R',
			90,
			Waypoint{2, -1},
		},
		{
			Waypoint{2, -1},
			'R',
			90,
			Waypoint{1, 2},
		},
	}
	for _, c := range cases {
		t.Run(string(c.action)+strconv.Itoa(c.angle), func(t *testing.T) {
			// when:
			result := turn(c.input, c.action, c.angle)

			// then:
			assert.Equal(t, result, c.expected)
		})
	}
}
