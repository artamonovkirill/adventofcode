package lobby

import (
	"fmt"
	"github.com/advendofcode/util"
	"regexp"
)

type Point struct {
	x, y, z int
}

type Lobby = map[Point]bool

func initialize(file string) Lobby {
	lines := util.Lines(file)
	tiles := Lobby{}
	for _, line := range lines {
		directions := regexp.MustCompile("(e|se|sw|w|nw|ne)").FindAllString(line, -1)
		p := Point{0, 0, 0}
		for _, d := range directions {
			p = move(p, d)
		}
		flipped := tiles[p]
		if flipped {
			delete(tiles, p)
		} else {
			tiles[p] = true
		}
	}
	return tiles
}

func move(p Point, direction string) Point {
	v := vector(direction)
	return Point{p.x + v.x, p.y + v.y, p.z + v.z}
}

var ns = map[string]Point{
	"ne": {0, 1, 1},
	"e":  {1, 1, 0},
	"se": {1, 0, -1},
	"sw": {0, -1, -1},
	"w":  {-1, -1, 0},
	"nw": {-1, 0, 1},
}

// https://homepages.inf.ed.ac.uk/rbf/CVonline/LOCAL_COPIES/AV0405/MARTIN/Hex.pdf
func vector(direction string) Point {
	v, ok := ns[direction]
	if !ok {
		panic("not implemented for " + direction)
	}
	return v
}

func advance(l Lobby, steps int) Lobby {
	if steps == 0 {
		return l
	}

	result := Lobby{}
	for p := range l {
		b := black(neighbours(p), l)
		if b == 1 || b == 2 {
			result[p] = true
		}
		for _, n := range neighbours(p) {
			if !l[n] {
				b := black(neighbours(n), l)
				if b == 2 {
					result[n] = true
				}
			}
		}
	}
	return advance(result, steps-1)
}

func black(ps []Point, l Lobby) int {
	result := 0
	for _, p := range ps {
		if l[p] {
			result++
		}
	}
	return result
}

func neighbours(p Point) []Point {
	result := make([]Point, len(ns))
	i := 0
	for _, n := range ns {
		result[i] = Point{p.x + n.x, p.y + n.y, p.z + n.z}
		i++
	}
	return result
}

func Solve() {
	init := initialize("2020/lobby/puzzle.txt")
	fmt.Println(len(init))
	fmt.Println(len(advance(init, 100)))
}
