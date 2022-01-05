package halting

import (
	"fmt"
	"github.com/advendofcode/util"
	"strconv"
)

func solve(file string) int {
	commands := util.Lines(file)
	for i, command := range commands {
		operation, argument := util.Split(command, " ")
		if operation == "acc" {
			continue
		}

		pointer := 0
		acc := 0
		executed := map[int]bool{}

		oldCommand := command

		if operation == "nop" {
			commands[i] = "jmp " + argument
		} else {
			commands[i] = "nop " + argument
		}

		for {
			if executed[pointer] == true {
				commands[i] = oldCommand
				break
			}
			if pointer >= len(commands) {
				return acc
			}
			executed[pointer] = true
			pointer, acc = process(commands, pointer, acc)
		}
	}
	return 0
}

func process(commands []string, i int, acc int) (int, int) {
	command := commands[i]
	operation, argument := util.Split(command, " ")
	switch operation {
	default:
		return i + 1, acc
	case "acc":
		return i + 1, adjust(argument)(acc)
	case "jmp":
		return adjust(argument)(i), acc
	}
}

func adjust(argument string) func(int) int {
	value, err := strconv.Atoi(argument[1:])
	if err != nil {
		panic(err)
	}
	sign := argument[0]
	if sign == '+' {
		return func(i int) int {
			return i + value
		}
	} else {
		return func(i int) int {
			return i - value
		}
	}
}

func Solve() {
	input := "2020/halting/puzzle.txt"
	fmt.Println(solve(input))
}
