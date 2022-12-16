package valves

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type valve struct {
	rate    int
	tunnels []string
}

type path struct {
	locations        []string
	releasePerMinute int
	totalRelease     int
	valves           map[string]bool
}

func max(file string, helpers int) int {
	valves := parse(file)
	var locations []string
	if helpers == 1 {
		locations = []string{"AA"}
	} else {
		locations = []string{"AA", "AA"}
	}
	paths := []path{{locations, 0, 0, map[string]bool{}}}
	time := 30 - (helpers-1)*4
	for i := 0; i < time; i++ {
		var newPaths []path
		for _, p := range paths {
			p0 := p.locations[0]
			v0 := valves[p0]
			if helpers > 1 {
				p1 := p.locations[1]
				v1 := valves[p1]
				// both move
				for _, a := range v0.tunnels {
					for _, b := range v1.tunnels {
						c := clone(p)
						c.locations = []string{a, b}
						c.totalRelease += p.releasePerMinute
						newPaths = append(newPaths, c)
					}
				}
				if p0 == p1 { //one has to move
					if v0.rate > 0 && !p.valves[p0] {
						for _, b := range v1.tunnels {
							c := clone(p)
							c.totalRelease += c.releasePerMinute
							c.releasePerMinute += v0.rate
							c.locations[1] = b
							c.valves[p0] = true
							newPaths = append(newPaths, c)
						}
					}
				} else {
					// both open
					if v0.rate > 0 && !p.valves[p0] && v1.rate > 0 && !p.valves[p1] {
						c := clone(p)
						c.totalRelease += c.releasePerMinute
						c.releasePerMinute += v0.rate + v1.rate
						c.valves[p0] = true
						c.valves[p1] = true
						newPaths = append(newPaths, c)
					} else if v0.rate > 0 && !p.valves[p0] { //you open
						for _, tunnel := range v1.tunnels {
							c := clone(p)
							c.totalRelease += c.releasePerMinute
							c.releasePerMinute += v0.rate
							c.locations = []string{p0, tunnel}
							c.valves[p0] = true
							newPaths = append(newPaths, c)
						}
					} else if v1.rate > 0 && !p.valves[p1] { //elephant opens
						for _, tunnel := range v0.tunnels {
							c := clone(p)
							c.totalRelease += c.releasePerMinute
							c.releasePerMinute += v1.rate
							c.locations = []string{tunnel, p1}
							c.valves[p1] = true
							newPaths = append(newPaths, c)
						}
					}
				}
			} else {
				// move
				for _, tunnel := range v0.tunnels {
					c := clone(p)
					c.locations = []string{tunnel}
					c.totalRelease += p.releasePerMinute
					newPaths = append(newPaths, c)
				}
				if v0.rate > 0 && !p.valves[p.locations[0]] {
					// open
					c := clone(p)
					c.totalRelease += c.releasePerMinute
					c.releasePerMinute += v0.rate
					c.valves[p.locations[0]] = true
					newPaths = append(newPaths, c)
				}
			}
		}
		sort.Slice(newPaths, func(i, j int) bool {
			return newPaths[i].totalRelease > newPaths[j].totalRelease
		})
		fmt.Println(i, newPaths[0].totalRelease)
		if len(newPaths) > 100000 {
			paths = newPaths[0:100000]
		} else {
			paths = newPaths
		}
	}
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].totalRelease > paths[j].totalRelease
	})
	return paths[0].totalRelease
}

func clone(p path) path {
	valves := make(map[string]bool)
	for k, v := range p.valves {
		valves[k] = v
	}
	return path{
		p.locations,
		p.releasePerMinute,
		p.totalRelease,
		valves}
}

func parse(file string) map[string]valve {
	lines := util.Lines(file)
	valves := make(map[string]valve)
	for _, line := range lines {
		parts := strings.Split(line, "; ")
		valveInfo := strings.Split(parts[0], " ")
		name := valveInfo[1]
		rate := util.Number(strings.ReplaceAll(valveInfo[4], "rate=", ""))
		tunnels := parseTunnels(parts[1])
		valves[name] = valve{rate: rate, tunnels: tunnels}
	}
	for k, v := range valves {
		if len(k) != 2 {
			panic("weird name " + k)
		}
		for _, tunnel := range v.tunnels {
			if len(tunnel) != 2 {
				panic("weird tunnel " + k)
			}
		}
	}
	return valves
}

func parseTunnels(input string) []string {
	input = strings.ReplaceAll(input, "tunnels lead to valves ", "")
	input = strings.ReplaceAll(input, "tunnel leads to valve ", "")
	return strings.Split(input, ", ")
}

func Solve() {
	fmt.Println(max("2022/valves/input.txt", 1))
	fmt.Println(max("2022/valves/input.txt", 2))
}
