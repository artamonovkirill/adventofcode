package report

import (
	"fmt"
	"github.com/advendofcode/util"
)

func solve(report []int) int {
	for i, a := range report {
		for j, b := range report[i+1:] {
			for _, c := range report[j+1:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	panic("no solution found")
}

func Solve() {
	report := util.Lines("2020/report/report.txt")
	fmt.Println(solve(report))
}
