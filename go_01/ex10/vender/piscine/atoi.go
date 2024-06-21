package piscine

func Atoi(s string) int {
	var sign int = 1
	var result int = 0

	for index, c := range s {
		if '0' <= c && c <= '9' {
			result = result*10 + int(c-'0')
		} else if index == 0 && c == '+' {
			continue
		} else if index == 0 && c == '-' {
			sign *= -1
		} else {
			return 0
		}
	}

	return sign * result
}
