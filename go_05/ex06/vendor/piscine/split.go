package piscine

func Split(s, sep string) []string {
	var result []string

	if s == "" {
		return []string{""}
	}
	if sep == "" {
		return []string{s}
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
		if i+sep_length <= source_length && s[i:i+sep_length] == sep {
			if s[start:i] != "" {
				result = append(result, s[start:i])
			}
			start = i + sep_length
		}
		if i == source_length-1 {
			if s[start:] != "" {
				result = append(result, s[start:])
			}
			break
		}
	}
	if result == nil {
		result = append(result, "")
	}
	return result
}
