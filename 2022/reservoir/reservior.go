package reservoir

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type coordinate struct {
	x, y int
}

type material int

const (
	air  = 0
	rock = 1
	sand = 2
)

func simulate(file string, bottom bool) int {
	cave := make(map[coordinate]material)
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			start := parse(parts[i])
			stop := parse(parts[i+1])
			if start.x != stop.x && start.y != stop.y {
				panic("not implemented")
			}
			if start.x == stop.x && start.y < stop.y {
				for y := start.y; y <= stop.y; y++ {
					cave[coordinate{start.x, y}] = rock
				}
			} else if start.x == stop.x && start.y > stop.y {
				for y := stop.y; y <= start.y; y++ {
					cave[coordinate{start.x, y}] = rock
				}
			} else if start.y == stop.y && start.x > stop.x {
				for x := stop.x; x <= start.x; x++ {
					cave[coordinate{x, start.y}] = rock
				}
			} else if start.y == stop.y && start.x < stop.x {
				for x := start.x; x <= stop.x; x++ {
					cave[coordinate{x, start.y}] = rock
				}
			} else {
				panic("not implemented for " + parts[i] + " and " + parts[i+1])
			}
		}
	}
	minimum := 2
	for k := range cave {
		if k.y > minimum {
			minimum = k.y + 2
		}
	}

	s := coordinate{500, 0}
	for i := 0; i < 10000000; i++ {
		if cave[coordinate{s.x, s.y + 1}] == air {
			s = coordinate{s.x, s.y + 1}
			if !bottom && s.y > minimum {
				return count(cave)
			}
		} else if cave[coordinate{s.x - 1, s.y + 1}] == air {
			s = coordinate{s.x - 1, s.y + 1}
		} else if cave[coordinate{s.x + 1, s.y + 1}] == air {
			s = coordinate{s.x + 1, s.y + 1}
		} else {
			cave[s] = sand
			if bottom && cave[coordinate{500, 0}] == sand {
				fmt.Println(toString(cave))
				return count(cave)
			}
			s = coordinate{500, 0}
		}
		if bottom && s.y == minimum-1 {
			cave[s] = sand
			if bottom && cave[coordinate{500, 0}] == sand {
				fmt.Println(toString(cave))
				return count(cave)
			}
			s = coordinate{500, 0}
		}
	}
	panic("no solution found")
}

func count(cave map[coordinate]material) int {
	result := 0
	for _, v := range cave {
		if v == sand {
			result += 1
		}
	}
	return result
}

func toString(cave map[coordinate]material) string {
	minY := 9999
	maxY := 0
	minX := 9999
	maxX := 0
	for k := range cave {
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
	}
	result := make([][]rune, maxY-minY+1)
	for y := 0; y <= maxY-minY; y++ {
		result[y] = make([]rune, maxX-minX+1)
		for x := 0; x <= maxX-minX; x++ {
			result[y][x] = '.'
		}
	}
	for k, v := range cave {
		if v == rock {
			result[k.y-minY][k.x-minX] = '#'
		} else {
			result[k.y-minY][k.x-minX] = 'o'
		}
	}
	lines := make([]string, len(result))
	for i, row := range result {
		lines[i] = string(row)
	}
	return strings.Join(lines, "\n")
}

func parse(input string) coordinate {
	parts := strings.Split(input, ",")
	return coordinate{util.Number(parts[0]), util.Number(parts[1])}
}

func Solve() {
	file := "2022/reservoir/input.txt"
	fmt.Println(simulate(file, false))
	fmt.Println(simulate(file, true))
}
