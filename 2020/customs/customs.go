package customs

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func one(input string) int {
	answers := make(map[int32]int, 26)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		for _, c := range line {
			answers[c] += 1
		}
	}
	result := 0
	for _, v := range answers {
		if v == len(lines) {
			result++
		}
	}
	return result
}

func many(input string) int {
	result := 0
	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		result += one(group)
	}
	return result
}

func Solve() {
	input := "2020/customs/puzzle.txt"
	groups := util.Text(input)
	fmt.Println(many(groups))
}
