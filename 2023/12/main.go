package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Solve(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, " ")
		visual := unfold(parts[0], "?")
		r := regexp.MustCompile(rx(unfold(parts[1], ",")))
		unknowns := strings.Count(visual, "?")
		matches := 0
		for i := 0; i < int(math.Pow(2, float64(unknowns))); i++ {
			b := binary(i, unknowns)
			j := 0
			var candidate []int32
			for _, c := range visual {
				if c == '?' {
					x := b[j]
					if x == '0' {
						candidate = append(candidate, '.')
					} else {
						candidate = append(candidate, '#')
					}
					j++
				} else {
					candidate = append(candidate, c)
				}
			}
			if r.MatchString(string(candidate)) {
				matches++
			}
		}
		result += matches
	}
	return result
}

func unfold(pattern, separator string) string {
	result := make([]string, 1)
	for i := 0; i < 1; i++ {
		result[i] = pattern
	}
	return strings.Join(result, separator)
}

func rx(counts string) string {
	result := "^\\.*"
	cs := strings.Split(counts, ",")
	for i, count := range cs {
		result += "#{" + count + "}"
		if i < len(cs)-1 {
			result += "\\.+"
		}
	}
	result += "\\.*$"
	return result
}

func binary(i, length int) string {
	result := strconv.FormatInt(int64(i), 2)
	for len(result) < length {
		result = "0" + result
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/12/input.txt"))
}
