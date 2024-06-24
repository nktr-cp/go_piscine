package piscine

import (
	"ft"
)

func CheckBase(base string) bool {
	runes := []rune(base)
	length := len(runes)
	// A base must contain at least 2 characters
	if len(runes) < 2 {
		return false
	}
	for i := 0; i < length-1; i++ {
		// A base should not contain + or - characters
		if runes[i] == '+' || runes[i] == '-' {
			return false
		}
		for j := i + 1; j < length; j++ {
			// Each character of a base must be unique
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}

func PrintNbrBaseRecur(nbr int, base string) {
	runes := []rune(base)
	length := len(runes)
	if nbr < 0 {
		ft.PrintRune('-')
		nbr = -nbr
	}
	if nbr >= length {
		PrintNbrBaseRecur(nbr/length, base)
	}
	ft.PrintRune(runes[nbr%length])
}

func PrintNbrBase(nbr int, base string) {
	if !CheckBase(base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}

	PrintNbrBaseRecur(nbr, base)
}
