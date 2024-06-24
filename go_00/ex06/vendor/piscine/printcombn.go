package piscine

import "ft"

func CombnRecurse(a [10]rune, index int, n int, start int) {
	if index == n {
		for i := 0; i < n; i++ {
			ft.PrintRune(a[i])
		}

		if (int)(a[0]-'0') != 10-n {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
	} else {
		for i := start; i < 10; i++ {
			a[index] = (rune)(i) + '0'
			CombnRecurse(a, index+1, n, i+1)
		}
	}
}

func PrintCombn(n int) {
	var a [10]rune

	CombnRecurse(a, 0, n, 0)
	ft.PrintRune('\n')
}
