package tube

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type cpu struct {
	value int
	cycle int
}

func strength(file string) int {
	c := cpu{value: 1}
	result := 0
	for _, line := range util.Lines(file) {
		c.cycle += 1
		result += record(c)
		if line != "noop" {
			c.cycle += 1
			result += record(c)
			c.value += util.Number(strings.Split(line, " ")[1])
		}
	}
	return result
}

func initalize() [][]rune {
	s := make([][]rune, 6)
	for i := 0; i < 6; i++ {
		row := make([]rune, 40)
		for j := 0; j < 40; j++ {
			row[j] = '.'
		}
		s[i] = row
	}
	return s
}

type screen [][]rune

func crt(file string) string {
	d := cpu{value: 1}
	s := initalize()
	for _, line := range util.Lines(file) {
		d.cycle += 1
		s = draw(d, s)
		if line != "noop" {
			d.cycle += 1
			s = draw(d, s)
			d.value += util.Number(strings.Split(line, " ")[1])
		}
	}
	result := make([]string, len(s))
	for i, line := range s {
		result[i] = string(line)
	}
	return strings.Join(result, "\n")
}

func draw(d cpu, s screen) screen {
	index := d.cycle - 1
	row := index / 40
	pos := index % 40
	if abs(pos-d.value) <= 1 {
		s[row][pos] = '#'
	}
	return s
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (d *cpu) strength() int {
	return d.cycle * d.value
}

func record(d cpu) int {
	switch d.cycle {
	case 20:
		return d.strength()
	case 60:
		return d.strength()
	case 100:
		return d.strength()
	case 140:
		return d.strength()
	case 180:
		return d.strength()
	case 220:
		return d.strength()
	default:
		return 0
	}
}

func Solve() {
	file := "2022/tube/input.txt"
	fmt.Println(strength(file))
	fmt.Println(crt(file))
}
