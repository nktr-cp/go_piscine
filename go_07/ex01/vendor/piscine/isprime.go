package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	unb := uint64(nb)
	for i := uint64(2); i*i <= unb; i++ {
		if unb%i == 0 {
			return false
		}
	}

	return true
}
