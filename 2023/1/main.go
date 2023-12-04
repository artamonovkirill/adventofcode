package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
)

var r = regexp.MustCompile("[1-9]|one|two|three|four|five|six|seven|eight|nine")

func Solve(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		first := r.FindString(line)
		var last string
		value := parse(first) * 10
		for i := 1; i <= len(line); i++ {
			tail := line[len(line)-i:]
			match := r.FindString(tail)
			if match != "" {
				last = match
				break
			}
		}
		value += parse(last)
		result += value
		fmt.Println(line, "- first:", first, "last:", last, "value:", value, "result:", result)
	}
	return result
}

func parse(input string) int {
	switch input {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		value := util.Number(input)
		return value
	}
}

func main() {
	fmt.Println(Solve("2023/1/input.txt"))
}
