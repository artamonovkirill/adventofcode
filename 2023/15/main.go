package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"strings"
)

func Hash(file string) int {
	result := 0
	for _, e := range strings.Split(util.Text(file), ",") {
		result += hash(e)
	}
	return result
}

func hash(input string) int {
	result := 0
	for _, c := range input {
		result += int(c)
		result *= 17
		result = result % 256
	}
	return result
}

type lens struct {
	label  string
	length int
	next   *lens
}

func lenses(input string) int {
	boxes := make([]*lens, 256)
	for _, e := range strings.Split(input, ",") {
		var label string
		if e[len(e)-1] == '-' {
			label = e[0 : len(e)-1]
		} else {
			label = e[0 : len(e)-2]
		}
		op := e[len(label)]
		i := hash(label)

		if op == '=' {
			length := int(e[len(label)+1] - '0')
			newLens := &lens{label: label, length: length}

			existing := boxes[i]
			if existing == nil {
				boxes[i] = newLens
			} else {
				// first in box matches
				if newLens.label == existing.label {
					newLens.next = existing.next
					boxes[i] = newLens
				} else {
					// next matches
					for existing != nil {
						if existing.next == nil {
							existing.next = newLens
							break
						} else if newLens.label == existing.next.label {
							newLens.next = existing.next.next
							existing.next = newLens
							break
						}
						existing = existing.next
					}
				}
			}
		} else {
			existing := boxes[i]
			if existing != nil {
				// first in box matches
				if existing.label == label {
					boxes[i] = existing.next
				} else {
					for existing != nil {
						if existing.next != nil && existing.next.label == label {
							existing.next = existing.next.next
						}
						existing = existing.next
					}
				}
			}
		}
	}

	result := 0
	for i, l := range boxes {
		index := 1
		for l != nil {
			power := (i + 1) * index * l.length
			result += power
			l = l.next
			index++
		}
	}
	return result
}

func Lenses(file string) int {
	return lenses(util.Text(file))
}

func main() {
	fmt.Println(Hash("2023/15/input.txt"))
	fmt.Println(Lenses("2023/15/input.txt"))
}
