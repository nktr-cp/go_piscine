package piscine

func SortWordArr(array []string) {
	array_len := 0
	for range array {
		array_len++
	}

	for i := 0; i < array_len-1; i++ {
		for j := i + 1; j < array_len; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
