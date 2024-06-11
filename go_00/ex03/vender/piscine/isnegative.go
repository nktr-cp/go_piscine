package piscine

import "ex03/vender/ft"

func IsNegative(nb int) {
	if nb < 0 {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
