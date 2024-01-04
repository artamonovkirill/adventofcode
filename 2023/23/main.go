package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"maps"
)

type point struct {
	x, y int
}

type path struct {
	visited map[point]bool
	current point
}

func Solve(file string) int {
	trails := parse(file)

	var starts []point
	for x, c := range trails[0] {
		if c == '.' {
			starts = append(starts, point{x, 0})
		}
	}
	if len(starts) > 1 {
		panic("where to start?")
	}

	var ends []point
	for x, c := range trails[len(trails)-1] {
		if c == '.' {
			ends = append(ends, point{x, len(trails) - 1})
		}
	}
	if len(ends) > 1 {
		panic("where to end?")
	}
	end := ends[0]

	start := path{current: starts[0], visited: make(map[point]bool)}
	start.visited[starts[0]] = true
	current := []path{start}
	result := 0
	for len(current) > 0 {
		var next []path
		for _, c := range current {
			for _, p := range []point{
				{c.current.x, c.current.y - 1},
				{c.current.x, c.current.y + 1},
				{c.current.x - 1, c.current.y},
				{c.current.x + 1, c.current.y},
			} {
				if p.x < 0 || p.y < 0 || p.x >= len(trails[0]) || p.y >= len(trails) {
					continue
				}
				if p == end {
					if len(c.visited)+1 > result {
						result = len(c.visited) + 1
					}
				}
				switch trails[p.y][p.x] {
				case '#':
					continue
				case '>':
					extended := path{
						current: p,
						visited: maps.Clone(c.visited),
					}
					extended.visited[p] = true
					slipped := point{p.x + 1, p.y}
					if !extended.visited[slipped] {
						extended.current = slipped
						extended.visited[slipped] = true
						next = append(next, extended)
					}
				case '<':
					extended := path{
						current: p,
						visited: maps.Clone(c.visited),
					}
					extended.visited[p] = true
					slipped := point{p.x - 1, p.y}
					if !extended.visited[slipped] {
						extended.current = slipped
						extended.visited[slipped] = true
						next = append(next, extended)
					}
				case '^':
					extended := path{
						current: p,
						visited: maps.Clone(c.visited),
					}
					extended.visited[p] = true
					slipped := point{p.x, p.y - 1}
					if !extended.visited[slipped] {
						extended.current = slipped
						extended.visited[slipped] = true
						next = append(next, extended)
					}
				case 'v':
					extended := path{
						current: p,
						visited: maps.Clone(c.visited),
					}
					extended.visited[p] = true
					slipped := point{p.x, p.y + 1}
					if !extended.visited[slipped] {
						extended.current = slipped
						extended.visited[slipped] = true
						next = append(next, extended)
					}
				default:
					if !c.visited[p] {
						extended := path{
							current: p,
							visited: maps.Clone(c.visited),
						}
						extended.visited[p] = true
						next = append(next, extended)
					}
				}
			}
		}
		current = next
	}

	return result - 1
}

func parse(file string) [][]int32 {
	lines := util.Lines(file)
	result := make([][]int32, len(lines))
	for y, line := range lines {
		result[y] = make([]int32, len(line))
		for x, c := range line {
			result[y][x] = c
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/23/input.txt"))
}
