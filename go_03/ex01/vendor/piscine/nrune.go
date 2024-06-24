package piscine

func NRune(s string, n int) rune {
	n--
	runes := []rune(s)

	if n < 0 || len(runes) <= n {
		return 0
	} else {
		return runes[n]
	}
}
