package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func Solve(file string) int {
	result := 0

	patterns := strings.Split(util.Text(file), "\n\n")
p:
	for _, pattern := range patterns {
		matrix := parse(pattern)
		originalH := firstOr(reflection(matrix), 0)
		originalV := firstOr(reflection(transpose(matrix)), 0)

		for y, line := range matrix {
			for x, c := range line {
				if c == '#' {
					matrix[y][x] = '.'
				} else {
					matrix[y][x] = '#'
				}

				hrz := reflection(matrix)
				vrt := reflection(transpose(matrix))
				if len(hrz) > 0 || len(vrt) > 0 {
					addition := 0
					for _, h := range hrz {
						if h != originalH {
							addition += 100 * h
						}
					}
					for _, v := range vrt {
						if v != originalV {
							addition += v
						}
					}
					if addition > 0 {
						result += addition
						continue p
					}
				}
				matrix[y][x] = c
				if toString(matrix) != pattern {
					panic("drift detected")
				}
			}
		}
		panic("no smudge found")
	}

	return result
}

func firstOr(matches []int, def int) int {
	if len(matches) == 0 {
		return def
	}
	return matches[0]
}

func toString(m [][]int32) string {
	lines := make([]string, len(m))
	for i, l := range m {
		lines[i] = string(l)
	}
	return strings.Join(lines, "\n")
}

func reflection(matrix [][]int32) []int {
	var matches []int
outer:
	for i := 1; i < len(matrix); i++ {
		top := matrix[0:i]
		bottom := matrix[i:]
		length := min(len(bottom), len(top))
		bottom = bottom[0:length]
		for j := 0; j < length; j++ {
			for x, c := range bottom[j] {
				if c != top[len(top)-j-1][x] {
					continue outer
				}
			}
		}
		matches = append(matches, len(top))
	}
	return matches
}

func parse(pattern string) [][]int32 {
	lines := strings.Split(pattern, "\n")
	result := make([][]int32, len(lines))
	for y, line := range lines {
		result[y] = make([]int32, len(line))
		for x, c := range line {
			result[y][x] = c
		}
	}
	return result
}

func transpose(m [][]int32) [][]int32 {
	result := make([][]int32, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		result[i] = make([]int32, len(m))
	}
	for y, line := range m {
		for x, c := range line {
			result[x][y] = c
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/13/input.txt"))
}
