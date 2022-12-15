package hill

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
)

type coordinate struct {
	x, y, z int
}

func path(file string) int {
	grid, end := parse(file)

	paths := make(map[coordinate][]coordinate)
	paths[end] = []coordinate{end}

	for {
		improvements := 0
		for k, v := range paths {
			for _, c := range around(k, grid) {
				if k.z-c.z <= 1 {
					existing, ok := paths[c]
					if !ok || len(existing) > len(v)+1 {
						improvements += 1
						paths[c] = append([]coordinate{c}, v...)
					}
				}
			}
		}
		if improvements == 0 {
			break
		}
	}

	var length []int
	for k, v := range paths {
		if k.z == 0 {
			length = append(length, len(v))
		}
	}
	sort.Ints(length)

	return length[0] - 1
}

func around(c coordinate, grid [][]coordinate) []coordinate {
	var result []coordinate
	if c.y-1 >= 0 {
		result = append(result, grid[c.y-1][c.x])
	}
	if c.y+1 < len(grid) {
		result = append(result, grid[c.y+1][c.x])
	}
	if c.x-1 >= 0 {
		result = append(result, grid[c.y][c.x-1])
	}
	if c.x+1 < len(grid[0]) {
		result = append(result, grid[c.y][c.x+1])
	}
	return result
}

func parse(file string) (grid [][]coordinate, end coordinate) {
	lines := util.Lines(file)
	grid = make([][]coordinate, len(lines))
	for y, line := range lines {
		grid[y] = make([]coordinate, len(line))
		for x, char := range []rune(line) {
			grid[y][x] = coordinate{x: x, y: y}
			switch char {
			case 'S':
				grid[y][x].z = 0
			case 'E':
				grid[y][x].z = 25
				end = grid[y][x]
			default:
				grid[y][x].z = int(char - 'a')
			}
		}
	}
	return grid, end
}

func Solve() {
	file := "2022/hill/input.txt"
	fmt.Println(path(file))
}
