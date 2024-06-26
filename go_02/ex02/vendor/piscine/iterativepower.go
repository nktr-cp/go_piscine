package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	intSize := 32 << (^uint(0) >> 63)
	if intSize == 32 && nb > 12 {
		return 0
	} else if intSize == 64 && nb > 20 {
		return 0
	}
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}

	return result
}
