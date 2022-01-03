package util

import (
	"bufio"
	"os"
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
