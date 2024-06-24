package piscine

import (
	"ex02/vender/ft"
	"os"
)

func RevParams() {
	args := os.Args
	for i := len(args) - 1; i >= 1; i-- {
		for _, c := range args[i] {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
