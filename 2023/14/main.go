package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type rock struct {
	x, y    int
	movable bool
}

func (r *rock) isMovable() bool {
	return r.movable
}

func Solve(file string) int {
	m := parse(file)

	n := 1_000_000_000
	var ns []int
	for i := 1; i <= n; i++ {
		cycle(m)
		intermediate := 0
		for j, row := range m {
			for _, r := range row {
				if r != nil && r.isMovable() {
					intermediate += len(m) - j
				}
			}
		}
		ns = append(ns, intermediate)
		if len(ns) > 2 {
		outer:
			for j := 1; j < len(ns)/3; j++ {
				tail := ns[len(ns)-j*3:]
				one := tail[0:j]
				two := tail[j : 2*j]
				three := tail[2*j : 3*j]
				for k, e := range one {
					if e != two[k] || e != three[k] {
						continue outer
					}
				}
				rest := n - len(ns)
				rest = rest % len(one)
				return one[rest-1]
			}
		}
	}

	result := 0
	for i, row := range m {
		for _, r := range row {
			if r != nil && r.isMovable() {
				result += len(m) - i
			}
		}
	}

	return result
}

func tiltUp(m [][]*rock) {
outer:
	for {
		for y := 1; y < len(m); y++ {
			restart := false
			for x, e := range m[y] {
				if e != nil && e.isMovable() && m[y-1][x] == nil {
					restart = true
					m[y-1][x] = e
					m[y][x] = nil
				}
			}
			if restart {
				continue outer
			}
		}
		break outer
	}
}

func tiltDown(m [][]*rock) {
outer:
	for {
		for y := 0; y < len(m)-1; y++ {
			restart := false
			for x, e := range m[y] {
				if e != nil && e.isMovable() && m[y+1][x] == nil {
					restart = true
					m[y+1][x] = e
					m[y][x] = nil
				}
			}
			if restart {
				continue outer
			}
		}
		break outer
	}
}

func tiltLeft(m [][]*rock) {
outer:
	for {
		for x := 1; x < len(m[0]); x++ {
			restart := false
			for y := 0; y < len(m); y++ {
				e := m[y][x]
				if e != nil && e.isMovable() && m[y][x-1] == nil {
					restart = true
					m[y][x-1] = e
					m[y][x] = nil
				}
			}
			if restart {
				continue outer
			}
		}
		break outer
	}
}

func tiltRight(m [][]*rock) {
outer:
	for {
		for x := 0; x < len(m[0])-1; x++ {
			restart := false
			for y := 0; y < len(m); y++ {
				e := m[y][x]
				if e != nil && e.isMovable() && m[y][x+1] == nil {
					restart = true
					m[y][x+1] = e
					m[y][x] = nil
				}
			}
			if restart {
				continue outer
			}
		}
		break outer
	}
}

func cycle(m [][]*rock) {
	tiltUp(m)
	tiltLeft(m)
	tiltDown(m)
	tiltRight(m)
}

func parse(file string) [][]*rock {
	lines := util.Lines(file)
	m := make([][]*rock, len(lines))
	for y, line := range lines {
		m[y] = make([]*rock, len(line))
		for x, c := range line {
			if c == 'O' {
				m[y][x] = &rock{x: x, y: y, movable: true}
			} else if c == '#' {
				m[y][x] = &rock{x: x, y: y}
			}
		}
	}
	return m
}

func toString(m [][]*rock) string {
	lines := make([]string, len(m))
	for i, row := range m {
		line := ""
		for _, r := range row {
			if r == nil {
				line += "."
			} else if r.isMovable() {
				line += "O"
			} else {
				line += "#"
			}
		}
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

func main() {
	fmt.Println(Solve("2023/14/input.txt"))
}
