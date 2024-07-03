package piscine

func Rot14(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + (r-'A'+14)%26
		} else if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + (r-'a'+14)%26
		}
	}
	return string(runes)
}
