package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type point struct {
	x, y int
}

func Solve(file string) int {
	x := 0
	y := 0
	border := []point{{x, y}}
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps := util.Number(parts[1])
		switch direction {
		case "R":
			for i := 1; i <= steps; i++ {
				x++
				border = append(border, point{x, y})
			}
		case "L":
			for i := 1; i <= steps; i++ {
				x--
				border = append(border, point{x, y})
			}
		case "D":
			for i := 1; i <= steps; i++ {
				y++
				border = append(border, point{x, y})
			}
		case "U":
			for i := 1; i <= steps; i++ {
				y--
				border = append(border, point{x, y})
			}
		default:
			panic("not implemented for " + direction)
		}
	}

	if border[len(border)-1] != border[0] {
		panic("how?")
	}
	border = append([]point{border[len(border)-2]}, border...)

	borderLine := make(map[point]string)
	for i := 1; i < len(border)-1; i++ {
		previous := border[i-1]
		self := border[i]
		next := border[i+1]
		if _, ok := borderLine[self]; ok {
			panic("already been here")
		}
		if previous.x == self.x {
			if self.x == next.x {
				borderLine[self] = "|"
			} else if previous.y < self.y && self.x < next.x {
				borderLine[self] = "L"
			} else if previous.y < self.y && self.x > next.x {
				borderLine[self] = "J"
			} else if previous.y > self.y && self.x < next.x {
				borderLine[self] = "F"
			} else if previous.y > self.y && self.x > next.x {
				borderLine[self] = "7"
			}
		} else if previous.y == self.y {
			if self.y == next.y {
				borderLine[self] = "-"
			} else if previous.x > self.x && self.y > next.y {
				borderLine[self] = "L"
			} else if previous.x > self.x && self.y < next.y {
				borderLine[self] = "F"
			} else if previous.x < self.x && self.y < next.y {
				borderLine[self] = "7"
			} else if previous.x < self.x && self.y > next.y {
				borderLine[self] = "J"
			}
		} else {
			panic("not implemented")
		}
	}

	minX := int(^uint(0) >> 1)
	minY := int(^uint(0) >> 1)
	maxX := 0
	maxY := 0
	for p := range borderLine {
		minX = min(minX, p.x)
		maxX = max(maxX, p.x)
		minY = min(minY, p.y)
		maxY = max(maxY, p.y)
	}

	m := make([][]string, maxY-minY+1)
	for i := range m {
		m[i] = make([]string, maxX-minX+1)
	}
	for k, v := range borderLine {
		m[k.y-minY][k.x-minX] = v
	}

	result := len(borderLine)

	for y, row := range m {
		in := false
		up := false
		down := false
		for x, e := range row {
			if borderLine[point{x + minX, y + minY}] != "" {
				if e == "F" || e == "7" || e == "|" {
					down = !down
				}
				if e == "L" || e == "J" || e == "|" {
					up = !up
				}
				if up && down {
					in = !in
					up = false
					down = false
				}
			} else if in {
				m[y][x] = "x"
				result++
			}
		}
	}
	fmt.Println(toString(m))

	return result
}

func toString(m [][]string) any {
	lines := make([]string, len(m))
	for y, row := range m {
		line := ""
		for _, c := range row {
			if c == "" {
				line += "."
			} else {
				line += c
			}
		}
		lines[y] = line
	}
	return strings.Join(lines, "\n")
}

func main() {
	fmt.Println(Solve("2023/18/input.txt"))
}
