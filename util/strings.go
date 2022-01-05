package util

import (
	"regexp"
)

func Split(input string, re string) (string, string) {
	parts := regexp.MustCompile(re).Split(input, 2)
	return parts[0], parts[1]
}
