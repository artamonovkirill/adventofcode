package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(" +")

func Solve(file string) int {
	result := 0
	lines := util.Lines(file)
	duration := util.Number(strings.Join(r.Split(lines[0], -1)[1:], ""))
	record := util.Number(strings.Join(r.Split(lines[1], -1)[1:], ""))
	fmt.Println(duration, record)
	for push := 1; push < duration; push++ {
		distance := push * (duration - push)
		if distance > record {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/6/input.txt"))
}
