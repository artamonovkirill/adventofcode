package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type Coordinate struct {
	x, y int
}

func Solve(file string, factor int) int {
	lines := util.Lines(file)

	var galaxies []Coordinate
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, Coordinate{x, y})
			}
		}
	}

	var pairs [][]Coordinate
	for i, first := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			second := galaxies[j]
			pairs = append(pairs, []Coordinate{first, second})
		}
	}

	xs := columns(lines)
	ys := rows(lines)
	result := 0
	for _, pair := range pairs {
		distance := 0
		for x := min(pair[0].x, pair[1].x); x <= max(pair[0].x, pair[1].x); x++ {
			if xs[x] == true {
				result += factor
			} else {
				result += 1
			}
		}
		for y := min(pair[0].y, pair[1].y); y <= max(pair[0].y, pair[1].y); y++ {
			if ys[y] == true {
				result += factor
			} else {
				result += 1
			}
		}
		result += distance - 2
	}
	return result
}

func expand(lines []string, factor int) []string {
	var result []string

	cs := columns(lines)
	rs := rows(lines)

	for y, line := range lines {
		var newLine []int32
		for i, c := range line {
			if cs[i] {
				for j := 0; j < factor-1; j++ {
					newLine = append(newLine, c)
				}
			}
			newLine = append(newLine, c)
		}
		str := string(newLine)
		if rs[y] {
			for j := 0; j < factor-1; j++ {
				result = append(result, str)
			}
		}
		result = append(result, str)
	}

	return result
}

func columns(lines []string) map[int]bool {
	result := make(map[int]bool)
outer:
	for i, c := range lines[0] {
		if c == '.' {
			for _, line := range lines[1:] {
				if line[i] != '.' {
					continue outer
				}
			}
			result[i] = true
		}
	}
	return result
}

func rows(lines []string) map[int]bool {
	result := make(map[int]bool)
	for i, line := range lines {
		if !strings.Contains(line, "#") {
			result[i] = true
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/11/input.txt", 2))
	fmt.Println(Solve("2023/11/input.txt", 1000000))
}
