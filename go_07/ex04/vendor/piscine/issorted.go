package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	elem_count := 0
	for range a {
		elem_count++
	}

	for i := 0; i < elem_count-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}
