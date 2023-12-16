package main

import (
	"fmt"
	"github.com/advendofcode/util"
)

type beam struct {
	x, y      int
	direction string
}

func Solve(file string) int {
	m := parse(file)
	result := 0
	for y := 0; y < len(m); y++ {
		current := []beam{{-1, y, "right"}}
		e := energy(current, m)
		if e > result {
			result = e
		}
	}
	for y := 0; y < len(m); y++ {
		current := []beam{{len(m[0]), y, "left"}}
		e := energy(current, m)
		if e > result {
			result = e
		}
	}
	for x := 0; x < len(m[0]); x++ {
		current := []beam{{x, -1, "down"}}
		e := energy(current, m)
		if e > result {
			result = e
		}
	}
	for x := 0; x < len(m[0]); x++ {
		current := []beam{{x, len(m), "up"}}
		e := energy(current, m)
		if e > result {
			result = e
		}
	}

	return result
}

func energy(current []beam, m [][]string) int {
	energized := make([][]map[string]bool, len(m))
	for y, line := range m {
		energized[y] = make([]map[string]bool, len(line))
		for x := range line {
			energized[y][x] = make(map[string]bool)
		}
	}

	for {
		var next []beam
		for _, b := range current {
			y := b.y
			x := b.x
			switch b.direction {
			case "right":
				x++
				if x < len(m[b.y]) {
					value := m[y][x]
					if !energized[y][x]["right"] {
						energized[y][x]["right"] = true
						switch value {
						case ".":
							next = append(next, beam{x, y, b.direction})
						case "-":
							next = append(next, beam{x, y, b.direction})
						case "/":
							next = append(next, beam{x, y, "up"})
						case "\\":
							next = append(next, beam{x, y, "down"})
						case "|":
							next = append(next, beam{x, y, "up"})
							next = append(next, beam{x, y, "down"})
						default:
							panic("not implemented")
						}
					}
				}
			case "left":
				x--
				if x >= 0 {
					value := m[y][x]
					if !energized[y][x]["left"] {
						energized[y][x]["left"] = true
						switch value {
						case ".":
							next = append(next, beam{x, y, b.direction})
						case "-":
							next = append(next, beam{x, y, b.direction})
						case "/":
							next = append(next, beam{x, y, "down"})
						case "\\":
							next = append(next, beam{x, y, "up"})
						case "|":
							next = append(next, beam{x, y, "up"})
							next = append(next, beam{x, y, "down"})
						default:
							panic("not implemented")
						}
					}
				}
			case "up":
				y--
				if y >= 0 {
					value := m[y][x]
					if !energized[y][x]["up"] {
						energized[y][x]["up"] = true
						switch value {
						case ".":
							next = append(next, beam{x, y, b.direction})
						case "|":
							next = append(next, beam{x, y, b.direction})
						case "\\":
							next = append(next, beam{x, y, "left"})
						case "/":
							next = append(next, beam{x, y, "right"})
						case "-":
							next = append(next, beam{x, y, "left"})
							next = append(next, beam{x, y, "right"})
						default:
							panic("not implemented")
						}
					}
				}
			case "down":
				y++
				if y < len(m) {
					value := m[y][x]
					if !energized[y][x]["down"] {
						energized[y][x]["down"] = true
						switch value {
						case ".":
							next = append(next, beam{x, y, b.direction})
						case "|":
							next = append(next, beam{x, y, b.direction})
						case "/":
							next = append(next, beam{x, y, "left"})
						case "\\":
							next = append(next, beam{x, y, "right"})
						case "-":
							next = append(next, beam{x, y, "left"})
							next = append(next, beam{x, y, "right"})
						default:
							panic("not implemented")
						}
					}
				}
			default:
				panic("not implemented")
			}
		}
		if len(next) == 0 {
			result := 0
			for _, line := range energized {
				for _, e := range line {
					if len(e) > 0 {
						result++
					}
				}
			}
			return result
		}
		current = next
	}
}

func parse(file string) [][]string {
	lines := util.Lines(file)
	result := make([][]string, len(lines))
	for y, line := range lines {
		result[y] = make([]string, len(line))
		for x, c := range line {
			result[y][x] = string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/16/input.txt"))
}
