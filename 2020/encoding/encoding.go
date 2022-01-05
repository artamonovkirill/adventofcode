package encoding

import (
	"github.com/advendofcode/util"
	"sort"
)

func check(file string, preamble int) int {
	ns := util.Numbers(file)
OUTER:
	for i, n := range ns {
		if i < preamble {
			continue
		}
		for _, c := range combinations(i-preamble, i) {
			if ns[c.a]+ns[c.b] == n {
				continue OUTER
			}
		}
		return n
	}
	return 0
}

type Pair struct {
	a int
	b int
}

func combinations(min int, max int) []Pair {
	var result []Pair
	for i := min; i < max-1; i++ {
		for j := i + 1; j < max; j++ {
			result = append(result, Pair{i, j})
		}
	}
	return result
}

func find(file string, target int) int {
	ns := util.Numbers(file)
	for i, n := range ns {
		cs := []int{n}
		sum := n
		j := i
		for sum < target {
			j++
			c := ns[j]
			cs = append(cs, c)
			sum += c
		}
		if len(cs) > 1 && sum == target {
			sort.Ints(cs)
			return cs[0] + cs[len(cs)-1]
		}
	}

	return 0
}
