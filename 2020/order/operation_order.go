package order

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
	"strconv"
	"strings"
)

func solve(expression string) int {
	brackets := regexp.MustCompile("\\([^()]+\\)")
	first := brackets.FindString(expression)
	if first != "" {
		replacement := solve(first[1 : len(first)-1])
		simplified := replace(expression, first, replacement)
		return solve(simplified)
	}

	addition := regexp.MustCompile("[0-9]+ \\+ [0-9]+")
	first = addition.FindString(expression)
	if first != "" {
		a, b := util.Split(first, " \\+ ")
		sum := util.Number(a) + util.Number(b)
		simplified := replace(expression, first, sum)
		return solve(simplified)
	}

	multiplication := regexp.MustCompile("[0-9]+ \\* [0-9]+")
	first = multiplication.FindString(expression)
	if first != "" {
		a, b := util.Split(first, " \\* ")
		product := util.Number(a) * util.Number(b)
		simplified := replace(expression, first, product)
		return solve(simplified)
	}
	return util.Number(expression)
}

func replace(expression string, first string, replacement int) string {
	return strings.Replace(expression, first, strconv.Itoa(replacement), 1)
}

func fold(xs []string, init int, f func(int, string) int) int {
	if len(xs) == 0 {
		return init
	}
	return fold(xs[1:], f(init, xs[0]), f)
}

func Solve() {
	input := "2020/order/puzzle.txt"
	lines := util.Lines(input)
	fmt.Println(fold(lines, 0, func(acc int, line string) int {
		return acc + solve(line)
	}))
}
