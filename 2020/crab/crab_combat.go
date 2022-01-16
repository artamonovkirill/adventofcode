package crab

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func toString(cards []int) string {
	return util.ToString(cards, ", ")
}

func solve(firstHand []int, secondHand []int) (int, int) {
	previous := map[string]map[string]bool{}
	for {
		if len(firstHand) == 0 {
			winner := 2
			return winner, score(secondHand)
		}
		if len(secondHand) == 0 {
			winner := 1
			return winner, score(firstHand)
		}

		if previous[toString(firstHand)][toString(secondHand)] {
			return 1, score(firstHand)
		}

		if _, ok := previous[toString(firstHand)]; !ok {
			previous[toString(firstHand)] = map[string]bool{}
		}
		previous[toString(firstHand)][toString(secondHand)] = true

		firstMove := firstHand[0]
		firstHand = firstHand[1:]
		secondMove := secondHand[0]
		secondHand = secondHand[1:]

		if len(firstHand) >= firstMove && len(secondHand) >= secondMove {
			player, _ := solve(deepCopy(firstHand, firstMove), deepCopy(secondHand, secondMove))
			if player == 1 {
				firstHand = append(firstHand, firstMove)
				firstHand = append(firstHand, secondMove)
			} else {
				secondHand = append(secondHand, secondMove)
				secondHand = append(secondHand, firstMove)
			}
		} else {
			if firstMove > secondMove {
				firstHand = append(firstHand, firstMove)
				firstHand = append(firstHand, secondMove)
			} else {
				secondHand = append(secondHand, secondMove)
				secondHand = append(secondHand, firstMove)
			}
		}
	}
}

func deepCopy(xs []int, n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = xs[i]
	}
	return result
}

func score(hand []int) int {
	result := 0
	for i := len(hand); i > 0; i-- {
		result += hand[len(hand)-i] * i
	}
	return result
}

func parse(file string) ([]int, []int) {
	text := util.Text(file)
	first, second := util.Split(text, "\n\n")
	return hand(first), hand(second)
}

func hand(text string) []int {
	lines := strings.Split(text, "\n")[1:]
	result := make([]int, len(lines))
	for i, l := range lines {
		result[i] = util.Number(l)
	}
	return result
}

func Solve() {
	f, s := parse("2020/crab/puzzle.txt")
	fmt.Println(solve(f, s))
}
