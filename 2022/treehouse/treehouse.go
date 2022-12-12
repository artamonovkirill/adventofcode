package treehouse

import (
	"fmt"
	"github.com/advendofcode/util"
)

func parse(file string) [][]int {
	lines := util.Lines(file)
	result := make([][]int, len(lines))
	for i, line := range lines {
		result[i] = make([]int, len(line))
		for j, cell := range []rune(line) {
			result[i][j] = util.Number(string(cell))
		}
	}
	return result
}

func Visible(file string) int {
	trees := parse(file)
	result := 0
	for row, line := range trees {
		for column := range line {
			if isOutside(row, column, trees) || visible(row, column, trees) {
				result += 1
			}
		}
	}
	return result
}

func isOutside(row int, column int, trees [][]int) bool {
	return row == 0 || column == 0 || row == len(trees)-1 || column == len(trees[row])-1
}

func visible(row int, column int, trees [][]int) bool {
	height := trees[row][column]
	top := true
	for i := 0; i < row; i++ {
		if trees[i][column] >= height {
			top = false
		}
	}
	bottom := true
	for i := row + 1; i < len(trees); i++ {
		if trees[i][column] >= height {
			bottom = false
		}
	}
	left := true
	for j := 0; j < column; j++ {
		if trees[row][j] >= height {
			left = false
		}
	}
	right := true
	for j := column + 1; j < len(trees[row]); j++ {
		if trees[row][j] >= height {
			right = false
		}
	}
	return top || bottom || left || right
}

func BestScore(file string) int {
	trees := parse(file)
	max := 0
	for row, line := range trees {
		for column := range line {
			score := Score(row, column, trees)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func Score(row int, column int, trees [][]int) int {
	height := trees[row][column]
	top := 0
	for i := row - 1; i >= 0; i-- {
		top += 1
		if trees[i][column] >= height {
			break
		}
	}
	bottom := 0
	for i := row + 1; i < len(trees); i++ {
		bottom += 1
		if trees[i][column] >= height {
			break
		}
	}
	left := 0
	for j := column - 1; j >= 0; j-- {
		left += 1
		if trees[row][j] >= height {
			break
		}
	}
	right := 0
	for j := column + 1; j < len(trees[row]); j++ {
		right += 1
		if trees[row][j] >= height {
			break
		}
	}
	return top * bottom * left * right
}

func Solve() {
	file := "2022/treehouse/input.txt"
	fmt.Println(Visible(file))
	fmt.Println(BestScore(file))
}
