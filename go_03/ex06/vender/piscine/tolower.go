package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			runes[i] = c - 'A' + 'a'
		}
	}

	return string(runes)
}
