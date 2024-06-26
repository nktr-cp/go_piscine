package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			runes[i] = c - 'a' + 'A'
		}
	}

	return string(runes)
}
