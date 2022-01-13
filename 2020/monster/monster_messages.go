package monster

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
	"strings"
)

func solve(file string) int {
	text := util.Text(file)
	rs, ms := util.Split(text, "\n\n")
	rules := parseRules(rs)
	regexes := convertToRegexes(rules)

	messages := strings.Split(ms, "\n")
	result := 0
	for i := 1; i <= 10; i++ {
		pattern := fmt.Sprintf("^%s+%s{%d}%s{%d}$", regexes[42], regexes[42], i, regexes[31], i)
		re := regexp.MustCompile(pattern)
		for _, message := range messages {
			if re.MatchString(message) {
				result++
			}
		}
	}
	return result
}

func parseRules(input string) map[int]string {
	lines := strings.Split(input, "\n")
	rules := make(map[int]string, len(lines))
	for _, line := range lines {
		i, r := util.Split(line, ": ")
		rules[util.Number(i)] = r
	}
	return rules
}

func convertToRegexes(rules map[int]string) map[int]string {
	regexes := make(map[int]string, len(rules))
	parse(42, rules, regexes)
	parse(31, rules, regexes)
	return regexes
}

func parse(id int, rules map[int]string, regexes map[int]string) string {
	rule := rules[id]
	if strings.Contains(rule, "|") {
		left, right := util.Split(rule, " \\| ")
		lefts := ""
		for _, r := range strings.Split(left, " ") {
			n := util.Number(r)
			re := parse(n, rules, regexes)
			lefts += re
		}
		rights := ""
		for _, r := range strings.Split(right, " ") {
			n := util.Number(r)
			re := parse(n, rules, regexes)
			rights += re
		}
		regexes[id] = "(" + lefts + "|" + rights + ")"
		return regexes[id]
	}
	if regexp.MustCompile("[0-9]+( [0-9]+)*").MatchString(rule) {
		result := ""
		for _, r := range strings.Split(rule, " ") {
			n := util.Number(r)
			re := parse(n, rules, regexes)
			result += re
		}
		regexes[id] = result
		return result
	}
	if rule == "\"a\"" {
		return "a"
	}
	if rule == "\"b\"" {
		return "b"
	}
	panic("not implemented")
}
