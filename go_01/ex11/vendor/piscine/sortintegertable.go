package piscine

// selection sort
func SortIntegerTable(s []int) {
	length := 0
	for range s {
		length++
	}

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			Swap(&s[i], &s[j])
		}
	}
}
