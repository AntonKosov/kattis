package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pile1 := read()
	pile2 := make([]int, 0, len(pile1))
	steps := 0
	for len(pile1) > 0 {
		steps++
		if len(pile2) > 0 && pile1[len(pile1)-1] == pile2[len(pile2)-1] {
			pile1 = pile1[:len(pile1)-1]
			pile2 = pile2[:len(pile2)-1]
			continue
		}

		pile2 = append(pile2, pile1[len(pile1)-1])
		pile1 = pile1[:len(pile1)-1]
	}

	if len(pile2) == 0 {
		fmt.Println(steps)
	} else {
		fmt.Println("impossible")
	}
}

func read() []int {
	reader := bufio.NewReaderSize(os.Stdin, 1_000_000)
	readLine(reader)
	line := readLine(reader)
	parts := strings.Split(line, " ")
	socks := make([]int, len(parts))
	for i, p := range parts {
		id, err := strconv.Atoi(p)
		if err != nil {
			panic(err.Error())
		}
		socks[i] = id
	}

	return socks
}

func readLine(reader *bufio.Reader) string {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	return string(line[:len(line)-1])
}
