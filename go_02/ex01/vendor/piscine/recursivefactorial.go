package piscine

func RecursiveFactorial(n int) int {
	if n < 0 || n > 12 {
		return 0
	}
	if n == 0 {
		return 1
	} else {
		return n * RecursiveFactorial((n - 1))
	}
}
