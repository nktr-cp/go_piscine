package piscine

func StrRev(s string) string {
	length := StrLen(s)
	var result = make([]rune, length)

	for i, c := range s {
		result[length-1-i] = c
	}

	return string(result)
}
