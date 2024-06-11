package piscine

import "ex02/vender/ft"

func PrintDigits() {
	for c := '0'; c <= '9'; c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
