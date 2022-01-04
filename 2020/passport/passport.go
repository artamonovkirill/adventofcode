package passport

import (
	"errors"
	"fmt"
	"github.com/advendofcode/util"
	mapset "github.com/deckarep/golang-set"
	"regexp"
	"sort"
	"strings"
)

var fields = map[string]string{
	"byr": "(19[2-9][0-9]|200[1-2])",
	"iyr": "20(1[0-9]|20)",
	"eyr": "20(2[0-9]|30)",
	"hgt": "(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)",
	"hcl": "#[0-9a-f]{6}",
	"ecl": "(amb|blu|brn|gry|grn|hzl|oth)",
	"pid": "[0-9]{9}",
}

func solve(file string) int {
	result := 0
	text := util.Text(file)
	entries := split(text, "\n\n")
	for _, entry := range entries {
		if valid(entry) {
			fmt.Println(entry + " is valid")
			result++
		}
	}
	unique(entries, "byr")
	return result
}

func unique(entries []string, field string) {
	result := mapset.NewSet()
	for _, entry := range entries {
		match, err := find(entry, field+":[^\n ]+")
		if valid(entry) {
			if err == nil {
				value := strings.ReplaceAll(match, field+":", "")
				result.Add(value)
			}
		}
	}
	var s []string
	for _, r := range result.ToSlice() {
		s = append(s, r.(string))
	}
	sort.Strings(s)
	fmt.Println(s)
}

func valid(entry string) bool {
	for k, v := range fields {
		full := fmt.Sprintf("%s:%s([ \n]|$)", k, v)
		_, err := find(entry, full)
		if err != nil {
			return false
		}
	}
	return true
}

func find(input string, re string) (string, error) {
	match := regexp.MustCompile(re).FindAllString(input, -1)
	if len(match) == 0 {
		return "", errors.New("No matches found for " + re + " in " + input)
	}
	if len(match) > 1 {
		panic("Too many matches found for " + re + " in " + input)
	}
	return match[0], nil
}

func split(input string, re string) []string {
	var result []string
	matches := regexp.MustCompile(re).Split(input, -1)
	for _, m := range matches {
		result = append(result, strings.ReplaceAll(m, "\n", " "))
	}
	return result
}

func Solve() {
	input := "2020/passport/puzzle.txt"
	fmt.Println(solve(input))
}
