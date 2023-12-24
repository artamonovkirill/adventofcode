package main

import (
	"fmt"
	"github.com/advendofcode/util"
)

type point struct {
	x, y int
}

func Solve(file string, steps int) int {
	lines := util.Lines(file)
	rocks := make([][]bool, len(lines))
	current := make(map[point]bool)
	for y, line := range lines {
		rocks[y] = make([]bool, len(line))
		for x, c := range line {
			if c == 'S' {
				current[point{x, y}] = true
			}
			rocks[y][x] = c == '#'
		}
	}
	for i := 0; i < steps; i++ {
		next := make(map[point]bool)
		for c := range current {
			if c.x > 0 {
				p := point{c.x - 1, c.y}
				if !rocks[p.y][p.x] {
					next[p] = true
				}
			}
			if c.x < len(rocks[0])-1 {
				p := point{c.x + 1, c.y}
				if !rocks[p.y][p.x] {
					next[p] = true
				}
			}
			if c.y > 0 {
				p := point{c.x, c.y - 1}
				if !rocks[p.y][p.x] {
					next[p] = true
				}
			}
			if c.y < len(rocks)-1 {
				p := point{c.x, c.y + 1}
				if !rocks[p.y][p.x] {
					next[p] = true
				}
			}
		}
		current = next
	}
	return len(current)
}

func main() {
	fmt.Println(Solve("2023/21/input.txt", 64))
}
