package main

import (
	"ft"
	"os"
	"piscine"
)

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func main() {
	length := 0
	for range os.Args {
		length++
	}

	if length == 1 {
		printStr("File name missing\n")
	} else if length > 2 {
		printStr("Too many arguments\n")
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			printStr("File does not exist\n")
		} else {
			piscine.DisplayFile(file)
			file.Close()
		}
	}
}
