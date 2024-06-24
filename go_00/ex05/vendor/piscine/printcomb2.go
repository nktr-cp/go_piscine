package piscine

import "ft"

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		var b int = a + 1
		for ; b <= 99; b++ {
			c := rune(a/10 + int('0'))
			d := rune(a%10 + int('0'))

			ft.PrintRune(c)
			ft.PrintRune(d)

			ft.PrintRune(' ')

			e := rune(b/10 + int('0'))
			f := rune(b%10 + int('0'))

			ft.PrintRune(e)
			ft.PrintRune(f)

			if !(a == 98 && b == 99) {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}
