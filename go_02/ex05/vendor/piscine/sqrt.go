package piscine

func Sqrt(n int) int {
	var pos int = 1
	for ; pos*pos < n; pos++ {
	}
	if pos*pos == n {
		return pos
	} else {
		return 0
	}
}
