package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	readInput()
}

func readInput() []string {
	buffer := make([]byte, 1_000_000)
	reader := bufio.NewReaderSize(os.Stdin, len(buffer))
	var sb strings.Builder
	for {
		read, err := reader.Read(buffer)
		sb.Write(buffer[:read])
		if err != nil {
			if err == io.EOF {
				return strings.Split(strings.ReplaceAll(sb.String(), "\r\n", "\n"), "\n")
			}
			panic(err)
		}
	}
}

/*func strToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return value
}*/
