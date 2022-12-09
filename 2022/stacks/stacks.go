package stacks

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func process(file string) string {
	var start []string
	var commands []string
	var size int
	for _, line := range util.Lines(file) {
		if strings.Contains(line, "[") {
			start = append(start, line)
		} else if strings.Contains(line, "move") {
			commands = append(commands, line)
		} else if strings.Contains(line, " 1   2   3") {
			size = util.Number(string(line[len(line)-1]))
		}
	}
	stack := initialize(start, size)
	apply(stack, commands)
	result := ""
	for _, column := range stack {
		result += column[0]
	}
	return result
}

func initialize(start []string, size int) [][]string {
	result := make([][]string, size)
	for _, line := range start {
		for i := 0; i < len(line); i += 4 {
			value := string(line[i+1])
			if value != " " {
				result[i/4] = push(result[i/4], value)
			}
		}
	}
	return result
}

func push(stack []string, value string) []string {
	if len(stack) == 0 {
		return []string{value}
	}
	return append(stack, value)
}

func apply(stack [][]string, commands []string) {
	for _, command := range commands {
		parts := strings.Split(command, " ")
		count := util.Number(parts[1])
		from := util.Number(parts[3]) - 1
		to := util.Number(parts[5]) - 1
		var load []string
		for i := 0; i < count; i++ {
			load = append(load, stack[from][0])
			stack[from] = stack[from][1:]
		}
		stack[to] = append(load, stack[to]...)
	}
}

func Solve() {
	file := "2022/stacks/input.txt"
	fmt.Println(process(file))
}
