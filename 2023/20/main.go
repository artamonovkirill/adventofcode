package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

type signal struct {
	from, to string
	high     bool
}

func Solve(file string, pushes int) int {
	receivers, modules := parse(file)

	highs := 0
	lows := 0
	for i := 0; i < pushes; i++ {
		signals := []signal{
			{from: "button", to: "broadcaster", high: false},
		}

		for len(signals) > 0 {
			for _, s := range signals {
				if s.high {
					highs++
				} else {
					lows++
				}
			}
			signals = process(receivers, modules, signals)
		}
	}
	return highs * lows
}

func Solve2(file string) map[string][]int {
	receivers, modules := parse(file)

	pushes := 0
	cycles := make(map[string][]int)
	for name := range modules["mg"].(map[string]bool) {
		cycles[name] = []int{}
	}
	for {
		pushes++

		cont := false
		for _, v := range cycles {
			if len(v) < 3 {
				cont = true
				break
			}
		}
		if !cont {
			return cycles
		}

		signals := []signal{
			{from: "button", to: "broadcaster", high: false},
		}

		for len(signals) > 0 {
			signals = process(receivers, modules, signals)
			for name, high := range modules["mg"].(map[string]bool) {
				if high {
					cycles[name] = append(cycles[name], pushes)
				}
			}
		}
	}
}

func process(receivers map[string][]string, modules map[string]interface{}, signals []signal) []signal {
	var nextSignals []signal
signal:
	for _, s := range signals {
		target := modules[s.to]
		rs := receivers[s.to]
		switch target.(type) {
		case bool:
			if !s.high {
				on := target.(bool)
				for _, r := range rs {
					nextSignals = append(nextSignals, signal{
						from: s.to,
						to:   r,
						high: !on,
					})
				}
				modules[s.to] = !on
			}
		case map[string]bool:
			conjunction := target.(map[string]bool)
			conjunction[s.from] = s.high
			for _, high := range conjunction {
				if !high {
					for _, r := range rs {
						nextSignals = append(nextSignals, signal{
							from: s.to,
							to:   r,
							high: true,
						})
					}
					continue signal
				}
			}
			for _, r := range rs {
				nextSignals = append(nextSignals, signal{
					from: s.to,
					to:   r,
					high: false,
				})
			}
		default:
			for _, r := range rs {
				nextSignals = append(nextSignals, signal{
					from: s.to,
					to:   r,
					high: s.high,
				})
			}
		}
	}
	return nextSignals
}

func parse(file string) (map[string][]string, map[string]interface{}) {
	receivers := make(map[string][]string)
	modules := make(map[string]interface{})
	var conjunctions []string
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, " -> ")
		id := parts[0]
		name := id[1:]
		rs := strings.Split(parts[1], ", ")
		if id == "broadcaster" {
			receivers[id] = rs
		} else if id[0] == '%' {
			receivers[name] = rs
			modules[name] = false
		} else if id[0] == '&' {
			conjunctions = append(conjunctions, name)
			receivers[name] = rs
			modules[name] = make(map[string]bool)
		} else {
			panic("not implemented for " + line)
		}
	}
	for name := range modules {
		for _, r := range receivers[name] {
			for _, c := range conjunctions {
				if r == c {
					modules[c].(map[string]bool)[name] = false
				}
			}
		}
	}
	return receivers, modules
}

func main() {
	fmt.Println(Solve("2023/20/input.txt", 1000))
	fmt.Println(Solve2("2023/20/input.txt"))
	fmt.Println(minDivisable(3947, 3793, 4003, 4019))
}

func minDivisable(a int, b int, c int, d int) int {
	i := 0
	for {
		i++
		value := d * i
		if value%a == 0 && value%b == 0 && value%c == 0 {
			return value
		}
	}
}
