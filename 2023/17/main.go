package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type point struct {
	x, y int
}

type edge struct {
	cost  int
	value string
}

type path struct {
	end   point
	value string
}

func Solve(file string) int {
	vs := values(file)

	g := graph(vs)

	best := make(map[path]int)
	var current []path
	for p, e := range g[0][0] {
		newPath := path{
			end: p, value: e.value,
		}
		best[newPath] = e.cost
		current = append(current, newPath)
	}
	for {
		var next []path
		for _, p := range current {
			for newEnd, e := range g[p.end.y][p.end.x] {
				currentDirection := p.value[len(p.value)-1]
				newCost := best[p] + e.cost
				newPath := path{
					end:   newEnd,
					value: e.value,
				}
				newDirection := e.value[len(e.value)-1]
				if newDirection == '>' || newDirection == '<' {
					if currentDirection == '^' || currentDirection == 'v' {
						if b, ok := best[newPath]; !ok || b > newCost {
							best[newPath] = newCost
							next = append(next, newPath)
						}
					}
				} else {
					if currentDirection == '<' || currentDirection == '>' {
						if b, ok := best[newPath]; !ok || b > newCost {
							best[newPath] = newCost
							next = append(next, newPath)
						}
					}
				}
			}
		}
		if len(next) == 0 {
			result := int(^uint(0) >> 1)
			for p, c := range best {
				if p.end.x == len(vs[0])-1 && p.end.y == len(vs)-1 {
					if result > c {
						result = c
					}
				}
			}
			return result
		}
		current = next
	}
}

type direction struct {
	fn    func(p point) point
	value string
}

var directions = []direction{
	{func(p point) point {
		return point{p.x + 1, p.y}
	}, ">",
	},
	{func(p point) point {
		return point{p.x - 1, p.y}
	}, "<",
	},
	{func(p point) point {
		return point{p.x, p.y + 1}
	}, "v",
	},
	{func(p point) point {
		return point{p.x, p.y - 1}
	}, "^",
	},
}

func graph(costs [][]int) [][]map[point]edge {
	result := make([][]map[point]edge, len(costs))
	for y, row := range costs {
		result[y] = make([]map[point]edge, len(row))
		for x := range row {
			edges := make(map[point]edge)
		direction:
			for _, d := range directions {
				cost := 0
				p := point{x, y}
				for i := 1; i <= 10; i++ {
					p = d.fn(p)
					if p.x >= 0 && p.y >= 0 && p.x < len(row) && p.y < len(costs) {
						cost += costs[p.y][p.x]
						if i >= 4 {
							edges[p] = edge{cost, strings.Repeat(d.value, i)}
						}
					} else {
						continue direction
					}
				}
			}
			result[y][x] = edges
		}
	}
	return result
}

func values(file string) [][]int {
	lines := util.Lines(file)
	result := make([][]int, len(lines))
	for y, line := range lines {
		result[y] = make([]int, len(line))
		for x, c := range line {
			result[y][x] = int(c - '0')
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/17/input.txt"))
}
