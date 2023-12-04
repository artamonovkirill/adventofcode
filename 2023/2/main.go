package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strconv"
	"strings"
)

type Game struct {
	red, green, blue int
}

func Valid(g Game, file string) int {
	result := 0
line:
	for _, line := range util.Lines(file) {
		picks := strings.Split(line, ": ")
		game := picks[0]
		id, err := strconv.Atoi(strings.Split(game, " ")[1])
		if err != nil {
			panic(err)
		}
		ps := strings.Split(picks[1], "; ")
		for _, p := range ps {
			xs := strings.Split(p, ", ")
			for _, x := range xs {
				ys := strings.Split(x, " ")
				color := ys[1]
				n, err := strconv.Atoi(ys[0])
				if err != nil {
					panic(err)
				}
				switch color {
				case "red":
					if g.red < n {
						continue line
					}
				case "green":
					if g.green < n {
						continue line
					}
				case "blue":
					if g.blue < n {
						continue line
					}
				default:
					panic("unknown: " + color)
				}
			}
		}
		result += id
	}
	return result
}

func Power(file string) int {
	result := 0
	for _, line := range util.Lines(file) {
		red := 0
		green := 0
		blue := 0
		picks := strings.Split(line, ": ")
		ps := strings.Split(picks[1], "; ")
		for _, p := range ps {
			xs := strings.Split(p, ", ")
			for _, x := range xs {
				ys := strings.Split(x, " ")
				color := ys[1]
				n, err := strconv.Atoi(ys[0])
				if err != nil {
					panic(err)
				}
				switch color {
				case "red":
					if red < n {
						red = n
					}
				case "green":
					if green < n {
						green = n
					}
				case "blue":
					if blue < n {
						blue = n
					}
				default:
					panic("unknown: " + color)
				}
			}
		}
		result += red * green * blue
	}
	return result
}

func main() {
	fmt.Println(Valid(Game{12, 13, 14}, "2023/2/input.txt"))
	fmt.Println(Power("2023/2/input.txt"))
}
