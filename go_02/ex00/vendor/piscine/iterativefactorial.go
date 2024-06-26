package piscine

// Errors (non possible values or overflows) will return 0.
func IterativeFactorial(n int) int {
	if n < 0 || n > 12 {
		return 0
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
