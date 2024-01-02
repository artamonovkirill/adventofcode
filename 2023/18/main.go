package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func Solve(file string, parse func(file string) []point) int {
	points := parse(file)

	area := 0
	perimeter := 0
	for i := 0; i < len(points)-1; i++ {
		area += points[i].x * points[i+1].y
		area -= points[i+1].x * points[i].y
		perimeter += (max(points[i+1].y, points[i].y)-min(points[i+1].y, points[i].y)+1)*(max(points[i+1].x, points[i].x)-min(points[i+1].x, points[i].x)+1) - 1
	}

	first := points[0]
	last := points[len(points)-1]
	area += last.x * first.y
	area -= first.x * last.y
	perimeter += (max(first.y, last.y)-min(first.y, last.y)+1)*(max(first.x, last.x)-min(first.x, last.x)+1) - 1
	area = area / 2
	inner := area - perimeter/2 + 1

	return inner + perimeter
}

func parse(file string) []point {
	x := 0
	y := 0
	var points []point
	for _, l := range util.Lines(file) {
		parts := strings.Split(l, " ")
		direction := parts[0]
		steps := util.Number(parts[1])
		switch direction {
		case "R":
			points = append(points, point{x + steps, y})
			x += steps
		case "L":
			points = append(points, point{x - steps, y})
			x -= steps
		case "D":
			points = append(points, point{x, y + steps})
			y += steps
		case "U":
			points = append(points, point{x, y - steps})
			y -= steps
		default:
			panic("not implemented")
		}
	}
	return points
}

func parse2(file string) []point {
	x := 0
	y := 0
	var points []point
	for _, l := range util.Lines(file) {
		instruction := strings.Split(l, " ")[2]
		direction := instruction[len(instruction)-2]
		sts, err := strconv.ParseUint(instruction[2:len(instruction)-2], 16, 62)
		if err != nil {
			panic(err)
		}
		steps := int(sts)
		if uint64(steps) != sts {
			panic("how?")
		}
		switch direction {
		case '0': //R
			points = append(points, point{x + steps, y})
			x += steps
		case '2': //L
			points = append(points, point{x - steps, y})
			x -= steps
		case '1': //D
			points = append(points, point{x, y + steps})
			y += steps
		case '3': //U
			points = append(points, point{x, y - steps})
			y -= steps
		default:
			panic("not implemented")
		}
	}
	return points
}

func main() {
	fmt.Println(Solve("2023/18/input.txt", parse))
	fmt.Println(Solve("2023/18/input.txt", parse2))
}
