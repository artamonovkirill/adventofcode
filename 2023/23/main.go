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
	end     point
	visited map[point]bool
	cost    int
	history []point
}

type edge struct {
	from, to point
	cost     int
	visited  map[point]bool
}

func Solve(file string) int {
	matrix := parse(file)
	ns := neighbours(matrix)

	var starts []int
	for x, n := range ns[0] {
		if len(n) == 1 {
			starts = append(starts, x)
		}
	}
	if len(starts) > 1 {
		panic("where to start?")
	}
	start := point{starts[0], 0}

	var ends []int
	for x, n := range ns[len(ns)-1] {
		if len(n) == 1 {
			ends = append(ends, x)
		}
	}
	if len(ends) > 1 {
		panic("where to start?")
	}
	end := point{ends[0], len(ns) - 1}

	g := graph(start, ns)

	var current []path
	for nxt, cost := range g[start] {
		current = append(current, path{
			cost:    cost,
			end:     nxt,
			visited: map[point]bool{start: true, nxt: true},
			history: []point{start, nxt},
		})
	}

	best := 0
	for len(current) > 0 {
		var next []path

		for _, c := range current {
			for nxt, cost := range g[c.end] {
				if !c.visited[nxt] {
					if nxt == end {
						if c.cost+cost > best {
							best = c.cost + cost
						}
					} else {
						visited := maps.Clone(c.visited)
						visited[nxt] = true
						next = append(next, path{
							cost:    c.cost + cost,
							end:     nxt,
							visited: visited,
							history: append(c.history, nxt),
						})
					}
				}
			}
		}

		current = next
	}

	return best
}

func graph(start point, ns [][][]point) map[point]map[point]int {
	var current []edge
	for _, n := range ns[start.y][start.x] {
		current = append(current, edge{start, n, 1, map[point]bool{
			start: true, n: true,
		}})
	}

	result := make(map[point]map[point]int)
	for len(current) > 0 {
		var next []edge
		for _, c := range current {
			to := c.to
			for _, n := range ns[to.y][to.x] {
				if c.visited[n] {
					// do nothing
				} else {
					if len(ns[n.y][n.x]) == 2 {
						visited := c.visited
						visited[n] = true
						next = append(next, edge{c.from, n, c.cost + 1, visited})
					} else if len(ns[n.y][n.x]) > 2 {
						if result[c.from][n] == 0 {
							if _, ok := result[c.from]; ok {
								result[c.from][n] = c.cost + 1
							} else {
								result[c.from] = map[point]int{n: c.cost + 1}
							}
							if _, ok := result[n]; ok {
								result[n][c.from] = c.cost + 1
							} else {
								result[n] = map[point]int{c.from: c.cost + 1}
							}
							for _, nn := range ns[n.y][n.x] {
								if !c.visited[nn] {
									if len(ns[nn.y][nn.x]) != 2 {
										panic("how?")
									}
									next = append(next, edge{n, nn, 1, map[point]bool{
										n: true, nn: true,
									}})
								}
							}
						}
					} else if len(ns[n.y][n.x]) == 1 {
						if result[c.from][n] == 0 {
							if _, ok := result[c.from]; ok {
								result[c.from][n] = c.cost + 1
							} else {
								result[c.from] = map[point]int{n: c.cost + 1}
							}
							if _, ok := result[n]; ok {
								result[n][c.from] = c.cost + 1
							} else {
								result[n] = map[point]int{c.from: c.cost + 1}
							}
						}
					} else {
						panic("not implemented")
					}
				}
			}
		}
		current = next
	}
	return result
}

func neighbours(matrix [][]int32) [][][]point {
	result := make([][][]point, len(matrix))
	for y, row := range matrix {
		result[y] = make([][]point, len(row))
		for x, c := range row {
			var ns []point
			if c != '#' {
				if y > 0 && matrix[y-1][x] != '#' {
					ns = append(ns, point{x, y - 1})
				}
				if y < len(matrix)-1 && matrix[y+1][x] != '#' {
					ns = append(ns, point{x, y + 1})
				}
				if x > 0 && matrix[y][x-1] != '#' {
					ns = append(ns, point{x - 1, y})
				}
				if x < len(matrix[0])-1 && matrix[y][x+1] != '#' {
					ns = append(ns, point{x + 1, y})
				}
			}
			result[y][x] = ns
		}
	}
	return result
}

func parse(file string) [][]int32 {
	lines := util.Lines(file)
	matrix := make([][]int32, len(lines))
	for y, line := range lines {
		matrix[y] = make([]int32, len(line))
		for x, c := range line {
			matrix[y][x] = c
		}
	}
	return matrix
}

func main() {
	fmt.Println(Solve("2023/23/input.txt"))
}
