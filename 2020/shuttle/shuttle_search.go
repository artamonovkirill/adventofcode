package shuttle

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

func find(timestamp int, busses string) int {
	bus := 0
	wait := timestamp
	for _, b := range strings.Split(busses, ",") {
		if b != "x" {
			n := util.Number(b)
			next := (timestamp/n)*n + n
			if next-timestamp < wait {
				bus = n
				wait = next - timestamp
			}
		}
	}
	return wait * bus
}

func Solve() {
	busses := "23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,449,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,x,x,x,x,x,x,29,x,991,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,17"
	fmt.Println(find(1008832, busses))
	// map[13:2 17:3 19:4 23:0 29:23 37:23 41:13 449:23 991:54]
	// 23 is the most common delay
	// 29, 37 and 449 are all primes
	// so final (result+23) should be divisible by 29*37*449 = 481777
	fmt.Println(earliest(busses, 481777, 23))
}

func earliest(input string, base int, delay int) int {
	multiplier := 1
	busses := map[int]int{}
	var ns []int
	for i, b := range strings.Split(input, ",") {
		if b != "x" {
			n := util.Number(b)
			busses[n] = i % n
			ns = append(ns, n)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	if base == 0 {
		base = ns[0]
	}
	fmt.Println(busses, ns)
outer:
	for {
		d := delay
		if delay == 0 {
			d = busses[base]
		}
		acc := base*multiplier - d
		if multiplier%100_000_000 == 0 {
			fmt.Println(multiplier, acc)
		}
		for _, b := range ns {
			if (acc+busses[b])%b != 0 {
				multiplier++
				continue outer
			}
		}
		return acc
	}
}
