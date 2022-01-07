package recitation

import "fmt"

func solve(start []int) int {
	used := map[int][]int{}
	for i, s := range start {
		used[s] = []int{i}
	}
	i := len(start)
	for i < 30000000 {
		u := used[last(start)]
		if len(u) <= 1 {
			start = append(start, 0)
			used[0] = append(used[0], i)
		} else {
			n := last(u) - u[len(u)-2]
			start = append(start, n)
			used[n] = append(used[n], i)
		}
		i++
	}
	return last(start)
}

func last(xs []int) int {
	return xs[len(xs)-1]
}

func Solve() {
	fmt.Println(solve([]int{0, 1, 5, 10, 3, 12, 19}))
}
