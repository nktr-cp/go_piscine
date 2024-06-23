package piscine

func IsAlpha(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			continue
		} else if c >= 'A' && c <= 'Z' {
			continue
		} else if c >= '0' && c <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
