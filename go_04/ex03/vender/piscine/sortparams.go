package piscine

import (
	"ex03/vender/ft"
	"os"
)

func SortParams() {
	args := os.Args
	var params []string

	for i, arg := range args {
		if i > 0 {
			params = append(params, arg)
		}
	}

	for i := 0; i < len(params)-1; i++ {
		for j := i + 1; j < len(params); j++ {
			if Compare(params[i], params[j]) == 1 {
				params[i], params[j] = params[j], params[i]
			}
		}
	}

	for _, arg := range params {
		for _, char := range arg {
			ft.PrintRune(char)
		}
		ft.PrintRune('\n')
	}
}
