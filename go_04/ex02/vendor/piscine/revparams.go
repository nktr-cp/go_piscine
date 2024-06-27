package piscine

import (
	"ft"
	"os"
)

func RevParams() {
	args := os.Args
	args_len := 0
	for range args {
		args_len++
	}
	for i := args_len - 1; i >= 1; i-- {
		for _, c := range args[i] {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
