package beacon

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type occupancy int

const (
	free occupancy = iota
	sensor
	beacon
	excluded
)

type coordinate struct {
	x, y int
}

func impossible(file string, y int) int {
	cs := make(map[coordinate]occupancy)
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, ": ")
		s := parse(strings.ReplaceAll(parts[0], "Sensor at ", ""))
		b := parse(strings.ReplaceAll(parts[1], "closest beacon is at ", ""))
		cs[s] = sensor
		cs[b] = beacon
		width := abs(s.x - b.x)
		d := distance(s, b)
		for x := s.x - width - 1000000; x <= s.x+width+1000000; x++ {
			c := coordinate{x, y}
			if c.y == y && distance(s, c) <= d {
				if cs[c] == free {
					cs[c] = excluded
				}
			}
		}
	}
	result := 0
	for k, v := range cs {
		if k.y == y && v == excluded {
			result += 1
		}
	}

	return result
}

func distance(a coordinate, b coordinate) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parse(s string) coordinate {
	parts := strings.Split(s, ", ")
	x := util.Number(strings.ReplaceAll(parts[0], "x=", ""))
	y := util.Number(strings.ReplaceAll(parts[1], "y=", ""))
	return coordinate{x, y}
}

func Solve() {
	file := "2022/beacon/input.txt"
	fmt.Println(impossible(file, 2_000_000))
	fmt.Println(frequency(file, 4_000_000))
}

type Sensor struct {
	x, y, distance int
}

func frequency(file string, max int) int {
	lines := util.Lines(file)
	sensors := make([]Sensor, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		s := parse(strings.ReplaceAll(parts[0], "Sensor at ", ""))
		b := parse(strings.ReplaceAll(parts[1], "closest beacon is at ", ""))
		sensors[i] = Sensor{s.x, s.y, distance(s, b)}
	}

	sort.Slice(sensors, func(i, j int) bool {
		return sensors[i].x < sensors[j].x
	})

	var ys []int
	var xs []int

one:
	for y := max; y >= 0; y-- {
		xMin := 0
		xMax := max
		for _, sensor := range sensors {
			if abs(sensor.y-y) > sensor.distance {
				continue
			}
			left := sensor.x - (sensor.distance - abs(sensor.y-y))
			right := sensor.x + (sensor.distance - abs(sensor.y-y))
			if right > xMax && xMax > left {
				xMax = left
			}
			if left < xMin && xMin < right {
				xMin = right
			}
			if xMin >= xMax {
				continue one
			}
		}
		ys = append(ys, y)
	}

two:
	for x := max; x >= 0; x-- {
		yMin := 0
		yMax := max
		for _, sensor := range sensors {
			if abs(sensor.x-x) > sensor.distance {
				continue
			}
			left := sensor.y - (sensor.distance - abs(sensor.x-x))
			right := sensor.y + (sensor.distance - abs(sensor.x-x))
			if right > yMax && yMax > left {
				yMax = left
			}
			if left < yMin && yMin < right {
				yMin = right
			}
			if yMin >= yMax {
				continue two
			}
		}
		xs = append(xs, x)
	}

	fmt.Printf("Candidates: x #%d, y #%d", len(xs), len(ys))

	for i, y := range ys {
		if i%100 == 0 {
			fmt.Println(i)
		}
	three:
		for _, x := range xs {
			c := coordinate{x, y}
			for _, sensor := range sensors {
				if distance(c, coordinate{sensor.x, sensor.y}) <= sensor.distance {
					continue three
				}
			}
			return c.x*4_000_000 + c.y
		}
	}
	panic("no solution found")
}
