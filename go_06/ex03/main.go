package main

import (
	"os"
	"piscine"
)

func main() {
	os.Exit(piscine.Cat(os.Args[1:]))
}
