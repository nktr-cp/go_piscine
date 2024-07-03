package piscine

func Unmatch(a []int) int {
	argLen := 0
	for range a {
		argLen++
	}

	for i := 0; i < argLen-1; i++ {
		for j := i + 1; j < argLen; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	for i := 0; i < argLen; i++ {
		if i < argLen-1 && a[i] == a[i+1] {
			i++
		} else {
			return a[i]
		}
	}
	return -1
}
