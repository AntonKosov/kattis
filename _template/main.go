package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	read()
}

func read() int {
	reader := bufio.NewReaderSize(os.Stdin, 1_000_000)
	readInt(reader)
	readLine(reader)
	return 0
}

func readInt(reader *bufio.Reader) int {
	line := readLine(reader)
	value, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return value
}

func readLine(reader *bufio.Reader) string {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	return string(line[:len(line)-1])
}
