package piscine

// Errors (non possible values or overflows) will return 0.
func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	intSize := 32 << (^uint(0) >> 63)
	if intSize == 32 && nb > 12 {
		return 0
	} else if intSize == 64 && nb > 20 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}
