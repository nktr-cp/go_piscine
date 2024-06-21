package main

import (
	"ex00/vender/ft"
	"ex00/vender/piscine"
)

func main() {
	ft.PrintRune(piscine.FirstRune("Hello!"))
	ft.PrintRune(piscine.FirstRune("Salut!"))
	ft.PrintRune(piscine.FirstRune("Ola!"))
	ft.PrintRune('\n')
}
