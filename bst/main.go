package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := read()

	initialSection := Section{Level: 0, First: 0, Last: len(nums) - 1}
	sections := make([]*Section, len(nums))
	for i := range sections {
		sections[i] = &initialSection
	}

	writer := bufio.NewWriterSize(os.Stdout, 1_000_000)
	defer writer.Flush()

	for counter := 0; len(nums) > 0; {
		num := nums[0]
		nums = nums[1:]
		section := sections[num]
		sections[num] = nil
		counter += section.Level
		writer.WriteString(fmt.Sprintf("%v\n", counter))
		section.Level++
		smallerSection := Section{Level: section.Level}
		if num-section.First > section.Last-num {
			smallerSection.First = num + 1
			smallerSection.Last = section.Last
			section.Last = num - 1
		} else {
			smallerSection.First = section.First
			smallerSection.Last = num - 1
			section.First = num + 1
		}
		for i := smallerSection.First; i <= smallerSection.Last; i++ {
			sections[i] = &smallerSection
		}
	}
}

func read() []int {
	lines := readInput()
	count := strToInt(lines[0])
	nums := make([]int, count)
	for i := 1; i < len(lines); i++ {
		nums[i] = strToInt(lines[i]) - 1
	}

	return nums
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

func strToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return value
}

type Section struct {
	Level int
	First int
	Last  int
}
