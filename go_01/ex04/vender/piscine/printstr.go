package piscine

import "ex04/vender/ft"

func PrintStr(s string) {
	// index is not necessary in this loop
	for _, c := range s {
		ft.PrintRune(c)
	}
}
