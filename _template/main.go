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
	buffer := make([]byte, 10_000_000)
	reader := bufio.NewReaderSize(os.Stdin, len(buffer))
	var sb strings.Builder
	sb.Grow(len(buffer))
	for {
		read, err := reader.Read(buffer)
		sb.Write(buffer[:read])
		if err != nil {
			if err == io.EOF {
				return strings.Split(sb.String(), "\n")
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
