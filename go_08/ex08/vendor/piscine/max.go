package piscine

func Max(arr []int) int {
	arrLen := 0
	for range arr {
		arrLen++
	}

	if arrLen == 0 {
		return 0
	}

	result := arr[0]
	for _, v := range arr {
		if v > result {
			result = v
		}
	}

	return result
}
