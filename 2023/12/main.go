package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
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
	var rvs []*regexp.Regexp
	for i := 0; i < len(visual); i++ {
		rvs = append(rvs, regexp.MustCompile(visual[0:i]))
	}
	rts := make(map[int]*regexp.Regexp)
	for i := 0; i < len(visual); i++ {
		v := visual[i:]
		rts[len(v)] = regexp.MustCompile(v)
	}

	for len(current) > 0 {
		var next []candidate
		mc := make(map[hash]int)
		for _, c := range current {
			remainingGroups := len(groups) - c.groups

			var ns []string
			if remainingGroups == 0 {
				ns = []string{"x"}
			} else if len(c.value) > 0 && c.value[len(c.value)-1] == '#' {
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
					if rts[len(addition)].MatchString(addition) && equal(nextGroups, groups) {
						result++
					}
				} else {
					value := c.value + addition
					if len(value) < len(visual) {
						nc := candidate{
							value:  value,
							groups: len(nextGroups),
						}
						if rvs[len(nc.value)].MatchString(nc.value) {
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
