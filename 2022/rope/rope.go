package rope

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type coordinate struct {
	x, y int
}

func newRope(length int) []*coordinate {
	rope := make([]*coordinate, length)
	for i := 0; i < length; i++ {
		rope[i] = &coordinate{}
	}
	return rope
}

func process(file string, length int) int {
	rope := newRope(length)
	visited := map[coordinate]bool{}
	tail := rope[len(rope)-1]
	visited[*tail] = true
	history := []coordinate{*tail}

	head := rope[0]

	for _, line := range util.Lines(file) {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance := util.Number(parts[1])
		for i := 0; i < distance; i++ {
			switch direction {
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			default:
				panic("not implemented for " + direction)
			}
			for i := 0; i < len(rope)-1; i++ {
				move(*rope[i], rope[i+1])
			}
			visited[*tail] = true
			history = append(history, *tail)
		}
	}
	fmt.Println(history)
	return len(visited)
}

func move(head coordinate, tail *coordinate) {
	if head.y == tail.y {
		if head.x-tail.x > 1 {
			tail.x += 1
		} else if head.x-tail.x < -1 {
			tail.x -= 1
		}
	} else if head.x == tail.x {
		if head.y-tail.y > 1 {
			tail.y += 1
		} else if head.y-tail.y < -1 {
			tail.y -= 1
		}
	} else {
		if abs(head.y-tail.y) == 1 && abs(head.x-tail.x) == 1 {
			return
		} else if head.y-tail.y >= 1 && head.x-tail.x >= 1 {
			tail.y += 1
			tail.x += 1
		} else if head.y-tail.y >= 1 && head.x-tail.x <= -1 {
			tail.y += 1
			tail.x -= 1
		} else if head.y-tail.y <= -1 && head.x-tail.x >= 1 {
			tail.y -= 1
			tail.x += 1
		} else if head.y-tail.y <= -1 && head.x-tail.x <= -1 {
			tail.y -= 1
			tail.x -= 1
		} else {
			fmt.Println(head.y-tail.y, head.x-tail.x)
			panic(fmt.Sprintf("not implemented for %v %v", head, *tail))
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Solve() {
	file := "2022/rope/input.txt"
	fmt.Println(process(file, 2))
	fmt.Println(process(file, 10))
}
