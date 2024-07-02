package piscine

import "ft"

func PrintNbrRecur(nbr uint64) {
	if nbr >= 10 {
		PrintNbrRecur(nbr / 10)
	}
	ft.PrintRune(rune('0' + nbr%10))
}

func PrintNbr(nbr int) {
	unbr := uint64(nbr)
	if nbr < 0 {
		ft.PrintRune('-')
		unbr = -unbr
	}
	PrintNbrRecur(unbr)
}
