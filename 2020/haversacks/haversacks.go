package haversacks

import (
	"github.com/advendofcode/util"
	"regexp"
	"strings"
)

func solve(file string) int {
	lines := util.Lines(file)
	bags := parse(lines)
	for container, contents := range bags {
		if container == "shiny gold" {
			return count(bags, contents)
		}
	}
	return 0
}

func parse(lines []string) map[string]map[string]int {
	bags := make(map[string]map[string]int, len(lines))
	for _, line := range lines {
		c, cs := util.Split(line, " contain ")
		container := strings.ReplaceAll(c, " bags", "")
		contents := strings.Split(cs, ", ")
		counts := make(map[string]int, len(contents))
		for _, content := range contents {
			if content[0:2] != "no" {
				count, color := util.Split(content, " ")
				color = regexp.MustCompile("[0-9]+ ").ReplaceAllString(color, "")
				color = regexp.MustCompile(" bags?[.]?").ReplaceAllString(color, "")
				counts[color] = util.Number(count)
			}
		}
		bags[container] = counts
	}
	return bags
}

func count(bags map[string]map[string]int, counts map[string]int) int {
	result := 0
	for color, c := range counts {
		result += c
		result += c * count(bags, bags[color])
	}
	return result
}
