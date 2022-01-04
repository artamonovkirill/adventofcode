package password

import (
	"regexp"
	"strconv"
)

func valid(input []string) int {
	result := 0
	for _, line := range input {
		rule, password := split(line, ": ")
		rng, char := split(rule, " ")
		fst, snd := split(rng, "-")

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

func split(input string, re string) (string, string) {
	parts := regexp.MustCompile(re).Split(input, -1)
	return parts[0], parts[1]
}
