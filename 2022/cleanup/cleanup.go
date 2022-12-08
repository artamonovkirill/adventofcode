package cleanup

import (
	"fmt"
	"strings"

	"github.com/advendofcode/util"
)

type Range struct {
	start int
	end   int
}

func (r Range) overlaps(other Range) bool {
	return (r.start <= other.start && other.start <= r.end) ||
		(r.start <= other.end && other.end <= r.end) ||
		(other.start <= r.start && r.start <= other.end) ||
		(other.start <= r.end && r.end <= other.end)
}

func count(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, ",")
		left, right := parse(parts[0]), parse(parts[1])
		if left.overlaps(right) {
			result++
		}
	}
	return result
}

func parse(input string) Range {
	parts := strings.Split(input, "-")
	return Range{start: util.Number(parts[0]), end: util.Number(parts[1])}
}

func Solve() {
	file := "2022/cleanup/input.txt"
	fmt.Println(count(file))
}
