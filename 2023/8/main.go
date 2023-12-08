package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
)

type Location struct {
	left, right string
}

var r = regexp.MustCompile("[0-9A-Z]{3}")

func Solve(file string) int {
	lines := util.Lines(file)
	commands := lines[0]

	routes := parse(lines[2:])

	steps := 0
	current := "AAA"

	for {
		next := commands[steps%len(commands)]
		if next == 'L' {
			current = routes[current].left
		} else {
			current = routes[current].right
		}
		steps++
		if current == "ZZZ" {
			return steps
		}
	}
}

func parse(lines []string) map[string]Location {
	routes := make(map[string]Location)
	for _, line := range lines {
		parts := r.FindAllString(line, -1)
		routes[parts[0]] = Location{
			parts[1],
			parts[2],
		}
	}
	return routes
}

func main() {
	fmt.Println(Solve("2023/8/input.txt"))
	fmt.Println(Solve2("2023/8/input.txt"))
}

func Solve2(file string) int {
	lines := util.Lines(file)
	commands := lines[0]

	routes := parse(lines[2:])

	var starts []string
	for name := range routes {
		if name[2] == 'A' {
			starts = append(starts, name)
		}
	}

	cycles := make(map[int]bool)
	for _, start := range starts {
		current := start
		steps := 0
		var results []int
		for {
			next := commands[steps%len(commands)]
			if next == 'L' {
				current = routes[current].left
			} else {
				current = routes[current].right
			}
			steps++
			if current[2] == 'Z' {
				results = append(results, steps)
			}
			if len(results) > 10 {
				for _, r := range results[1:] {
					if r%results[0] != 0 {
						panic("no cycle found")
					}
				}
				cycles[results[0]] = true
				break
			}
		}
	}

	var step int
	for k := range cycles {
		step = k
		break
	}

steps:
	for i := step; i < 296596217528783955; i += step {
		for k := range cycles {
			if i%k != 0 {
				continue steps
			}
		}
		return i
	}
	panic("no solution found")
}
