package piscine

import "os"

func DisplayFile(file *os.File) {
	buf := []byte{0}

	for {
		_, err := file.Read(buf)
		if err != nil {
			return
		}
		os.Stdout.Write(buf)
	}
}
