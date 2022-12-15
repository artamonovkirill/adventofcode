package monkeys

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type monkey struct {
	n           int
	inspections int
	items       []map[int]int
	divisor     int
	onTrue      int
	onFalse     int
	operation   func(int) int
}

func business(file string, rounds int) int {
	var divisors []int
	lines := util.Lines(file)
	for i, line := range lines {
		if i%7 == 3 {
			line = strings.ReplaceAll(line, "  Test: divisible by ", "")
			divisors = append(divisors, util.Number(line))
		}
	}
	monkeys := parse(lines, divisors)
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.inspections = m.inspections + 1
				for k, v := range item {
					item[k] = m.operation(v) % k
				}
				var to *monkey
				if item[m.divisor] == 0 {
					to = monkeys[m.onTrue]
				} else {
					to = monkeys[m.onFalse]
				}
				to.items = append(to.items, item)
			}
			m.items = []map[int]int{}
		}
	}
	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspections
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	return inspections[0] * inspections[1]
}

func parse(lines []string, divisors []int) []*monkey {
	var monkeys []*monkey
	for i, line := range lines {
		switch i % 7 {
		case 0:
			monkeys = append(monkeys, &monkey{
				n: i / 7,
			})
		case 1:
			line = strings.ReplaceAll(line, "Starting items:", "")
			line = strings.ReplaceAll(line, " ", "")
			items := strings.Split(line, ",")
			monkeys[i/7].items = make([]map[int]int, len(items))
			for j, item := range items {
				monkeys[i/7].items[j] = make(map[int]int)
				for _, d := range divisors {
					monkeys[i/7].items[j][d] = util.Number(item) % d
				}
			}
		case 2:
			line = strings.ReplaceAll(line, "Operation:", "")
			line = strings.ReplaceAll(line, " ", "")
			if line == "new=old*old" {
				monkeys[i/7].operation = func(i int) int {
					return i * i
				}
			} else if strings.Contains(line, "+") {
				parts := strings.Split(line, "+")
				monkeys[i/7].operation = func(i int) int {
					return i + util.Number(parts[1])
				}
			} else if strings.Contains(line, "*") {
				parts := strings.Split(line, "*")
				monkeys[i/7].operation = func(i int) int {
					return i * util.Number(parts[1])
				}
			} else {
				panic("not implemented for " + line)
			}
		case 3:
			line = strings.ReplaceAll(line, "  Test: divisible by ", "")
			monkeys[i/7].divisor = util.Number(line)
		case 4:
			line = strings.ReplaceAll(line, "    If true: throw to monkey ", "")
			monkeys[i/7].onTrue = util.Number(line)
		case 5:
			line = strings.ReplaceAll(line, "    If false: throw to monkey ", "")
			monkeys[i/7].onFalse = util.Number(line)
		case 6:
			continue
		default:
			panic("not implemented for " + line)
		}
	}
	return monkeys
}

func Solve() {
	file := "2022/monkeys/input.txt"
	fmt.Println(business(file, 10000))
}
