package main

import (
	"ex11/vender/piscine"
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
