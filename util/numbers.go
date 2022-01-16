package util

import (
	"strconv"
	"strings"
)

func Number(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func ToString(xs []int, separator string) string {
	result := make([]string, len(xs))
	for i, x := range xs {
		result[i] = strconv.Itoa(x)
	}
	return strings.Join(result, separator)
}
