package main

import (
	"ft"
	"piscine"
)

func printStrln(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
	ft.PrintRune('\n')
}

func main() {
	if piscine.ComCheck() {
		printStrln("Alert!!!")
	}
}
