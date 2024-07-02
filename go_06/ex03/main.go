package main

import (
	"os"
	"piscine"
)

func main() {
	piscine.Cat(os.Args[1:])
}
