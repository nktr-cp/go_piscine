package main

import (
	"os"
	"piscine"
)

func main() {
	piscine.Ztail(os.Args[1:])
}
