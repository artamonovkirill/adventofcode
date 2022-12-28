package flow

import (
	"fmt"
	"github.com/advendofcode/util"
	"path"
	"runtime"
	"strings"
)

type coordinate struct {
	x, y int
}

func proceed(file string, count int) int {
	rocks := parse()
	pattern := util.Text(file)
	cave := [][]bool{make([]bool, 7)}
	for x := 0; x < 7; x++ {
		cave[0][x] = true
	}

	j := -1
outer:
	for n := 0; n < count; n++ {
		height := len(cave)
		i := n % len(rocks)
		rock := rocks[i]
		var moves []rune
		rock = update(rock, func(c coordinate) coordinate {
			return coordinate{c.x, c.y + 3 + height}
		})
		for {
			j++
			direction := pattern[j%len(pattern)]
			moves = append(moves, rune(direction))
			switch direction {
			case '>':
				if right(rock) < 6 {
					movable := true
					for _, c := range rock {
						if occupied(cave, coordinate{c.x + 1, c.y}) {
							movable = false
						}
					}
					if movable {
						rock = update(rock, func(c coordinate) coordinate {
							return coordinate{c.x + 1, c.y}
						})
					}
				}
			case '<':
				if left(rock) > 0 {
					movable := true
					for _, c := range rock {
						if occupied(cave, coordinate{c.x - 1, c.y}) {
							movable = false
						}
					}
					if movable {
						rock = update(rock, func(c coordinate) coordinate {
							return coordinate{c.x - 1, c.y}
						})
					}
				}
			default:
				panic("not implemented for " + string(direction))
			}
			for _, c := range rock {
				if occupied(cave, coordinate{c.x, c.y - 1}) {
					for _, c := range rock {
						if c.y >= len(cave) {
							cave = append(cave, make([]bool, 7))
						}
						cave[c.y][c.x] = true
					}
					rock = rocks[i]

					continue outer
				}
			}
			rock = update(rock, func(c coordinate) coordinate {
				return coordinate{c.x, c.y - 1}
			})
		}
	}
	return len(cave) - 1
}

func occupied(cave [][]bool, c coordinate) bool {
	if len(cave) <= c.y {
		return false
	}
	row := cave[c.y]
	return len(row) > c.x && row[c.x]
}

func update(rock []coordinate, f func(c coordinate) coordinate) []coordinate {
	result := make([]coordinate, len(rock))
	for i, c := range rock {
		result[i] = f(c)
	}
	return result
}

func right(cs []coordinate) int {
	result := 0
	for _, c := range cs {
		if c.x > result {
			result = c.x
		}
	}
	return result
}

func left(cs []coordinate) int {
	result := 7
	for _, c := range cs {
		if c.x < result {
			result = c.x
		}
	}
	return result
}

func parse() [][]coordinate {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot read file")
	}
	var result [][]coordinate
	y := 0
	var rock []coordinate
	for _, line := range util.Lines(path.Dir(filename) + "/rocks.txt") {
		if line == "" {
			result = append(result, rock)
			y = 0
			rock = []coordinate{}
			continue
		}
		for x, c := range []rune(line) {
			if c == '#' {
				rock = append(rock, coordinate{x + 2, y})
			}
		}
		y++
	}
	result = append(result, rock)
	return result
}

func toString(cave [][]bool, rock []coordinate) string {
	maxY := len(cave)
	for _, c := range rock {
		if c.y > maxY {
			maxY = c.y
		}
	}
	result := make([][]rune, maxY+1)
	for y := range result {
		result[y] = make([]rune, 7)
		for x := range result[y] {
			result[y][x] = '.'
		}
	}

	for y, row := range cave {
		for x, o := range row {
			if o {
				result[maxY-y][x] = '#'
			}
		}
	}
	for _, c := range rock {
		result[maxY-c.y][c.x] = '@'
	}

	lines := make([]string, len(result))
	for i, row := range result {
		lines[i] = string(row)
	}
	return strings.Join(lines, "\n")
}

func rockToString(rock []coordinate) string {
	return toString([][]bool{}, rock)
}

func Solve() {
	fmt.Println(proceed("2022/flow/input.txt", 2022))
	//fmt.Println(predict("2022/flow/input.txt", 1000000000000))
}

func predict(file string, n int) int {
	var block int
	var blockHeight int
	var start int
	if strings.Contains(file, "example.txt") {
		// for 10000 n, there are 43 levels / 25 blocks before repetitions
		// the rest has 285 repetitions of 35 blocks / 53 levels
		start = 25
		block = 35
		blockHeight = 53
	} else {
		// height 7925
		// for 10000 n, there are x levels / x blocks before repetitions
		// the rest has 21 repetitions of x blocks / x levels
		blockHeight = 2777
		panic("not implemented")
	}
	rest := n - start
	repeats := rest/block - 1
	tail := rest % block
	calc := proceed(file, start+block+tail)
	return repeats*blockHeight + calc
}
