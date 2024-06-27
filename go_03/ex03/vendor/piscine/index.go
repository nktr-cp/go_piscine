package piscine

func Index(s string, toFind string) int {
	target_length := 0
	for range toFind {
		target_length++
	}
	source_length := 0
	for range s {
		source_length++
	}
	for i := 0; i+target_length <= source_length; i++ {
		if s[i:i+target_length] == toFind {
			return i
		}
	}

	return -1
}
