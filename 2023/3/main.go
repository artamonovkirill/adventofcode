package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile("[0-9]+")

func Parts(file string) int {
	result := 0
	lines := util.Lines(file)
	for i, line := range lines {
		numbers := r.FindAllString(line, -1)
		indexes := r.FindAllStringIndex(line, -1)
	numbers:
		for n, ixs := range indexes {
			number, err := strconv.Atoi(numbers[n])
			if err != nil {
				panic(err)
			}
			for y := max(i-1, 0); y < min(i+2, len(lines)); y++ {
				for x := max(ixs[0]-1, 0); x < min(ixs[1]+1, len(line)); x++ {
					c := lines[y][x]
					if (c < '0' || c > '9') && c != '.' {
						result += number
						continue numbers
					}
				}
			}
		}

	}
	return result
}

func main() {
	fmt.Println(Parts("2023/3/input.txt"))
	fmt.Println(Gears("2023/3/input.txt"))
}

type Coordinate struct {
	x, y int
}

func Gears(file string) int {
	result := 0
	gears := make(map[Coordinate][]int)

	lines := util.Lines(file)
	for i, line := range lines {
		numbers := r.FindAllString(line, -1)
		indexes := r.FindAllStringIndex(line, -1)
		for n, ixs := range indexes {
			number, err := strconv.Atoi(numbers[n])
			if err != nil {
				panic(err)
			}
			for y := max(i-1, 0); y < min(i+2, len(lines)); y++ {
				for x := max(ixs[0]-1, 0); x < min(ixs[1]+1, len(line)); x++ {
					c := lines[y][x]
					if c == '*' {
						coordinate := Coordinate{x, y}
						v, ok := gears[coordinate]
						if ok {
							gears[coordinate] = append(v, number)
						} else {
							gears[coordinate] = []int{number}
						}
					}
				}
			}
		}
	}
	for _, v := range gears {
		if len(v) == 2 {
			result += v[0] * v[1]
		}
	}

	return result
}
