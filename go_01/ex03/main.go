package main

import (
	"ex03/vender/piscine"
	"fmt"
)

func main() {
	a := 13
	b := 2

	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
