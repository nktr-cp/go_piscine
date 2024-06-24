package piscine

import (
	"ex01/vender/ft"
	"os"
)

func PrintParams() {
	args := os.Args
	for i, arg := range args {
		if i != 0 {
			for _, char := range arg {
				ft.PrintRune(char)
			}
			ft.PrintRune('\n')
		}
	}
}
