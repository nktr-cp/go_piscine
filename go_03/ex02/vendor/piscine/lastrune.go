package piscine

func LastRune(s string) rune {
	length := 0
	for range s {
		length++
	}
	runes := []rune(s)
	return runes[length-1]
}
