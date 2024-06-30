package main

import (
	"ft"
	"os"
)

type boolean int

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	} else {
		return 0
	}
}

const yes, no boolean = 1, 0
const EvenMsg = "I have an even number of arguments"
const OddMsg = "I have an odd number of arguments"

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	lengthOfArg := 0
	for range os.Args {
		lengthOfArg++
	}
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
