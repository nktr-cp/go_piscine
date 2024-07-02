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
			return
		}
		os.Stdout.Write(buf)
	}
}

func catFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		printStr("File does not exist\n")
	} else {
		buf := []byte{0}

		for {
			_, err := file.Read(buf)
			if err != nil {
				return
			}
			os.Stdout.Write(buf)
		}
	}
}

func Cat(args []string) {
	arg_len := 0
	for range args {
		arg_len++
	}
	if arg_len == 0 {
		catStdin()
	} else {
		for i := 0; i < arg_len; i++ {
			catFile(args[i])
		}
	}
}
