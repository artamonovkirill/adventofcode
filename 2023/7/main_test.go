package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	// expect:
	assert.Equal(t, Solve("example.txt"), 5905)
}

func TestRank(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int
	}{
		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPairs},
		{"A23A4", Pair},
		{"23456", HighestCard},

		{"32T3K", Pair},
		{"KK677", TwoPairs},
		{"KTJJT", FourOfAKind},
		{"T55J5", FourOfAKind},
		{"QQQJA", FourOfAKind},

		{"A9999", FourOfAKind},

		{"AAAAJ", FiveOfAKind},
		{"AAAJJ", FiveOfAKind},
		{"AAJJJ", FiveOfAKind},
		{"AJJJJ", FiveOfAKind},
		{"JJJJJ", FiveOfAKind},

		{"234JA", Pair},
		{"23JJA", ThreeOfAKind},
		{"2JJJA", FourOfAKind},
	} {
		t.Run(tc.input, func(t *testing.T) {
			// expect:
			assert.Equal(t, rank(tc.input), tc.expected)
		})
	}
}
