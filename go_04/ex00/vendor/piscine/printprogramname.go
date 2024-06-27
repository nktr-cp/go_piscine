package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	fullPath := os.Args[0]
	start := 0
	length := 0
	for range fullPath {
		length++
	}
	for i := length - 1; i >= 0; i-- {
		if fullPath[i] == '/' || fullPath[i] == '\\' {
			start = i + 1
			break
		}
	}
	programName := fullPath[start:]
	for _, r := range programName {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
