package main

import (
	"os"
	"piscine"
)

func main() {
	os.Exit(piscine.Ztail(os.Args[1:]))
}
