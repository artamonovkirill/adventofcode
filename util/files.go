package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

// Lines reads file line by line
func Lines(file string) []string {
	f, err := os.Open(file)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func Numbers(file string) []int {
	lines := Lines(file)
	result := make([]int, len(lines))
	for i, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		result[i] = n
	}
	return result
}

func Text(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(content)
}
