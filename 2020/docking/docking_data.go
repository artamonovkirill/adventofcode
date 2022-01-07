package docking

import (
	"fmt"
	"github.com/advendofcode/util"
	"strconv"
	"strings"
)

func solve(file string) int {
	memory := map[int]int{}
	text := util.Text(file)
	chunks := strings.Split(text, "mask = ")
	for _, chunk := range chunks {
		if chunk != "" {
			applyCommands(chunk, memory)
		}
	}
	result := 0
	for _, v := range memory {
		result += v
	}
	return result
}

func applyCommands(chunk string, memory map[int]int) {
	lines := strings.Split(chunk, "\n")
	mask := lines[0]
	if len(mask) != 36 {
		panic("Not implemented for " + mask)
	}
	commands := lines[1:]
	for _, command := range commands {
		if command != "" {
			applyCommand(command, mask, memory)
		}
	}
}

func applyCommand(command string, mask string, memory map[int]int) {
	parts := strings.Split(command, " = ")
	address := util.Number(strings.ReplaceAll(
		strings.ReplaceAll(parts[0], "mem[", ""),
		"]", ""))
	value := util.Number(parts[1])
	binary := fmt.Sprintf("%036b", address)
	addresses := []string{""}
	for i := 0; i < len(binary); i++ {
		addresses = next(addresses, binary[i:i+1], mask[i:i+1])
	}
	for _, a := range addresses {
		n, err := strconv.ParseInt(a, 2, 64)
		if err != nil {
			panic(err)
		}
		memory[int(n)] = value
	}
}

func next(addresses []string, a string, m string) []string {
	var result []string
	for _, address := range addresses {
		switch m {
		case "0":
			result = append(result, address+a)
		case "1":
			result = append(result, address+"1")
		case "X":
			result = append(result, address+"0")
			result = append(result, address+"1")
		default:
			panic("not implemented for " + m)
		}
	}
	return result
}
