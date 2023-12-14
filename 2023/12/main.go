package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type candidate struct {
	value  string
	groups int
}

func Solve(file string, factor int) int {
	result := 0
	for i, line := range util.Lines(file) {
		fmt.Println("processing", i)
		result += Matches(line, factor)
	}
	return result
}

type hash struct {
	groups, length int
}

func Matches(line string, factor int) int {
	result := 0
	parts := strings.Split(line, " ")
	visual := regex(unfold(parts[0], "?", factor))
	groups := parse(unfold(parts[1], ",", factor))

	current := []candidate{
		{},
	}
	for len(current) > 0 {
		var next []candidate
		mc := make(map[hash]int)
		for _, c := range current {
			remainingGroups := len(groups) - c.groups

			var ns []string
			if remainingGroups == 0 {
				ns = []string{"x"}
			} else {
				switch visual[len(c.value)] {
				case 'x':
					ns = []string{"x"}
				case '#':
					ns = []string{"#"}
				case '.':
					ns = []string{"x", "#"}
				default:
					panic("unexpected first char")
				}
			}

			for _, char := range ns {
				var addition string
				var nextGroups []int

				if char == "#" && remainingGroups > 0 {
					nextGroup := groups[c.groups]
					nextGroups = groups[0 : c.groups+1]
					addition = strings.Repeat(char, nextGroup)
					if remainingGroups > 1 {
						addition += "x"
					}
				} else {
					nextGroups = groups[0:c.groups]
					addition = char
				}

				if len(addition)+len(c.value) == len(visual) {
					if matches(addition, visual[len(c.value):]) && equal(nextGroups, groups) {
						result++
					}
				} else {
					value := c.value + addition
					if len(value) < len(visual) {
						nc := candidate{
							value:  value,
							groups: len(nextGroups),
						}
						if matches(nc.value, visual[0:len(nc.value)]) {
							h := hash{nc.groups, len(nc.value)}
							mc[h] = mc[h] + 1
							next = append(next, nc)
						}
					}
				}
			}
		}
		if len(next) > 100_000 {
			fmt.Println("  ", len(next), "to go,", "unique", len(mc))
		}
		current = next
	}

	return result
}

func matches(value string, regex string) bool {
	for i := 0; i < len(value); i++ {
		if regex[i] != '.' && value[i] != regex[i] {
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

func regex(input string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			input,
			".", "x"),
		"?", ".")
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
