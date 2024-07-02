package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	fmt.Printf("%#v\n", piscine.Split(s, "u?"))
	fmt.Printf("%#v\n", piscine.Split(s, "H"))
	fmt.Printf("%#v\n", piscine.Split(s, "how"))
	fmt.Printf("%#v\n", piscine.Split(s, "?"))
	fmt.Printf("%#v\n", piscine.Split(s, "e"))
	fmt.Printf("%#v\n", piscine.Split("", "hello"))
	fmt.Printf("%#v\n", piscine.Split("", ""))
	fmt.Printf("%#v\n", piscine.Split("llllll", "l"))
}
