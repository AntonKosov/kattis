package main

import (
	"bufio"
	"os"
)

func main() {
	n, m := read()

	idx := len(n) - m + 1
	if idx < 0 {
		leadingZeros := -idx
		newN := make([]byte, leadingZeros+len(n))
		for i := 0; i < leadingZeros; i++ {
			newN[i] = '0'
		}
		copy(newN[leadingZeros:], n)
		n = newN
		idx = 0
	}

	whole, decimal := n[:idx], n[idx:]
	if len(whole) == 0 {
		whole = []byte{'0'}
	}

	for len(decimal) > 0 && decimal[len(decimal)-1] == '0' {
		decimal = decimal[:len(decimal)-1]
	}

	output(whole, decimal)
}

func read() ([]byte, int) {
	reader := bufio.NewReaderSize(os.Stdin, 2_000_000)
	n, _ := reader.ReadBytes('\n')
	m, _ := reader.ReadBytes('\n')

	return n[:len(n)-1], len(m) - 1
}

func output(whole, decimal []byte) {
	os.Stdout.Write(whole)
	if len(decimal) > 0 {
		os.Stdout.WriteString(".")
		os.Stdout.Write(decimal)
	}
	os.Stdout.WriteString("\n")
}
