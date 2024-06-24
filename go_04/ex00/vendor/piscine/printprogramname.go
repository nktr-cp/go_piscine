package piscine

import (
	"ft"
	"os"
)

func Split(str string, sep rune) []string {
	var parts []string
	start := 0
	for i, char := range str {
		if char == sep {
			parts = append(parts, str[start:i])
			start = i + 1
		}
	}
	parts = append(parts, str[start:])
	return parts
}

func PrintProgramName() {
	fullPath := os.Args[0]
	parts := Split(fullPath, '/')
	programName := parts[len(parts)-1]

	for _, char := range programName {
		ft.PrintRune(char)
	}
	ft.PrintRune('\n')
}
