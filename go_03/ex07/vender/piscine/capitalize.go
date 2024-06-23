package piscine

func IsAlphaNumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	} else if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	is_first := true
	runes := []rune(s)
	for i, c := range s {
		if is_first {
			if IsAlphaNumeric(c) {
				is_first = false
			}
			if c >= 'a' && c <= 'z' {
				runes[i] = c - 'a' + 'A'
			}
		} else {
			if c >= 'A' && c <= 'Z' {
				runes[i] = c - 'A' + 'a'
			}
			if !IsAlphaNumeric(c) {
				is_first = true
			}
		}
	}
	return string(runes)
}
