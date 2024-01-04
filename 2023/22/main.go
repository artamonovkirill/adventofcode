package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"math"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type brick struct {
	name     string
	elements map[point]bool
}

func Solve(file string) int {
	bricks, stack := parse(file)
	fall(bricks, stack)
	return len(findRedundant(bricks, stack))
}

func Solve2(file string) int {
	bs, s := parse(file)
	fall(bs, s)
	redundant := findRedundant(bs, s)

	result := 0
	for i, b := range bs {
		if !redundant[b.name] {
			bricks, stack := parse(file)
			fall(bricks, stack)

			rest := append(bricks[:i], bricks[i+1:]...)
			for e := range b.elements {
				delete(stack, e)
			}

			fallen := fall(rest, stack)
			result += fallen
		}
	}

	return result
}

func fall(bricks []brick, stack map[point]string) int {
	i := 0
	fallen := make(map[string]bool)
falling:
	for {
		if i == len(bricks) {
			return len(fallen)
		}

		b := bricks[i]
		newBrick := brick{
			name:     b.name,
			elements: make(map[point]bool),
		}
		for e := range b.elements {
			lower := point{e.x, e.y, e.z - 1}
			newBrick.elements[lower] = true
			if e.z == 1 || (stack[lower] != "" && stack[lower] != b.name) {
				i++
				continue falling
			}
		}
		bricks[i] = newBrick
		fallen[newBrick.name] = true
		for e := range b.elements {
			stack[e] = ""
		}
		for e := range newBrick.elements {
			stack[e] = newBrick.name
		}
	}
}

func minZ(b brick) int {
	result := math.MaxInt
	for e := range b.elements {
		if e.z < result {
			result = e.z
		}
	}
	return result
}

func parse(file string) ([]brick, map[point]string) {
	lines := util.Lines(file)
	bricks := make([]brick, len(lines))
	stack := make(map[point]string)
	for i, line := range lines {
		parts := strings.Split(line, "~")
		start := strings.Split(parts[0], ",")
		end := strings.Split(parts[1], ",")
		startX := util.Number(start[0])
		endX := util.Number(end[0])
		startY := util.Number(start[1])
		endY := util.Number(end[1])
		startZ := util.Number(start[2])
		endZ := util.Number(end[2])
		name := strconv.Itoa(i)
		b := brick{
			name:     name,
			elements: make(map[point]bool),
		}
		for x := min(startX, endX); x <= max(startX, endX); x++ {
			for y := min(startY, endY); y <= max(startY, endY); y++ {
				for z := min(startZ, endZ); z <= max(startZ, endZ); z++ {
					b.elements[point{x, y, z}] = true
					stack[point{x, y, z}] = name
				}
			}
		}
		bricks[i] = b
	}
	sort.Slice(bricks, func(i, j int) bool {
		return minZ(bricks[i]) < minZ(bricks[j])
	})
	return bricks, stack
}

func findRedundant(bricks []brick, stack map[point]string) map[string]bool {
	supports := make(map[string]map[string]bool)

	for _, b := range bricks {
		for e := range b.elements {
			higher := point{e.x, e.y, e.z + 1}
			if stack[higher] != "" && stack[higher] != b.name {
				if _, ok := supports[stack[higher]]; !ok {
					supports[stack[higher]] = make(map[string]bool)
				}
				supports[stack[higher]][b.name] = true
			}
		}
	}
	indestructible := make(map[string]bool)
	for _, s := range supports {
		if len(s) == 1 {
			for name := range s {
				indestructible[name] = true
			}
		}
	}

	result := make(map[string]bool)
	for _, b := range bricks {
		if !indestructible[b.name] {
			result[b.name] = true
		}
	}

	return result
}

func main() {
	fmt.Println(Solve("2023/22/input.txt"))
	fmt.Println(Solve2("2023/22/input.txt"))
}
