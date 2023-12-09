package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func PredictFuture(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		values := parse(line)
		result += predictNext(values)
	}
	return result
}

func predictNext(values []int) int {
	current := values
	data := [][]int{current}
outer:
	for {
		var next []int
		for i := 1; i < len(current); i++ {
			next = append(next, current[i]-current[i-1])
		}
		data = append(data, next)
		for _, n := range next {
			if n != 0 {
				current = next
				continue outer
			}
		}
		for i := len(data) - 2; i >= 0; i-- {
			prev := data[i]
			curr := data[i+1]
			data[i] = append(prev, prev[len(prev)-1]+curr[len(curr)-1])
		}
		initial := data[0]
		return initial[len(initial)-1]
	}
}

func parse(line string) []int {
	var result []int
	for _, e := range strings.Split(line, " ") {
		result = append(result, util.Number(e))
	}
	return result
}

func ExtrapolateHistory(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		values := parse(line)
		result += extrapolatePrevious(values)
	}
	return result
}

func extrapolatePrevious(values []int) int {
	current := values
	data := [][]int{current}
outer:
	for {
		var next []int
		for i := 1; i < len(current); i++ {
			next = append(next, current[i]-current[i-1])
		}
		data = append(data, next)
		for _, n := range next {
			if n != 0 {
				current = next
				continue outer
			}
		}
		hist := make([]int, len(data))
		hist[len(data)-1] = 0
		for i := len(data) - 2; i >= 0; i-- {
			prev := hist[i+1]
			curr := data[i]
			hist[i] = curr[0] - prev
		}
		return hist[0]
	}
}

func main() {
	fmt.Println(PredictFuture("2023/9/input.txt"))
	fmt.Println(ExtrapolateHistory("2023/9/input.txt"))
}
