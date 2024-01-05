package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type stone struct {
	x, y, vx, vy float32
}

func Solve(file string, min, max float32) int {
	lines := util.Lines(file)
	stones := make([]stone, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " @ ")
		point := strings.Split(parts[0], ", ")
		vs := strings.Split(parts[1], ", ")
		stones[i] = stone{
			x:  float32(util.Number(point[0])),
			y:  float32(util.Number(point[1])),
			vx: float32(util.Number(strings.ReplaceAll(vs[0], " ", ""))),
			vy: float32(util.Number(strings.ReplaceAll(vs[1], " ", ""))),
		}
	}

	result := 0
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			a := stones[i]
			b := stones[j]
			// a.x + a.vx * x1 = b.x + b.vx * x2
			// a.y + a.vy * x1 = b.y + b.vy * x2
			// x1 = (b.y - a.y + b.vy * x2) / a.vy
			// x1 = b.y / a.vy - a.y / a.vy + b.vy / a.vy * x2
			// a.x + a.vx * (b.y / a.vy - a.y / a.vy + b.vy / a.vy * x2) = b.x + b.vx * x2
			// a.x + a.vx*b.y/a.vy - a.vx*a.y/a.vy + a.vx*b.vy/a.vy*x2 = b.x + b.vx*x2
			// x2 * (b.vx - a.vx*b.vy/a.vy) = a.x + a.vx*b.y/a.vy - a.vx*a.y/a.vy - b.x
			x2 := (a.x + a.vx*b.y/a.vy - a.vx*a.y/a.vy - b.x) / (b.vx - a.vx*b.vy/a.vy)
			x1 := (b.y - a.y + b.vy*x2) / a.vy
			x := b.x + b.vx*x2
			y := b.y + b.vy*x2
			if min <= x && x <= max && min <= y && y <= max && x1 >= 0 && x2 >= 0 {
				result++
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Solve("2023/24/input.txt", 200000000000000, 400000000000000))
}
