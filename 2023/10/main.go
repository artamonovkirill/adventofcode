package main

import (
	"fmt"
	"github.com/advendofcode/util"
)

type Coordinate struct {
	x, y int
}
type Point struct {
	value                 string
	c                     Coordinate
	up, left, right, down bool
}

func LongestDistance(file string) int {
	start, points := parse(file)
	p := border(start, points)
	return len(p) / 2
}

func border(start Point, points [][]Point) map[Coordinate]bool {
	paths := make([][]Point, 2)
	var current []Point

	if start.up {
		current = append(current, points[start.c.y-1][start.c.x])
	}
	if start.down {
		current = append(current, points[start.c.y+1][start.c.x])
	}
	if start.left {
		current = append(current, points[start.c.y][start.c.x-1])
	}
	if start.right {
		current = append(current, points[start.c.y][start.c.x+1])
	}

	for {
		if len(current) != 2 {
			panic("expected current to have 2 elements, got" + fmt.Sprintln(current))
		}
		paths[0] = append(paths[0], current[0])
		paths[1] = append(paths[1], current[1])
		var next []Point
		for i, c := range current {
			if c.up {
				n := points[c.c.y-1][c.c.x]
				if n.value != "S" && !contains(paths[i], n) {
					next = append(next, n)
				}
			}
			if c.down {
				n := points[c.c.y+1][c.c.x]
				if n.value != "S" && !contains(paths[i], n) {
					next = append(next, n)
				}
			}
			if c.left {
				n := points[c.c.y][c.c.x-1]
				if n.value != "S" && !contains(paths[i], n) {
					next = append(next, n)
				}
			}
			if c.right {
				n := points[c.c.y][c.c.x+1]
				if n.value != "S" && !contains(paths[i], n) {
					next = append(next, n)
				}
			}
		}
		if next[0] == next[1] {
			result := make(map[Coordinate]bool)
			result[start.c] = true
			result[next[0].c] = true
			for _, ps := range paths {
				for _, p := range ps {
					result[p.c] = true
				}
			}
			return result
		}
		current = next
	}
}

func parse(file string) (Point, [][]Point) {
	lines := util.Lines(file)
	var start Point

	m := make([][]Point, len(lines))

	for y, line := range lines {
		m[y] = make([]Point, len(line))
		for x, c := range line {
			p := Point{c: Coordinate{x, y}, value: string(c)}
			switch c {
			case 'S':
				p.left = connectsLeft(x, y, lines)
				p.right = connectsRight(x, y, lines)
				p.up = connectsUp(x, y, lines)
				p.down = connectsDown(x, y, lines)
				start = p
			case '-':
				p.left = connectsLeft(x, y, lines)
				p.right = connectsRight(x, y, lines)
			case 'L':
				p.up = connectsUp(x, y, lines)
				p.right = connectsRight(x, y, lines)
			case '|':
				p.up = connectsUp(x, y, lines)
				p.down = connectsDown(x, y, lines)
			case 'F':
				p.right = connectsRight(x, y, lines)
				p.down = connectsDown(x, y, lines)
			case '7':
				p.left = connectsLeft(x, y, lines)
				p.down = connectsDown(x, y, lines)
			case 'J':
				p.left = connectsLeft(x, y, lines)
				p.up = connectsUp(x, y, lines)
			case '.':
			default:
				panic("not implemented for " + fmt.Sprintln(x, y, string(c)))
			}
			m[y][x] = p
		}
	}
	return start, m
}

func contains(ps []Point, point Point) bool {
	for _, p := range ps {
		if p == point {
			return true
		}
	}
	return false
}

func connectsDown(x, y int, lines []string) bool {
	if y < len(lines)-1 {
		down := lines[y+1][x]
		if down == '|' || down == 'J' || down == 'L' || down == 'S' {
			return true
		}
	}
	return false
}

func connectsUp(x, y int, lines []string) bool {
	if y > 0 {
		up := lines[y-1][x]
		if up == '|' || up == '7' || up == 'F' || up == 'S' {
			return true
		}
	}
	return false
}

func connectsRight(x, y int, lines []string) bool {
	if x < len(lines[y])-1 {
		right := lines[y][x+1]
		if right == 'J' || right == '-' || right == '7' || right == 'S' {
			return true
		}
	}
	return false
}

func connectsLeft(x, y int, lines []string) bool {
	if x > 0 {
		left := lines[y][x-1]
		if left == 'L' || left == '-' || left == 'F' || left == 'S' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(LongestDistance("2023/10/input.txt"))
	fmt.Println(Insides("2023/10/input.txt"))
}

func Insides(file string) int {
	start, m := parse(file)
	p := border(start, m)
	result := 0
	for _, row := range m {
		in := false
		up := false
		down := false
		for _, e := range row {
			if p[e.c] {
				if e.down {
					down = !down
				}
				if e.up {
					up = !up
				}
				if up && down {
					in = !in
					up = false
					down = false
				}
			} else if in {
				result++
			}
		}
	}
	return result
}
