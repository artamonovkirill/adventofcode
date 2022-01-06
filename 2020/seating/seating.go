package seating

import (
	"github.com/advendofcode/util"
	"strings"
)

func parse(file string) [][]string {
	lines := util.Lines(file)
	result := make([][]string, len(lines))
	for i, line := range lines {
		result[i] = make([]string, len(line))
		for j, e := range line {
			result[i][j] = string(e)
		}
	}
	return result
}

func advance(old [][]string, steps int) [][]string {
	if steps == 0 {
		return old
	}
	new := make([][]string, len(old))
	for i, row := range old {
		new[i] = make([]string, len(row))
		for j, e := range row {
			ns := neighbours(old, i, j)
			new[i][j] = update(e, ns)
		}
	}
	return advance(new, steps-1)
}

func update(seat string, neighbours map[string]int) string {
	switch seat {
	case "L":
		if neighbours["#"] == 0 {
			return "#"
		}
	case "#":
		if neighbours["#"] >= 5 {
			return "L"
		}
	}
	return seat
}

var ns = [][]int{
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func neighbours(state [][]string, i int, j int) map[string]int {
	result := make(map[string]int, 3)
	for _, n := range ns {
		result[neighbour(n, i, j, state)]++
	}
	return result
}

func neighbour(neighbour []int, i int, j int, state [][]string) string {
	multiplier := 1
	for {
		y := neighbour[0]*multiplier + i
		x := neighbour[1]*multiplier + j
		if x >= 0 && y >= 0 && y < len(state) && x < len(state[0]) {
			v := state[y][x]
			if v == "." {
				multiplier += 1
			} else {
				return v
			}
		} else {
			return "."
		}
	}
}

func toString(state [][]string) string {
	rows := make([]string, len(state))
	for i, row := range state {
		rows[i] = strings.Join(row, "")
	}
	return strings.Join(rows, "\n")
}

func solve(file string) int {
	old := parse(file)
	for {
		new := advance(old, 1)
		if toString(old) == toString(new) {
			result := 0
			for _, row := range old {
				for _, e := range row {
					if e == "#" {
						result++
					}
				}
			}
			return result
		}
		old = new
	}
}
