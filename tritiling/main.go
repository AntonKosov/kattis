package main

import "fmt"

func main() {
	widths := read()
	ans := Calc(widths)
	for _, a := range ans {
		fmt.Println(a)
	}
}

func Calc(widths []int) []int {
	cache := make([][]int, 0, len(widths)+1)
	firstCol := make([]int, 8)
	firstCol[7] = 1
	cache = append(cache, firstCol)
	answers := make([]int, 0, len(widths))
	for _, width := range widths {
		answers = append(answers, calc(width, &cache))
	}
	return answers
}

func calc(width int, cache *[][]int) int {
	// 01234567
	// _x_x_x_x
	// __xx__xx
	// ____xxxx
	for len(*cache) <= width {
		prev := (*cache)[len(*cache)-1]
		col := make([]int, 8)
		col[0] = prev[7]
		col[1] = prev[6]
		// It's impossible to generate 2 and 5, so they don't contribute to the result.
		//col[2] = prev[5]
		//col[5] = prev[2]
		col[3] = prev[4] + prev[7] // ignore
		col[4] = prev[3]
		col[6] = prev[1] + prev[7]
		col[7] = prev[0] + prev[3] + prev[6]
		*cache = append(*cache, col)
	}

	return (*cache)[width][7]
}

func read() (widths []int) {
	for {
		var width int
		fmt.Scanf("%d", &width)
		if width < 0 {
			return widths
		}
		widths = append(widths, width)
	}
}
