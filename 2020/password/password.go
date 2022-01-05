package password

import (
	"github.com/advendofcode/util"
	"strconv"
)

func valid(input []string) int {
	result := 0
	for _, line := range input {
		rule, password := util.Split(line, ": ")
		rng, char := util.Split(rule, " ")
		fst, snd := util.Split(rng, "-")

		if (password[atoi(fst)-1] == char[0]) != (password[atoi(snd)-1] == char[0]) {
			result++
		}
	}
	return result
}

func atoi(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
