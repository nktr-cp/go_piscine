package piscine

import "ft"

func IsValidPosition(pattern [8]int, current int, check int) bool {
	for i := 0; i < current; i++ {
		if pattern[i] == check {
			return false
		} else if pattern[i]+current-i == check {
			return false
		} else if pattern[i]-current+i == check {
			return false
		}
	}
	return true
}

func EightQueensRecursion(pattern [8]int, current int) {
	if current == 8 {
		for i := 0; i < 8; i++ {
			ft.PrintRune(rune(pattern[i] + 1 + '0'))
		}
		ft.PrintRune('\n')
		return
	}

	for i := 0; i < 8; i++ {
		if IsValidPosition(pattern, current, i) {
			pattern[current] = i
			EightQueensRecursion(pattern, current+1)
		}
	}
}

func EightQueens() {
	var pattern [8]int

	EightQueensRecursion(pattern, 0)
}
