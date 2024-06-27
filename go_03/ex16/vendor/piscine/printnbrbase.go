package piscine

import (
	"ft"
)

func CheckBase(base string) bool {
	length := 0
	for range base {
		length++
	}
	runes := []rune(base)
	// A base must contain at least 2 characters
	if length < 2 {
		return false
	}
	for i := 0; i < length; i++ {
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

func PrintNbrBaseRecur(nbr uint64, base string) {
	length := uint64(0)
	for range base {
		length++
	}
	runes := []rune(base)
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

	unbr := uint64(nbr)
	if nbr < 0 {
		ft.PrintRune('-')
		unbr = -unbr
	}

	PrintNbrBaseRecur(unbr, base)
}
