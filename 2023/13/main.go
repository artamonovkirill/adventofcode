package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func Solve(file string) int {
	result := 0

	patterns := strings.Split(util.Text(file), "\n\n")
	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
	outer:
		for i := 1; i < len(lines); i++ {
			top := lines[0:i]
			bottom := lines[i:]
			length := min(len(bottom), len(top))
			bottom = bottom[0:length]
			for j := 0; j < length; j++ {
				if bottom[j] != top[len(top)-j-1] {
					continue outer
				}
			}
			result += 100 * len(top)
		}

		columns := transpose(strings.Split(pattern, "\n"))
	outer2:
		for i := 1; i < len(columns); i++ {
			top := columns[0:i]
			bottom := columns[i:]
			length := min(len(bottom), len(top))
			bottom = bottom[0:length]
			for j := 0; j < length; j++ {
				if bottom[j] != top[len(top)-j-1] {
					continue outer2
				}
			}
			result += len(top)
		}
	}

	return result
}

func transpose(lines []string) []string {
	transposed := make([][]int32, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		transposed[i] = make([]int32, len(lines))
	}
	for y, line := range lines {
		for x, c := range line {
			transposed[x][y] = c
		}
	}
	result := make([]string, len(transposed))
	for i, line := range transposed {
		result[i] = string(line)
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/13/input.txt"))
}
