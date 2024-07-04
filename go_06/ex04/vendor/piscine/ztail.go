package piscine

import "os"

type ZtailAttrs struct {
	byteSize int
	flg      bool
	files    []string
	fileNum  int
	buff     []byte
}

const bufferSize = 30000

func atoi(str string) int64 {
	var nbr int64
	for _, c := range str {
		if c < '0' || c > '9' {
			return -1
		}
		nbr = nbr*10 + int64(c-'0')
		if nbr >= 1<<60 {
			return 1 << 60
		}
	}
	return nbr
}

func fillOutStr(outStr *[]byte, data ZtailAttrs, doneBytes int) {
	if doneBytes < data.byteSize {
		*outStr = append(*outStr, data.buff[0])
	} else {
		for i := 0; i < data.byteSize-1; i++ {
			(*outStr)[i] = (*outStr)[i+1]
		}
		(*outStr)[data.byteSize-1] = data.buff[0]
	}
}

func tailStdin(data ZtailAttrs) {
	outStr := make([]byte, 0, data.byteSize)
	var doneBytes int
	for {
		n, err := os.Stdin.Read(data.buff)
		if n <= 0 || err != nil {
			break
		}
		fillOutStr(&outStr, data, doneBytes)
		doneBytes += n
	}
	if doneBytes > data.byteSize {
		os.Stdout.Write(outStr)
	} else {
		os.Stdout.Write(outStr[:doneBytes])
	}
}

func printErr(filename string) {
	os.Stdout.WriteString("ztail: " + filename + ": No such file or directory\n")
}

func countBytes(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	var fileBytes int = 0

	for {
		bytesRead, err := file.Read(buffer)
		if bytesRead <= 0 || err != nil {
			break
		}
		fileBytes += int(bytesRead)
	}

	return fileBytes
}

func printFile(data ZtailAttrs, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		printErr(filename)
		return
	}
	defer file.Close()

	bytesRest := countBytes(filename)
	bytesDone := 0
	buffer := make([]byte, 1)

	for {
		bytesRead, err := file.Read(buffer)
		if bytesRead <= 0 || err != nil {
			break
		}
		if bytesDone+data.byteSize >= bytesRest {
			os.Stdout.Write(buffer[:bytesRead])
		}
		bytesDone += bytesRead
	}
}

func printConnect(data *ZtailAttrs, idx int) {
	if idx > 0 {
		os.Stdout.WriteString("\n")
	}
	os.Stdout.WriteString("==> " + data.files[idx] + " <==\n")
}

func printFiles(data ZtailAttrs) {
	for i, filename := range data.files {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			printErr(filename)
			continue
		}

		if data.fileNum > 1 {
			printConnect(&data, i)
		}
		printFile(data, filename)
	}
}

func Ztail(args []string) {
	arg_len := 0
	for range args {
		arg_len++
	}

	if arg_len == 0 || args[0] != "-c" || arg_len == 1 {
		os.Exit(1)
	}

	data := ZtailAttrs{
		byteSize: int(atoi(args[1])),
		flg:      false,
		files:    args[2:],
		fileNum:  arg_len - 2,
		buff:     make([]byte, 1),
	}

	if data.byteSize == -1 {
		os.Stdout.WriteString("ztail: illegal offset -- " + args[0] + "\n")
		os.Exit(1)
	}

	if arg_len == 2 {
		tailStdin(data)
	} else {
		printFiles(data)
	}
}
