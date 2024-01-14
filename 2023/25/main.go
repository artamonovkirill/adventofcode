package main

import (
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type edge struct {
	from, to string
}

func Solve(file string) int {
	var edges []edge
	var vertices []string
	graph := make(map[string][]string)
	for _, line := range util.Lines(file) {
		parts := strings.Split(line, ": ")
		from := parts[0]
		graph[from] = []string{}
		vertices = append(vertices, from)
		tos := strings.Split(parts[1], " ")
		for _, to := range tos {
			graph[to] = []string{}
			vertices = append(vertices, to)
			edges = append(edges, newEdge(from, to))
		}
	}
	if len(edges) != len(unique(edges)) {
		panic("how?")
	}
	vertices = unique(vertices)

	for _, e := range edges {
		graph[e.from] = append(graph[e.from], e.to)
		graph[e.to] = append(graph[e.to], e.from)
	}

	counts := make(map[edge]int)
	for vertice := range graph {
		paths := make(map[string][]edge)
		visited := map[string]bool{
			vertice: true,
		}
		current := []string{vertice}

		for len(current) > 0 {
			var next []string

			for _, c := range current {
				for _, e := range graph[c] {
					if !visited[e] {
						visited[e] = true
						next = append(next, e)
						paths[e] = append(paths[c], newEdge(c, e))
					}
				}
			}

			current = next
		}

		for _, path := range paths {
			for _, e := range path {
				counts[e]++
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return counts[edges[i]] > counts[edges[j]]
	})

	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			for k := j + 1; k < len(edges); k++ {
				visited := make(map[string]bool)
				excluded := []edge{edges[i], edges[j], edges[k]}
				var group []string
				var groups [][]string
				for _, vertice := range vertices {
					if !visited[vertice] {
						visited[vertice] = true
						group = append(group, vertice)

						current := []string{vertice}

						for len(current) > 0 {
							var next []string

							for _, c := range current {
								for _, e := range graph[c] {
									if !visited[e] && !contains(excluded, newEdge(c, e)) {
										visited[e] = true
										next = append(next, e)
										group = append(group, e)
									}
								}
							}

							current = next
						}

						groups = append(groups, group)
						group = []string{}
					}
				}
				if len(groups) == 2 {
					return len(groups[0]) * len(groups[1])
				}
			}
		}
	}

	panic("no solution found")
}

func contains(xs []edge, e edge) bool {
	for _, x := range xs {
		if x == e {
			return true
		}
	}
	return false
}

func newEdge(from string, to string) edge {
	if from < to {
		return edge{from, to}
	} else if to < from {
		return edge{to, from}
	} else {
		panic("how?")
	}
}

func unique[T string | edge](xs []T) []T {
	set := make(map[T]bool)
	for _, x := range xs {
		set[x] = true
	}
	var result []T
	for k := range set {
		result = append(result, k)
	}
	return result
}

func main() {
	fmt.Println(Solve("2023/25/input.txt"))
}
