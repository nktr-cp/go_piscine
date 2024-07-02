package piscine

func AtoiBase(s string, base string) int {
	result := 0
	base_length := len(base)
	for _, c := range s {
		index := 0
		for i, b := range base {
			if c == b {
				index = i
				break
			}
		}

		result = result*base_length + index
	}

	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	nbr_i := AtoiBase(nbr, baseFrom)
	base_len := 0
	for range baseTo {
		base_len++
	}
	result := ""
	for nbr_i > 0 {
		result = string(baseTo[nbr_i%base_len]) + result
		nbr_i /= base_len
	}
	return result
}
