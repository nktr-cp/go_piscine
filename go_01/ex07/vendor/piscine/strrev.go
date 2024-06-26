package piscine

func StrRev(s string) string {
	result := ""
	length := StrLen(s)

	for i := length - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}
