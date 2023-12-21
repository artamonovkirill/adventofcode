package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type part struct {
	x, m, a, s int
}

func Solve(file string) int {
	input := strings.Split(util.Text(file), "\n\n")
	workflows := parseWorkflows(input[0])
	parts := parseParts(input[1])
	result := 0
part:
	for _, p := range parts {
		current := workflows["in"]
	inf:
		for {
		fn:
			for _, fn := range current {
				out := fn(p)
				switch out {
				case "":
					continue fn
				case "R":
					continue part
				case "A":
					result += p.x + p.m + p.a + p.s
					continue part
				default:
					current = workflows[out]
					continue inf
				}
			}
		}
	}
	return result
}

func parseParts(input string) []part {
	lines := strings.Split(input, "\n")
	result := make([]part, len(lines))
	for i, line := range lines {
		p := part{}
		for _, values := range strings.Split(line[1:len(line)-1], ",") {
			parts := strings.Split(values, "=")
			switch parts[0] {
			case "x":
				p.x = util.Number(parts[1])
			case "m":
				p.m = util.Number(parts[1])
			case "a":
				p.a = util.Number(parts[1])
			case "s":
				p.s = util.Number(parts[1])
			default:
				panic("not implemented")
			}
		}
		result[i] = p
	}
	return result
}

func parseWorkflows(input string) map[string][]func(part) string {
	result := make(map[string][]func(part) string)
	for _, line := range strings.Split(input, "\n") {
		chunks := strings.Split(line[0:len(line)-1], "{")
		name := chunks[0]
		fns := strings.Split(chunks[1], ",")
		functions := make([]func(part) string, len(fns))
		for i, fn := range fns {
			parts := strings.Split(fn, ":")
			if len(parts) == 1 {
				functions[i] = func(p part) string {
					return parts[0]
				}
			} else {
				matcher := parts[0]
				target := parts[1]
				switch matcher[0:2] {
				case "a<":
					functions[i] = func(p part) string {
						if p.a < util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "a>":
					functions[i] = func(p part) string {
						if p.a > util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "m<":
					functions[i] = func(p part) string {
						if p.m < util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "m>":
					functions[i] = func(p part) string {
						if p.m > util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "s<":
					functions[i] = func(p part) string {
						if p.s < util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "s>":
					functions[i] = func(p part) string {
						if p.s > util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "x<":
					functions[i] = func(p part) string {
						if p.x < util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				case "x>":
					functions[i] = func(p part) string {
						if p.x > util.Number(matcher[2:]) {
							return target
						} else {
							return ""
						}
					}
				default:
					panic("not implemented for " + matcher)
				}
			}
		}
		result[name] = functions
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/19/input.txt"))
}
