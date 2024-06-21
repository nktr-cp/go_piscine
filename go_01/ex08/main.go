package main

import (
	"ex08/vender/piscine"
	"fmt"
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
