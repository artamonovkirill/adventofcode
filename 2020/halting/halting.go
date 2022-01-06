package halting

import (
	"github.com/advendofcode/util"
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
	value := util.Number(argument[1:])
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
