package distress

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
)

func count(file string) int {
	lines := util.Lines(file)
	result := 0
	for i := 0; i < len(lines); i += 3 {
		a := parse(lines[i])
		b := parse(lines[i+1])
		if compare(a, b) == right {
			result += i/3 + 1
		}
	}
	return result
}

func arrange(file string) int {
	var packets [][]interface{}
	for _, line := range util.Lines(file) {
		if line != "" {
			packets = append(packets, parse(line))
		}
	}
	packets = append(packets, parse("[[2]]"))
	packets = append(packets, parse("[[6]]"))
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == right
	})
	result := 1
	for i, packet := range packets {
		if len(packet) == 1 {
			switch packet[0].(type) {
			case []interface{}:
				p := packet[0].([]interface{})
				if len(p) == 1 {
					switch p[0] {
					case 2:
						result *= i + 1
					case 6:
						result *= i + 1
					}
				}
			}
		}
	}
	return result
}

type array struct {
	parent interface{}
	value  []interface{}
}

func parse(s string) []interface{} {
	var current array
	stack := ""
	init := len(s)
	for len(s) > 0 {
		switch s[0] {
		case ',':
			if len(stack) > 0 {
				current.value = append(current.value, util.Number(stack))
				stack = ""
			}
		case '[':
			if init != len(s) {
				current = array{parent: current, value: []interface{}{}}
			}
		case ']':
			if len(stack) > 0 {
				current.value = append(current.value, util.Number(stack))
				stack = ""
			}
			if current.parent == nil {
				return current.value
			}
			parent := current.parent.(array)
			parent.value = append(parent.value, current.value)
			current = parent
		default:
			if s[0] >= '0' && s[0] <= '9' {
				stack += s[0:1]
			} else {
				panic("not implemented for " + s)
			}
		}
		s = s[1:]
	}
	return current.value
}

const (
	right = iota
	unknown
	wrong
)

func compare(as []interface{}, bs []interface{}) int {
	if len(bs) == 0 && len(as) > 0 {
		return wrong
	}
	if len(as) == 0 && len(bs) > 0 {
		return right
	}
	for i, a := range as {
		if i == len(bs) {
			return wrong
		}
		b := bs[i]
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				if a.(int) > b.(int) {
					return wrong
				}
				if a.(int) < b.(int) {
					return right
				}
			default:
				result := compare([]interface{}{a.(int)}, b.([]interface{}))
				if result != unknown {
					return result
				}
			}
		default:
			switch b.(type) {
			case int:
				result := compare(a.([]interface{}), []interface{}{b.(int)})
				if result != unknown {
					return result
				}
			default:
				result := compare(a.([]interface{}), b.([]interface{}))
				if result != unknown {
					return result
				}
			}
		}
	}
	if len(as) < len(bs) {
		return right
	}
	return unknown
}

func Solve() {
	file := "2022/distress/input.txt"
	fmt.Println(count(file))
	fmt.Println(arrange(file))
}
