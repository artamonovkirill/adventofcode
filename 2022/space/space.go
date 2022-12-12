package space

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type directory struct {
	contents map[string]interface{}
	parent   interface{}
}

func smallest(file string) int {
	var root directory
	var current *directory
	for _, line := range util.Lines(file) {
		switch line {
		case "$ cd /":
			root = directory{contents: make(map[string]interface{})}
			current = &root
		case "$ ls":
			continue
		case "$ cd ..":
			newCurrent := current.parent.(*directory)
			current = newCurrent
		default:
			if line[0] == '$' {
				parts := strings.Split(line, " ")
				newCurrent := current.contents[parts[2]].(directory)
				current = &newCurrent
			} else {
				parts := strings.Split(line, " ")
				if parts[0] == "dir" {
					current.contents[parts[1]] = directory{
						contents: make(map[string]interface{}),
						parent:   current,
					}
				} else {
					current.contents[parts[1]] = util.Number(parts[0])
				}
			}
		}
	}
	result := &Result{}
	sum(root, result)
	sort.Ints(result.value)
	free := 70_000_000 - result.value[len(result.value)-1]
	for _, size := range result.value {
		if size+free >= 30_000_000 {
			return size
		}
	}
	panic("no solution found")
}

type Result struct {
	value []int
}

func sum(dir directory, acc *Result) int {
	size := 0
	for _, v := range dir.contents {
		switch v.(type) {
		case int:
			size += v.(int)
		case directory:
			size += sum(v.(directory), acc)
		default:
			panic(fmt.Sprintf("not implemented for %T", v))
		}
	}
	acc.value = append(acc.value, size)
	return size
}

func Solve() {
	file := "2022/space/input.txt"
	fmt.Println(smallest(file))
}
