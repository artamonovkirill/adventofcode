package util

import (
	"bufio"
	"os"
	"strconv"
)

// Lines reads file line by line
func Lines(file string) []int {
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
	var result []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
