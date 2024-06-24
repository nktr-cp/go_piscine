package piscine

func CheckBase(base string) bool {
	runes := []rune(base)
	length := len(runes)
	// A base must contain at least 2 characters
	if len(runes) < 2 {
		return false
	}
	for i := 0; i < length-1; i++ {
		// A base should not contain + or - characters
		if runes[i] == '+' || runes[i] == '-' {
			return false
		}
		for j := i + 1; j < length; j++ {
			// Each character of a base must be unique
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !CheckBase(base) {
		return 0
	}

	result := 0
	base_length := len(base)
	for _, c := range s {
		index := 0
		for i, b := range base {
			if c == b {
				index = i
				break
			}
		}

		result = result*base_length + index
	}

	return result
}
