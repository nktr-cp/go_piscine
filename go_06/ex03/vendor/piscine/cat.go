package piscine

import (
	"ft"
	"os"
)

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func catStdin() {
	buf := []byte{0}
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			return
		}
		os.Stdout.Write(buf)
	}
}

func catFile(filename string, status *int) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		*status = 1
		if err.Error() != "EOF" {
			printStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		return
	} else {
		buf := []byte{0}

		for {
			_, err := file.Read(buf)
			if err != nil {
				if err.Error() != "EOF" {
					*status = 1
					printStr("ERROR: " + err.Error() + "\n")
				}
				return
			}
			os.Stdout.Write(buf)
		}
	}
}

func Cat(args []string) int {
	arg_len := 0
	for range args {
		arg_len++
	}
	status := 0
	if arg_len == 0 {
		catStdin()
	} else {
		for i := 0; i < arg_len; i++ {
			catFile(args[i], &status)
		}
	}
	return status
}
