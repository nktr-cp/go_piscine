package piscine

import "ex01/vender/ft"

func PrintReverseAlphabet() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
