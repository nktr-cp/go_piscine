package piscine

func FindNextPrime(nb int) int {
	result := nb
	for !IsPrime(result) {
		result++
	}

	return result
}
