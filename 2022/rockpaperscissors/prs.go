package rockpaperscissors

import (
	"fmt"
	"strings"

	"github.com/advendofcode/util"
)

type hand int

const (
	rock     hand = 1
	paper         = 2
	scissors      = 3
)

func score(my hand, other hand) int {
	if my == other {
		return 3
	}
	switch my {
	case rock:
		switch other {
		case paper:
			return 0
		case scissors:
			return 6
		}
	case paper:
		switch other {
		case scissors:
			return 0
		case rock:
			return 6
		}
	case scissors:
		switch other {
		case rock:
			return 0
		case paper:
			return 6
		}
	}
	panic(fmt.Sprintf("not supported: %d, %d", my, other))
}

func win(h hand) hand {
	switch h {
	case rock:
		return paper
	case paper:
		return scissors
	case scissors:
		return rock
	}
	panic("not supported")
}

func draw(h hand) hand {
	return h
}

func loose(h hand) hand {
	switch h {
	case rock:
		return scissors
	case paper:
		return rock
	case scissors:
		return paper
	}
	panic("not supported")
}

func mapping() map[string]func(hand) hand {
	return map[string]func(hand) hand{
		"Z": win,
		"Y": draw,
		"X": loose,
	}
}

func parse(hand string) hand {
	switch hand {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	default:
		panic("not supported: " + hand)
	}
}

func best(file string) int {
	mapping := mapping()
	result := 0
	for _, l := range util.Lines(file) {
		hands := strings.Split(l, " ")
		other := parse(hands[0])
		my := mapping[hands[1]](other)
		result += score(my, other) + int(my)
	}
	return result
}

func Solve() {
	file := "2022/rockpaperscissors/input.txt"
	fmt.Println(best(file))
}
