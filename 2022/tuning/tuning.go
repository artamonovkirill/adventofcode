package tuning

import (
	"fmt"
	"github.com/advendofcode/util"
)

func tune(input string, length int) int {
	ri := []rune(input)
	for i := range ri {
		m := make(map[rune]bool)
		for j := 0; j < length; j++ {
			m[ri[i+j]] = true
		}
		if len(m) == length {
			return i + length
		}
	}
	panic("no uniq sequence found")
}

func Solve() {
	file := "2022/tuning/input.txt"
	fmt.Println(tune(util.Text(file), 4))
	fmt.Println(tune(util.Text(file), 14))
}
