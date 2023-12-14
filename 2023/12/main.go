package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func Solve(file string, factor int) int {
	result := 0
	for _, line := range util.Lines(file) {
		result += Matches(line, factor)
	}
	return result
}

func Matches(line string, factor int) int {
	result := 0
	parts := strings.Split(line, " ")
	visual := unfold(parts[0], "?", factor)
	groups := parse(unfold(parts[1], ",", factor))

	ms := make([]map[int]int, len(visual)+1)
	for i := range ms {
		ms[i] = make(map[int]int)
	}
	ms[0][0] = 1

	for i := 0; i < len(visual); i++ {
		current := ms[i]
		for g, v := range current {
			remainingGroups := len(groups) - g

			var ns []string
			if remainingGroups == 0 {
				ns = []string{"."}
			} else {
				switch visual[i] {
				case '.':
					ns = []string{"."}
				case '#':
					ns = []string{"#"}
				case '?':
					ns = []string{".", "#"}
				default:
					panic("unexpected first char")
				}
			}

			for _, char := range ns {
				var addition string
				var nextGroups []int

				if char == "#" && remainingGroups > 0 {
					nextGroup := groups[g]
					nextGroups = groups[0 : g+1]
					addition = strings.Repeat(char, nextGroup)
					if remainingGroups > 1 {
						addition += "."
					}
				} else {
					nextGroups = groups[0:g]
					addition = char
				}

				value := i + len(addition)
				if value <= len(visual) && matches(addition, visual[i:value]) && (value < len(visual) || equal(nextGroups, groups)) {
					ms[value][len(nextGroups)] = ms[value][len(nextGroups)] + v
				}
			}
		}
	}

	for _, v := range ms[len(visual)] {
		result += v
	}

	return result
}

func matches(value string, regex string) bool {
	for i := 0; i < len(value); i++ {
		if regex[i] != '?' && value[i] != regex[i] {
			return false
		}
	}
	return true
}

func equal(as []int, bs []int) bool {
	if len(as) != len(bs) {
		return false
	}
	for i, a := range as {
		if bs[i] != a {
			return false
		}
	}
	return true
}

func parse(input string) []int {
	var result []int
	for _, n := range strings.Split(input, ",") {
		result = append(result, util.Number(n))
	}
	return result
}

func unfold(pattern, separator string, factor int) string {
	result := make([]string, factor)
	for i := 0; i < factor; i++ {
		result[i] = pattern
	}
	return strings.Join(result, separator)
}

func main() {
	fmt.Println(Solve("2023/12/input.txt", 1))
	fmt.Println(Solve("2023/12/input.txt", 5))
}
