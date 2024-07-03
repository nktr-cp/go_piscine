package main

import (
	"fmt"
	"piscine"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting;", piscine.Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
