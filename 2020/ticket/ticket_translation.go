package ticket

import (
	"fmt"
	"github.com/advendofcode/util"
	"reflect"
	"regexp"
	"strings"
)

func invalidTickets(file string) int {
	text := util.Text(file)
	chunks := strings.Split(text, "\n\n")
	rules := parseRules(chunks[0])
	tickets := parseTickets(chunks[2])

	result := 0
	for _, ticket := range tickets {
	ticket:
		for _, v := range ticket {
			for _, rule := range rules {
				if rule(v) {
					continue ticket
				}
			}
			result += v
		}
	}
	return result
}

func solve(file string) map[string]int {
	text := util.Text(file)
	chunks := strings.Split(text, "\n\n")
	rules := parseRules(chunks[0])
	ticket := parseTickets(chunks[1])[0]
	tickets := parseTickets(chunks[2])

	validTickets := valid(tickets, rules)

	values := transpose(validTickets)

	candidates := findCandidates(values, rules)
	fields := make(map[string]int, len(tickets[0]))

	for len(fields) < len(ticket) {
		next, i := find(candidates)
		fields[next] = i
		candidates = remove(candidates, next)
	}

	result := make(map[string]int, len(fields))
	for t, i := range fields {
		result[t] = ticket[i]
	}
	return result
}

func remove(candidates map[int][]string, value string) map[int][]string {
	result := make(map[int][]string, len(candidates)-1)
	for i, types := range candidates {
		if !reflect.DeepEqual(types, []string{value}) {
			var filtered []string
			for _, t := range types {
				if t != value {
					filtered = append(filtered, t)
				}
			}
			result[i] = filtered
		}
	}
	return result
}

func find(candidates map[int][]string) (string, int) {
	for i, types := range candidates {
		if len(types) == 1 {
			return types[0], i
		}
	}
	return "", 0
}

func findCandidates(values [][]int, rules map[string]func(int) bool) map[int][]string {
	result := make(map[int][]string, len(values))
	for i, vs := range values {
	rule:
		for t, rule := range rules {
			for _, v := range vs {
				if !rule(v) {
					continue rule
				}
			}
			result[i] = append(result[i], t)
		}
	}
	return result
}

func valid(tickets [][]int, rules map[string]func(int) bool) [][]int {
	var result [][]int
ticket:
	for _, ticket := range tickets {
		for _, v := range ticket {
			if !check(v, rules) {
				continue ticket
			}
		}
		result = append(result, ticket)
	}
	return result
}

func transpose(input [][]int) [][]int {
	result := make([][]int, len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if result[j] == nil {
				result[j] = make([]int, len(input))
			}
			result[j][i] = input[i][j]
		}
	}
	return result
}

func check(value int, rules map[string]func(int) bool) bool {
	for _, rule := range rules {
		if rule(value) {
			return true
		}
	}
	return false
}

func parseTickets(input string) [][]int {
	lines := strings.Split(input, "\n")[1:]
	result := make([][]int, len(lines))
	for i, line := range lines {
		ns := strings.Split(line, ",")
		r := make([]int, len(ns))
		for j, n := range ns {
			r[j] = util.Number(n)
		}
		result[i] = r
	}
	return result
}

func parseRules(input string) map[string]func(int) bool {
	lines := strings.Split(input, "\n")
	result := make(map[string]func(int) bool, len(lines))
	for _, line := range lines {
		t, ranges := util.Split(line, ": ")
		first, second := util.Split(ranges, " or ")
		result[t] = func(i int) bool {
			return parseRange(first)(i) || parseRange(second)(i)
		}
	}
	return result
}

func parseRange(input string) func(int) bool {
	low, high := util.Split(input, "-")
	return func(i int) bool {
		return util.Number(low) <= i && i <= util.Number(high)
	}
}

func Solve() {
	input := "2020/ticket/puzzle.txt"
	fmt.Println(invalidTickets(input))
	fields := solve(input)
	result := 1
	for t, v := range fields {
		if regexp.MustCompile("^departure.*").MatchString(t) {
			result *= v
		}
	}
	fmt.Println(result)
}
