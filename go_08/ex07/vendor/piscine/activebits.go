package piscine

func ActiveBits(n int) int {
	bitSize := 32 << (^uint(0) >> 63)

	count := 0
	for i := 0; i < bitSize; i++ {
		if n&(1<<uint(i)) > 0 {
			count++
		}
	}

	return count
}
