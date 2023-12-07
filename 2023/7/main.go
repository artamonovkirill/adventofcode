package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type player struct {
	hand string
	lex  string
	rank int
	bid  int
}

const (
	HighestCard  int = 0
	Pair             = 1
	TwoPairs         = 2
	ThreeOfAKind     = 3
	FullHouse        = 4
	FourOfAKind      = 5
	FiveOfAKind      = 6
)

func Solve(file string) int {
	result := 0
	players := parse(file)
	sort.Slice(players, func(i, j int) bool {
		a := players[i]
		b := players[j]
		if a.rank == b.rank {
			return a.lex < b.lex
		}
		return a.rank < b.rank
	})
	for i, p := range players {
		rank := i + 1
		result += rank * p.bid
	}
	return result
}

func rank(hand string) int {
	groups := make(map[int32]int)
	jokers := 0
	for _, c := range hand {
		if c == 'J' {
			jokers++
		} else {
			groups[c]++
		}
	}

	if len(groups) <= 1 {
		return FiveOfAKind
	}

	combinations := make(map[int]int)
	for _, v := range groups {
		combinations[v]++
	}

	if combinations[4] == 1 {
		return FourOfAKind
	}

	if combinations[3] == 1 {
		if jokers == 1 {
			return FourOfAKind
		}
		if combinations[2] == 1 {
			return FullHouse
		}
		return ThreeOfAKind
	}

	if combinations[2] == 2 {
		if jokers == 1 {
			return FullHouse
		}
		return TwoPairs
	}

	if combinations[2] == 1 {
		if jokers == 1 {
			return ThreeOfAKind
		}
		if jokers == 2 {
			return FourOfAKind
		}
		return Pair
	}

	switch jokers {
	case 1:
		return Pair
	case 2:
		return ThreeOfAKind
	case 3:
		return FourOfAKind
	default:
		return HighestCard
	}
}

func parse(file string) []player {
	var result []player
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, " ")
		p := player{parts[0], normalize(parts[0]), rank(parts[0]), util.Number(parts[1])}
		result = append(result, p)
	}
	return result
}

func normalize(hand string) string {
	result := ""
	for _, c := range hand {
		switch c {
		case '2':
			result += "02"
		case '3':
			result += "03"
		case '4':
			result += "04"
		case '5':
			result += "05"
		case '6':
			result += "06"
		case '7':
			result += "07"
		case '8':
			result += "08"
		case '9':
			result += "09"
		case 'T':
			result += "10"
		case 'J':
			result += "00"
		case 'Q':
			result += "12"
		case 'K':
			result += "13"
		case 'A':
			result += "14"
		default:
			panic("unknown: " + string(c))
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/7/input.txt"))
}
