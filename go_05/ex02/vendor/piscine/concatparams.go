package piscine

func ConcatParams(args []string) string {
	arg_len := 0
	for range args {
		arg_len++
	}
	result := ""
	for i := 0; i < arg_len; i++ {
		result += args[i]
		if i != arg_len-1 {
			result += "\n"
		}
	}
	return result
}
