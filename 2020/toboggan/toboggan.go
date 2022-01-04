package toboggan

import (
	"github.com/advendofcode/util"
)

func hits(mp []string, right int, down int) int {
	result := 0
	width := len(mp[0])
	x := 0
	y := 0
	for {
		x = (x + right) % width
		y += down
		if y >= len(mp) {
			return result
		}
		if mp[y][x] == '#' {
			result++
		}
	}
}

type pattern struct {
	right int
	down  int
}

func solve(file string) int {
	result := 1
	lines := util.Lines(file)

	patterns := []pattern{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, p := range patterns {
		result *= hits(lines, p.right, p.down)
	}
	return result
}
