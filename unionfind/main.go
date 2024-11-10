package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	counts := strings.Split(input[0], " ")
	vertices, operations := strToInt(counts[0]), strToInt(counts[1])
	union := make([]int, vertices)
	for i := range union {
		union[i] = i
	}

	writer := bufio.NewWriterSize(os.Stdout, 10_000_000)
	defer writer.Flush()

	for i := 1; i <= operations; i++ {
		vals := strings.Split(input[i], " ")
		r1, r2 := root(union, strToInt(vals[1])), root(union, strToInt(vals[2]))
		if vals[0] == "=" {
			union[r1] = r2
			continue
		}

		if r1 == r2 {
			writer.WriteString("yes\n")
		} else {
			writer.WriteString("no\n")
		}
	}
}

func root(union []int, node int) int {
	if union[node] == node {
		return node
	}

	r := root(union, union[node])
	union[node] = r

	return r
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

func strToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return value
}
