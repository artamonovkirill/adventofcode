package rucksack

import (
	"fmt"
	"unicode"

	"github.com/advendofcode/util"
)

func duplicates(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		first := line[0 : len(line)/2]
		second := line[len(line)/2:]
		if len(first) != len(second) {
			panic("expected to equal: " + first + ", " + second)
		}
		packed := packed(first)
		duplicates := map[rune]bool{}
		for _, item := range []rune(second) {
			if packed[item] {
				duplicates[item] = true
			}
		}
		for d := range duplicates {
			result += priority(d)
		}
	}
	return result
}

func packed(items string) map[rune]bool {
	result := map[rune]bool{}
	for _, item := range []rune(items) {
		result[item] = true
	}
	return result
}

func identifiers(file string) int {
	result := 0
	lines := util.Lines(file)
	for i := 0; i < len(lines); i += 3 {
		a, b, c := packed(lines[i]), packed(lines[i+1]), packed(lines[i+2])
		for item := range a {
			if b[item] && c[item] {
				result += priority(item)
			}
		}
	}
	return result
}

func priority(letter rune) int {
	if unicode.IsLower(letter) {
		return int(letter - 'a' + 1)
	}
	return int(letter - 'A' + 27)
}

func Solve() {
	file := "2022/rucksack/input.txt"
	fmt.Println(duplicates(file))
	fmt.Println(identifiers(file))
}
