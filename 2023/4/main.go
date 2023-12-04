package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"math"
	"strings"
)

func Points(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		matches := 0
		card := strings.Split(strings.Split(line, ": ")[1], " | ")
		winners := strings.Split(card[0], " ")
		values := strings.Split(card[1], " ")
		for _, winner := range winners {
			if winner != "" {
				for _, value := range values {
					if winner == value {
						matches++
					}
				}
			}
		}
		if matches > 0 {
			score := pow2(matches)
			result += score
		}
	}
	return result
}

func pow2(matches int) int {
	return int(math.Pow(float64(2), float64(matches-1)))
}

func main() {
	fmt.Println(Points("2023/4/input.txt"))
	fmt.Println(Cards("2023/4/input.txt"))
}

func Cards(file string) int {
	result := 0
	lines := util.Lines(file)
	var current []int
	for i := range lines {
		current = append(current, i)
	}
	for {
		var next []int
		for _, i := range current {
			line := lines[i]
			matches := 0
			card := strings.Split(strings.Split(line, ": ")[1], " | ")
			winners := strings.Split(card[0], " ")
			values := strings.Split(card[1], " ")
			for _, winner := range winners {
				if winner != "" {
					for _, value := range values {
						if winner == value {
							matches++
						}
					}
				}
			}
			for j := 0; j < matches; j++ {
				next = append(next, i+j+1)
			}
		}
		result += len(current)
		if len(next) == 0 {
			return result
		}
		current = next
		fmt.Println(len(current), "to go")
	}
}
