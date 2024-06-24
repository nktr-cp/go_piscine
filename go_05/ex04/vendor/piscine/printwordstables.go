package piscine

import "ft"

func PrintWordsTables(table []string) {
	for _, str := range table {
		for _, c := range str {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
