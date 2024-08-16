package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list := read()
	for _, ms := range list {
		fmt.Println(getHash(ms))
	}
}

func getHash(ms MagicSequence) uint64 {
	buckets := buildSeqSection(ms)
	keys := make([]uint64, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	hash := uint64(0)
	for len(keys) > 0 {
		key := keys[0]
		keys = keys[1:]
		count := buckets[key]
		values := make([]uint64, 0, count)
		valuesIdx := make(map[uint64]int, count)
		for i := 0; i < count; i++ {
			hash = (hash*ms.x + key) % ms.y
			if idx, ok := valuesIdx[hash]; ok {
				seqLen := len(values) - idx
				hash = values[idx+(count-idx-1)%seqLen]
				break
			}
			values = append(values, hash)
			valuesIdx[hash] = len(values) - 1
		}
	}

	return hash
}

func buildSeqSection(ms MagicSequence) map[uint64]int {
	seq := make([]uint64, 0, ms.n)
	seq = append(seq, ms.a)
	res := make(map[uint64]int, ms.n)
	res[ms.a] = 1
	encountered := map[uint64]int{ms.a: 0}
	for i := 1; i < ms.n; i++ {
		v := (seq[i-1]*ms.b + ms.a) % ms.c
		if startIdx, ok := encountered[v]; ok {
			seqLen := i - startIdx
			for j := 0; j < seqLen && i+j < ms.n; j++ {
				v := seq[startIdx+j]
				rest := ms.n - i - j
				res[v] += rest / seqLen
				if rest%seqLen != 0 {
					res[v]++
				}
			}
			break
		}
		encountered[v] = i
		seq = append(seq, v)
		res[v] = 1
	}

	return res
}

func read() []MagicSequence {
	reader := bufio.NewReaderSize(os.Stdin, 1_000_000)
	count := readInt(reader)
	list := make([]MagicSequence, count)
	for i := 0; i < count; i++ {
		var ms MagicSequence
		ms.n = readInt(reader)
		line := readLine(reader)
		abc := strings.Split(line, " ")
		ms.a = uint64(parseInt(abc[0]))
		ms.b = uint64(parseInt(abc[1]))
		ms.c = uint64(parseInt(abc[2]))
		line = readLine(reader)
		xy := strings.Split(line, " ")
		ms.x = uint64(parseInt(xy[0]))
		ms.y = uint64(parseInt(xy[1]))

		list[i] = ms
	}

	return list
}

func parseInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
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

type MagicSequence struct {
	a, b, c uint64
	n       int
	x, y    uint64
}
