package adapter

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
)

type Link struct {
	from int
	to   int
}

func solve(file string) int {
	ns := util.Numbers(file)
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	ns = append(ns, 0)
	bs := make(map[int]bool, len(ns))
	for _, n := range ns {
		bs[n] = true
	}

	paths := map[int]int{
		ns[0]: 1,
	}
	for _, n := range ns {
		for i := 1; i <= 3; i++ {
			if bs[n+i] {
				paths[n] += paths[n+i]
			}
		}
	}
	return paths[0]
}

func Solve() {
	input := "2020/adapter/puzzle.txt"
	fmt.Println(solve(input))
}
