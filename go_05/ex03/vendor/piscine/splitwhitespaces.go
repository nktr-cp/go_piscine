package piscine

func SplitWhiteSpaces(str string) []string {
	result := []string{}
	start_pos := 0
	for start_pos < len(str) {
		end_pos := start_pos
		for end_pos < len(str) && str[end_pos] != ' ' && str[end_pos] != '\t' && str[end_pos] != '\n' {
			end_pos++
		}
		result = append(result, str[start_pos:end_pos])
		start_pos = end_pos + 1
	}

	return result
}
