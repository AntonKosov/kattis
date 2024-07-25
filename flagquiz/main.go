package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flags := read()
	bestFlags := findBestFlags(flags)
	for _, flag := range bestFlags {
		fmt.Println(flag)
	}
}

func findBestFlags(flags []string) []string {
	attrs := make([][]string, len(flags))
	for i, flag := range flags {
		attrs[i] = strings.Split(flag, ", ")
	}

	changes := make([]int, len(flags))
	for i, attrs1 := range attrs {
		for j := i + 1; j < len(attrs); j++ {
			ch := getChanges(attrs1, attrs[j])
			changes[i] = max(changes[i], ch)
			changes[j] = max(changes[j], ch)
		}
	}

	bestIdx := minIndeces(changes)
	bestFlags := make([]string, len(bestIdx))
	for i, idx := range minIndeces(changes) {
		bestFlags[i] = flags[idx]
	}

	return bestFlags
}

func getChanges(attr1, attr2 []string) int {
	changes := 0
	for i := range attr1 {
		if attr1[i] != attr2[i] {
			changes++
		}
	}

	return changes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minIndeces(changes []int) []int {
	minValue := changes[0]
	res := []int{}
	for i, p := range changes {
		if minValue > p {
			minValue = p
			res = res[:0]
		}

		if minValue == p {
			res = append(res, i)
		}
	}

	return res
}

func read() []string {
	reader := bufio.NewReaderSize(os.Stdin, 1_000_000)
	readLine(reader)
	count := readInt(reader)
	flags := make([]string, count)
	for i := 0; i < count; i++ {
		flags[i] = readLine(reader)
	}

	return flags
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
