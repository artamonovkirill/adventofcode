package jigsaw

import (
	"errors"
	"fmt"
	"github.com/advendofcode/util"
	"reflect"
	"regexp"
	"strings"
)

type Image = map[int]map[int]Tile

type Pixels struct {
	value [][]bool
}
type Edge = []bool

type Direction int

const (
	Left = iota
	Right
	Up
	Down
)

func (d Direction) ToString() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	default:
		panic("not implemented")
	}
}

type Tile struct {
	ID int
	I  int
}

type Point struct {
	x int
	y int
}

func free(i Image, p Point) []Direction {
	var result []Direction
	if upper, ok := i[p.y-1]; ok {
		_, ok = upper[p.x]
		if !ok {
			result = append(result, Up)
		}
	} else {
		result = append(result, Up)
	}
	if lower, ok := i[p.y+1]; ok {
		_, ok = lower[p.x]
		if !ok {
			result = append(result, Down)
		}
	} else {
		result = append(result, Down)
	}
	if _, ok := i[p.y][p.x-1]; !ok {
		result = append(result, Left)
	}
	if _, ok := i[p.y][p.x+1]; !ok {
		result = append(result, Right)
	}
	return result
}

func pixelsToString(ps Pixels) string {
	rows := make([]string, len(ps.value))
	for j, r := range ps.value {
		row := ""
		for _, v := range r {
			if v {
				row += "#"
			} else {
				row += "."
			}
		}
		rows[j] = row
	}
	return strings.Join(rows, "\n")
}

func corners(file string) int {
	tiles := parse(file)
	ids := keys(tiles)
	head := ids[0]
	init := Image{0: map[int]Tile{0: {head, 0}}}
	image, err := arrange(init, remove(ids, head), tiles)

	if err != nil {
		panic(err)
	}

	result := 1
	for _, c := range findCorners(image) {
		result *= c
	}
	return result
}

func monsterFree(file string, monsterFile string) int {
	tiles := parse(file)
	ids := keys(tiles)
	head := ids[0]
	init := Image{0: map[int]Tile{0: {head, 0}}}
	arrangement, err := arrange(init, remove(ids, head), tiles)
	if err != nil {
		panic(err)
	}
	image := join(arrangement, tiles)
	monster := parsePixels(util.Lines(monsterFile))
	monsters := map[int]map[int]bool{}
	for _, r := range alternate(monster) {
		for y := 0; y < len(image.value)-len(r.value)+1; y++ {
			for x := 0; x < len(image.value[y])-len(r.value[0])+1; x++ {
				for _, p := range matches(image, r, Point{x, y}) {
					if _, ok := monsters[p.y]; !ok {
						monsters[p.y] = map[int]bool{}
					}
					monsters[p.y][p.x] = true
				}
			}
		}
	}
	return on(image) - count(monsters)
}

func count(monsters map[int]map[int]bool) int {
	result := 0
	for _, r := range monsters {
		result += len(r)
	}
	return result
}

func on(ps Pixels) int {
	result := 0
	for _, r := range ps.value {
		for _, v := range r {
			if v {
				result++
			}
		}
	}
	return result
}

func matches(image Pixels, r Pixels, p Point) []Point {
	var result []Point
	for ym := 0; ym < len(r.value); ym++ {
		row := image.value[p.y+ym]
		monsterRow := r.value[ym]
		for xm := 0; xm < len(monsterRow); xm++ {
			if monsterRow[xm] {
				if !row[p.x+xm] {
					return result
				}
			}
		}
	}
	for ym := 0; ym < len(r.value); ym++ {
		row := image.value[p.y+ym]
		monsterRow := r.value[ym]
		for xm := 0; xm < len(monsterRow); xm++ {
			if monsterRow[xm] {
				if row[p.x+xm] {
					result = append(result, Point{p.x + xm, p.y + ym})
				}
			}
		}
	}
	return result
}

func arrange(image Image, ids []int, tiles map[int][]Pixels) (Image, error) {
	if len(ids) == 0 {
		corners := findCorners(image)
		if len(corners) == 4 {
			return image, nil
		}
		return nil, errors.New("too many corners")
	}
	for y, row := range image {
		for x, tile := range row {
			for _, id := range ids {
				for i, version := range tiles[id] {
					matches := tiles[tile.ID][tile.I].align(version)
					p := Point{x, y}
					for _, direction := range intersect(free(image, p), matches) {
						candidate := move(p, direction)
						for d, n := range neighbours(candidate, image) {
							if !aligns(version, tiles[n.ID][n.I], d) {
								panic("not implemented")
							}
						}
						extended := add(image, candidate, id, i)
						image, err := arrange(extended, remove(ids, id), tiles)
						if err == nil {
							return image, nil
						}
					}
				}
			}
		}
	}
	return nil, errors.New("no IDs matched")
}

func findCorners(i Image) []int {
	var result []int
	for y, row := range i {
		for x, v := range row {
			if len(free(i, Point{x, y})) == 2 {
				result = append(result, v.ID)
			}
		}
	}
	return result
}

func aligns(a, b Pixels, direction Direction) bool {
	for _, d := range a.align(b) {
		if d == direction {
			return true
		}
	}
	return false
}

func neighbours(p Point, i Image) map[Direction]Tile {
	result := map[Direction]Tile{}
	if t, ok := i[p.y][p.x-1]; ok {
		result[Left] = t
	}
	if t, ok := i[p.y][p.x+1]; ok {
		result[Right] = t
	}
	if upper, ok := i[p.y-1]; ok {
		if t, ok := upper[p.x]; ok {
			result[Up] = t
		}
	}
	if lower, ok := i[p.y+1]; ok {
		if t, ok := lower[p.x]; ok {
			result[Down] = t
		}
	}
	return result
}

