package piscine

import "os"

// import "fmt" // DEBUG

func atoi(str string) int {
	str_len := 0
	for range str {
		str_len++
	}
	res := 0
	neg := false
	i := 0

	// Handle sign
	if i < str_len && (str[i] == '+' || str[i] == '-') {
		if str[i] == '-' {
			neg = true
		}
		i++
	}

	// Convert digits
	intSize := 32 << (^uint(0) >> 63)
	var maxInt, minInt int
	if intSize == 32 {
		maxInt = 1<<31 - 1
		minInt = -1 << 31
	} else {
		maxInt = 1<<63 - 1
		minInt = -1 << 63
	}
	for i < str_len && str[i] >= '0' && str[i] <= '9' {
		digit := int(str[i] - '0')
		if !neg && (res > (maxInt-digit)/10) {
			return maxInt
		}
		if neg && (-res < (minInt+digit)/10) {
			return minInt
		}
		res = res*10 + digit
		i++
	}

	if neg {
		res *= -1
	}

	return res
}

func validNumber(str string) bool {
	for i, r := range str {
		if i == 0 && (r == '+' || r == '-') {
			continue
		} else if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func validOperator(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "/" || str == "%" {
		return true
	}
	return false
}

func checkArgs(args []string) bool {
	argsLen := 0
	for range args {
		argsLen++
	}
	if argsLen != 3 {
		return false
	}

	return validNumber(args[0]) && validOperator(args[1]) && validNumber(args[2])
}

func detectAddOverflow(n1, n2 int) bool {
	intSize := 32 << (^uint(0) >> 63)
	var maxInt, minInt int
	if intSize == 32 {
		maxInt = 1<<31 - 1
		minInt = -1 << 31
	} else {
		maxInt = 1<<63 - 1
		minInt = -1 << 63
	}

	if n1 > 0 && n2 > 0 {
		if n1 > maxInt-n2 {
			return true
		}
	} else if n1 < 0 && n2 < 0 {
		if n1 < minInt-n2 {
			return true
		}
	}
	return false
}

func detectMulOverflow(n1, n2 int) bool {
	intSize := 32 << (^uint(0) >> 63)
	var maxInt, minInt int
	if intSize == 32 {
		maxInt = 1<<31 - 1
		minInt = -1 << 31
	} else {
		maxInt = 1<<63 - 1
		minInt = -1 << 63
	}

	if n1 > 0 && n2 > 0 {
		if n1 > maxInt/n2 {
			return true
		}
	} else if n1 < 0 && n2 < 0 {
		if n1 < maxInt/n2 {
			return true
		}
	} else if n1 > 0 && n2 < 0 {
		if n2 < minInt/n1 {
			return true
		}
	} else if n1 < 0 && n2 > 0 {
		if n1 < minInt/n2 {
			return true
		}
	}
	return false
}

func Doop(args []string) int {
	if !checkArgs(args) {
		os.Exit(0)
	}

	n1 := atoi(args[0])
	op := args[1]
	n2 := atoi(args[2])

	if op == "+" {
		if detectAddOverflow(n1, n2) {
			os.Exit(0)
		}
		return n1 + n2
	} else if op == "-" {
		if detectAddOverflow(n1, -n2) {
			os.Exit(0)
		}
		return n1 - n2
	} else if op == "*" {
		if detectMulOverflow(n1, n2) {
			os.Exit(0)
		}
		return n1 * n2
	} else if op == "/" {
		if n2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			os.Exit(0)
		}
		return n1 / n2
	} else if op == "%" {
		if n2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			os.Exit(0)
		}
		return n1 % n2
	}
	return -1 // Should never reach this
}
