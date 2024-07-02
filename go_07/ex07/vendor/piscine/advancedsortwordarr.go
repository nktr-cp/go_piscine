package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	elem_count := 0
	for range a {
		elem_count++
	}

	for i := 0; i < elem_count-1; i++ {
		for j := i + 1; j < elem_count; j++ {
			if f(a[i], a[j]) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