func move(p Point, d Direction) Point {
	switch d {
	case Up:
		return Point{p.x, p.y - 1}
	case Down:
		return Point{p.x, p.y + 1}
	case Left:
		return Point{p.x - 1, p.y}
	case Right:
		return Point{p.x + 1, p.y}
	default:
		panic("not implemented")
	}
}

func add(image Image, p Point, id int, i int) Image {
	cp := copyMap(image)
	if _, ok := cp[p.y]; !ok {
		cp[p.y] = map[int]Tile{}
	}
	cp[p.y][p.x] = Tile{id, i}
	return cp
}

func copyMap(image Image) Image {
	result := make(Image, len(image))
	for y, row := range image {
		result[y] = make(map[int]Tile, len(row))
		for x, tile := range row {
			result[y][x] = tile
		}
	}
	return result
}

func intersect(a, b []Direction) []Direction {
	m := map[Direction]bool{}

	for _, item := range a {
		m[item] = true
	}

	var result []Direction
	for _, item := range b {
		if _, ok := m[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

func remove(xs []int, x int) []int {
	var result []int
	for _, v := range xs {
		if v != x {
			result = append(result, v)
		}
	}
	return result
}

func keys(m map[int][]Pixels) []int {
	var result []int
	for k := range m {
		result = append(result, k)
	}
	return result
}

func parse(file string) map[int][]Pixels {
	text := util.Text(file)
	sections := strings.Split(text, "\n\n")
	tiles := make(map[int][]Pixels, len(sections))
	for _, image := range sections {
		lines := strings.Split(image, "\n")
		header := lines[0]
		id := util.Number(regexp.MustCompile("[0-9]+").FindString(header))
		tiles[id] = alternate(parsePixels(lines[1:]))
	}
	return tiles
}

func alternate(image Pixels) []Pixels {
	result := make([]Pixels, 12)
	result[0] = image
	result[1] = result[0].rotate()
	result[2] = result[1].rotate()
	result[3] = result[2].rotate()
	result[4] = flipHorizontally(result[0])
	result[5] = flipHorizontally(result[1])
	result[6] = flipHorizontally(result[2])
	result[7] = flipHorizontally(result[3])
	result[8] = flipVertically(result[0])
	result[9] = flipVertically(result[1])
	result[10] = flipVertically(result[2])
	result[11] = flipVertically(result[3])
	return result
}

func NewPixels(len int) Pixels {
	return Pixels{make([][]bool, len)}
}

func parsePixels(lines []string) Pixels {
	result := NewPixels(len(lines))
	for i, line := range lines {
		row := make([]bool, len(line))
		for j, c := range line {
			row[j] = c == '#'
		}
		result.value[i] = row
	}
	return result
}

func (ps Pixels) rotate() Pixels {
	result := NewPixels(len(ps.value[0]))
	for i, row := range ps.value {
		for j, v := range row {
			x := len(ps.value) - 1 - i
			y := j
			if result.value[y] == nil {
				result.value[y] = make([]bool, len(ps.value))
			}
			result.value[y][x] = v
		}
	}
	return result
}

func flipHorizontally(ps Pixels) Pixels {
	result := NewPixels(len(ps.value))
	for i, r := range ps.value {
		result.value[i] = make([]bool, len(r))
		for j, v := range r {
			result.value[i][len(r)-1-j] = v
		}
	}
	return result
}

func flipVertically(ps Pixels) Pixels {
	result := NewPixels(len(ps.value))
	for i, r := range ps.value {
		result.value[len(ps.value)-1-i] = make([]bool, len(r))
		for j, v := range r {
			result.value[len(ps.value)-1-i][j] = v
		}
	}
	return result
}

func (ps Pixels) align(to Pixels) []Direction {
	var result []Direction
	if reflect.DeepEqual(left(ps), right(to)) {
		result = append(result, Left)
	}
	if reflect.DeepEqual(right(ps), left(to)) {
		result = append(result, Right)
	}
	if reflect.DeepEqual(ps.value[0], to.value[len(to.value)-1]) {
		result = append(result, Up)
	}
	if reflect.DeepEqual(ps.value[len(to.value)-1], to.value[0]) {
		result = append(result, Down)
	}
	return result
}

func left(ps Pixels) Edge {
	result := make(Edge, len(ps.value))
	for i, r := range ps.value {
		result[i] = r[0]
	}
	return result
}

func right(ps Pixels) Edge {
	result := make(Edge, len(ps.value))
	for i, r := range ps.value {
		result[i] = r[len(r)-1]
	}
	return result
}

func Solve() {
	input := "2020/jigsaw/puzzle.txt"
	monster := "2020/jigsaw/monster.txt"
	fmt.Println(corners(input))
	fmt.Println(monsterFree(input, monster))
}

func join(image Image, tiles map[int][]Pixels) Pixels {
	var height int
	for _, row := range tiles {
		for _, tile := range row {
			height = len(tile.value)
		}
	}

	minY := 0
	maxY := 0
	for y := range image {
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}

	minX := 0
	maxX := 0
	for x := range image[0] {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
	}

	result := Pixels{}
	for y := minY; y <= maxY; y++ {
		for i := 1; i < height-1; i++ {
			var row []bool
			for x := minX; x <= maxX; x++ {
				t := image[y][x]
				for j := 1; j < len(tiles[t.ID][t.I].value)-1; j++ {
					row = append(row, tiles[t.ID][t.I].value[i][j])
				}
			}
			result.value = append(result.value, row)
		}
	}
	return result
}
