package main

import (
	"ex01/vender/piscine"
	"fmt"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
