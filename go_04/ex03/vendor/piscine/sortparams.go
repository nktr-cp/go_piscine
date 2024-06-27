package piscine

import (
	"ft"
	"os"
)

func SortParams() {
	args := os.Args[1:]
	args_len := 0
	for range args {
		args_len++
	}

	for i := 0; i < args_len; i++ {
		for j := i + 1; j < args_len; j++ {
			if Compare(args[i], args[j]) == 1 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for _, arg := range args {
		for _, char := range arg {
			ft.PrintRune(char)
		}
		ft.PrintRune('\n')
	}
}
