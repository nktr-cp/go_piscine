package piscine

import "os"

const maxInt int = int(^uint(0) >> 1)
const minInt int = -maxInt - 1

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
	for i < str_len && str[i] >= '0' && str[i] <= '9' {
		digit := int(str[i] - '0')
		if !neg && (res > (maxInt-digit)/10) {
			os.Exit(1)
		}
		if neg && (-res < (minInt+digit)/10) {
			os.Exit(1)
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
	str_len := 0
	for range str {
		str_len++
	}
	if str_len == 0 || ((str[0] == '+' || str[0] == '-') && (str_len == 1)) {
		return false
	}
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
		os.Exit(1)
	}

	n1 := atoi(args[0])
	op := args[1]
	n2 := atoi(args[2])

	switch op {
	case "+":
		if detectAddOverflow(n1, n2) {
			os.Exit(1)
		}
		return n1 + n2
	case "-":
		if detectAddOverflow(n1, -n2) {
			os.Exit(1)
		}
		return n1 - n2
	case "*":
		if detectMulOverflow(n1, n2) {
			os.Exit(1)
		}
		return n1 * n2
	case "/":
		if n2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			os.Exit(1)
		}
		if n1 == minInt && n2 == -1 {
			os.Exit(1)
		}
		return n1 / n2
	case "%":
		if n2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			os.Exit(1)
		}
		return n1 % n2
	default:
		return -1 // Should never reach this
	}
}
