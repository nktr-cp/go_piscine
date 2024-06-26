package piscine

func Join(strs []string, sep string) string {
	length := 0
	for range strs {
		length++
	}
	result := ""
	for i, s := range strs {
		result += s
		if i != length-1 {
			result += sep
		}
	}
	return result
}
