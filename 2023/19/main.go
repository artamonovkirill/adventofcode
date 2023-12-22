package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type interval struct {
	left, right int
}

type group struct {
	x, m, a, s interval
	workflow   string
}

func Solve(file string) int {
	input := strings.Split(util.Text(file), "\n\n")
	workflows := parseWorkflows(input[0])
	result := 0
	current := []group{
		{
			x:        interval{1, 4000},
			m:        interval{1, 4000},
			a:        interval{1, 4000},
			s:        interval{1, 4000},
			workflow: "in",
		},
	}
	var accepted []group
	for {
		var next []group
		for _, g := range current {
			for _, ng := range workflows[g.workflow](g) {
				switch ng.workflow {
				case "R":
				case "A":
					accepted = append(accepted, ng)
				default:
					next = append(next, ng)
				}
			}
		}
		if len(next) == 0 {
			for _, acc := range accepted {
				result += length(acc.x) * length(acc.m) * length(acc.a) * length(acc.s)
			}
			return result
		}
		current = next
	}
}

func length(i interval) int {
	return i.right - i.left + 1
}

func parseWorkflows(input string) map[string]func(group) []group {
	result := make(map[string]func(group) []group)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		chunks := strings.Split(line[0:len(line)-1], "{")
		name := chunks[0]
		matchers := chunks[1]
		result[name] = func(input group) []group {
			var gs []group
			current := input
			for _, matcher := range strings.Split(matchers, ",") {
				parts := strings.Split(matcher, ":")
				if len(parts) == 1 {
					gs = append(gs, group{
						x: current.x, m: current.m, a: current.a, s: current.s,
						workflow: parts[0],
					})
				} else {
					m := parts[0]
					target := parts[1]
					value := util.Number(m[2:])
					switch m[0:2] {
					case "a<":
						if current.a.right >= value {
							gs = append(gs, group{
								x: current.x, m: current.m, a: interval{current.a.left, value - 1}, s: current.s,
								workflow: target,
							})
							current = group{
								x: current.x, m: current.m, a: interval{value, current.a.right}, s: current.s,
								workflow: current.workflow,
							}
						}
					case "a>":
						if current.a.left <= value {
							gs = append(gs, group{
								x: current.x, m: current.m, a: interval{value + 1, current.a.right}, s: current.s,
								workflow: target,
							})
							current = group{
								x: current.x, m: current.m, a: interval{current.a.left, value}, s: current.s,
								workflow: current.workflow,
							}
						}
					case "m<":
						if current.m.right >= value {
							gs = append(gs, group{
								x: current.x, m: interval{current.m.left, value - 1}, a: current.a, s: current.s,
								workflow: target,
							})
							current = group{
								x: current.x, m: interval{value, current.m.right}, a: current.a, s: current.s,
								workflow: current.workflow,
							}
						}
					case "m>":
						if current.m.left <= value {
							gs = append(gs, group{
								x: current.x, m: interval{value + 1, current.m.right}, a: current.a, s: current.s,
								workflow: target,
							})
							current = group{
								x: current.x, m: interval{current.m.left, value}, a: current.a, s: current.s,
								workflow: current.workflow,
							}
						}
					case "s<":
						if current.s.right >= value {
							gs = append(gs, group{
								x: current.x, m: current.m, a: current.a, s: interval{current.s.left, value - 1},
								workflow: target,
							})
							current = group{
								x: current.x, m: current.m, a: current.a, s: interval{value, current.s.right},
								workflow: current.workflow,
							}
						}
					case "s>":
						if current.s.left <= value {
							gs = append(gs, group{
								x: current.x, m: current.m, a: current.a, s: interval{value + 1, current.s.right},
								workflow: target,
							})
							current = group{
								x: current.x, m: current.m, a: current.a, s: interval{current.s.left, value},
								workflow: current.workflow,
							}
						}
					case "x<":
						if current.x.right >= value {
							gs = append(gs, group{
								x: interval{current.x.left, value - 1}, m: current.m, a: current.a, s: current.s,
								workflow: target,
							})
							current = group{
								x: interval{value, current.x.right}, m: current.m, a: current.a, s: current.s,
								workflow: current.workflow,
							}
						}
					case "x>":
						if current.x.left <= value {
							gs = append(gs, group{
								x: interval{value + 1, current.x.right}, m: current.m, a: current.a, s: current.s,
								workflow: target,
							})
							current = group{
								x: interval{current.x.left, value}, m: current.m, a: current.a, s: current.s,
								workflow: current.workflow,
							}
						}
					default:
						panic("not implemented for " + matcher + target)
					}
				}
			}
			return gs
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/19/input.txt"))
}
