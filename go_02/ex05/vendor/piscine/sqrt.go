package piscine

func Sqrt(n int) int {
	if n < 0 {
		return 0
	}
	un := uint64(n)
	var pos uint64
	for ; pos*pos < un; pos++ {
	}
	if pos*pos == un {
		return int(pos)
	} else {
		return 0
	}
}
