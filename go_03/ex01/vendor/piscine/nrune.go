package piscine

func NRune(s string, n int) rune {
	length := 0
	for range s {
		length++
	}
	runes := []rune(s)
	n--

	if n < 0 || length <= n {
		return 0
	} else {
		return runes[n]
	}
}
