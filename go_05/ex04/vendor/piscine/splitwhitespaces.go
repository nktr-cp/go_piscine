package piscine

func SplitWhiteSpaces(str string) []string {
	str_len := 0
	for range str {
		str_len++
	}
	result := []string{}
	start_pos := 0
	for start_pos < str_len {
		end_pos := start_pos
		for end_pos < str_len && str[end_pos] != ' ' && str[end_pos] != '\t' && str[end_pos] != '\n' {
			end_pos++
		}
		result = append(result, str[start_pos:end_pos])
		start_pos = end_pos + 1
		for start_pos < str_len && (str[start_pos] == ' ' || str[start_pos] == '\t' || str[start_pos] == '\n') {
			start_pos++
		}
	}

	return result
}
