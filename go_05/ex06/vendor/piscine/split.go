package piscine

func Split(s, sep string) []string {
	var result []string

	if sep == "" {
		result = append(result, s)
		return result
	}

	source_length := 0
	for range s {
		source_length++
	}
	sep_length := 0
	for range sep {
		sep_length++
	}
	start := 0
	for i := 0; i < source_length; i++ {
		if i == source_length-1 {
			result = append(result, s[start:])
			break
		}
		if s[i:i+sep_length] == sep {
			result = append(result, s[start:i])
			start = i + sep_length
		}
	}
	return result
}
