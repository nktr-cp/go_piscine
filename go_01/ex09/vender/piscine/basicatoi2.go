package piscine

func BasicAtoi2(s string) int {
	var result int = 0

	for _, c := range s {
		if '0' <= c && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			return 0
		}
	}

	return result
}
