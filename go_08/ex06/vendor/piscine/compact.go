package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, str := range *ptr {
		if str == "" {
			continue
		} else {
			count++
		}
	}

	compact := make([]string, count)

	count = 0
	for _, str := range *ptr {
		if str == "" {
			continue
		} else {
			compact[count] = str
			count++
		}
	}

	*ptr = compact
	return count
}
