package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type rock struct {
	x, y    int
	movable bool
}

func (r *rock) isMovable() bool {
	return r.movable
}

func Solve(file string) int {
	lines := util.Lines(file)
	m := make([][]*rock, len(lines))
	for y, line := range lines {
		m[y] = make([]*rock, len(line))
		for x, c := range line {
			if c == 'O' {
				m[y][x] = &rock{x: x, y: y, movable: true}
			} else if c == '#' {
				m[y][x] = &rock{x: x, y: y}
			}
		}
	}

outer:
	for {
		for y := 1; y < len(m); y++ {
			restart := false
			for x, e := range m[y] {
				if e != nil && e.isMovable() && m[y-1][x] == nil {
					restart = true
					m[y-1][x] = e
					m[y][x] = nil
				}
			}
			if restart {
				continue outer
			}
		}
		break outer
	}

	result := 0
	for i, row := range m {
		for _, r := range row {
			if r != nil && r.isMovable() {
				result += len(m) - i
			}
		}
	}

	return result
}

func toString(m [][]*rock) string {
	lines := make([]string, len(m))
	for i, row := range m {
		line := ""
		for _, r := range row {
			if r == nil {
				line += "."
			} else if r.isMovable() {
				line += "O"
			} else {
				line += "#"
			}
		}
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

func main() {
	fmt.Println(Solve("2023/14/input.txt"))
}
