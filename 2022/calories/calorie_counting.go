package calories

import (
	"fmt"
	"github.com/advendofcode/util"
)

func max(file string) int {
	max := 0
	current := 0
	for _, l := range util.Lines(file) {
		if l == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += util.Number(l)
		}
	}
	return max
}

func topThree(file string) int {
	top := []int{0, 0, 0}
	current := 0
	for _, l := range util.Lines(file) {
		if l == "" {
			if current > top[0] {
				top[2] = top[1]
				top[1] = top[0]
				top[0] = current
			} else if current > top[1] {
				top[2] = top[1]
				top[1] = current
			} else if current > top[2] {
				top[2] = current
			}
			current = 0
		} else {
			current += util.Number(l)
		}
	}
	if current > top[0] {
		top[2] = top[1]
		top[1] = top[0]
		top[0] = current
	} else if current > top[1] {
		top[2] = top[1]
		top[1] = current
	} else if current > top[2] {
		top[2] = current
	}
	return top[0] + top[1] + top[2]
}

func Solve() {
	file := "2022/calories/input.txt"
	fmt.Println(max(file))
	fmt.Println(topThree(file))
}
