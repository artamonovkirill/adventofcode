package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type Mapping struct {
	next    string
	entries []Entry
}

type Range struct {
	start, end int
}

type Entry struct {
	destination, source Range
}

func Solve(file string) int {
	text := strings.Split(util.Text(file), "\n\n")
	current := parseSeeds(text[0])
	mappings := parseMappings(text)
	currentType := "seed"
	for {
		mapping := mappings[currentType]
		var next []Range
		fmt.Println("from", currentType, "to", mapping.next)
	outer:
		for _, c := range current {
			rest := c
			for _, m := range mapping.entries {
				diff := m.destination.start - m.source.start
				if m.source.start <= rest.start && rest.end <= m.source.end {
					next = append(next, Range{rest.start + diff, rest.end + diff})
					continue outer
				} else if m.source.start > rest.end {
					continue
				} else if m.source.end < rest.start {
					continue
				} else if m.source.start <= rest.start && m.source.end <= rest.end {
					next = append(next, Range{rest.start + diff, m.source.end + diff})
					rest = Range{m.source.end + 1, rest.end}
				} else if m.source.start >= rest.start && m.source.end >= rest.end {
					next = append(next, Range{m.source.start + diff, rest.end + diff})
					rest = Range{rest.start, m.source.start - 1}
				} else {
					fmt.Println(c, m)
					panic("now what?")
				}
				if rest.start >= rest.end {
					panic("what now?")
				}
			}
			next = append(next, rest)
		}
		currentType = mapping.next
		current = next
		if currentType == "location" {
			return Min(next)
		}
	}
}

func parseMappings(text []string) map[string]Mapping {
	maps := make(map[string]Mapping)
	for i := 1; i < len(text); i++ {
		input := strings.Split(text[i], "\n")
		header := strings.Split(strings.ReplaceAll(input[0], " map:", ""), "-to-")
		entries := make([]Entry, len(input)-1)
		for j := 1; j < len(input); j++ {
			entry := strings.Split(input[j], " ")
			length := util.Number(entry[2])
			destination := util.Number(entry[0])
			source := util.Number(entry[1])
			entries[j-1] = Entry{
				Range{destination, destination + length - 1},
				Range{source, source + length - 1},
			}
		}
		sort.Slice(entries, func(x, y int) bool {
			return entries[x].source.start < entries[y].source.start
		})
		m := Mapping{
			next:    header[1],
			entries: entries,
		}
		maps[header[0]] = m
	}
	return maps
}

func Min(xs []Range) int {
	result := xs[0].start
	for i := 1; i < len(xs); i++ {
		if xs[i].start < result {
			result = xs[i].start
		}
	}
	return result
}

func parseSeeds(input string) []Range {
	var result []Range
	parts := strings.Split(strings.Split(input, ": ")[1], " ")
	for i := 0; i <= len(parts)/2; i += 2 {
		start := util.Number(parts[i])
		length := util.Number(parts[i+1])
		result = append(result, Range{start, start + length - 1})
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/5/input.txt"))
}
