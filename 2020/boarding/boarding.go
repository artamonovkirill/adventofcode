package boarding

import (
	"fmt"
	"github.com/advendofcode/util"
	"strconv"
)

type Pass struct {
	Column int
	Row    int
	ID     int
}

func parse(input string) Pass {
	r := row(input[0:7])
	c := column(input[7:])
	id := r*8 + c
	return Pass{Row: r, Column: c, ID: id}
}

func column(input string) int {
	return choose(input, 0, 7, 'L')
}

func row(input string) int {
	return choose(input, 0, 127, 'F')
}

func choose(input string, low int, high int, down uint8) int {
	if low == high {
		return low
	}
	if input[0] == down {
		return choose(input[1:], low, (low+high)/2, down)
	} else {
		return choose(input[1:], (low+high)/2+1, high, down)
	}
}

func Solve() {
	input := "2020/boarding/puzzle.txt"
	passes := util.Lines(input)
	plane(passes)
	fmt.Println(82*8 + 3)
}

func plane(passes []string) {
	plane := make(map[int]map[int]string, 128)
	for _, pass := range passes {
		p := parse(pass)
		if plane[p.Row] == nil {
			plane[p.Row] = make(map[int]string, 8)
		}
		plane[p.Row][p.Column] = "X"
	}

	for i := 0; i < 128; i++ {
		print(strconv.Itoa(i) + ": ")
		for j := 0; j < 8; j++ {
			v := "."
			if plane[i][j] != "" {
				v = plane[i][j]
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}
