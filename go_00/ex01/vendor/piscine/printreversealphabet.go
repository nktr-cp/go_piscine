package piscine

import "ft"

func PrintReverseAlphabet() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
