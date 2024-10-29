package piscine

import (
	"github.com/01-edu/z01"
)

const N = 8

var position = [N]int{}

func getInte(r int) rune {
	tab := []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	return tab[r]
}

func isSafe(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]
		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func solve(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			z01.PrintRune(getInte(position[i] + 1))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

func EightQueens() {
	solve(0)
}
