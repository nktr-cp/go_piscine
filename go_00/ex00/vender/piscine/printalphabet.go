package piscine

import "ex00/vender/ft"

func PrintAlphabet() {
	for c := 'a'; c <= 'z'; c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
