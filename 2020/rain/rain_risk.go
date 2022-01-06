package rain

import (
	"fmt"
	"github.com/advendofcode/util"
	"math"
	"strconv"
)

type Waypoint struct {
	n int
	e int
}

type Ship struct {
	n int
	e int
	w Waypoint
}

func solve(commands []string) int {
	s := Ship{0, 0, Waypoint{1, 10}}
	for _, c := range commands {
		s = process(s, c)
	}
	return distance(s)
}

func distance(s Ship) int {
	return int(math.Abs(float64(s.e)) + math.Abs(float64(s.n)))
}

func process(s Ship, command string) Ship {
	action := command[0]
	value, err := strconv.Atoi(command[1:])
	if err != nil {
		panic(err)
	}
	switch action {
	case 'F':
		return forward(s, value)
	case 'N':
		return Ship{s.n, s.e, Waypoint{s.w.n + value, s.w.e}}
	case 'W':
		return Ship{s.n, s.e, Waypoint{s.w.n, s.w.e - value}}
	case 'E':
		return Ship{s.n, s.e, Waypoint{s.w.n, s.w.e + value}}
	case 'S':
		return Ship{s.n, s.e, Waypoint{s.w.n - value, s.w.e}}
	case 'R':
		return Ship{s.n, s.e, turn(s.w, action, value)}
	case 'L':
		return Ship{s.n, s.e, turn(s.w, action, value)}
	default:
		panic("not implemented for " + command)
	}
}

func turn(w Waypoint, action uint8, angle int) Waypoint {
	if angle%90 != 0 {
		panic("not implemented for angle" + strconv.Itoa(angle))
	}
	switch angle {
	case 90:
		if action == 'L' {
			return Waypoint{w.e, -w.n}
		} else {
			return Waypoint{-w.e, w.n}
		}
	case 180:
		return turn(turn(w, action, 90), action, 90)
	case 270:
		return turn(turn(w, action, 180), action, 90)
	}
	panic("not implemented for " + strconv.Itoa(angle))
}

func forward(s Ship, distance int) Ship {
	return Ship{
		n: s.n + distance*s.w.n,
		e: s.e + distance*s.w.e,
		w: s.w,
	}
}

func Solve() {
	lines := util.Lines("2020/rain/puzzle.txt")
	fmt.Println(solve(lines))
}
