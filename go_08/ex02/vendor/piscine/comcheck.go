package piscine

import "os"

func ComCheck() bool {
	args := os.Args[1:]

	for _, arg := range args {
		if arg == "42" {
			return true
		} else if arg == "piscine" {
			return true
		} else if arg == "piscine 42" {
			return true
		}
	}

	return false
}
