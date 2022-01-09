package conway

import (
	"github.com/advendofcode/util"
)

type Line = map[int]bool
type Square = map[int]Line
type Cube = map[int]Square
type Tesseract = map[int]Cube

func lit(t Tesseract) int {
	result := 0
	do(t, func(p Point, v bool) {
		if v {
			result++
		}
	})
	return result
}

func parse(file string) Tesseract {
	lines := util.Lines(file)
	result := make(Tesseract, 1)
	result[0] = make(Cube, 1)
	result[0][0] = make(Square, len(lines))
	for y, line := range lines {
		result[0][0][y] = make(Line, len(line))
		for x, v := range line {
			if v == '#' {
				result[0][0][y][x] = true
			} else {
				result[0][0][y][x] = false
			}
		}
	}
	return result
}

type Point struct {
	x int
	y int
	z int
	w int
}

func neighbours(p Point) []Point {
	ns := []int{-1, 0, 1}
	var result []Point
	for _, x := range ns {
		for _, y := range ns {
			for _, z := range ns {
				for _, w := range ns {
					if x != 0 || y != 0 || z != 0 || w != 0 {
						result = append(result, Point{
							p.x + x,
							p.y + y,
							p.z + z,
							p.w + w,
						})
					}
				}
			}
		}
	}
	return result
}

func advance(t Tesseract, steps int) Tesseract {
	if steps == 0 {
		return t
	}
	update := Tesseract{}
	do(extend(t), func(p Point, v bool) {
		set(update, p, nextValue(v, litNeighbours(p, t)))
	})
	return advance(update, steps-1)
}

func litNeighbours(p Point, t Tesseract) int {
	lit := 0
	for _, n := range neighbours(p) {
		if get(t, n) {
			lit++
		}
	}
	return lit
}

func nextValue(value bool, lit int) bool {
	if value && !(lit == 2 || lit == 3) {
		return false
	}
	if !value && lit == 3 {
		return true
	}
	return value
}

func extend(t Tesseract) Tesseract {
	result := Tesseract{}
	do(t, func(p Point, v bool) {
		for _, n := range neighbours(p) {
			set(result, n, get(t, n))
		}
	})
	return result
}

func get(t Tesseract, p Point) bool {
	return t[p.w][p.z][p.y][p.x]
}

func set(t Tesseract, p Point, v bool) {
	if t[p.w] == nil {
		t[p.w] = Cube{}
	}
	if t[p.w][p.z] == nil {
		t[p.w][p.z] = Square{}
	}
	if t[p.w][p.z][p.y] == nil {
		t[p.w][p.z][p.y] = Line{}
	}
	t[p.w][p.z][p.y][p.x] = v
}

func do(t Tesseract, f func(Point, bool)) {
	for w, c := range t {
		for z, s := range c {
			for y, l := range s {
				for x, v := range l {
					f(Point{x, y, z, w}, v)
				}
			}
		}
	}
}
