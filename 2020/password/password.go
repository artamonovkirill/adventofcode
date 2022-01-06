package password

import (
	"github.com/advendofcode/util"
)

func valid(input []string) int {
	result := 0
	for _, line := range input {
		rule, password := util.Split(line, ": ")
		rng, char := util.Split(rule, " ")
		fst, snd := util.Split(rng, "-")

		if (password[util.Number(fst)-1] == char[0]) != (password[util.Number(snd)-1] == char[0]) {
			result++
		}
	}
	return result
}
